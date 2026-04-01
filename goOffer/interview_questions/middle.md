# ❓ Вопросы для Middle Go Developer

## 1. Продвинутая конкурентность

### 1.1 Паттерны с каналами

**Вопрос:** Реализуйте Worker Pool

```go
type Job func() error

func WorkerPool(workers int, jobs <-chan Job) {
    var wg sync.WaitGroup
    
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                if err := job(); err != nil {
                    log.Printf("Worker %d: %v", id, err)
                }
            }
        }(i)
    }
    
    wg.Wait()
}
```

**Вопрос:** Что такое fan-out/fan-in?

```go
// Fan-out - распределение работы между горутинами
func fanOut(data []int, workers int) []<-chan int {
    channels := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        ch := make(chan int)
        channels[i] = ch
        go func(ch chan<- int, offset int) {
            defer close(ch)
            for j := offset; j < len(data); j += workers {
                ch <- data[j]
            }
        }(ch, i)
    }
    return channels
}

// Fan-in - сбор результатов
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for v := range ch {
                out <- v
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

---

### 1.2 Context

**Вопрос:** Когда использовать context.WithTimeout?

```go
func fetchWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    return http.DefaultClient.Do(req)
}
```

**Вопрос:** Как отменить группу горутин?

```go
func processWithCancel(ctx context.Context, data []int) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    
    var wg sync.WaitGroup
    errCh := make(chan error, len(data))
    
    for _, item := range data {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            if err := process(val); err != nil {
                cancel() // Отменяем все при ошибке
                errCh <- err
            }
        }(item)
    }
    
    // Ждём завершения или отмены
    go func() {
        wg.Wait()
        close(errCh)
    }()
    
    select {
    case err := <-errCh:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

---

## 2. Управление памятью

### 2.1 Escape Analysis

**Вопрос:** Что такое escape analysis?

```go
// Переменная "убегает" на heap если:
// 1. Возвращается из функции
// 2. Используется в горутине
// 3. Сохраняется в глобальной переменной

// ❌ Убегает на heap
func create() *int {
    x := 42
    return &x  // x уходит на heap
}

// ✅ Остаётся на stack
func createValue() int {
    x := 42
    return x  // x на стеке
}
```

**Вопрос:** Как посмотреть escape analysis?

```bash
go build -gcflags="-m" main.go
```

---

### 2.2 Оптимизация аллокаций

**Вопрос:** Как уменьшить количество аллокаций?

```go
// ❌ Много аллокаций
func concat(strings []string) string {
    result := ""
    for _, s := range strings {
        result += s  // Новая аллокация каждый раз!
    }
    return result
}

// ✅ Используем strings.Builder
func concat(strings []string) string {
    var builder strings.Builder
    for _, s := range strings {
        builder.WriteString(s)
    }
    return builder.String()
}
```

**Вопрос:** Что такое sync.Pool?

```go
// Пул объектов для повторного использования
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func processData(data []byte) ([]byte, error) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    
    buf.Reset()
    buf.Write(data)
    // ... обработка
    
    result := make([]byte, buf.Len())
    copy(result, buf.Bytes())
    return result, nil
}
```

---

## 3. Интерфейсы и паттерны

### 3.1 Интерфейсы

**Вопрос:** Почему маленькие интерфей лучше?

```go
// ❌ Большое интерфей сложно реализовать
type DataStore interface {
    Get(key string) (string, error)
    Set(key, value string) error
    Delete(key string) error
    Exists(key string) (bool, error)
    Keys() []string
    Values() []string
    Len() int
    Clear() error
}

// ✅ Маленькие интерфей (как в stdlib)
type Getter interface {
    Get(key string) (string, error)
}

type Setter interface {
    Set(key, value string) error
}

type Deleter interface {
    Delete(key string) error
}
```

**Вопрос:** Что такое io.Reader и io.Writer?

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Композиция интерфейсов
type ReadWriter interface {
    Reader
    Writer
}
```

---

### 3.2 Паттерны

**Вопрос:** Реализуйте Singleton

```go
type Singleton struct {
    data string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{data: "initialized"}
    })
    return instance
}
```

**Вопрос:** Реализуйте Builder

```go
type Server struct {
    host string
    port int
    tls  bool
}

type ServerBuilder struct {
    server *Server
}

func NewServerBuilder() *ServerBuilder {
    return &ServerBuilder{server: &Server{}}
}

