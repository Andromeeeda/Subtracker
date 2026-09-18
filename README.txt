
# Трекер подписок 
Сервис позволяет мониторить свои подписки даты их списания и аналитику подписок 

 Технологический стек
Категория	Технология
🔧 Контроль версий	Git
🐹 Язык разработки	Golang
🌐 API	REST API
🗄 База данных	PostgreSQL
🔄 Миграции	Database Migrations
🌍 Конфигурация	Переменные окружения
📝 Мониторинг	Логгирование
⚙️ Сборка	Makefile
🐳 Контейнеризация	Docker + Docker Compose

## REST API
# Аутентификация
  - **POST /auth/register** - регистрация нового пользователя
  - Body: `{"username": "string", "password": "string"}`
  - Success (201): `{"token": "string"}`
  - Error (400/409): `{"error": "string"}`

- **POST /auth/login** - вход пользователя
  - Body: `{"username": "string", "password": "string"}`
  - Success (200): `{"token": "string"}`
  - Error (400/401): `{"error": "string"}

  # CRUD операции 
  # Операции над пользователями 
  - GET /users/me - получить пользователя
  - Success (200): `{"id","username",}`
  - Error (401)

   - POST/users/me/change_password - изменение пароля с проверкой 
  - Body: `{"current_password" : "string", "new_password": "string"}
  - Success (200): `{"{}"}
  - Error (400,401): `{"error": "string"}

  - DELETE /users/me - удаление пользователя 
  - Body: `{"password" : "string}
  - Success(200): `{}`
  - Error(400,401):  `{"error": "string"}

  # Операции над подписками 
  GET /subscriptions - получить все подписки 
  -Success(200): `{"subscriptions": [{"id": "uuid", "title": "string", "cost": 15.99, "date": "2026-10-01"}]}`

  GET /subscriptions/{id} - получить конкретную подписку 
  - Success(200), `{"id": "uuid", "title": "string", "cost": 15.99, "date": "2026-10-01"}`,
  - Erorr(400,401,404):  `{"error": "string"}

  POST /subscriptions - создать подписку 
  Body: `{"title":"string","cost":"float","date": "date"}`
  Succed(201): `{"id":"uuid","title":"string","cost":"float","date": "date"}`
  Error(400,500):  `{"error": "string"}

  
 PATCH /subscriptions/{id}  - изменить подписку частично
 - Body: `{"title": string}`  //как пример
 -Succed(200): {"id": "uuid" , "title": string,"cost":"float","date": "date"} 
 - Error(400,404,500): `{"error": "string"}


   DELETE /subscriptions/{id} - удалить подписку 
   -Success(200): `{}`
   -Erorr(404,500): `{"error": "string"}` 

   GET /subscriptions/stats  - получить статистику по подпискам 







