import httpx


class SessionExpired(Exception):
    pass


class ApiError(Exception):
    def __init__(self, status: int, message: str):
        self.status = status
        self.message = message
        super().__init__(f"[{status}] {message}")


class SubTrackerClient:
    def __init__(self, base_url: str):
        self._client = httpx.AsyncClient(base_url=base_url, timeout=10.0,trust_env=False,)

    async def close(self):
        await self._client.aclose()

    # ---------- auth ----------

    async def register(self, username: str, password: str) -> None:
        r = await self._client.post(
            "/auth/register",
            json={"username": username, "password": password},
        )
        if r.status_code == 409:
            raise ApiError(409, "Пользователь с таким именем уже существует")
        if r.status_code >= 400:
            raise self._err(r)

    async def login(self, username: str, password: str) -> str:
        r = await self._client.post(
            "/auth/login",
            json={"username": username, "password": password},
        )
        if r.status_code == 401:
            raise ApiError(401, "Неверный логин или пароль")
        if r.status_code >= 400:
            raise self._err(r)
        return r.json()["session_id"]

    # ---------- subscriptions ----------

    async def list_subscriptions(self, session_id: str) -> list[dict]:
        r = await self._client.get("/subscriptions", headers=self._auth(session_id))
        self._check(r)
        return r.json().get("subscriptions", [])

    async def create_subscription(
        self, session_id: str, title: str, cost: float, date: str
    ) -> dict:
        r = await self._client.post(
            "/subscriptions",
            headers=self._auth(session_id),
            json={"title": title, "cost": cost, "date": date},
        )
        self._check(r)
        if r.status_code >= 400:
            raise self._err(r)
        return r.json()

    async def delete_subscription(self, session_id: str, sub_id: str) -> None:
        r = await self._client.delete(
            f"/subscriptions/{sub_id}",
            headers=self._auth(session_id),
        )
        self._check(r)
        if r.status_code == 404:
            raise ApiError(404, "Подписка не найдена")
        r.raise_for_status()

    # ---------- helpers ----------

    @staticmethod
    def _auth(session_id: str) -> dict[str, str]:
        return {"Authorization": f"Bearer {session_id}"}

    @staticmethod
    def _check(r: httpx.Response):
        if r.status_code == 401:
            raise SessionExpired()

    @staticmethod
    def _err(r: httpx.Response) -> ApiError:
        try:
            msg = r.json().get("error", r.text)
        except Exception:
            msg = r.text
        return ApiError(r.status_code, msg)