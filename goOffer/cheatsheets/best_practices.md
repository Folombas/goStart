# ✅ Лучшие практики Go (Best Practices)

## 1. Именование

### Переменные
```go
// ✅ Короткие имена для коротких областей видимости
var i int
var user *User

// ✅ Описательные имена для глобальных переменных
var maxRetries = 3
var defaultTimeout = 30 * time.Second

// ❌ Избегайте бессмысленных имён
var d int  // Что это?
var data interface{}  // Слишком общо
```

### Функции и методы
```go
// ✅ Глаголы для функций
func GetUser(id int) (*User, error)
func CalculateTotal(items []Item) float64

// ✅ Существительные для методов получения
func (u *User) Name() string

// ✅ Глаголы для методов действия
func (u *User) UpdateEmail(email string) error
```

### Интерфейсы
```go
// ✅ Односложные имена для маленьких интерфей
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }

// ✅ -er суффикс для одно-методных интерфейсов
type Formatter interface { Format() string }
type Validator interface { Validate() error }

// ❌ Избегайте интерфейсов с суффиксом Interface
type ReaderInterface  // Плохо
```

## 2. Обработка ошибок

### ✅ Правильно
```go
// Проверяйте ошибки сразу
file, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("open file: %w", err)
}
defer file.Close()

// Обёртывайте с контекстом
if err := process(); err != nil {
    return fmt.Errorf("process data: %w", err)
}

// Используйте sentinel errors для ожидаемых ошибок
var ErrNotFound = errors.New("not found")

if user == nil {
    return ErrNotFound
}
```

### ❌ Неправильно
```go
// Игнорирование ошибок
file, _ := os.Open("file.txt")

// Panic для обычных ошибок
if err != nil {
    panic(err)  // Только для критических ошибок!
}

// Возврат ошибки без контекста
return err  // Что случилось?
```

## 3. Конкурентность

### Горутины
```go
// ✅ Знайте, когда горутина завершится
func Process(ctx context.Context, data []int) {
    var wg sync.WaitGroup
    for _, item := range data {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            process(val)
        }(item)
    }
    wg.Wait()
}

// ✅ Используйте context для отмены
func Start(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return  // Корректное завершение
            default:
                work()
            }
        }
    }()
}

// ❌ Утечка горутин
func Start() {
    go func() {
        for {
            work()  // Никогда не завершится
        }
    }()
}
```

### Каналы
```go
// ✅ Закрывайте канал только отправитель
func Producer(ch chan<- int) {
    defer close(ch)  // Отправитель закрывает
    ch <- 1
    ch <- 2
}

// ✅ Используйте buffered каналы осознанно
ch := make(chan int, 100)  // Когда знаете размер

// ✅ Range для чтения
for msg := range ch {
    process(msg)
}

// ❌ Закрытие канала получателя
for i := 0; i < 10; i++ {
    close(ch)  // Panic!
}

// ❌ Отправка в закрытый канал
close(ch)
ch <- 1  // Panic!
```

### Select
```go
// ✅ Timeout с select
select {
case result := <-ch:
    handle(result)
case <-time.After(5 * time.Second):
    return errors.New("timeout")
}

// ✅ Non-blocking операция
select {
case ch <- data:
    // Отправлено
default:
    // Канал полон, не блокируем
}

// ✅ Правильный приоритет
select {
case <-ctx.Done():
    return ctx.Err()  // Важнее всего
case result := <-ch:
    handle(result)
}
```

## 4. Работа с памятью

### Слайсы
```go
// ✅ Предварительное выделение容量
slice := make([]int, 0, 100)
for i := 0; i < 100; i++ {
    slice = append(slice, i)  // Без реаллокаций
}

// ✅ Копируйте при возврате части большого слайса
func GetHeader(data []byte) []byte {
    result := make([]byte, 100)
    copy(result, data[:100])
    return result  // data может освободиться
}

// ❌ Утечка памяти через слайс
func GetFirst10(data []byte) []byte {
    return data[:10]  // data не освободится
}
```

### Map
```go
// ✅ Предварительное выделение
m := make(map[string]int, 100)

// ✅ Очистка для переиспользования
for k := range m {
    delete(m, k)
}
// или
m = make(map[string]int)
```

### Strings
```go
// ✅ strings.Builder для конкатенации
var builder strings.Builder
for _, s := range strings {
    builder.WriteString(s)
}
result := builder.String()

// ✅ strings.Join для слайсов
result := strings.Join(parts, ", ")

// ❌ Многократная конкатенация
result := ""
for _, s := range strings {
    result += s  // O(n²) аллокаций!
}
```

## 5. Интерфейсы

### ✅ Правильно
```go
// Маленькие интерфей
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Принимайте интерфейсы, возвращайте структуры
func Process(r io.Reader) (*Result, error) {
    // ...
}

// Размещайте интерфейсы где они используются
// (не там где объявлены)
package service

type repository interface {  // Внутри пакета
    Get(id int) (*User, error)
}
```

### ❌ Неправильно
```go
// Большие интерфей
type DataStore interface {
    Get() error
    Set() error
    Delete() error
    Update() error
    List() ([]Item, error)
    Count() int
    // ... ещё 10 методов
}

// Возврат интерфейса
func GetUser() IUser {  // Зачем?
    return &User{}
}
```

