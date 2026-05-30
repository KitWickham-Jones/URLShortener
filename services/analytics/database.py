import asyncpg 

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

async def get_slugs():
	async with pool.acquire() as conn:
		return await conn.fetch(
			"SELECT DISTINCT slug FROM clicks"
		)

async def get_click_data(slug: str):
	async with pool.acquire() as conn:
		return await conn.fetchval(
			"SELECT COUNT(*) FROM clicks WHERE slug =$1", slug
		)

async def close_db():
	await pool.close()
