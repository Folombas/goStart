# Новые примеры Go

Эта папка содержит 8 новых примеров, демонстрирующих различные возможности Go.

## 📚 Список примеров

### 1. Generics (Дженерики)
**Путь:** `cmd/generics_main/`

Демонстрирует работу с дженериками (Go 1.18+):
- Стеки с дженериками
- Ограничения типов (constraints)
- Функции Map и Filter
- Пары значений (Pair)

**Запуск:**
```bash
cd project
go run cmd/generics_main/main.go
```

---

### 2. Context
**Путь:** `cmd/context_main/`

Примеры работы с контекстом:
- Базовый контекст с отменой
- Контекст с таймаутом
- Контекст с дедлайном
- Передача значений
- Цепочка отмены
- Управление воркерами

**Запуск:**
```bash
cd project
go run cmd/context_main/main.go
```

---

### 3. Errors (Обработка ошибок)
**Путь:** `cmd/errors_main/`

Продвинутая обработка ошибок:
- Кастомные типы ошибок
- Обёртывание ошибок
- Ошибки с метаданными
- Лучшие практики

**Запуск:**
```bash
cd project
go run cmd/errors_main/main.go
```

---

### 4. Channels & Goroutines
**Путь:** `cmd/channels_main/`

Конкурентность в Go:
- Базовые каналы
- Буферизированные каналы
- Направление каналов
- Select оператор
- Worker Pool
- Pipeline
- Race Condition и Mutex

**Запуск:**
```bash
cd project
go run cmd/channels_main/main.go
```

---

### 5. JSON
**Путь:** `cmd/json_main/`

Работа с JSON:
- Marshal/Unmarshal
- Теги JSON
- Работа с map
- Encoder/Decoder
- Валидация JSON
- Кастомная сериализация

**Запуск:**
```bash
cd project
go run cmd/json_main/main.go
```

---

### 6. File I/O
**Путь:** `cmd/fileio_main/`

Работа с файлами:
- Чтение/запись файлов
- Построчное чтение/запись
- Работа с директориями
- Пути (filepath)
- Копирование файлов
- Проверка существования

**Запуск:**
```bash
cd project
go run cmd/fileio_main/main.go
```

---

### 7. HTTP Client
**Путь:** `cmd/httpclient_main/`

HTTP запросы:
- GET/POST запросы
- Параметры URL
- JSON в запросах
- Формы
- Кастомный клиент
- Контекст и таймауты
- Заголовки
- Работа с GitHub API

**Запуск:**
```bash
cd project
go run cmd/httpclient_main/main.go
```

**Требует:** подключения к интернету

---

### 8. Testing
**Путь:** `cmd/testing_main/`

Тестирование в Go:
- Базовые тесты
- Table-driven тесты
- Тестовые хелперы
- Setup/Teardown
- Моки и заглушки
- Бенчмарки
- Code Coverage
- Race Detector

**Запуск:**
```bash
cd project
go run cmd/testing_main/main.go
```

---

## 🚀 Быстрый запуск всех примеров

```bash
cd /home/gofer/godev/projects/goStart/project

# Последовательный запуск
go run cmd/generics_main/main.go
go run cmd/context_main/main.go
go run cmd/errors_main/main.go
go run cmd/channels_main/main.go
go run cmd/json_main/main.go
go run cmd/fileio_main/main.go
go run cmd/httpclient_main/main.go
go run cmd/testing_main/main.go
```

## 📁 Структура

```
project/
├── cmd/
│   ├── generics_main/
│   ├── context_main/
│   ├── errors_main/
│   ├── channels_main/
│   ├── json_main/
│   ├── fileio_main/
│   ├── httpclient_main/
│   └── testing_main/
└── data_types/
    ├── generics/
    ├── context_pkg/
    ├── errors_pkg/
    ├── channels/
    ├── json_pkg/
    ├── fileio/
    ├── httpclient/
    └── testing_pkg/
```

## 📝 Заметки

- Все примеры написаны в образовательных целях
- Каждый пример содержит подробные комментарии
- Примеры используют современные возможности Go
- Код следует лучшим практикам Go

## 🎯 Изучаемые темы

- ✅ Generics (дженерики)
- ✅ Context и отмена операций
- ✅ Продвинутая обработка ошибок
- ✅ Каналы и горутины
- ✅ Работа с JSON
- ✅ Файловый I/O
- ✅ HTTP клиент
- ✅ Тестирование и бенчмарки
