# 📋 Шпаргалка по синтаксису Go

## Объявление переменных

```go
var x int = 10          // Явное объявление
var y = 20              // С выводом типа
z := 30                 // Короткое объявление
var a, b = 1, 2         // Множественное
var (                   // Групповое
    name string
    age  int
)
```

## Типы данных

```go
// Числовые
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128

// Строки и байты
string, byte, rune

// Логический
bool

// Специальные
nil, any (interface{}), never
```

## Zero values

```go
int     → 0
float   → 0.0
bool    → false
string  → ""
slice   → nil
map     → nil
pointer → nil
```

## Массивы и слайсы

```go
// Массив (фиксированный размер)
arr := [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3, 4}  // Вывод размера

// Слайс (динамический)
slice := []int{1, 2, 3}
slice2 := make([]int, 5, 10)  // len=5, cap=10

// Операции
slice[1:3]      // Срез [1, 3)
slice[:2]       // С начала до 2
slice[2:]       // От 2 до конца
append(s, 1, 2) // Добавление
copy(dst, src)  // Копирование
len(s), cap(s)  // Длина и вместимость
```

## Map

```go
m := map[string]int{"a": 1, "b": 2}
m2 := make(map[string]int)

m["c"] = 3           // Добавление
val := m["a"]        // Чтение
delete(m, "b")       // Удаление
val, ok := m["x"]    // Проверка
len(m)               // Размер
```

## Функции

```go
// Обычная функция
func add(a, b int) int {
    return a + b
}

// Множественные результаты
func div(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Именованные результаты
func get() (x int, y int) {
    x = 1
    y = 2
    return  // naked return
}

// Variadic функция
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

## Методы и структуры

```go
type Person struct {
    Name string
    Age  int
}

// Value receiver
func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// Pointer receiver
func (p *Person) Birthday() {
    p.Age++
}
```

## Интерфейсы

```go
type Speaker interface {
    Speak() string
}

// Реализация неявная
type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

// Type assertion
var s Speaker = Dog{}
dog := s.(Dog)

// Type switch
switch v := s.(type) {
case Dog:
    fmt.Println("It's a dog")
default:
    fmt.Println("Unknown")
}
```

## Control flow

```go
// If
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// If с инициализацией
if val, ok := m["key"]; ok {
    fmt.Println(val)
}

// For (while)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// For each
for i, v := range slice {
    fmt.Println(i, v)
}

// For без условия (while true)
for {
    if condition {
        break
    }
}

// Switch
switch day {
case "Mon", "Tue":
    fmt.Println("weekday")
case "Sat", "Sun":
    fmt.Println("weekend")
default:
    fmt.Println("unknown")
}

// Switch без условия
switch {
case x > 10:
    fmt.Println("big")
case x < 0:
    fmt.Println("negative")
}
```

## Указатели

```go
x := 10
p := &x      // Указатель на x
fmt.Println(*p)  // Разыменование: 10
*p = 20      // Изменение x через указатель
```

## Defer, Panic, Recover

```go
// Defer - отложенное выполнение
defer file.Close()
defer mu.Unlock()

// Panic - аварийное завершение
panic("something went wrong")

// Recover - перехват panic
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

## Горутины и каналы

```go
// Горутина
go func() {
    fmt.Println("running")
}()

// Канал
ch := make(chan int)        // Unbuffered
ch2 := make(chan int, 10)   // Buffered

ch <- 42        // Отправка
val := <-ch     // Получение
close(ch)       // Закрытие

// Select
select {
case val := <-ch1:
    fmt.Println("from ch1:", val)
case ch2 <- data:
    fmt.Println("sent to ch2")
case <-time.After(time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("non-blocking")
}
```

## Context

```go
// Создание
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)
ctx, cancel := context.WithTimeout(ctx, time.Second)
ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Hour))
ctx = context.WithValue(ctx, "key", "value")

// Использование
val := ctx.Value("key")
<-ctx.Done()  // Ожидание отмены
err := ctx.Err()
```

## Работа с ошибками

