# 📖 Основы Go (Go Basics)

## 1. Что такое Go?

**Go (Golang)** — компилируемый, статически типизированный язык программирования, разработанный Google в 2007 году.

### Ключевые особенности:
- ✅ **Простота** — минималистичный синтаксис, легко учить
- ✅ **Быстрота** — компиляция за секунды, быстрое выполнение
- ✅ **Конкурентность** — встроенная поддержка горутин и каналов
- ✅ **Статическая типизация** — типы проверяются на этапе компиляции
- ✅ **Сборка мусора** — автоматическое управление памятью
- ✅ **Стандартная библиотека** — богатая библиотека "из коробки"

---

## 2. Переменные и типы данных

### Объявление переменных

```go
// 1. Явное объявление
var age int = 25
var name string = "Alice"

// 2. С выводом типа
var age = 25        // int
var name = "Alice"  // string

// 3. Короткое объявление (только внутри функций)
age := 25
name := "Alice"

// 4. Множественное объявление
var x, y = 10, 20
a, b := 1, 2

// 5. Групповое объявление
var (
    name    string
    age     int
    isAdult bool
)
```

### Базовые типы

```go
// Числовые
int, int8, int16, int32, int64     // Знаковые целые
uint, uint8, uint16, uint32, uint64 // Беззнаковые целые
float32, float64                    // Числа с плавающей точкой
complex64, complex128               // Комплексные числа

// Строки и байты
string  // Неизменяемая последовательность байт
byte    // Алиас для uint8
rune    // Алиас для int32 (символ Unicode)

// Логический
bool    // true или false

// Специальные
nil     // Аналог null
any     // Алиас для interface{} (принимает любой тип)
```

### Zero Values (значения по умолчанию)

```go
int     → 0
float64 → 0.0
bool    → false
string  → "" (пустая строка)
slice   → nil
map     → nil
pointer → nil
```

---

## 3. Константы

```go
// Объявление констант
const Pi = 3.14159
const MaxUsers = 100

// Типизированные константы
const Pi float64 = 3.14159

// Групповое объявление
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
    Thursday       // 4
    Friday         // 5
    Saturday       // 6
)

// iota автоматически увеличивается
const (
    _  = iota      // 0 (пропускаем)
    KB = 1 << iota // 1 << 1 = 2
    MB             // 1 << 2 = 4
    GB             // 1 << 3 = 8
    TB             // 1 << 4 = 16
)
```

---

## 4. Массивы и слайсы

### Массивы

```go
// Фиксированный размер (часть типа!)
var arr [5]int                    // [0 0 0 0 0]
arr2 := [3]int{1, 2, 3}          // [1 2 3]
arr3 := [...]int{1, 2, 3, 4}     // [4]int{1 2 3 4}

// Многомерные массивы
var matrix [3][3]int

// Копирование (копируется весь массив!)
arr4 := arr2  // Копия [1 2 3]
```

### Слайсы (динамические массивы)

```go
// Создание
slice := []int{1, 2, 3}
slice2 := make([]int, 5)      // len=5, cap=5
slice3 := make([]int, 3, 10)  // len=3, cap=10

// len() - текущее количество элементов
// cap() - вместимость до реальнойлокации

// Срез (создаёт слайс на тот же массив)
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]  // [2 3 4], len=3, cap=4
s2 := arr[:3]  // [1 2 3]
s3 := arr[2:]  // [3 4 5]

// Append (может вызвать реальлокацию)
slice = append(slice, 4)
slice = append(slice, 5, 6, 7)
slice2 := append(slice1, slice2...)  // Распаковка

// Copy
dst := make([]int, len(src))
n := copy(dst, src)  // Копирует min(len(dst), len(src))
```

### Важные особенности слайсов

```go
// ⚠️ Слайс - дескриптор (указатель на массив + len + cap)
slice1 := []int{1, 2, 3}
slice2 := slice1  // Копия дескриптора, не данных!
slice2[0] = 100
fmt.Println(slice1[0])  // 100!

// ⚠️ Append может изменить исходный слайс
slice1 := []int{1, 2, 3}
slice2 := append(slice1, 4)
// slice1 и slice2 могут указывать на один массив!

// ✅ Как избежать проблем
slice2 := append([]int(nil), slice1...)  // Копия
```

---

## 5. Map (хэш-таблицы)

```go
// Создание
m := map[string]int{"a": 1, "b": 2}
m2 := make(map[string]int)

// Операции
m["c"] = 3           // Добавление
val := m["a"]        // Чтение (0 если ключа нет)
delete(m, "b")       // Удаление
len(m)               // Размер

// Проверка существования
val, ok := m["key"]
if ok {
    // Ключ существует
}

// Итерация (порядок не гарантирован!)
for key, value := range m {
    fmt.Println(key, value)
}

// ⚠️ Map не потокобезопасна!
// Используйте sync.Map или mutex
```

---

## 6. Функции

### Объявление

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
    return  // naked return (возвращает x=1, y=2)
}

// Variadic (переменное число аргументов)
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)           // 6
nums := []int{1, 2, 3}
sum(nums...)           // Распаковка слайса
```

### Анонимные функции и замыкания

```go
// Анонимная функция
double := func(x int) int {
    return x * 2
}
fmt.Println(double(5))  // 10

// Замыкание (захват переменных)
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter := makeCounter()
fmt.Println(counter())  // 1
fmt.Println(counter())  // 2
fmt.Println(counter())  // 3
```

---

## 7. Структуры

```go
// Объявление структуры
type Person struct {
    Name string
    Age  int
    City string
}

