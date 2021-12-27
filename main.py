from fastapi import FastAPI
import psycopg as pg

# testing
class DBConnector:

    def __init__(self, dbname, user, password, address):

        self.dbname = dbname
        self.user = user
        self.password = password
        self.address = address

        self.conn = self.connect()
        self.cursor = self.conn.cursor()

        self.queries = {"get_table": ["SELECT * FROM ", " ORDER BY id"]}
        
    def connect(self):

        conn = None

        try:
            conn = pg.connect(dbname=self.dbname, user=self.user, password=self.password,
                       host=self.address[0], port=self.address[1])
        except Exception as err:
            print(err)

        return conn

    def get_table(self, tableName):
        
        logins = []

        query = self.queries["get_table"][0] + tableName + self.queries["get_table"][1]
        
        self.cursor.execute(query)
        records = self.cursor.fetchall()

        for row in records:
            logins.append({"id": row[0], "username": row[1], "password": row[2]})

        return logins


dbc = DBConnector(dbname="d2d1ljqhqhl34q",
                    user="udmehkiskcczbm",
                    password="d4f6d3d3a48a96f498f7829d75ef285bd9777989c15a135aa5a72903fc86127e",
                    address=("ec2-54-161-164-220.compute-1.amazonaws.com", "5432"))

app = FastAPI()


@app.get("/")
def get_root():
    return {"Hello": "World!"}


@app.get("/table")
def get_root(table_name: str):
    return dbc.get_table(table_name)


