from dataclasses import dataclass
from model.message_type import MessageType


@dataclass
class MessageBody:
    content: str
    type: MessageType
