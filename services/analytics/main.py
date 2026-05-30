import asyncio
from fastapi import FastAPI
from contextlib import asynccontextmanager
from database import init_db, close_db, get_click_data
from consumer import start_consumer


@asynccontextmanager
async def lifespan(app: FastAPI):
	await init_db("postgresql://admin:secret@localhost:5432/urlshortener")
	asyncio.create_task(start_consumer())
	yield
	await close_db()

app = FastAPI(lifespan=lifespan)

@app.get("/analytics/{slug}")
async def click_data(slug: str):
	count =  await get_click_data(slug)
	return {"slug": slug, "clicks": count}