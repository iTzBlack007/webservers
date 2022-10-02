# Run this via
# python3 -m uvicorn main:app

from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles

app = FastAPI()

app.mount("/", StaticFiles(directory="."), name="static")
