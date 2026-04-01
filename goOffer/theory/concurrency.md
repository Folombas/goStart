# 🔄 Конкурентность в Go (Concurrency)

## 1. Горутины

### Что такое горутина?

**Горутина** — лёгкий поток управления, управляемый runtime Go (не ОС).

```go
// Запуск горутины
go func() {
    fmt.Println("Running in goroutine")
}()

// Именованная функция
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
}

go worker(1)  // Запуск горутины
```

### Отличия от потоков ОС

| Горутины | Потоки ОС |
|----------|-----------|
| ~2KB стек (растёт динамически) | ~1MB стек |
| Управляются runtime Go | Управляются ОС |
| Дешёвое переключение | Дорогое переключение |
| Тысячи на поток ОС | Ограничено ОС |

### Планировщик Go

```
Горутины → M:N планировщик → Системные потоки
           (Go Runtime)      (OS Threads)
```

- **M** — количество системных потоков
- **N** — количество горутин
- Планировщик сам решает, когда переключать горутины

---

## 2. Каналы

### Что такое канал?

**Канал** — типобезопасный способ связи между горутинами.

```go
// Создание канала
ch := make(chan int)        // Unbuffered (синхронный)
ch2 := make(chan int, 10)   // Buffered (асинхронный)

// Отправка
ch <- 42  // Блокируется пока не получат

// Получение
val := <-ch  // Блокируется пока не отправят

// Закрытие канала (только отправителем!)
close(ch)

// Проверка закрытия
val, ok := <-ch
if !ok {
    // Канал закрыт и пуст
}
```

### Unbuffered vs Buffered

```go
// Unbuffered канал (синхронный)
ch := make(chan int)
// Отправка блокируется до получения
// Получение блокируется до отправки

// Buffered канал (асинхронный)
ch := make(chan int, 3)
// Отправка блокируется только когда буфер полон
ch <- 1  // Не блокируется
ch <- 2  // Не блокируется
ch <- 3  // Не блокируется
ch <- 4  // Блокируется!
```

### Range и закрытие

```go
// Range по каналу (автоматически закрывается)
for msg := range ch {
    fmt.Println("Received:", msg)
}

// Закрытие канала отправителем
func producer(ch chan<- int) {
    defer close(ch)  // Закрываем когда закончили
    for i := 0; i < 10; i++ {
        ch <- i
    }
}
```

### Направление каналов

```go
// Двунаправленный
func process(ch chan int) {
    ch <- 1      // Отправка
    val := <-ch  // Получение
}

// Только отправка (chan<-)
func producer(ch chan<- int) {
    ch <- 1
    ch <- 2
}

// Только получение (<-chan)
func consumer(ch <-chan int) {
    val := <-ch
    fmt.Println(val)
}
```

---

## 3. Select

### Оператор select

**Select** позволяет ждать несколько каналов одновременно.

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
}
```

### Timeout с select

```go
select {
case result := <-ch:
    fmt.Println("Result:", result)
case <-time.After(time.Second):
    fmt.Println("Timeout!")
}
```

### Non-blocking операция

```go
select {
case ch <- data:
    fmt.Println("Sent")
default:
    fmt.Println("Channel full, non-blocking")
}

select {
case val := <-ch:
    fmt.Println("Received:", val)
default:
    fmt.Println("Channel empty, non-blocking")
}
```

### Приоритет в select

```go
// Select выбирает случайный готовый case
// Если готовы несколько — выбирает случайно

select {
case <-ctx.Done():  // Важнее всего
    return ctx.Err()
case result := <-ch:
    handle(result)
}
```

---

## 4. Sync примитивы

### WaitGroup

**WaitGroup** ждёт завершения набора горутин.

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("Worker", id)
    }(i)
}

wg.Wait()  // Ждём завершения всех
```

### Mutex

**Mutex** обеспечивает взаимное исключение.

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

### RWMutex

**RWMutex** — mutex с разделением на чтение/запись.

```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

// Чтение (множественные читатели)
func (c *Cache) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[key]
}

// Запись (эксклюзивный писатель)
func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}
```

### Once

**Once** выполняет код только один раз.

```go
var (
    instance *Config
    once     sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        instance = loadConfig()
    })
    return instance
}
```

### Cond

**Cond** — условная переменная для сигнализации.

```go
type Queue struct {
    mu    sync.Mutex
    cond  *sync.Cond
    items []int
}

func NewQueue() *Queue {
    q := &Queue{}
    q.cond = sync.NewCond(&q.mu)
    return q
}

func (q *Queue) Push(item int) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = append(q.items, item)
    q.cond.Signal()  // Будим одного ждущего
}

func (q *Queue) Pop() int {
    q.mu.Lock()
    defer q.mu.Unlock()
    
    for len(q.items) == 0 {
        q.cond.Wait()  // Ждём пока не появится элемент
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item
}
```

---

## 5. Atomic операции

**Atomic** — атомарные операции для примитивных типов.

