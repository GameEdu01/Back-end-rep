from cryptography import fernet # importing all the external packages
from fastapi import FastAPI
from fastapi.param_functions import Query
from psycopg import sql
from pydantic import BaseModel
from cryptography.fernet import Fernet
import psycopg as pg
import requests
import random
import string
import json
import time


key = b'yURsNoRMdtDMy8QUj-05B64K-5cvNaJ-VNxvQZOu154=' # Key used to encrypt secrets, will be hidden in the .env in future
fernet = Fernet(key)
      

class Config:

    def __init__(self, dir):
        
        self.configDir = dir

        self.configData = {}
        with open(dir, 'r') as file:
            self.configData = json.load(file)


class Secret:

    def get_secret(lenght, start):

        secret = start
        for i in range(lenght):
            if not i == 0 and not i == lenght-1:
                if random.randint(1, 6) == 1:
                    secret += "."
                    continue
            secret += random.choice(string.ascii_letters)

        dt = time.time()

        return secret, dt


class DBConnector: # Connetion to the database class

    def __init__(self, dbname, user, password, address):

        self.dbname = dbname
        self.user = user
        self.password = password
        self.address = address

        self.conn = self.connect()
        self.cursor = self.conn.cursor()

        self.queries = {"get_table": ["SELECT * FROM ", " ORDER BY id"],
                        "value_exists": ["SELECT EXISTS(SELECT 1 FROM ", " WHERE ", " = '", "')"],
                        "get_columns": ["SELECT column_name FROM information_schema.columns WHERE table_schema = 'public' AND table_name = '", "'"],
                        "get_user": ["SELECT * FROM users WHERE username = '", "'"]}
        
    def connect(self): # getting the connection

        conn = None

        try:
            conn = pg.connect(dbname=self.dbname, user=self.user, password=self.password,
                       host=self.address[0], port=self.address[1])
        except Exception as err:
            print(err)

        return conn

    def get_table(self, tableName): # get table from database by name, returns json
        
        logins = []

        queryRequestList = self.queries["get_table"]
        query = queryRequestList[0] + tableName + queryRequestList[1]
        
        self.cursor.execute(query)
        records = self.cursor.fetchall()

        queryRequestList = self.queries["get_columns"]
        query = queryRequestList[0] + tableName + queryRequestList[1]

        self.cursor.execute(query)
        columns = self.cursor.fetchall()

        for row in records:
            rowJson = {}
            for i in range(len(row)):
                rowJson[str(columns[i][0])] = (row[i])
            logins.append(rowJson)

        return logins

    def value_exists(self, tableName, value, key): # check if value exist in certain place

        valueExists = False

        queryRequestList = self.queries["value_exists"]
        query = queryRequestList[0] + tableName + queryRequestList[1] + key + queryRequestList[2] + value + queryRequestList[3]

        self.cursor.execute(query)
        valueExists = self.cursor.fetchall()[0][0]

        return valueExists

    def get_user_data(self, username): # get user from users by it's username

        queryRequestList = self.queries["get_user"]
        query = queryRequestList[0] + username + queryRequestList[1]

        self.cursor.execute(query)
        userData = self.cursor.fetchall()[0]
        userDataJson = {}

        queryRequestList = self.queries["get_columns"]
        query = queryRequestList[0] + "users" + queryRequestList[1]

        self.cursor.execute(query)
        columns = self.cursor.fetchall()

        for i in range(len(userData)):
            userDataJson[str(columns[i][0])] = userData[i]

        return userDataJson


# NOTE i am planning to make a single function instead of user_exists() and email_exists()
# PS DONE


class APIConnecter: # Requests to already existing api class, but i think everything should've been written in the fastapi

    def __init__(self, address):

        self.address = address
        
        self.postLocations = {"login": "api/user/login"}


config = Config("config.json")
app = FastAPI()
apic = APIConnecter(r"https://gameedu.herokuapp.com/")
dbc = DBConnector(dbname="d2d1ljqhqhl34q",
                  user="udmehkiskcczbm",
                  password="d4f6d3d3a48a96f498f7829d75ef285bd9777989c15a135aa5a72903fc86127e",
                  address=("ec2-54-161-164-220.compute-1.amazonaws.com", "5432"))


class Login(BaseModel): # Login basic model
    username: str
    password: str


class SignUp(BaseModel): # SignUp basic model
    username: str
    password: str
    email: str


class GetUser(BaseModel):
    username: str
    password: str


