from enum import Enum


class MessageType(Enum):
    """
    Enum for message types.

    PlainText: plain/text

    Json: json

    Html: html

    Markdown: markdown

    Xml: xml
    """

    PLAIN_TEXT = "plain/text"
    JSON = "json"
    HTML = "html"
    MARKDOWN = "markdown"
    XML = "xml"
