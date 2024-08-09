# Сервис домов

`config.yaml` - минимальный файл конфигурации для базы данных и приложения

## Запуск приложения

**Для корректной работы сервиса необходимо указать пароль от базы данных в файле `config.yaml` и  `docker-compose.yaml`-
`postgres`.**

**Миграции в ручную накатывать не нужно, все происходит автоматически.**

**Затем можно запустить сервис командой:**

```bash
docker compose up 
```

**Методы:**

1. **Регистрация нового пользователя**

   **Пример запроса:**
   ```bash
    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{
      "email": "test@test.com",
      "password": "123qweasd",
      "user_type": "moderator"
      }'
   ```

2. **Aутентификация пользователя**. <p>
   Немного изменил механизм логина. Вместо id пользователя вводится email и пароль.

   **Пример запроса:**
   ```bash
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{
      "email": "test@test.com",
      "password": "123qweasd",
      }'
   ```

3. ***Создание квартиры*** <p>
   *Для этого эндпоинта необходимо проставлять хэдер Authorization в формате: "Bearer + TOKEN"*

   **Пример запроса:**
   ```bash
    curl -X POST http://localhost:8080/flat/create \
      -H "Content-Type: application/json" \
    -H "Authorization": "Bearer TOKEN" \
      -d '{
      "house_id": 12345,
      "price": 10000,
      "rooms": 4
     }'
   ```

4. ***Создание дома*** <p>
   *Для этого эндпоинта необходимо проставлять хэдер Authorization в формате: "Bearer + TOKEN"*

   **Пример запроса:**
   ```bash
    curl -X POST http://localhost:8080/house/create \
      -H "Content-Type: application/json" \
      -H "Authorization": "Bearer TOKEN" \
      -d '{
      "address": "Лесная улица, 7, Москва, 125196",
      "year": 2000,
      "developer": "Мэрия города"
   }'
   ```

5. ***Обновление статуса квартиры*** <p>
   *Для этого эндпоинта необходимо проставлять хэдер Authorization в формате: "Bearer + TOKEN"*

   **Пример запроса:**
   ```bash
    curl -X POST http://localhost:8080/house/create \
      -H "Content-Type: application/json" \
      -H "Authorization": "Bearer TOKEN" \
      -d '{
      "id": 123456,
      "status": "approved"
   }'
   ```

## Дополнения к решению

1. ***Для логина пользователя вместо id и пароля применяется почта и пароль***
2. ***Добавил сохранения номера квартиры в базу, который так же указывается при создании квартиры. В схеме OpenAPI их не было.***