import asyncio
import logging

from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode

from bot.config import BOT_TOKEN,API_BASE_URL

from bot.handlers import startHandlers
from bot.handlers import auth,subscriptions 
from bot.api_client import SubTrackerClient

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
)
logger = logging.getLogger(__name__)


async def main():
    logger.info("Запуск бота...")
    bot = Bot(
        token=BOT_TOKEN,
        default=DefaultBotProperties(parse_mode=ParseMode.HTML),
    )

    me = await bot.get_me()
    logger.info("Бот запущен: @%s (id=%s)", me.username, me.id)

    dp = Dispatcher()

    client = SubTrackerClient(API_BASE_URL)
    
    dp["client"] = client


    dp.include_router(startHandlers.router)
    dp.include_router(auth.router)
    dp.include_router(subscriptions.router)

    logger.info("Начинаю polling...")
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())