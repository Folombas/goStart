# ❓ Вопросы для Junior Go Developer

## 1. Основы языка

### 1.1 Переменные и типы данных

**Вопрос:** Какие способы объявления переменных вы знаете?

```go
// 1. Явное объявление с типом
var age int = 25

// 2. Вывод типа
var name = "Alice"

// 3. Короткое объявление
city := "Moscow"

// 4. Множественное объявление
var x, y = 10, 20
a, b := 1, 2

// 5. Групповое объявление
var (
    name string
    age  int
)
```

**Вопрос:** Какие значения по умолчанию у типов в Go?

```go
int     → 0
float64 → 0.0
bool    → false
string  → ""
slice   → nil
map     → nil
pointer → nil
```

---

### 1.2 Слайсы и массивы

**Вопрос:** В чём разница между массивом и слайсом?

```go
// Массив - фиксированный размер, значение
arr := [3]int{1, 2, 3}
arr2 := arr // Копирование всего массива

// Слайс - динамический размер, ссылка
slice := []int{1, 2, 3}
slice2 := slice // Копирование дескриптора
```

**Вопрос:** Что такое len и cap у слайса?

```go
slice := make([]int, 3, 5)
len(slice) → 3  // текущее количество элементов
cap(slice) → 5  // вместимость до реальнойлокации
```

**Вопрос:** Что будет при append сверх capacity?

```go
// Go создаст новый массив большей вместимости
// и скопирует туда данные (аллокация!)
```

---

### 1.3 Map

**Вопрос:** Как работает map в Go?

```go
// Hash table внутри
// O(1) средняя сложность доступа
// Не гарантирует порядок элементов
// Потокобезопасна только с mutex/rwmutex
```

**Вопрос:** Что вернёт получение несуществующего ключа?

```go
m := map[string]int{"a": 1}
val := m["b"]  // 0 (zero value для int)

// Проверка существования:
val, ok := m["b"]  // val=0, ok=false
```

---

### 1.4 Функции и методы

**Вопрос:** Что такое receiver и какие они бывают?

```go
type Counter struct {
    value int
}

// Value receiver (копия)
func (c Counter) GetValue() int {
    return c.value
}

// Pointer receiver (ссылка)
func (c *Counter) Increment() {
    c.value++
}
```

**Вопрос:** Когда использовать pointer receiver?

```go
// Pointer receiver когда:
// 1. Нужно изменить receiver
// 2. Receiver большой (избегаем копирования)
// 3. Для консистентности (если хоть один метод pointer)
```

---

### 1.5 Интерфейсы

**Вопрос:** Что такое интерфейс в Go?

```go
// Интерфейс - набор методов
type Speaker interface {
    Speak() string
}

// Тип реализует интерфейс неявно
type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

// Dog реализует Speaker автоматически
```

**Вопрос:** Что такое empty interface?

```go
// interface{} принимает любое значение
// Аналог any в других языках
func PrintAny(v interface{}) {
    fmt.Println(v)
}

// С Go 1.18 можно использовать any
func PrintAny2(v any) {
    fmt.Println(v)
}
```

---

## 2. Обработка ошибок

**Вопрос:** Как обрабатывать ошибки в Go?

```go
result, err := SomeFunction()
if err != nil {
    // Обработка ошибки
    return err
}
// Использование result
```

**Вопрос:** Что такое error в Go?

```go
type error interface {
    Error() string
}

// Кастомная ошибка
type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}
```

**Вопрос:** Чем отличается panic от ошибки?

```go
// error - ожидаемая ситуация (файл не найден)
// panic - критическая ошибка (index out of range)
// panic прерывает выполнение, error - нет
```

---

## 3. Конкурентность

**Вопрос:** Что такое горутина?

```go
// Горутина - лёгкий поток управления
// Управляется runtime Go
// Размер стека ~2KB (растёт динамически)
// Дешевле системных потоков

go func() {
    fmt.Println("Running in goroutine")
}()
```

**Вопрос:** Что такое канал?

```go
// Канал - типобезопасный способ связи горутин
ch := make(chan int)

// Отправка
ch <- 42

// Получение
val := <-ch
```

**Вопрос:** Buffered vs unbuffered каналы?

```go
// Unbuffered (синхронный)
ch := make(chan int)
// Блокируется пока не будет получатель

// Buffered (асинхронный)
ch := make(chan int, 10)
// Буферизует до 10 значений
```

---

## 4. Практические задачи

### Задача 1: Разворот слайса
```go
func Reverse(slice []int) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}
```

### Задача 2: Уникальные элементы
```go
func Unique(slice []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    
    for _, v := range slice {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }
    return result
}
```

### Задача 3: Подсчёт слов
```go
func CountWords(text string) map[string]int {
    words := strings.Fields(text)
    counts := make(map[string]int)
    
    for _, word := range words {
        counts[word]++
    }
    return counts
}
```

---

## 5. Типичные ошибки Junior

### 1. Неправильная работа с слайсами
```go
// ❌ Плохо: слайс ссылается на большой массив
func GetFirst10(data []byte) []byte {
    return data[:10]  // data не освободится
}

// ✅ Хорошо: копируем нужную часть
func GetFirst10(data []byte) []byte {
    result := make([]byte, 10)
    copy(result, data)
    return result
}
```

### 2. Забывают проверять ошибки
```go
// ❌ Плохо
file, _ := os.Open("file.txt")

// ✅ Хорошо
file, err := os.Open("file.txt")
if err != nil {
    return err
}
```

### 3. Race condition
```go
// ❌ Плохо: гонка данных
var counter int
go func() { counter++ }()

// ✅ Хорошо: mutex
var mu sync.Mutex
var counter int
go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

### 4. Утечка горутин
```go
// ❌ Плохо: горутина никогда не завершится
func Start() {
    go func() {
        for {
            // Бесконечный цикл
        }
    }()
}

// ✅ Хорошо: контекст для отмены
func Start(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                // Работа
            }
        }
    }()
}
```

---

## 6. Советы для собеседования

### Что нужно знать точно:
1. ✅ Разница слайс/массив/мап
2. ✅ Обработка ошибок
3. ✅ Горутины и каналы (базово)
4. ✅ Интерфейсы
5. ✅ Методы и receiver

### На что обратят внимание:
- Понимаете ли вы, что делает код
- Можете ли объяснить свой выбор
- Знаете ли про обработку ошибок
- Пишете ли чистый код

### Можно не знать (для Junior):
- Глубокое понимание GC
- Оптимизацию производительности
- Сложные паттерны конкурентности
- Внутренности runtime

---

**Удачи на собеседовании! 🍀**
