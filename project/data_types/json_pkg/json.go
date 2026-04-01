package json_pkg

import (
	"encoding/json"
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

// ========== БАЗОВЫЕ СТРУКТУРЫ ==========

// Person - базовая структура для примеров
type Person struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Age   int      `json:"age,omitempty"` // omitempty - не включать, если 0
	Hobby []string `json:"hobby,omitempty"`
}

// User - структура с разными тегами
type User struct {
	ID       int     `json:"user_id"`
	Username string  `json:"username"`
	Password string  `json:"-"`              // - не включать в JSON
	Salary   float64 `json:"salary,string"`  // string - кодировать как строку
	Active   bool    `json:"active"`
}

// Config - структура с вложенными объектами
type Config struct {
	AppName    string `json:"app_name"`
	Version    string `json:"version"`
	Debug      bool   `json:"debug"`
	Database   DBConfig `json:"database"`
	Features   []string `json:"features"`
}

// DBConfig - вложенная структура
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ========== MARSHAL (кодирование в JSON) ==========

// DemoMarshal - демонстрация кодирования в JSON
func DemoMarshal() {
	fmt.Println("1️⃣  Marshal (кодирование в JSON):")
	fmt.Println()

	person := Person{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
		Hobby: []string{"чтение", "программирование", "велоспорт"},
	}

	// Кодируем в JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	fmt.Printf("   JSON: %s\n", string(jsonData))
	fmt.Println()

	// MarshalIndent - красивое форматирование
	fmt.Println("   С форматированием (MarshalIndent):")
	jsonIndent, _ := json.MarshalIndent(person, "   ", "  ")
	fmt.Printf("   %s\n", string(jsonIndent))
	fmt.Println()
}

// DemoMarshalTags - демонстрация тегов JSON
func DemoMarshalTags() {
	fmt.Println("2️⃣  Теги JSON:")
	fmt.Println()

	user := User{
		ID:       42,
		Username: "john_doe",
		Password: "secret123", // Не попадёт в JSON из-за тега "-"
		Salary:   5000.50,
		Active:   true,
	}

	jsonData, _ := json.MarshalIndent(user, "   ", "  ")
	fmt.Printf("   User с тегами:\n")
	fmt.Printf("   %s\n", string(jsonData))
	fmt.Println()
	fmt.Println("   Обратите внимание:")
	fmt.Println("   - Password отсутствует (тег \"-\")")
	fmt.Println("   - Salary закодирован как строка (тег \",string\")")
	fmt.Println("   - user_id вместо ID (кастомное имя)")
	fmt.Println()
}

// ========== UNMARSHAL (декодирование из JSON) ==========

// DemoUnmarshal - демонстрация декодирования из JSON
func DemoUnmarshal() {
	fmt.Println("3️⃣  Unmarshal (декодирование из JSON):")
	fmt.Println()

	jsonString := `{
		"id": 2,
		"name": "Bob",
		"email": "bob@example.com",
		"age": 30,
		"hobby": ["футбол", "музыка"]
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	fmt.Printf("   Декодированный Person:\n")
	fmt.Printf("   ID: %d\n", person.ID)
	fmt.Printf("   Name: %s\n", person.Name)
	fmt.Printf("   Email: %s\n", person.Email)
	fmt.Printf("   Age: %d\n", person.Age)
	fmt.Printf("   Hobby: %v\n", person.Hobby)
	fmt.Println()
}

// DemoUnmarshalComplex - декодирование сложных структур
func DemoUnmarshalComplex() {
	fmt.Println("4️⃣  Декодирование сложных структур:")
	fmt.Println()

	jsonString := `{
		"app_name": "MyApp",
		"version": "1.0.0",
		"debug": true,
		"database": {
			"host": "localhost",
			"port": 5432,
			"username": "admin",
			"password": "secret"
		},
		"features": ["auth", "logging", "cache"]
	}`

	var config Config
	err := json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	fmt.Printf("   Config:\n")
	fmt.Printf("   AppName: %s\n", config.AppName)
	fmt.Printf("   Version: %s\n", config.Version)
	fmt.Printf("   Debug: %v\n", config.Debug)
	fmt.Printf("   Database.Host: %s\n", config.Database.Host)
	fmt.Printf("   Database.Port: %d\n", config.Database.Port)
	fmt.Printf("   Features: %v\n", config.Features)
	fmt.Println()
}

// ========== РАБОТА С MAP ==========

// DemoMapJSON - работа с map[string]interface{}
func DemoMapJSON() {
	fmt.Println("5️⃣  Работа с map[string]interface{}:")
	fmt.Println()

	// Когда структура неизвестна заранее
	jsonString := `{
		"name": "Unknown Product",
		"price": 99.99,
		"in_stock": true,
		"tags": ["sale", "new"],
		"details": {
			"weight": "1.5kg",
			"color": "blue"
		}
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	fmt.Printf("   Распаковка в map:\n")
	for key, value := range data {
		fmt.Printf("   %s: %v (%T)\n", key, value, value)
	}
	fmt.Println()

	// Доступ к вложенным значениям с проверкой типа
	if details, ok := data["details"].(map[string]interface{}); ok {
		fmt.Printf("   Вложенный объект details:\n")
		for k, v := range details {
			fmt.Printf("     %s: %v\n", k, v)
		}
	}
	fmt.Println()
}

// ========== JSON ENCODER / DECODER ==========

// DemoEncoderDecoder - использование Encoder и Decoder
func DemoEncoderDecoder() {
	fmt.Println("6️⃣  Encoder и Decoder:")
	fmt.Println()

	people := []Person{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35},
	}

	// Используем Encoder для записи в strings.Builder
	var builder strings.Builder
	encoder := json.NewEncoder(&builder)
	encoder.SetIndent("", "  ")

	fmt.Println("   Encoder (запись в builder):")
	for _, person := range people {
		if err := encoder.Encode(person); err != nil {
			fmt.Printf("   Ошибка: %v\n", err)
		}
	}
	fmt.Printf("   %s\n", builder.String())

	// Используем Decoder для чтения
	fmt.Println("   Decoder (чтение из строки):")
	jsonString := `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`
	decoder := json.NewDecoder(strings.NewReader(jsonString))

	var decoded []Person
	if err := decoder.Decode(&decoded); err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	for _, p := range decoded {
		fmt.Printf("   %+v\n", p)
	}
	fmt.Println()
}

