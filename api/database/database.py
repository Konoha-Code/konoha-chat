import os
import psycopg2


class Database:
    _db = None

    def _get_instance(self):
        if self._db is None:
            database = os.environ.get("POSTGRES_DB")
            username = os.environ.get("POSTGRES_USER")
            password = os.environ.get("POSTGRES_PASSWORD")
            host = os.environ.get("POSTGRES_HOST")
            port = os.environ.get("POSTGRES_PORT")

            self._db = psycopg2.connect(
                f"dbname={database} user={username} password={password} host={host} port={port}"
            )
        return self._db

    def select(self, query, params=None):
        return self._execute(query, params)

    def insert(self, query, params=None):
        return self._execute(f"{query} returning id", params, inserting=True)

    def update(self, query, params=None):
        return self._execute(query, params, updating=True)

    def _execute(self, query, params=None, inserting=False, updating=False):
        cursor = self._get_instance().cursor()
        cursor.execute(query, params)

        if inserting:
            insert_id = cursor.fetchone()[0]
            self._get_instance().commit()
            return {"id": f"{insert_id}"}
        elif updating:
            affected_rows = cursor.rowcount
            self._get_instance().commit()
            return {"id": f"{affected_rows}"}
        else:
            return cursor.fetchall()
