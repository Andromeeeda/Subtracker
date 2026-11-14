from aiogram import F, Router
from aiogram.filters import Command
from aiogram.fsm.context import FSMContext
from aiogram.fsm.state import State, StatesGroup
from aiogram.types import (
    CallbackQuery, InlineKeyboardButton, InlineKeyboardMarkup, Message,
)

from bot.api_client import ApiError, SessionExpired, SubTrackerClient
from bot.session import get_session_id

router = Router()


# ---------- FSM для /add ----------

class AddSub(StatesGroup):
    title = State()
    cost = State()
    date = State()


# ---------- helpers ----------

def _fmt(subs: list[dict]) -> str:
    if not subs:
        return "У вас пока нет подписок.\nДобавить: /add"
    lines = ["📋 <b>Ваши подписки:</b>\n"]
    for s in subs:
        lines.append(f"• <b>{s['title']}</b> — {s['cost']:.2f} ₽ (до {s['date']})")
    return "\n".join(lines)


def _kb(subs: list[dict]) -> InlineKeyboardMarkup:
    return InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text=f"🗑 {s['title']}", callback_data=f"del:{s['id']}")]
        for s in subs
    ])


# ---------- /subs ----------

@router.message(Command("subs"))
async def cmd_subs(message: Message, client: SubTrackerClient):
    sid = get_session_id(message.chat.id)
    if not sid:
        await message.answer("Сначала /login")
        return

    try:
        subs = await client.list_subscriptions(sid)
    except SessionExpired:
        await message.answer("Сессия невалидна. Выполните /login.")
        return
    except ApiError as e:
        await message.answer(f"❌ {e.message}")
        return

    await message.answer(_fmt(subs), parse_mode="HTML",
                         reply_markup=_kb(subs) if subs else None)


# ---------- /add (FSM) ----------

@router.message(Command("add"))
async def cmd_add(message: Message, state: FSMContext):
    if not get_session_id(message.chat.id):
        await message.answer("Сначала /login")
        return
    await state.set_state(AddSub.title)
    await message.answer("Название подписки? (например, Netflix)\n/cancel — отмена")


@router.message(AddSub.title)
async def add_title(message: Message, state: FSMContext):
    title = (message.text or "").strip()
    if not title:
        await message.answer("Название не может быть пустым. Попробуйте ещё раз.")
        return
    if len(title) > 255:
        await message.answer("Слишком длинное название (макс. 255).")
        return

    await state.update_data(title=title)
    await state.set_state(AddSub.cost)
    await message.answer("Стоимость? (например, 15.99)\n/cancel — отмена")


@router.message(AddSub.cost)
async def add_cost(message: Message, state: FSMContext):
    text = (message.text or "").strip().replace(",", ".")
    try:
        cost = float(text)
        if cost < 0:
            raise ValueError
    except ValueError:
        await message.answer("Введите положительное число, например 15.99")
        return

    await state.update_data(cost=cost)
    await state.set_state(AddSub.date)
    await message.answer("Дата списания? (формат <code>YYYY-MM-DD</code>, например 2026-10-01)\n/cancel — отмена",
                         parse_mode="HTML")


@router.message(AddSub.date)
async def add_date(message: Message, state: FSMContext, client: SubTrackerClient):
    from datetime import datetime
    text = (message.text or "").strip()
    try:
        datetime.strptime(text, "%Y-%m-%d")
    except ValueError:
        await message.answer("Неверный формат. Нужно <code>YYYY-MM-DD</code>, например 2026-10-01.",
                             parse_mode="HTML")
        return

    data = await state.get_data()
    await state.clear()

    sid = get_session_id(message.chat.id)
    if not sid:
        await message.answer("Сессия потерялась. /login")
        return

    try:
        sub = await client.create_subscription(
            sid, data["title"], data["cost"], text
        )
    except SessionExpired:
        await message.answer("Сессия невалидна. /login")
        return
    except ApiError as e:
        await message.answer(f"❌ {e.message}")
        return

    await message.answer(
        f"✅ Подписка добавлена:\n"
        f"<b>{sub['title']}</b> — {sub['cost']:.2f} ₽ (до {sub['date']})",
        parse_mode="HTML",
    )


@router.message(Command("cancel"))
async def cmd_cancel(message: Message, state: FSMContext):
    if await state.get_state() is None:
        await message.answer("Нечего отменять.")
        return
    await state.clear()
    await message.answer("Отменено.")


# ---------- /del ----------

@router.callback_query(F.data.startswith("del:"))
async def cb_delete(callback: CallbackQuery, client: SubTrackerClient):
    chat_id = callback.message.chat.id
    sid = get_session_id(chat_id)
    if not sid:
        await callback.answer("Сначала /login", show_alert=True)
        return

    sub_id = callback.data.split(":", 1)[1]
    try:
        await client.delete_subscription(sid, sub_id)
    except SessionExpired:
        await callback.answer("Сессия невалидна, /login", show_alert=True)
        return
    except ApiError as e:
        await callback.answer(f"Ошибка: {e.message}", show_alert=True)
        return

    await callback.answer("Удалено")

    try:
        subs = await client.list_subscriptions(sid)
        await callback.message.edit_text(
            _fmt(subs), parse_mode="HTML",
            reply_markup=_kb(subs) if subs else None,
        )
    except Exception:
        pass