```go
var counter int64

// Инкремент
atomic.AddInt64(&counter, 1)

// Загрузка
val := atomic.LoadInt64(&counter)

// Сохранение
atomic.StoreInt64(&counter, 42)

// Swap
old := atomic.SwapInt64(&counter, 100)

// Compare And Swap
swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
if swapped {
    fmt.Println("Swapped successfully")
}

// LoadPointer / StorePointer для указателей
var ptr unsafe.Pointer
atomic.StorePointer(&ptr, unsafe.Pointer(&value))
```

---

## 6. Context

### Что такое Context?

**Context** — для отмены операций, таймаутов, передачи значений.

### Создание контекста

```go
// Базовый контекст
ctx := context.Background()

// С отменой
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // Обязательно!

// С таймаутом
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// С дедлайном
ctx, cancel := context.WithDeadline(context.Background(), 
    time.Now().Add(time.Hour))
defer cancel()

// С значением
ctx = context.WithValue(ctx, "user_id", 123)
```

### Использование контекста

```go
func DoWork(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():  // Отмена/таймаут
            return ctx.Err()
        default:
            // Работа
            if err := doStep(); err != nil {
                return err
            }
        }
    }
}

// HTTP с контекстом
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := http.DefaultClient.Do(req)
```

### Правила работы с Context

```go
// ✅ Передавайте context первым параметром
func GetUser(ctx context.Context, id int) (*User, error)

// ✅ Не храните context в структуре
type Service struct {
    // ctx context.Context  // Плохо!
}

// ✅ Всегда вызывайте cancel
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

// ✅ Проверяйте ctx.Done()
select {
case <-ctx.Done():
    return ctx.Err()
default:
}
```

---

## 7. Паттерны конкурентности

### Worker Pool

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

// Использование
jobs := make(chan Job, 100)
go WorkerPool(5, jobs)

for i := 0; i < 100; i++ {
    taskNum := i
    jobs <- func() error {
        return process(taskNum)
    }
}
close(jobs)
```

### Fan-Out / Fan-In

```go
// Fan-out - распределение работы
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

### Pipeline (Конвейер)

```go
// Generator
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 1: удвоение
func doubler(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}

// Stage 2: возведение в квадрат
func squarer(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Использование
ch := generator(1, 2, 3, 4)
ch = doubler(ch)
ch = squarer(ch)

for result := range ch {
    fmt.Println(result)  // 4, 16, 36, 64
}
```

### ErrGroup

```go
// Go 1.20+ или golang.org/x/sync/errgroup
var g errgroup.Group

for i := 0; i < 10; i++ {
    g.Go(func() error {
        return doWork()
    })
}

if err := g.Wait(); err != nil {
    log.Fatal(err)
}

// С контекстом (отмена при первой ошибке)
ctx, cancel := context.WithCancel(context.Background())
g, ctx := errgroup.WithContext(ctx)

g.Go(func() error {
    return doWork(ctx)
})
```

---

## 8. Race Condition

### Что такое гонка данных?

**Race Condition** — когда несколько горутин обращаются к общим данным, и хотя бы одна пишет.

```go
// ❌ Race condition!
var counter int
for i := 0; i < 100; i++ {
    go func() {
        counter++  // Неатомарная операция!
    }()
}

// ✅ С mutex
var mu sync.Mutex
var counter int
for i := 0; i < 100; i++ {
    go func() {
        mu.Lock()
        counter++
        mu.Unlock()
    }()
}

// ✅ С atomic
var counter int64
for i := 0; i < 100; i++ {
    go func() {
        atomic.AddInt64(&counter, 1)
    }()
}
```

### Обнаружение гонок

```bash
# Запуск с race detector
go test -race
go run -race main.go

# Вывод:
# WARNING: DATA RACE
# Read at 0x00c000014098 by goroutine 7:
# Previous write at 0x00c000014098 by goroutine 6:
```

---

## 9. Типичные ошибки

### Утечка горутин

```go
// ❌ Горутина никогда не завершится
func Start() {
    go func() {
        for {
            work()
        }
    }()
}

// ✅ С контекстом
func Start(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                work()
            }
        }
    }()
}
```

### Отправка в закрытый канал

```go
// ❌ Panic!
close(ch)
ch <- 1  // panic: send on closed channel

// ✅ Закрывайте только отправитель
func producer(ch chan<- int) {
    defer close(ch)
    ch <- 1
}
```

### Неправильный размер буфера

```go
// ❌ Мёртвая блокировка
ch := make(chan int)
ch <- 1  // Блокируется навсегда!

// ✅ Buffered или горутина
ch := make(chan int, 1)
ch <- 1  // OK

// или
go func() { ch <- 1 }()
```

---

## 10. Best Practices

### ✅ Делайте
- Используйте context для отмены
- Закрывайте каналы только отправители
- Проверяйте ошибки
- Используйте defer для unlock
- Применяйте race detector

### ❌ Не делайте
- Не игнорируйте ошибки
- Не создавайте горутины без контроля
- Не используйте глобальные переменные
- Не забывайте про cancel
- Не паникуйте в production

---

**Конкурентность — мощная возможность Go, но требует осторожности!** 🔄
