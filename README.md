# Salute Speech API
Простой доступ к API Salute Speech для Golang. Распознавание и Синтер речи. Самый простой доступ по REST.

### Установка
```
go get github.com/saintbyte/salute_speech_api
```

### Использование
```
import (
salute_speech_api "github.com/saintbyte/salute_speech_api"
)
```
Некоторые вещи придется проверят самому , такие как длина распозноваемого аудио

Задать переменную среды SALUTE_SPEECH_AUTH_DATA из "Авторизационные данные" в настройка проекта который находиться в пространстве.

Возможно задать пути к времененым файлами с авторизацией. По умолчанию файлы будут валяться в текущей папке.
* SALUTE_TOKEN_FILE - Путь к файлу с токеном. Пример: /tmp/.salute_speech_token
* SALUTE_EXPIRES_FILE - Путь к файлу с временем устаревания токена . Пример: /tmp/.salute_speech_expires

Есть пример .env - .env_examle

И подробнее о функциях: https://pkg.go.dev/github.com/saintbyte/salute_speech_api

### Примеры
см. в директории examples