class AdminUserRequest(BaseModel):
    username: str
    password: str
    user_username: str


@app.post("/fastapi/") # Hello world!
def get_root(): 
    return {"Hello": "World!"}


@app.post("/api/table") # get table from the database by it's name
def get_table(table_name: str, getuser: GetUser):
    
    if not dbc.value_exists("users", getuser.username, "username"):
        return {"message": "This user does not exist"}

    user = dbc.get_user_data(getuser.username)
    
    if not user:
        return {"message": "Incorrect username or password"}
    if not fernet.decrypt(user["password"].encode()).decode() == getuser.password:
        return {"message": "Incorrect username or password"}
    if not user["is_superuser"]:
        return {"message": "This user is not a superuser, so he can't access it"}

    return dbc.get_table(table_name)

@app.post("/api/get_user_with_admin")
def get_user_with_admin(aur: AdminUserRequest):

    if not dbc.value_exists("users", aur.username, "username"):
        return {"message": "This admin user does not exist"}

    user = dbc.get_user_data(aur.username)
    
    if not user:
        return {"message": "Incorrect admin username or password"}
    if not fernet.decrypt(user["password"].encode()).decode() == aur.password:
        return {"message": "Incorrect admin username or password"}
    if not user["is_superuser"]:
        return {"message": "You don't have access to this request!"} 

    if not dbc.value_exists("users", aur.user_username, "username"):
        return {"message": "This user does not exist"}

    user = dbc.get_user_data(aur.user_username)

    return {aur.user_username: user}

@app.post("/api/login_with_go") # Login with Igor's go api, not synced with the rest of fastapi i wrote (takes a lot of time to run)
def login(login: Login):

    request = requests.post(apic.address + apic.postLocations["login"],
                            data={"username": login.username,
                                  "password": login.password})
    if request.ok:
        return request.json()

    return {"message": "There was an error login in, %s".format(request.status_code)}

# I created new "users" table for the signup function, i hope you like it so we can keep it for now

@app.post("/api/get_user")  # TODO make a demo of gettting user without password and email
def get_user(getuser: GetUser): # get user, you must know the password in order to access it

    if not dbc.value_exists("users", getuser.username, "username"):
        return {"message": "This user does not exist"}

    user = dbc.get_user_data(getuser.username)
    
    if not user:
        return {"message": "Incorrect username or password"}
    if not fernet.decrypt(user["password"].encode()).decode() == getuser.password:
        return {"message": "Incorrect username or password"}

    return {getuser.username: user}


@app.post("/api/get_user_demo")
def get_user_demo(username: str):

    if not dbc.value_exists("users", username, "username"):
        return {"message": "This user does not exist"}

    user = dbc.get_user_data(username)

    del user["password"]
    del user["token"]
    del user["session"]
    del user["session_expire"]

    return {username: user}


@app.post("/api/signup") # Signing up, all secrity measures are in there, basic syntax check is also included
def signup(signup: SignUp): 

    if " " in signup.email or not "@" in signup.email:
        return {"message": "Please use an appropriate email address"}
    if " " in signup.username:
        return {"message": "Username can not contain spaces"}
    if " " in signup.password:
        return {"message": "Password can not contain spaces"}

    userExists = dbc.value_exists("users", signup.username, "username")
    emailExists = dbc.value_exists("users", signup.email, "email")
    if userExists or emailExists:
        return {"message": "This username or email already exists!"}

    signup.password = fernet.encrypt(signup.password.encode()).decode()

    token, _ = Secret.get_secret(60, "TTT")
    session, time_created = Secret.get_secret(30, "")

    dbc.cursor.execute("""INSERT INTO users (username, password, email, token, is_superuser, session, session_expire)
        VALUES ('{}', '{}', '{}', '{}', false, '{}', {})""".format(signup.username, signup.password, signup.email, token, session, time_created + config.configData["session_lenght"]))

    dbc.conn.commit() # Session expire is a time.time() object, i think it's easier to work with it

    return {"message": f"{signup.username} was added to database",
            "token": token}


# If you want to run this script localy on your machine:
# uvicorn api:app --reload      

# Then go into http://127.0.0.1:800/docs and you'll see the documentation

# If even after all you saw here you still don't want to integrate whole api into fastapi, please make your
# own documentation (and again, FastAPI makes it automatically), or just explain me a bit more about 
# your api, so it's easier for me to understand how to work with it 

# UPDATE: nope, i am rewritng all of it with fastapi anyway :D