from cryptography import fernet # importing all the external packages
from fastapi import FastAPI
from fastapi.param_functions import Query
from pydantic import BaseModel
import psycopg as pg
from cryptography.fernet import Fernet
import requests


key = b'yURsNoRMdtDMy8QUj-05B64K-5cvNaJ-VNxvQZOu154=' # Key used to encrypt secrets, will be hidden in the .env in future
fernet = Fernet(key)


class DBConnector: # Connetion to the database class

    def __init__(self, dbname, user, password, address):

        self.dbname = dbname
        self.user = user
        self.password = password
        self.address = address

        self.conn = self.connect()
        self.cursor = self.conn.cursor()

        self.queries = {"get_table": ["SELECT * FROM ", " ORDER BY id"],
                        "username_exists": ["SELECT EXISTS(SELECT 1 FROM ", " WHERE username = '", "')"],
                        "email_exists": ["SELECT EXISTS(SELECT 1 FROM ", " WHERE email = '", "')"],
                        "get_columns": ["SELECT column_name FROM information_schema.columns WHERE table_schema = 'public' AND table_name = '", "'"]}
        
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

        query = self.queries["get_table"][0] + tableName + self.queries["get_table"][1]
        
        self.cursor.execute(query)
        records = self.cursor.fetchall()

        query = self.queries["get_columns"][0] + tableName + self.queries["get_columns"][1]

        self.cursor.execute(query)
        columns = self.cursor.fetchall()

        for row in records:
            rowJson = {}
            for i in range(len(row)):
                rowJson[str(columns[i][0])] = (row[i])
            logins.append(rowJson)

        return logins

    def user_exists(self, tableName, username): # check if username exist in certain database

        userExists = False

        query = self.queries["username_exists"][0] + tableName + self.queries["username_exists"][1] + username + self.queries["username_exists"][2]

        self.cursor.execute(query)

        userExists = self.cursor.fetchall()[0][0]

        return userExists

    def email_exists(self, tableName, email): # check if email exist in certain database

        emailExists = False

        query = self.queries["email_exists"][0] + tableName + self.queries["email_exists"][1] + email + self.queries["email_exists"][2]

        self.cursor.execute(query)

        emailExists = self.cursor.fetchall()[0][0]

        return emailExists


# NOTE i am planning to make a single function instead of user_exists() and email_exists()


class APIConnecter: # Requests to already existing api class, but i think everything should've been written in the fastapi

    def __init__(self, address):

        self.address = address
        
        self.postLocations = {"login": "api/user/login"}


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


@app.get("/fastapi/") # Hello world!
def get_root(): 
    return {"Hello": "World!"}


@app.post("/fastapi/table") # get table from the database by it's name
def get_table(table_name: str):
    return dbc.get_table(table_name)


@app.post("/fastapi/login_with_go") # Login with Igor's go api, not synced with the rest of fastapi i wrote (takes a lot of time to run)
def login(login: Login):

    request = requests.post(apic.address + apic.postLocations["login"],
                            data={"username": login.username,
                                  "password": login.password})
    if request.ok:
        return request.json()

    return {"message": "There was an error login in, %s".format(request.status_code)}

# I created new "users" table for the signup function, i hope you like it so we can keep it for now

@app.post("/fastapi/signup") # Signing up, all secrity measures are in there, basic syntax check is also included
def signup(signup: SignUp):

    if " " in signup.email or not "@" in signup.email:
        return {"message": "Please use an appropriate email address"}
    if " " in signup.username:
        return {"message": "Username can not contain spaces"}
    if " " in signup.password:
        return {"message": "Password can not contain spaces"}

    userExists = dbc.user_exists("users", signup.username)
    emailExists = dbc.email_exists("users", signup.email)
    if userExists or emailExists:
        return {"message": "This username or email already exists!"}

    signup.password = fernet.encrypt(signup.password.encode()).decode()

    token = Fernet.generate_key().decode()

    dbc.cursor.execute("""INSERT INTO users (username, password, email, token)
                            VALUES ('{}', '{}', '{}', '{}')""".format(signup.username, signup.password, signup.email, token))

    dbc.conn.commit()

    return {"message": f"{signup.username} was added to database",
            "token": token}


# If you want to run this script localy on your machine:
# uvicorn app:app --reload      

# Then go into http://127.0.0.1:800/docs and you'll see the documentation

# If even after all you saw here you still don't want to integrate whole api into fastapi, please make your
# own documentation (and again, FastAPI makes it automatically), or just explain me a bit more about 
# your api, so it's easier for me to understand how to work with it 