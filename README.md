# Audio Book AI

Backend-сервис, который превращает загруженный текстовый файл в аудиокнигу: считает слова, режет текст на части, озвучивает каждую часть и склеивает результат в один аудиофайл.

Пока только backend API (без фронтенда) — тестируется через `curl`/Postman.

## Как это работает

1. Пользователь логинится через Google (`POST /api/auth/google`), получает JWT.
2. Загружает `.txt`-файл (`POST /api/book/create`) — в ответ приходит `id` и статус `pending`.
3. В фоне (отдельный воркер) книга проходит пайплайн:
   `извлечение текста → подсчёт слов → разбивка на части → озвучка каждой части → склейка в один файл`.
4. Пользователь опрашивает статус (`GET /api/book/:id`) и, когда готово, скачивает аудио (`GET /api/book/:id/audio`).

Книги привязаны к пользователю — доступ к чужой книге по `id` возвращает `404`.

## Стек

- **Go 1.26**, Clean Architecture (структура скопирована с внутреннего проекта `mobility-backend`)
- **Echo** — HTTP-фреймворк
- **GORM + PostgreSQL** — хранение данных
- **Atlas** — миграции, генерируются из GORM-моделей (руками SQL не пишем)
- **Redis + asynq** — очередь фоновых задач (обработка книги идёт в отдельном процессе-воркере)
- **Wire + wiregenx** — генерация DI-контейнера по `// @inject`-аннотациям на конструкторах
- **zap** — логирование
- **Docker Compose** — локальный запуск всего стека (Postgres, Redis, web, async-воркер)

## Структура проекта

```
cmd/                     точки входа: main.go (-mode http|task-worker), http/, async/
di/wire/                 декларации Wire + сгенерированный provider.go
migration/               Atlas: конфиг + сгенерированные .sql миграции
docker/                  Dockerfile + docker-compose.dev.yml
env/                     .env.local (не в git)
src/
  core/
    domain/
      entity/            бизнес-сущности (Book, User, ...) + enum
      ports/              интерфейсы: repository, ttsport, ingestport, chunkport,
                          audioport, taskport, security, gateway, httpport
    application/
      response/           единый JSON-конверт для ответов/ошибок
      services/           переиспользуемые сервисы (выдача/проверка JWT)
      usecases/           бизнес-логика: bookusecases/, authusecases/
  entrypoint/
    http/                 HTTP-хендлеры, группы роутов, middleware
    asynctask/             обработчик фоновой задачи (озвучка книги)
  infrastructure/          конкретные реализации портов: Postgres, Echo-адаптер,
                          TTS-заглушка, Google OAuth, asynq, файловое хранилище...
```

Слои общаются только через интерфейсы (`ports/`) — это специально, чтобы можно было
подменить любую конкретную реализацию (например TTS-провайдера) без переписывания
остального кода.

## TTS (озвучка)

Сейчас подключена **заглушка** (`src/infrastructure/tts/mock_provider.go`): она честно
генерирует тихий WAV с коротким бипом на границе каждого куска текста — не пустышка,
а настоящий проигрываемый файл, чтобы вся остальная цепочка (склейка, отдача файла)
была реально проверена.

Причина заглушки: у `uzbekvoice.ai` (референс для этого проекта) нет публичной
документации API. Когда появится доступ к реальному TTS (uzbekvoice.ai или другому),
достаточно добавить новую реализацию интерфейса `ttsport.TTSProvider` — остальной код
менять не придётся.

## Google-вход

`POST /api/auth/google` принимает `id_token` от Google Sign-In и проверяет его через
официальную библиотеку `google.golang.org/api/idtoken` — это настоящая проверка подписи
по публичным ключам Google, не заглушка. Чтобы она реально заработала, нужно задать
`GOOGLE_CLIENT_ID` в `env/.env.local` (сейчас пусто — до этого любой токен будет
отклонён с ошибкой "invalid google id token", это ожидаемо).

## Переменные окружения (`env/.env.local`)

| Переменная | Назначение |
|---|---|
| `PORT` | порт HTTP-сервера |
| `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE` | подключение к Postgres |
| `REDIS_ADDR`, `REDIS_PASSWORD` | подключение к Redis |
| `DATA_DIR` | папка на диске для загруженных файлов/аудио |
| `MAX_CHUNK_CHARS` | максимальный размер куска текста перед озвучкой |
| `TTS_PROVIDER` | какой TTS-провайдер использовать (сейчас только `mock`) |
| `GOOGLE_CLIENT_ID` | OAuth Client ID из Google Cloud Console — нужно вписать самостоятельно |
| `JWT_SECRET` | секрет для подписи JWT (обязательно сменить перед продом) |
| `JWT_ACCESS_EXPIRE_MINUTES` / `JWT_REFRESH_EXPIRE_MINUTES` | время жизни токенов |

## Запуск

```bash
export PATH="$HOME/go/bin:$PATH"   # wire, wiregenx

# 1. Собрать DI-контейнер (нужно после любого изменения // @inject конструкторов)
cd di && make wire-build && cd ..

# 2. Сгенерировать миграцию из GORM-моделей (нужно после изменения моделей)
cd migration && make -f makemigration db-diff-dev && cd ..

# 3. Поднять всё окружение (Postgres, Redis, web, async-воркер)
cd docker && docker compose -f docker-compose.dev.yml --env-file ../env/.env.local up --build -d && cd ..

# 4. Применить миграции к базе
cd migration && make -f migrate.local migrate-apply && cd ..
```

## Проверка API

```bash
curl -s localhost:8080/api/health

# логин (нужен настоящий id_token от Google, когда задан GOOGLE_CLIENT_ID)
curl -s -X POST -H 'Content-Type: application/json' \
  -d '{"id_token":"<google_id_token>"}' localhost:8080/api/auth/google
# -> {"access_token": "...", "refresh_token": "...", "user": {...}}

TOKEN=<access_token из ответа выше>

curl -s -H "Authorization: Bearer $TOKEN" localhost:8080/api/auth/me

curl -s -H "Authorization: Bearer $TOKEN" -F "file=@sample.txt" localhost:8080/api/book/create
# -> {"id": 1, "status": "pending"}

curl -s -H "Authorization: Bearer $TOKEN" localhost:8080/api/book/1
# повторять, пока status не станет "done"

curl -s -H "Authorization: Bearer $TOKEN" -o audiobook.wav localhost:8080/api/book/1/audio
```

## API

| Метод | Путь | Auth | Описание |
|---|---|---|---|
| `GET` | `/api/health` | нет | проверка живости сервиса |
| `POST` | `/api/auth/google` | нет | вход по Google ID-токену, выдаёт JWT |
| `POST` | `/api/auth/refresh` | нет | обновление access-токена по refresh-токену |
| `GET` | `/api/auth/me` | Bearer | текущий пользователь |
| `POST` | `/api/book/create` | Bearer | загрузить `.txt`-книгу |
| `GET` | `/api/book/:id` | Bearer | статус обработки книги |
| `GET` | `/api/book/:id/audio` | Bearer | скачать готовое аудио |

## Планы (не сделано)

- Реальный TTS-провайдер вместо заглушки
- Поддержка `.epub`/`.pdf`
- Список книг пользователя (`GET /api/book`)
- CI/CD (не переносили из mobility-backend — там завязано на приватную инфраструктуру)
