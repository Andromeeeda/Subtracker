import logging
from aiogram import Router 
from aiogram.filters import Command
from aiogram.types import Message

router = Router()

logger = logging.getLogger(__name__)

@router.message(Command("start"))
async def cmd_start(message: Message):
     logger.info("Получено /start от chat_id=%s", message.chat.id)
     await message.answer("Привет! Я бот SubTracker.\n"
     "/register логин пароль — создать аккаунт\n"
        "/login логин пароль — войти\n"
        "/logout — выйти из бота\n\n"
        "<b>Подписки:</b>\n"
        "/subs — список\n"
        "/add — добавить\n"
        "🗑 — удалить (кнопка под подпиской)",
        parse_mode="HTML",)

@router.message(Command("help"))
async def cmd_help(message: Message):
    await message.answer("Команды:\n/start — начать\n/help — помощь")