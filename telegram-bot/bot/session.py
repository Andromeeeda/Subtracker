import json
from pathlib import Path

_PATH = Path("sessions.json")
_sessions: dict[int, str] = {}


def _load() -> None:
    global _sessions
    if _PATH.exists():
        try:
            raw = json.loads(_PATH.read_text(encoding="utf-8"))
            _sessions = {int(k): v for k, v in raw.items()}
        except Exception:
            _sessions = {}


def _save() -> None:
    _PATH.write_text(
        json.dumps({str(k): v for k, v in _sessions.items()}, ensure_ascii=False, indent=2),
        encoding="utf-8",
    )


def get_session_id(chat_id: int) -> str | None:
    return _sessions.get(chat_id)


def set_session_id(chat_id: int, session_id: str) -> None:
    _sessions[chat_id] = session_id
    _save()


def clear_session_id(chat_id: int) -> None:
    _sessions.pop(chat_id, None)
    _save()


_load()