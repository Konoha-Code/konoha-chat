from fastapi import FastAPI
from model.message_body import MessageBody

from controller.controller import Controller

app = FastAPI()


@app.get("/ping")
async def ping():
    return {"message": "pong"}


@app.get("/users/")
async def get_users():
    """
    Get all users from the database
    """
    return Controller.get_all_users()


@app.get("/messages/")
async def get_messages(
    sender_id: int = None, recipient_id: int = None, unreaded_only: bool = False
):
    """
    Get all messages from the database for a specific user to another user if none is provided, get all messages from the user to all other users,
    if sender or recipient is invalid, return an empty list
    """
    return Controller.get_messages(sender_id, recipient_id, unreaded_only)


@app.post("/message/")
async def post_message(sender_id: int, recipient_id: int, message: MessageBody):
    """
    Post a message to the database
    """
    return Controller.post_message(sender_id, recipient_id, message)


@app.post("/user/")
async def post_user(name: str):
    """
    Post a user to the database
    """
    return Controller.post_user(name)


@app.patch("/set_read/")
async def set_read(messages_id: list):
    """
    Set a message or list of message as read, invalid message id is ignored, if the message is already read, it is ignored
    """
    return Controller.set_read(messages_id)
