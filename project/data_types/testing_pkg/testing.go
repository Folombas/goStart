package testing_pkg

import (
	"errors"
	"fmt"
	"strings"
)

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

// ========== ФУНКЦИИ ДЛЯ ТЕСТИРОВАНИЯ ==========

// Add - сложение двух чисел
func Add(a, b int) int {
	return a + b
}

// Divide - деление с ошибкой
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

// Greet - приветствие
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// IsValidEmail - простая валидация email
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// User - структура пользователя
type User struct {
	Name  string
	Email string
	Age   int
}

// NewUser - создание пользователя с валидацией
func NewUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("имя не может быть пустым")
	}
	if !IsValidEmail(email) {
		return nil, errors.New("невалидный email")
	}
	if age < 0 {
		return nil, errors.New("возраст не может быть отрицательным")
	}

	return &User{Name: name, Email: email, Age: age}, nil
}

// IsAdult - проверка совершеннолетия
func (u *User) IsAdult() bool {
	return u.Age >= 18
}

// ========== КАЛЬКУЛЯТОР ДЛЯ ТЕСТИРОВАНИЯ ==========

// Calculator - калькулятор
type Calculator struct {
	lastResult int
}

// NewCalculator создаёт калькулятор
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add складывает два числа
func (c *Calculator) Add(a, b int) int {
	c.lastResult = a + b
	return c.lastResult
}

// Subtract вычитает
func (c *Calculator) Subtract(a, b int) int {
	c.lastResult = a - b
	return c.lastResult
}

// Multiply умножает
func (c *Calculator) Multiply(a, b int) int {
	c.lastResult = a * b
	return c.lastResult
}

// Divide делит
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	c.lastResult = a / b
	return c.lastResult, nil
}

// GetLastResult возвращает последний результат
func (c *Calculator) GetLastResult() int {
	return c.lastResult
}

// Reset сбрасывает результат
func (c *Calculator) Reset() {
	c.lastResult = 0
}

// ========== ДЕМО ФУНКЦИИ ==========