// Создание
p1 := Person{Name: "Alice", Age: 25, City: "Moscow"}
p2 := Person{"Bob", 30, "SPb"}  // Позиционно
p3 := Person{Name: "Charlie"}   // Частично (Age="", City=0)

// Доступ к полям
fmt.Println(p1.Name)  // Alice
p1.Age = 26

// Вложенные структуры
type Company struct {
    Name    string
    Director Person  // Вложенная структура
}

// Анонимные поля (встраивание)
type Employee struct {
    Person  // Встраивание (как наследование)
    Salary  int
}

e := Employee{
    Person: Person{Name: "Dave", Age: 35},
    Salary: 5000,
}
fmt.Println(e.Name)  // Dave (автоматический доступ)
```

---

## 8. Методы

```go
type Counter struct {
    value int
}

// Value receiver (копия)
func (c Counter) GetValue() int {
    return c.value
}

// Pointer receiver (ссылка, может изменять)
func (c *Counter) Increment() {
    c.value++
}

func (c *Counter) Add(n int) {
    c.value += n
}

// Использование
c := Counter{value: 0}
c.Increment()      // &c автоматически
fmt.Println(c.GetValue())  // 1
```

### Когда использовать pointer receiver?

```go
// ✅ Pointer receiver:
// 1. Нужно изменить receiver
// 2. Receiver большой (избегаем копирования)
// 3. Для консистентности

// ✅ Value receiver:
// 1. Не нужно изменять receiver
// 2. Receiver маленький (int, string)
// 3. Иммутабельные структуры
```

---

## 9. Интерфейсы

### Объявление и реализация

```go
// Интерфейс - набор методов
type Speaker interface {
    Speak() string
}

// Тип реализует интерфейс неявно
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

// Dog автоматически реализует Speaker
var s Speaker = Dog{Name: "Rex"}
fmt.Println(s.Speak())  // Woof!
```

### Empty Interface

```go
// interface{} принимает любой тип (как any в Java/Python)
func PrintAny(v interface{}) {
    fmt.Println(v)
}

// С Go 1.18 можно использовать any
func PrintAny2(v any) {
    fmt.Println(v)
}

// Type assertion (извлечение типа)
var i interface{} = "hello"
s := i.(string)  // Извлекаем string

// Type assertion с проверкой
s, ok := i.(string)
if ok {
    // i действительно string
}

// Type switch
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Популярные интерфейы из stdlib

```go
// Stringer - строковое представление
type Stringer interface {
    String() string
}

// Error - ошибки
type Error interface {
    Error() string
}

// Reader - чтение данных
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer - запись данных
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

---

## 10. Обработка ошибок

### Базовая обработка

```go
result, err := SomeFunction()
if err != nil {
    // Обработка ошибки
    return err
}
// Использование result
```

### Создание ошибок

```go
// Простая ошибка
err := errors.New("something went wrong")

// Форматированная ошибка
err := fmt.Errorf("user %d not found", userID)

// Обёртывание (Go 1.13+)
err := fmt.Errorf("database error: %w", underlyingErr)
```

### Кастомные ошибки

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// Использование
if email == "" {
    return &ValidationError{"email", "cannot be empty"}
}
```

### Проверка ошибок

```go
// errors.Is - сравнение
if errors.Is(err, ErrNotFound) {
    // Обработка конкретной ошибки
}

// errors.As - извлечение типа
var ve *ValidationError
if errors.As(err, &ve) {
    // err содержит ValidationError
}

// errors.Unwrap - распаковка
underlying := errors.Unwrap(err)
```

---

## 11. Control Flow

### If-else

```go
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
```

### For

```go
// Классический for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// For как while
for x < 10 {
    x++
}

// Бесконечный цикл
for {
    if condition {
        break
    }
}

// For each (range)
for i, v := range slice {
    fmt.Println(i, v)
}

// Только индексы
for i := range slice {
    fmt.Println(i)
}

// Только значения
for _, v := range slice {
    fmt.Println(v)
}

// Map
for key, value := range m {
    fmt.Println(key, value)
}
```

### Switch

```go
// Обычный switch
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

// Switch с падением (fallthrough)
switch x {
case 1:
    fmt.Println("one")
    fallthrough  // Выполнится следующий case!
case 2:
    fmt.Println("two or one")
}
```

---

## 12. Указатели

```go
x := 10
p := &x      // Указатель на x (тип *int)

fmt.Println(*p)  // 11 (разыменование)
*p = 20          // Изменение x через указатель
fmt.Println(x)   // 20

// Нулевой указатель
var p *int = nil

// new - выделение памяти
p := new(int)  // *int, *p = 0
*p = 42
```

---

## 13. Defer, Panic, Recover

### Defer

```go
// Отложенное выполнение (LIFO)
defer file.Close()    // Выполнится последним
defer mu.Unlock()     // Выполнится предпоследним

// Defer с аргументами (вычисляются сразу)
defer fmt.Println(x)  // x вычисляется сейчас

// Defer с функцией (выполнится позже)
defer func() {
    fmt.Println("cleanup")
}()
```

### Panic и Recover

```go
// Panic - аварийное завершение
panic("something went wrong")

// Recover - перехват panic
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()

panic("test")  // Не завершит программу
```

---

## 14. Generics (Go 1.18+)

### Базовые дженерики

```go
// Функция с дженериками
func Identity[T any](v T) T {
    return v
}

result := Identity[int](42)
result := Identity("hello")  // Вывод типа
```

### Ограничения (Constraints)

```go
// Интерфейс как ограничение
type Number interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}
```

### Структуры с дженериками

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}
```

---

**Это базовые основы Go! Далее изучайте конкурентность, работу с памятью и продвинутые темы.** 📚
