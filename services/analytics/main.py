import asyncio
from fastapi import FastAPI
from contextlib import asynccontextmanager
from database import init_db, close_db, get_click_data, get_slugs
from consumer import start_consumer
import os
from dotenv import load_dotenv

load_dotenv()
DATABASE_URL = os.getenv("DATABASE_URL")

@asynccontextmanager
async def lifespan(app: FastAPI):
	await init_db(DATABASE_URL)
	asyncio.create_task(start_consumer())
	yield
	await close_db()

app = FastAPI(lifespan=lifespan)


@app.get("/")
async def list_slugs():
	slugs = await get_slugs()
	return {"Slug List" : slugs}

@app.get("/analytics/{slug}")
async def click_data(slug: str):
	count =  await get_click_data(slug)
	return {"slug": slug, "clicks": count}