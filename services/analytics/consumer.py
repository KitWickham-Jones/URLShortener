import asyncio
import json
import os
from database import insert_click
from aiokafka import AIOKafkaConsumer

async def start_consumer():
	consumer = AIOKafkaConsumer(
		"click-events",
		bootstrap_servers=os.getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
	)
	await consumer.start()
	async for msg in consumer:
		data = json.loads(msg.value.decode("utf-8"))
		await insert_click(data["slug"])

if __name__ == "__main__":
    asyncio.run(start_consumer())
