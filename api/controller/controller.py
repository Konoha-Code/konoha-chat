from fastapi import HTTPException
from model.message_body import MessageBody
from model.message_type import MessageType
from database.database import Database
from model.user import User
from model.message import Message


class Controller:
    @staticmethod
    def get_all_users():
        try:
            database = Database()

            users = database.select("SELECT * FROM tb_user")

            return [User(*user) for user in users]
        except Exception as e:
            print(e)
            raise HTTPException(status_code=500, detail="Internal server error")

    @staticmethod
    def get_messages(
        sender_id: int = None, recipient_id: int = None, unreaded_only: bool = False
    ):
        try:
            database = Database()

            query = "SELECT * FROM tb_message"
            param_tuple = ()
            where_tuple = ()

            if sender_id is not None:
                where_tuple += ("sender_id = %s",)
                param_tuple += (sender_id,)
            if recipient_id is not None:
                where_tuple += ("recipient_id = %s",)
                param_tuple += (recipient_id,)

            if len(where_tuple) > 0:
                query += " WHERE "
                query += " AND ".join(where_tuple)

            if unreaded_only:
                if len(where_tuple) > 0:
                    query += " AND "
                else:
                    query += " WHERE "

                query += "dt_readed_at IS NULL"

            messages = database.select(query, param_tuple)

            return [Message(*message) for message in messages]
        except Exception as e:
            print(e)
            raise HTTPException(status_code=500, detail="Internal server error")

    @staticmethod
    def _user_exists(name: str):
        try:
            database = Database()

            return database.select("SELECT * FROM tb_user WHERE id = %s", (name,))
        except Exception as e:
            print(e)
            raise HTTPException(status_code=500, detail="Internal server error")

    @staticmethod
    def post_message(sender_id: int, recipient_id: int, message: MessageBody):
        if sender_id == recipient_id:
            raise HTTPException(
                status_code=400, detail="Sender and recipient cannot be the same"
            )

        if message.content.strip() == "":
            raise HTTPException(status_code=400, detail="Message cannot be empty")

        if message.type not in MessageType:
            raise HTTPException(
                status_code=400, detail="Message type must be one of the following: "
                + ", ".join(MessageType)
            )

        if not Controller._user_exists(sender_id):
            raise HTTPException(
                status_code=400, detail=f"Sender with id {sender_id} does not exist"
            )

        if not Controller._user_exists(recipient_id):
            raise HTTPException(
                status_code=400,
                detail=f"Recipient with id {recipient_id} does not exist",
            )

        try:
            database = Database()

            return database.insert(
                "INSERT INTO tb_message (sender_id, recipient_id, tx_message, id_message_type) VALUES (%s, %s, %s, %s)",
                (sender_id, recipient_id, message.content, message.type.value),
            )
        except Exception as e:
            print(e)
            raise HTTPException(status_code=500, detail="Internal server error")

    @staticmethod
    def post_user(name: str):
        try:
            database = Database()

            return database.insert("INSERT INTO tb_user (tx_name) VALUES (%s)", (name,))
        except Exception as e:
            print(e)
            raise HTTPException(status_code=500, detail="Internal server error")

    @staticmethod
    def set_read(messages_id: list):
        try:
            database = Database()

            return database.update(
                "UPDATE tb_message SET dt_readed_at = current_timestamp WHERE id IN %s and dt_readed_at is null",
                (tuple(messages_id),),
            )
        except Exception as e:
            print(e)
            raise HTTPException(status_code=500, detail="Internal server error")