// DemoBasicTests - демонстрация базовых тестов
func DemoBasicTests() {
	fmt.Println("1️⃣  Базовые тесты (примеры кода):")
	fmt.Println()

	fmt.Println(`   Файл: calculator_test.go`)
	fmt.Println()
	fmt.Println(`   func TestAdd(t *testing.T) {`)
	fmt.Println(`       result := Add(2, 3)`)
	fmt.Println(`       if result != 5 {`)
	fmt.Println(`           t.Errorf("ожидалось 5, получено %d", result)`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   func TestDivide(t *testing.T) {`)
	fmt.Println(`       result, err := Divide(10, 2)`)
	fmt.Println(`       if err != nil {`)
	fmt.Println(`           t.Fatal(err)`)
	fmt.Println(`       }`)
	fmt.Println(`       if result != 5 {`)
	fmt.Println(`           t.Errorf("ожидалось 5, получено %d", result)`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
}

// DemoTableDrivenTests - table-driven тесты
func DemoTableDrivenTests() {
	fmt.Println("2️⃣  Table-driven тесты:")
	fmt.Println()

	fmt.Println(`   func TestCalculatorOperations(t *testing.T) {`)
	fmt.Println(`       tests := []struct {`)
	fmt.Println(`           name     string`)
	fmt.Println(`           a, b     int`)
	fmt.Println(`           op       string`)
	fmt.Println(`           expected int`)
	fmt.Println(`           wantErr  bool`)
	fmt.Println(`       }{`)
	fmt.Println(`           {"addition", 5, 3, "add", 8, false},`)
	fmt.Println(`           {"subtraction", 10, 4, "subtract", 6, false},`)
	fmt.Println(`           {"multiplication", 3, 7, "multiply", 21, false},`)
	fmt.Println(`           {"division", 20, 4, "divide", 5, false},`)
	fmt.Println(`           {"division by zero", 5, 0, "divide", 0, true},`)
	fmt.Println(`       }`)
	fmt.Println()
	fmt.Println(`       calc := NewCalculator()`)
	fmt.Println(`       for _, tt := range tests {`)
	fmt.Println(`           t.Run(tt.name, func(t *testing.T) {`)
	fmt.Println(`               var result int`)
	fmt.Println(`               var err error`)
	fmt.Println()
	fmt.Println(`               switch tt.op {`)
	fmt.Println(`               case "add":`)
	fmt.Println(`                   result = calc.Add(tt.a, tt.b)`)
	fmt.Println(`               case "subtract":`)
	fmt.Println(`                   result = calc.Subtract(tt.a, tt.b)`)
	fmt.Println(`               case "multiply":`)
	fmt.Println(`                   result = calc.Multiply(tt.a, tt.b)`)
	fmt.Println(`               case "divide":`)
	fmt.Println(`                   result, err = calc.Divide(tt.a, tt.b)`)
	fmt.Println(`               }`)
	fmt.Println()
	fmt.Println(`               if tt.wantErr && err == nil {`)
	fmt.Println(`                   t.Error("ожидалась ошибка")`)
	fmt.Println(`               }`)
	fmt.Println(`               if result != tt.expected {`)
	fmt.Println(`                   t.Errorf("ожидалось %d, получено %d", tt.expected, result)`)
	fmt.Println(`               }`)
	fmt.Println(`           })`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
}

// DemoTestHelpers - тестовые хелперы
func DemoTestHelpers() {
	fmt.Println("3️⃣  Тестовые хелперы:")
	fmt.Println()

	fmt.Println(`   // createTestUser создаёт тестового пользователя`)
	fmt.Println(`   func createTestUser(t *testing.T, name, email string, age int) *User {`)
	fmt.Println(`       t.Helper() // Важно! Указывает что это хелпер`)
	fmt.Println()
	fmt.Println(`       user, err := NewUser(name, email, age)`)
	fmt.Println(`       if err != nil {`)
	fmt.Println(`           t.Fatalf("не удалось создать пользователя: %v", err)`)
	fmt.Println(`       }`)
	fmt.Println(`       return user`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   func TestUserIsAdult(t *testing.T) {`)
	fmt.Println(`       adult := createTestUser(t, "Alice", "alice@example.com", 25)`)
	fmt.Println(`       if !adult.IsAdult() {`)
	fmt.Println(`           t.Error("пользователь 25 лет должен быть совершеннолетним")`)
	fmt.Println(`       }`)
	fmt.Println()
	fmt.Println(`       minor := createTestUser(t, "Bob", "bob@example.com", 15)`)
	fmt.Println(`       if minor.IsAdult() {`)
	fmt.Println(`           t.Error("пользователь 15 лет не должен быть совершеннолетним")`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
}

// DemoTestSetup - setup и teardown
func DemoTestSetup() {
	fmt.Println("4️⃣  Setup и Teardown:")
	fmt.Println()

	fmt.Println(`   func TestMain(m *testing.M) {`)
	fmt.Println(`       // Setup - выполняется перед всеми тестами`)
	fmt.Println(`       fmt.Println("Инициализация тестов...")`)
	fmt.Println(`       os.Setenv("TEST_MODE", "true")`)
	fmt.Println()
	fmt.Println(`       // Запуск тестов`)
	fmt.Println(`       code := m.Run()`)
	fmt.Println()
	fmt.Println(`       // Teardown - выполняется после всех тестов`)
	fmt.Println(`       fmt.Println("Очистка после тестов...")`)
	fmt.Println(`       os.Unsetenv("TEST_MODE")`)
	fmt.Println()
	fmt.Println(`       os.Exit(code)`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   // Setup/Teardown для отдельного теста:`)
	fmt.Println()
	fmt.Println(`   func TestWithSetup(t *testing.T) {`)
	fmt.Println(`       // Setup`)
	fmt.Println(`       db := setupTestDB(t)`)
	fmt.Println(`       defer teardownTestDB(t, db)`)
	fmt.Println()
	fmt.Println(`       // Тест`)
	fmt.Println(`       // ...`)
	fmt.Println(`   }`)
	fmt.Println()
}

// DemoMocking - моки и заглушки
func DemoMocking() {
	fmt.Println("5️⃣  Моки и заглушки:")
	fmt.Println()

	fmt.Println(`   // Интерфейс для базы данных`)
	fmt.Println(`   type Database interface {`)
	fmt.Println(`       GetUser(id int) (*User, error)`)
	fmt.Println(`       SaveUser(user *User) error`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   // MockDatabase - мок для тестирования`)
	fmt.Println(`   type MockDatabase struct {`)
	fmt.Println(`       GetUserFunc func(id int) (*User, error)`)
	fmt.Println(`       SaveUserFunc func(user *User) error`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   func (m *MockDatabase) GetUser(id int) (*User, error) {`)
	fmt.Println(`       return m.GetUserFunc(id)`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   func (m *MockDatabase) SaveUser(user *User) error {`)
	fmt.Println(`       return m.SaveUserFunc(user)`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   // Тест с моком:`)
	fmt.Println(`   func TestUserService(t *testing.T) {`)
	fmt.Println(`       mockDB := &MockDatabase{`)
	fmt.Println(`           GetUserFunc: func(id int) (*User, error) {`)
	fmt.Println(`               if id == 1 {`)
	fmt.Println(`                   return &User{Name: "Test"}, nil`)
	fmt.Println(`               }`)
	fmt.Println(`               return nil, errors.New("not found")`)
	fmt.Println(`           },`)
	fmt.Println(`       }`)
	fmt.Println()
	fmt.Println(`       user, err := mockDB.GetUser(1)`)
	fmt.Println(`       if err != nil {`)
	fmt.Println(`           t.Fatal(err)`)
	fmt.Println(`       }`)
	fmt.Println(`       if user.Name != "Test" {`)
	fmt.Println(`           t.Errorf("ожидалось Test, получено %s", user.Name)`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
}

// DemoBenchmarks - бенчмарки
func DemoBenchmarks() {
	fmt.Println("6️⃣  Бенчмарки (тесты производительности):")
	fmt.Println()

	fmt.Println(`   func BenchmarkAdd(b *testing.B) {`)
	fmt.Println(`       for i := 0; i < b.N; i++ {`)
	fmt.Println(`           Add(2, 3)`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   func BenchmarkCalculator(b *testing.B) {`)
	fmt.Println(`       calc := NewCalculator()`)
	fmt.Println(`       b.ResetTimer()`)
	fmt.Println(`       for i := 0; i < b.N; i++ {`)
	fmt.Println(`           calc.Add(100, 200)`)
	fmt.Println(`       }`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   // Запуск бенчмарков:`)
	fmt.Println(`   go test -bench=. -benchmem`)
	fmt.Println()
	fmt.Println(`   // Пример вывода:`)
	fmt.Println(`   BenchmarkAdd-8-100000000-1.2ns`)
	fmt.Println(`   BenchmarkCalculator-8-50000000-2.5ns`)
	fmt.Println()
}

// DemoCoverage - покрытие тестами
func DemoCoverage() {
	fmt.Println("7️⃣  Покрытие тестами (Code Coverage):")
	fmt.Println()

	fmt.Println(`   // Запуск с покрытием:`)
	fmt.Println(`   go test -cover`)
	fmt.Println()
	fmt.Println(`   // Покрытие с отчётом:`)
	fmt.Println(`   go test -coverprofile=coverage.out`)
	fmt.Println(`   go tool cover -html=coverage.out`)
	fmt.Println()
	fmt.Println(`   // Покрытие по пакетам:`)
	fmt.Println(`   go test -coverpkg=./...`)
	fmt.Println()
	fmt.Println(`   // Пример вывода:`)
	fmt.Println(`   coverage: 85.7% of statements`)
	fmt.Println()
	fmt.Println("   Советы по хорошему покрытию:")
	fmt.Println("   - Стремитесь к 70-80% покрытию")
	fmt.Println("   - 100% покрытие не всегда целесообразно")
	fmt.Println("   - Тестируйте важную бизнес-логику")
	fmt.Println("   - Не забывайте про тесты ошибок")
	fmt.Println()
}

// DemoRaceDetector - детектор гонок
func DemoRaceDetector() {
	fmt.Println("8️⃣  Детектор гонок (Race Detector):")
	fmt.Println()

	fmt.Println(`   // Запуск с детектором гонок:`)
	fmt.Println(`   go test -race`)
	fmt.Println()
	fmt.Println(`   // Пример теста с гонкой:`)
	fmt.Println(`   func TestRaceCondition(t *testing.T) {`)
	fmt.Println(`       counter := 0`)
	fmt.Println(`       var wg sync.WaitGroup`)
	fmt.Println()
	fmt.Println(`       for i := 0; i < 100; i++ {`)
	fmt.Println(`           wg.Add(1)`)
	fmt.Println(`           go func() {`)
	fmt.Println(`               defer wg.Done()`)
	fmt.Println(`               counter++ // Гонка!`)
	fmt.Println(`           }()`)
	fmt.Println(`       }`)
	fmt.Println(`       wg.Wait()`)
	fmt.Println(`   }`)
	fmt.Println()
	fmt.Println(`   // Вывод race detector:`)
	fmt.Println(`   WARNING: DATA RACE`)
	fmt.Println(`   Read at 0x00c000014098 by goroutine 7:`)
	fmt.Println(`   Previous write at 0x00c000014098 by goroutine 6:`)
	fmt.Println()
	fmt.Println("   Решение: используйте mutex или атомарные операции")
	fmt.Println()
}

// DemoBestPractices - лучшие практики
func DemoBestPractices() {
	fmt.Println("9️⃣  Лучшие практики тестирования:")
	fmt.Println()

	fmt.Println("   ✅ Имена тестов: Test<Function>_<Scenario>_<Expected>")
	fmt.Println("      Пример: TestDivide_ByZero_ReturnsError")
	fmt.Println()
	fmt.Println("   ✅ Используйте t.Run() для подтестов")
	fmt.Println()
	fmt.Println("   ✅ Table-driven тесты для однотипных случаев")
	fmt.Println()
	fmt.Println("   ✅ t.Helper() для хелпер функций")
	fmt.Println()
	fmt.Println("   ✅ defer для cleanup")
	fmt.Println()
	fmt.Println("   ✅ Не пропускайте ошибки в тестах")
	fmt.Println()
	fmt.Println("   ✅ Тестируйте не только успех, но и ошибки")
	fmt.Println()
	fmt.Println("   ✅ Используйте -race в CI/CD")
	fmt.Println()
	fmt.Println("   ✅ Mock внешние зависимости")
	fmt.Println()
	fmt.Println("   ✅ Интеграционные тесты отдельно от юнит-тестов")
	fmt.Println()
}

// DemoTesting - основная демонстрационная функция
func DemoTesting() {
	PrintHeader("🧪 Тестирование в Go")

	DemoBasicTests()
	DemoTableDrivenTests()
	DemoTestHelpers()
	DemoTestSetup()
	DemoMocking()
	DemoBenchmarks()
	DemoCoverage()
	DemoRaceDetector()
	DemoBestPractices()

	PrintFooter()
}
