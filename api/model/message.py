from dataclasses import dataclass
from datetime import datetime
from model.message_body import MessageBody


@dataclass
class Message:
    id: int
    sender_id: int
    recipient_id: int
    message: MessageBody
    dt_created_at: datetime
    dt_readed_at: datetime

    def __init__(
        self,
        id,
        sender_id,
        recipient_id,
        tx_message,
        id_message_type,
        dt_created_at,
        dt_readed_at,
    ) -> None:
        self.id = id
        self.sender_id = sender_id
        self.recipient_id = recipient_id
        self.message = MessageBody(tx_message, id_message_type)
        self.dt_created_at = dt_created_at.isoformat()
        self.dt_readed_at = dt_readed_at.isoformat() if dt_readed_at else None