## 6. Структуры кода

### Организация пакетов
```go
// ✅ Стандартная структура
cmd/
internal/
pkg/
api/
configs/
migrations/

// ✅ Один пакет = одна ответственность
package auth  // Только аутентификация
package users // Только пользователи

// ❌ God object пакет
package common  // Всё подряд
package utils   // Бесполезные функции
```

### Зависимости
```go
// ✅ Зависимости через интерфейсы
type Service struct {
    repo Repository  // Интерфейс
}

// ✅ Внедрение зависимостей
func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

// ❌ Жёсткие зависимости
type Service struct {
    db *sql.DB  // Конкретная реализация
}
```

## 7. Тестирование

### ✅ Правильно
```go
// Table-driven тесты
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.expected {
                t.Errorf("Add(%d, %d) = %d", tt.a, tt.b, got)
            }
        })
    }
}

// Используйте t.Helper()
func assertEqual(t *testing.T, a, b int) {
    t.Helper()  // Важно для трассировки
    if a != b {
        t.Errorf("%d != %d", a, b)
    }
}

// Параллельные тесты
func TestParallel(t *testing.T) {
    t.Parallel()
    // ...
}
```

### ❌ Неправильно
```go
// Тест без проверки ошибок
func TestSomething(t *testing.T) {
    result, _ := SomeFunction()  // Игнор ошибки
    if result != expected {
        t.Fail()
    }
}

// Один большой тест
func TestAll(t *testing.T) {
    // 100 строк кода
    // Проверяет всё подряд
}
```

## 8. Производительность

### ✅ Оптимизация
```go
// ✅ Используйте sync.Pool для частых аллокаций
var pool = sync.Pool{
    New: func() interface{} {
        return &bytes.Buffer{}
    },
}

buf := pool.Get().(*bytes.Buffer)
defer pool.Put(buf)

// ✅ Предварительное выделение
slice := make([]int, 0, 1000)

// ✅ strings.Builder
var builder strings.Builder
builder.Grow(1000)  // Предварительный размер

// ✅ Итерация по индексу (быстрее чем range для больших данных)
for i := 0; i < len(slice); i++ {
    use(slice[i])
}
```

### ❌ Избегайте
```go
// ❌ Конкатенация в цикле
result := ""
for i := 0; i < 1000; i++ {
    result += strconv.Itoa(i)
}

// ❌ Частые аллокации в горячем коде
for {
    buffer := make([]byte, 1024)  // Каждый раз новая
    read(buffer)
}

// ❌ Неиспользуемые переменные
_ = unusedVariable  // Компилятор предупредит
```

## 9. Контекст

### ✅ Правильно
```go
// ✅ Передавайте context первым параметром
func GetUser(ctx context.Context, id int) (*User, error)

// ✅ Всегда проверяйте ctx.Done()
func Process(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            work()
        }
    }
}

// ✅ Используйте context.WithTimeout
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()
```

### ❌ Неправильно
```go
// ❌ Хранение context в структуре
type Service struct {
    ctx context.Context  // Плохо!
}

// ❌ Игнорирование отмены
func Process(ctx context.Context) {
    longOperation()  // Не проверяет ctx.Done()
}

// ❌ context.Background() в библиотеках
func Get(ctx context.Context) {
    // Используйте переданный ctx, не создавайте новый
}
```

## 10. Generics

### ✅ Когда использовать
```go
// ✅ Для типовых структур данных
type Stack[T any] struct {
    items []T
}

// ✅ Для общих алгоритмов
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// ✅ С ограничениями
type Number interface {
    int | int32 | int64 | float32 | float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}
```

### ❌ Когда не использовать
```go
// ❌ Когда нужен только один тип
func ProcessInts(nums []int) []int  // Лучше чем generics

// ❌ Для усложнения кода
func Process[T VeryComplexConstraint](v T) T

// ❌ Когда interface{} достаточно
func Print(v any)  // Проще чем generics
```

## 11. Code Style

### Форматирование
```bash
# ✅ Всегда используйте gofmt
gofmt -w .

# ✅ Или goimports (добавляет импорты)
goimports -w .

# ✅ Запускайте в pre-commit хуке
```

### Линтеры
```bash
# ✅ Используйте golangci-lint
golangci-lint run

# ✅ Настройте .golangci.yml
linters:
  enable:
    - govet
    - errcheck
    - staticcheck
    - unused
```

### Комментарии
```go
// ✅ Комментируйте WHY, не WHAT
// Используем map для O(1) доступа
m := make(map[string]int)

// ✅ Документируйте экспортируемые API
// GetUser возвращает пользователя по ID
// Возвращает ErrNotFound если пользователь не найден
func GetUser(id int) (*User, error)

// ❌ Избыточные комментарии
i++  // Увеличиваем i на 1
```

---

## Чек-лист перед коммитом

- [ ] Код отформатирован (`gofmt`)
- [ ] Линтеры проходят (`golangci-lint`)
- [ ] Тесты написаны и проходят
- [ ] Ошибки обрабатываются
- [ ] Нет утечек горутин
- [ ] Context используется правильно
- [ ] Нет race condition
- [ ] Код следует best practices

---

**Помните: хороший код — это код, который легко читать и поддерживать!** 📖