// ========== ВАЛИДАЦИЯ JSON ==========

// DemoJSONValidation - валидация JSON
func DemoJSONValidation() {
	fmt.Println("7️⃣  Валидация JSON:")
	fmt.Println()

	validJSON := `{"name": "Alice", "age": 25}`
	invalidJSON := `{"name": "Bob", "age": }`

	// Проверка валидности JSON
	fmt.Println("   Проверка валидности:")

	var v interface{}
	err := json.Unmarshal([]byte(validJSON), &v)
	fmt.Printf("   Valid JSON: %v\n", err == nil)

	err = json.Unmarshal([]byte(invalidJSON), &v)
	fmt.Printf("   Invalid JSON: %v (ошибка: %v)\n", err != nil, err)
	fmt.Println()

	// json.Valid для быстрой проверки
	fmt.Println("   Быстрая проверка с json.Valid:")
	fmt.Printf("   Valid: %v\n", json.Valid([]byte(validJSON)))
	fmt.Printf("   Invalid: %v\n", json.Valid([]byte(invalidJSON)))
	fmt.Println()
}

// ========== CUSTOM MARSHAL/UNMARSHAL ==========

// CustomDate - тип с кастомной сериализацией
type CustomDate struct {
	Year  int
	Month int
	Day   int
}

// MarshalJSON - кастомная сериализация
func (d CustomDate) MarshalJSON() ([]byte, error) {
	dateStr := fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
	return json.Marshal(dateStr)
}

// UnmarshalJSON - кастомная десериализация
func (d *CustomDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	fmt.Sscanf(dateStr, "%d-%d-%d", &d.Year, &d.Month, &d.Day)
	return nil
}

// DemoCustomMarshal - кастомная сериализация
func DemoCustomMarshal() {
	fmt.Println("8️⃣  Кастомная сериализация (MarshalJSON/UnmarshalJSON):")
	fmt.Println()

	date := CustomDate{Year: 2024, Month: 1, Day: 15}

	jsonData, _ := json.Marshal(date)
	fmt.Printf("   CustomDate в JSON: %s\n", string(jsonData))

	var newDate CustomDate
	json.Unmarshal(jsonData, &newDate)
	fmt.Printf("   JSON в CustomDate: %+v\n", newDate)
	fmt.Println()
}

// DemoJSON - основная демонстрационная функция
func DemoJSON() {
	PrintHeader("📦 Работа с JSON в Go")

	DemoMarshal()
	DemoMarshalTags()
	DemoUnmarshal()
	DemoUnmarshalComplex()
	DemoMapJSON()
	DemoEncoderDecoder()
	DemoJSONValidation()
	DemoCustomMarshal()

	PrintFooter()
}
