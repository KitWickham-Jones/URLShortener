import asyncio
from aiokafka import AIOKafkaConsumer

async def start_consumer():
	consumer = AIOKafkaConsumer(
		"click-events",
		bootstrap_servers="localhost:9092"
	)
	await consumer.start()
	async for msg in consumer:
		print(msg.value)

asyncio.run(start_consumer())