
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
  - Success (201): `{"succesfully text": "string"}`
  - Error (400/409,500): `{"error": "string"}`

- **POST /auth/login** - вход пользователя
  - Body: `{"username": "string", "password": "string"}`
  - Success (200): `{"session_id": "string"}`
  - Error (400/401,500): `{"error": "string"}

  # CRUD операции 
  # Операции над пользователями 
  - GET /users/me - получить пользователя
  - Success (200): `{"id","username",}`
  - Error (401,500)

   - POST/users/me/password - изменение пароля с проверкой 
  - Body: `{"current_password" : "string", "new_password": "string"}
  - Success (200): `{"{}"}
  - Error (400,401,500): `{"error": "string"}

  - DELETE /users/me - удаление пользователя 
  - Body: `{"password" : "string}
  - Success(200): `{}`
  - Error(400,401,500):  `{"error": "string"}

  # Операции над подписками 
  GET /subscriptions - получить все подписки 
  -Success(200): `{"subscriptions": [{"id": "uuid", "title": "string", "cost": 15.99, "date": "2026-10-01"}]}`

  GET /subscriptions/{id} - получить конкретную подписку 
  - Success(200), `{"id": "uuid", "title": "string", "cost": 15.99, "date": "2026-10-01"}`,
  - Erorr(400,401,404,500):  `{"error": "string"}

  POST /subscriptions - создать подписку 
  Body: `{"title":"string","cost":"float","date": "date"}`
  Succed(201): `{"id":"uuid","title":"string","cost":"float","date": "date"}`
  Error(400,409,500):  `{"error": "string"}

  
 PATCH /subscriptions/{id}  - изменить подписку частично
 - Body: `{"title": string}`  //как пример
 -Succed(200): {"id": "uuid" , "title": string,"cost":"float","date": "date"} 
 - Error(400,404,500): `{"error": "string"}


   DELETE /subscriptions/{id} - удалить подписку 
   -Success(200): `{}`
   -Erorr(400,404,500): `{"error": "string"}` 

   GET /subscriptions/stats  - получить статистику по подпискам 


  Реализовать фичи:
  1. middleware для аутентификации и прочего +
  2. Дописать круды для пользователей + 
  3. Сделать логи +
  4. Чуть переделать обработку ошибок + 
  5. разработать окружение проекта 
  6. Сделать бд Постгес
  7. Написать фронтенд + 

  //протестировать GET /subscriptions там чуть проблема 
  //в DELETE user нужно проверить удаляется ли сессия так как сам аккаунт по сути удаляется и удалить для этой сессии 
  // и этого пользователя его подписки 

  //Исправить что при выходе из аккаунта удаляются все подписки пользователя 
  


