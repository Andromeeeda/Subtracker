from aiogram import Router
from aiogram.filters import Command
from aiogram.types import Message

from bot.api_client import ApiError, SubTrackerClient
from bot.session import clear_session_id, get_session_id, set_session_id

router = Router()


@router.message(Command("register"))
async def cmd_register(message: Message, client: SubTrackerClient):
    args = (message.text or "").split(maxsplit=2)
    if len(args) != 3:
        await message.answer(
            "Формат: <code>/register логин пароль</code>\n"
            "⚠️ Пароль должен быть не короче 6 символов.",
            parse_mode="HTML",
        )
        return

    _, username, password = args

    try:
        await message.delete()
    except Exception:
        pass

    try:
        await client.register(username, password)
    except ApiError as e:
        await message.answer(f"❌ {e.message}")
        return

    await message.answer(
        "✅ Аккаунт создан!\n"
        "Теперь войдите: <code>/login логин пароль</code>",
        parse_mode="HTML",
    )


@router.message(Command("login"))
async def cmd_login(message: Message, client: SubTrackerClient):
    args = (message.text or "").split(maxsplit=2)
    if len(args) != 3:
        await message.answer(
            "Формат: <code>/login логин пароль</code>",
            parse_mode="HTML",
        )
        return

    _, username, password = args

    try:
        await message.delete()
    except Exception:
        pass

    try:
        sid = await client.login(username, password)
    except ApiError as e:
        await message.answer(f"❌ {e.message}")
        return

    set_session_id(message.chat.id, sid)
    await message.answer(
        f"✅ Вы вошли как <b>{username}</b>.\n"
        "Сессия сохранена — /login больше не нужен.",
        parse_mode="HTML",
    )


@router.message(Command("logout"))
async def cmd_logout(message: Message):
    clear_session_id(message.chat.id)
    await message.answer(
        "✅ Вы вышли из аккаунта в боте.\n"
    )