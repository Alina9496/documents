# Banner service
# Пример использования API

## Регистрация пользователя

**Метод:** POST  
**URL:** http://localhost:8080/api/register  
**Заголовки:**
- `admin_token`: токен администратора

**Параметры формы:**
- `login`: Логин пользователя.
- `pswd`: Пароль пользователя.
**Заголовок:**
- `token`: Токен пользователя.

Пример использования cURL:

```bash
curl --location 'http://localhost:8080/api/register' \
--header 'admin_token: admin_token' \
--form 'login="user@mail.ru"' \
--form 'pswd="Passw_345"'
```

---

## Аутентификация пользователя

**Метод:** POST  
**URL:** http://localhost:8080/api/auth  

**Параметры формы:**
- `login`: Логин пользователя (электронная почта).
- `pswd`: Пароль пользователя.

Пример использования cURL:

```bash
curl --location 'http://localhost:8080/api/auth' \
--form 'login="asdDADd123"' \
--form 'pswd="Passw_345"'
```

---

Эти примеры показывают, как использовать команды cURL для регистрации нового пользователя и аутентификации существующего пользователя через API.

## Загрузка документа

**Метод:** POST  
**URL:** http://localhost:8080/api/docs  

**Параметры формы:**
- `meta`: JSON строка с метаданными документа.
  - `name`: Имя файла.
  - `file`: Флаг, указывающий, что загружается файл.
  - `public`: Флаг, определяющий публичность документа.
  - `token`: Токен пользователя.
  - `mime`: MIME-тип файла.
  - `grant`: Массив логинов пользователей, которым предоставляется доступ к документу.
- `file`: Путь к файлу на локальной машине.
**Заголовок:**
- `token`: Токен пользователя.

Пример использования cURL:

```bash
curl --location 'http://localhost:8080/api/docs' \
--header 'admin_token: admin_token' \
--form 'meta="{  \"name\": \"photo.jpg\",  \"file\": true,  \"public\": false,  \"token\": \"JTTLEqyIO1r6HIvSOESB\",  \"mime\": \"image/jpg\",  \"grant\": [    \"login\" ,\"login2\"  ]}"' \
--form 'file=@"/path"'
```

---

В этом примере показано, как загрузить документ с использованием команды cURL. Метаданные документа передаются в виде строки JSON, а сам файл передается через параметр `file`.

## Получение документа

**Метод:** GET  
**URL:** http://localhost:8080/api/docs/{document_id}


**Путь:**
- `{document_id}`: Идентификатор документа.
**Заголовок:**
- `token`: Токен пользователя.


Пример использования cURL:

```bash
curl --location 'http://localhost:8080/api/docs/1a394bd7-b384-4415-abfa-953ae26b3a4f'\
--header 'token: JTTLEqyIO1r6HIvSOESB'
```

---

Этот запрос используется для получения информации о документе по его уникальному идентификатору.

## Список документов

**Метод:** GET  
**URL:** http://localhost:8080/api/docs  

**Параметры запроса:**
- `login`: Логин пользователя.
- `limit`: Лимит количества возвращаемых документов.
- `key`: Ключ фильтра (например, `mime`).
- `value`: Значение фильтра (например, `image/jpg`).

**Заголовок:**
- `token`: Токен пользователя.

Пример использования cURL:

```bash
curl --location 'http://localhost:8080/api/docs?login=User12344&limit=10&key=mime&value=image%2Fjpg' \
--header 'token: JTTLEqyIO1r6HIvSOESB'
```

---

Этот запрос используется для получения списка документов, соответствующих указанным фильтрам.

## Удаление документа

**Метод:** DELETE  
**URL:** http://localhost:8080/api/docs/{document_id}  

**Путь:**
- `{document_id}`: Идентификатор документа.

**Заголовок:**
- `token`: Токен пользователя.

Пример использования cURL:

```bash
curl --location --request DELETE 'http://localhost:8080/api/docs/fbc46988-6c86-4add-b3d7-25254796da44' \
--header 'token: JTTLEqyIO1r6HIvSOESB'
```

---

Этот запрос используется для удаления документа с указанным идентификатором