func (b *ServerBuilder) Host(h string) *ServerBuilder {
    b.server.host = h
    return b
}

func (b *ServerBuilder) Port(p int) *ServerBuilder {
    b.server.port = p
    return b
}

func (b *ServerBuilder) TLS(enabled bool) *ServerBuilder {
    b.server.tls = enabled
    return b
}

func (b *ServerBuilder) Build() *Server {
    return b.server
}

// Использование:
server := NewServerBuilder().
    Host("localhost").
    Port(8080).
    TLS(true).
    Build()
```

---

## 4. Работа с ошибками

### 4.1 Обёртывание ошибок

**Вопрос:** Чем отличается %v от %w?

```go
// ❌ %v - просто форматирует
err := fmt.Errorf("database error: %v", underlyingErr)
errors.Is(err, underlyingErr)  // false

// ✅ %w - оборачивает для errors.Is/As
err := fmt.Errorf("database error: %w", underlyingErr)
errors.Is(err, underlyingErr)  // true
```

**Вопрос:** Как извлечь ошибку определённого типа?

```go
type NotFoundError struct {
    Resource string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Resource)
}

func handle(err error) {
    var notFound *NotFoundError
    if errors.As(err, &notFound) {
        log.Printf("Resource not found: %s", notFound.Resource)
    }
}
```

---

## 5. Тестирование

### 5.1 Table-driven тесты

**Вопрос:** Напишите table-driven тест

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 0, 0},
        {"mixed", -5, 10, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### 5.2 Mock интерфейсов

**Вопрос:** Как замокать зависимость?

```go
// Интерфейс
type Database interface {
    GetUser(id int) (*User, error)
}

// Мок для тестов
type MockDatabase struct {
    GetUserFunc func(id int) (*User, error)
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    return m.GetUserFunc(id)
}

// Тест
func TestUserService(t *testing.T) {
    mockDB := &MockDatabase{
        GetUserFunc: func(id int) (*User, error) {
            if id == 1 {
                return &User{Name: "Test"}, nil
            }
            return nil, errors.New("not found")
        },
    }
    
    service := NewUserService(mockDB)
    user, err := service.GetUser(1)
    
    if err != nil {
        t.Fatal(err)
    }
    if user.Name != "Test" {
        t.Errorf("expected Test, got %s", user.Name)
    }
}
```

---

## 6. Производительность

### 6.1 Бенчмарки

**Вопрос:** Как написать бенчмарк?

```go
func BenchmarkConcatPlus(b *testing.B) {
    strings := []string{"a", "b", "c", "d", "e"}
    for i := 0; i < b.N; i++ {
        _ = concatPlus(strings)
    }
}

func BenchmarkConcatBuilder(b *testing.B) {
    strings := []string{"a", "b", "c", "d", "e"}
    for i := 0; i < b.N; i++ {
        _ = concatBuilder(strings)
    }
}

// Запуск:
// go test -bench=. -benchmem
```

### 6.2 Профилирование

**Вопрос:** Как использовать pprof?

```go
import (
    _ "net/http/pprof"
    "net/http"
)

func main() {
    // Запуск HTTP сервера для pprof
    go http.ListenAndServe("localhost:6060", nil)
    
    // ... основной код
}

// Использование:
// go tool pprof http://localhost:6060/debug/pprof/heap
// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

---

## 7. Типичные задачи Middle

### Задача 1: Rate Limiter
```go
type RateLimiter struct {
    tokens   int
    maxTokens int
    refill   time.Duration
    lastRefill time.Time
    mu       sync.Mutex
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    elapsed := now.Sub(rl.lastRefill)
    
    // Добавляем токены
    rl.tokens += int(elapsed / rl.refill)
    if rl.tokens > rl.maxTokens {
        rl.tokens = rl.maxTokens
    }
    rl.lastRefill = now
    
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    return false
}
```

### Задача 2: Concurrent Cache
```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}
```

---

## 8. Что нужно знать Middle

### Обязательно:
- ✅ Паттерны конкурентности (worker pool, fan-out/in)
- ✅ Context и отмена операций
- ✅ Обработка и обёртывание ошибок
- ✅ Тестирование (table-driven, mock)
- ✅ Базовое профилирование
- ✅ Интерфейсы и композиция

### Желательно:
- Понимание работы GC
- Escape analysis
- Оптимизация аллокаций
- gRPC/Protobuf
- Работа с БД (sqlx, gorm)

---

**Удачи в подготовке! 🚀**
