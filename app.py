from cryptography import fernet
from fastapi import FastAPI
from fastapi.param_functions import Query
from pydantic import BaseModel
import psycopg as pg
from cryptography.fernet import Fernet
import requests
from requests.exceptions import ReadTimeout


key = b'yURsNoRMdtDMy8QUj-05B64K-5cvNaJ-VNxvQZOu154=' # Will be hidden in the future
fernet = Fernet(key)

class DBConnector: # Connetion to the database

    def __init__(self, dbname, user, password, address):

        self.dbname = dbname
        self.user = user
        self.password = password
        self.address = address

        self.conn = self.connect()
        self.cursor = self.conn.cursor()

        self.queries = {"get_table": ["SELECT * FROM ", " ORDER BY id"],
                        "username_exists": ["SELECT EXISTS(SELECT 1 FROM ", " WHERE username = '", "')"]}
        
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

        for row in records:
            logins.append({"id": row[0], "username": row[1], "password": row[2]})

        return logins

    def user_exists(self, tableName, username): # check if username exist in certain database

        userExists = False

        query = self.queries["username_exists"][0] + tableName + self.queries["username_exists"][1] + username + self.queries["username_exists"][2]

        self.cursor.execute(query)

        userExists = self.cursor.fetchall()[0][0]

        return userExists


class APIConnecter: # Requests to already existing api, but i think everything should've been written in the fastapi

    def __init__(self, address):

        self.address = address
        
        self.postLocations = {"login": "api/user/login"}


app = FastAPI()
apic = APIConnecter(r"https://gameedu.herokuapp.com/")
dbc = DBConnector(dbname="d2d1ljqhqhl34q",
                  user="udmehkiskcczbm",
                  password="d4f6d3d3a48a96f498f7829d75ef285bd9777989c15a135aa5a72903fc86127e",
                  address=("ec2-54-161-164-220.compute-1.amazonaws.com", "5432"))


class Login(BaseModel):
    username: str
    password: str


class SignUp(BaseModel):
    username: str
    password: str
    email: str


@app.get("/fastapi/")
def get_root(): 
    return {"Hello": "World!"}


@app.post("/fastapi/table")
def get_table(table_name: str):
    return dbc.get_table(table_name)


@app.post("/fastapi/login_with_go")
def login(login: Login):

    request = requests.post(apic.address + apic.postLocations["login"],
                            data={"username": login.username,
                                  "password": login.password})
    if request.ok:
        return request.json()

    return {"message": "There was an error login in, %s".format(request.status_code)}


@app.post("/fastapi/signup")
def signup(signup: SignUp):

    userExists = dbc.user_exists("users", signup.username)
    if userExists:
        return {"message": "This username already exists!"}

    signup.password = fernet.encrypt(signup.password.encode())
    signup.password = signup.password.decode()

    dbc.cursor.execute("""INSERT INTO users (username, password, email)
                            VALUES ('{}', '{}', '{}')""".format(signup.username, signup.password, signup.email))

    dbc.conn.commit()

    return {"message": f"{signup.username} was added to database.users"}