```go
// Базовая обработка
if err != nil {
    return err
}

// Создание ошибки
err := errors.New("error message")
err := fmt.Errorf("formatted: %v", underlying)

// Обёртывание (Go 1.13+)
err := fmt.Errorf("context: %w", underlying)

// Проверка
errors.Is(err, target)      // Сравнение
errors.As(err, &target)     // Извлечение типа
errors.Unwrap(err)          // Распаковка
```

## Generics (Go 1.18+)

```go
// Функция с дженериками
func Identity[T any](v T) T {
    return v
}

// Ограничения
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

// Структура с дженериками
type Stack[T any] struct {
    items []T
}
```

## Работа с JSON

```go
// Структура с тегами
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Password string `json:"-"`
    Email    string `json:"email,omitempty"`
}

// Marshal
data, _ := json.Marshal(user)
json, _ := json.MarshalIndent(user, "", "  ")

// Unmarshal
err := json.Unmarshal(data, &user)

// Encoder/Decoder
enc := json.NewEncoder(w)
enc.Encode(user)

dec := json.NewDecoder(r)
dec.Decode(&user)
```

## IO

```go
// Чтение файла
data, _ := os.ReadFile("file.txt")

// Запись файла
os.WriteFile("file.txt", data, 0644)

// Потоковое чтение
file, _ := os.Open("file.txt")
defer file.Close()
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}

// Потоковая запись
file, _ := os.Create("file.txt")
writer := bufio.NewWriter(file)
writer.WriteString("hello")
writer.Flush()
```

## HTTP

```go
// Сервер
http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
})
http.ListenAndServe(":8080", nil)

// Клиент
resp, _ := http.Get("https://api.example.com/data")
defer resp.Body.Close()
json.NewDecoder(resp.Body).Decode(&result)

// POST с JSON
data := map[string]string{"key": "value"}
jsonData, _ := json.Marshal(data)
resp, _ := http.Post("https://api.example.com/data", "application/json", 
    bytes.NewBuffer(jsonData))
```

## Тестирование

```go
// Базовый тест
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}

// Table-driven тест
func TestMath(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if result := Add(tt.a, tt.b); result != tt.expected {
                t.Errorf("Add(%d, %d) = %d", tt.a, tt.b, result)
            }
        })
    }
}

// Бенчмарк
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

## Sync примитивы

```go
// Mutex
var mu sync.Mutex
mu.Lock()
// критическая секция
mu.Unlock()

// RWMutex
var rw sync.RWMutex
rw.RLock()   // Чтение
rw.RUnlock()
rw.Lock()    // Запись
rw.Unlock()

// WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // работа
}()
wg.Wait()

// Once
var once sync.Once
once.Do(func() {
    // выполнится один раз
})

// Cond
cond := sync.NewCond(&sync.Mutex{})
cond.L.Lock()
cond.Wait()
cond.Signal()
cond.Broadcast()
cond.L.Unlock()
```

## Атомарные операции

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

// CompareAndSwap
swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
```

## Time

```go
// Текущее время
now := time.Now()

// Форматирование
now.Format("2006-01-02 15:04:05")

// Парсинг
t, _ := time.Parse("2006-01-02", "2024-01-15")

// Duration
duration := 5 * time.Second
time.Sleep(duration)

// Timer
timer := time.NewTimer(time.Second)
<-timer.C

// Ticker
ticker := time.NewTicker(time.Second)
for range ticker.C {
    // каждую секунду
}
ticker.Stop()
```

## Path и OS

```go
// Пути
path := filepath.Join("home", "user", "file.txt")
dir := filepath.Dir(path)
base := filepath.Base(path)
ext := filepath.Ext(path)
abs, _ := filepath.Abs("relative/path")

// Директории
os.Mkdir("dir", 0755)
os.MkdirAll("a/b/c", 0755)
entries, _ := os.ReadDir("dir")
os.RemoveAll("dir")

// Переменные окружения
os.Setenv("KEY", "value")
val := os.Getenv("KEY")
os.Unsetenv("KEY")
```

---
