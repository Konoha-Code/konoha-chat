#!/bin/bash

pip install -r requirements.txt && uvicorn app:app --reload --port 8080 --host 0.0.0.0