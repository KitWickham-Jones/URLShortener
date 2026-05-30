import asyncio
import json
from database import init_db, insert_click
from aiokafka import AIOKafkaConsumer

async def start_consumer():
	await init_db("postgresql://admin:secret@localhost:5432/urlshortener")
	consumer = AIOKafkaConsumer(
		"click-events",
		bootstrap_servers="localhost:9092"
	)
	await consumer.start()
	async for msg in consumer:
		data = json.loads(msg.value.decode("utf-8"))
		await insert_click(data["slug"])


asyncio.run(start_consumer())