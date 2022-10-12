CREATE TABLE tb_user (
    id SERIAL PRIMARY KEY,
    tx_name TEXT NOT NULL
);

CREATE TABLE tb_message (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    recipient_id INTEGER NOT NULL,
    tx_message TEXT NOT NULL,
    id_message_type TEXT NOT NULL DEFAULT 'plain/text',
    dt_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    dt_readed_at TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT fk_sender_id FOREIGN KEY (sender_id) REFERENCES tb_user (id),
    CONSTRAINT fk_recipient_id FOREIGN KEY (recipient_id) REFERENCES tb_user (id)
);