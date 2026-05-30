import asyncio
import json
from database import init_db, insert_click
from aiokafka import AIOKafkaConsumer

async def start_consumer():
	consumer = AIOKafkaConsumer(
		"click-events",
		bootstrap_servers="localhost:9092"
	)
	await consumer.start()
	async for msg in consumer:
		data = json.loads(msg.value.decode("utf-8"))
		await insert_click(data["slug"])

if __name__ == "__main__":
    asyncio.run(start_consumer())
