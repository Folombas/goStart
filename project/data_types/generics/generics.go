package generics

import "fmt"

// PrintHeader выводит заголовок раздела
func PrintHeader(title string) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Printf("  %s\n", title)
	fmt.Println("========================================")
	fmt.Println()
}

// PrintFooter выводит подвал
func PrintFooter() {
	fmt.Println()
	fmt.Println("📚 Изучайте Go дальше!")
	fmt.Println()
}

// ========== БАЗОВЫЕ ДЖЕНЕРИКИ ==========

// Stack - стек с использованием дженериков
// T - любой тип данных
type Stack[T any] struct {
	items []T
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop удаляет и возвращает последний элемент
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

// IsEmpty проверяет, пуст ли стек
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size возвращает размер стека
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// ========== ДЖЕНЕРИКИ С ОГРАНИЧЕНИЯМИ ==========

// Number - ограничение для числовых типов
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Sum - суммирует все числа в слайсе
func Sum[T Number](numbers []T) T {
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

// Average - вычисляет среднее значение
func Average[T Number](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := Sum(numbers)
	return float64(sum) / float64(len(numbers))
}

// ========== ДЖЕНЕРИКИ С ИНТЕРФЕЙСАМИ ==========

// Stringer - ограничение для типов с методом String()
type Stringer interface {
	String() string
}

// PrintAll - печатает все элементы, реализующие Stringer
func PrintAll[T Stringer](items []T) {
	for i, item := range items {
		fmt.Printf("  [%d] %s\n", i, item.String())
	}
}

// Person - пример структуры с методом String()
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d лет)", p.Name, p.Age)
}

// ========== ДЖЕНЕРИКИ С НЕСКОЛЬКИМИ ТИПАМИ ==========

// Pair - пара значений разных типов
type Pair[K, V any] struct {
	Key   K
	Value V
}

// NewPair создаёт новую пару
func NewPair[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{Key: key, Value: value}
}

// String возвращает строковое представление пары
func (p Pair[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", p.Key, p.Value)
}

// ========== ФУНКЦИИ МАП И ФИЛЬТР ==========

// Map - применяет функцию к каждому элементу слайса
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter - фильтрует элементы слайса по предикату
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// ========== ДЕМО ФУНКЦИИ ==========

// DemoGenerics - основная демонстрационная функция
func DemoGenerics() {
	PrintHeader("📦 Дженерики (Generics) в Go")

	// 1. Базовые дженерики - Stack
	fmt.Println("1️⃣  Стеки с дженериками:")
	fmt.Println("   Создаём стек для int и стек для string")
	fmt.Println()

	intStack := &Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("   Int Stack:")
	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("     Pop: %d\n", val)
		}
	}
	fmt.Println()

	strStack := &Stack[string]{}
	strStack.Push("Hello")
	strStack.Push("Go")
	strStack.Push("Generics")

	fmt.Println("   String Stack:")
	for !strStack.IsEmpty() {
		if val, ok := strStack.Pop(); ok {
			fmt.Printf("     Pop: %s\n", val)
		}
	}
	fmt.Println()

	// 2. Дженерики с ограничениями - числовые операции
	fmt.Println("2️⃣  Числовые операции с дженериками:")
	fmt.Println()

	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.5, 2.5, 3.5, 4.5}

	fmt.Printf("   Ints: %v\n", ints)
	fmt.Printf("   Sum: %d\n", Sum(ints))
	fmt.Printf("   Average: %.2f\n", Average(ints))
	fmt.Println()

	fmt.Printf("   Floats: %v\n", floats)
	fmt.Printf("   Sum: %.1f\n", Sum(floats))
	fmt.Printf("   Average: %.2f\n", Average(floats))
	fmt.Println()

	// 3. Дженерики с интерфейсами
	fmt.Println("3️⃣  Дженерики с интерфейсами:")
	fmt.Println()

	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	PrintAll(people)
	fmt.Println()

	// 4. Пары с разными типами
	fmt.Println("4️⃣  Пары (Pair) с разными типами:")
	fmt.Println()

	pair1 := NewPair("name", "Alice")
	pair2 := NewPair(1, 100)
	pair3 := NewPair(true, "yes")

	fmt.Printf("   Pair 1: %s\n", pair1)
	fmt.Printf("   Pair 2: %s\n", pair2)
	fmt.Printf("   Pair 3: %s\n", pair3)
	fmt.Println()

	// 5. Map и Filter
	fmt.Println("5️⃣  Функции Map и Filter:")
	fmt.Println()

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Map: умножить каждое число на 2
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("   Original: %v\n", numbers)
	fmt.Printf("   Doubled:  %v\n", doubled)
	fmt.Println()

	// Filter: только чётные числа
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("   Evens: %v\n", evens)
	fmt.Println()

	// Filter: только числа > 5
	greaterThan5 := Filter(numbers, func(n int) bool {
		return n > 5
	})
	fmt.Printf("   Greater than 5: %v\n", greaterThan5)
	fmt.Println()

	PrintFooter()
}
