import asyncpg 
import asyncio


async def init_db(dsn: str):
	global pool
	pool = await asyncpg.create_pool(dsn)


async def insert_click(slug: str):
	async with pool.acquire() as conn:
		await conn.execute(
			"INSERT INTO clicks (slug) VALUES ($1)",
			slug
		)
	print(f"Inserted click for {slug}")

# async def main():
# 	await init_db("postgresql://admin:secret@localhost:5432/urlshortener")
# 	await insert_click("abc12")


# asyncio.run(main())