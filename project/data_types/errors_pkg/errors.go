package errors_pkg

import (
	"errors"
	"fmt"
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

// ========== БАЗОВЫЕ ОШИБКИ ==========

// ValidationError - ошибка валидации
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field '%s' - %s", e.Field, e.Message)
}

// DatabaseError - ошибка базы данных
type DatabaseError struct {
	Operation string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s: %v", e.Operation, e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// ========== CUSTOM ERROR С METADATA ==========

// APIError - ошибка API с кодом и метаданными
type APIError struct {
	Code       int
	Message    string
	Details    map[string]interface{}
	StatusCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

func (e *APIError) GetStatusCode() int {
	return e.StatusCode
}

// ========== ОШИБКИ С ОБЁРТЫВАНИЕМ ==========

// BusinessError - бизнес-ошибка с обёрткой
type BusinessError struct {
	Operation string
	Reason    string
	Err       error
}

func (e *BusinessError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("business error during %s (%s): %v", e.Operation, e.Reason, e.Err)
	}
	return fmt.Sprintf("business error during %s: %s", e.Operation, e.Reason)
}

func (e *BusinessError) Unwrap() error {
	return e.Err
}

// ========== ФУНКЦИИ ДЛЯ РАБОТЫ С ОШИБКАМИ ==========

// NewValidationError создаёт ошибку валидации
func NewValidationError(field, message string) error {
	return &ValidationError{Field: field, Message: message}
}

// NewDatabaseError создаёт ошибку БД
func NewDatabaseError(operation string, err error) error {
	return &DatabaseError{Operation: operation, Err: err}
}

// NewAPIError создаёт ошибку API
func NewAPIError(code int, message string, statusCode int) *APIError {
	return &APIError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Details:    make(map[string]interface{}),
	}
}

// WithDetail добавляет деталь к ошибке API
func (e *APIError) WithDetail(key string, value interface{}) *APIError {
	e.Details[key] = value
	return e
}

// NewBusinessError создаёт бизнес-ошибку
func NewBusinessError(operation, reason string, err error) error {
	return &BusinessError{Operation: operation, Reason: reason, Err: err}
}

// ========== ДЕМО ФУНКЦИИ ==========

// DemoBasicErrors - демонстрация базовых ошибок
func DemoBasicErrors() {
	fmt.Println("1️⃣  Базовые ошибки:")
	fmt.Println()

	// Простая ошибка
	err := errors.New("простая ошибка")
	fmt.Printf("   Простая ошибка: %v\n", err)
	fmt.Println()

	// Форматированная ошибка
	name := ""
	if name == "" {
		err := fmt.Errorf("имя не может быть пустым")
		fmt.Printf("   Форматированная ошибка: %v\n", err)
	}
	fmt.Println()

	// Ошибка с контекстом
	userID := 42
	err = fmt.Errorf("пользователь %d не найден: %w", userID, errors.New("not found"))
	fmt.Printf("   Ошибка с контекстом: %v\n", err)
	fmt.Println()
}

// DemoCustomErrors - демонстрация кастомных ошибок
func DemoCustomErrors() {
	fmt.Println("2️⃣  Кастомные типы ошибок:")
	fmt.Println()

	// Ошибка валидации
	err := NewValidationError("email", "должен быть валидным email адресом")
	fmt.Printf("   ValidationError: %v\n", err)
	fmt.Printf("   Type: %T\n\n", err)

	// Ошибка базы данных
	dbErr := NewDatabaseError("SELECT users", errors.New("connection timeout"))
	fmt.Printf("   DatabaseError: %v\n", dbErr)
	fmt.Printf("   Type: %T\n\n", dbErr)

	// Ошибка API
	apiErr := NewAPIError(1001, "Недостаточно средств", 400)
	apiErr.WithDetail("balance", 100).WithDetail("required", 250)
	fmt.Printf("   APIError: %v\n", apiErr)
	fmt.Printf("   StatusCode: %d\n", apiErr.GetStatusCode())
	fmt.Printf("   Details: %v\n", apiErr.Details)
	fmt.Println()
}

// DemoErrorWrapping - демонстрация обёртывания ошибок
func DemoErrorWrapping() {
	fmt.Println("3️⃣  Обёртывание ошибок:")
	fmt.Println()

	// Базовая ошибка
	baseErr := errors.New("файл не существует")

	// Обёртываем с контекстом
	wrappedErr := fmt.Errorf("чтение конфига: %w", baseErr)
	fmt.Printf("   Wrapped error: %v\n", wrappedErr)
	fmt.Println()

	// Проверяем тип ошибки
	fmt.Println("   Проверка типа ошибки:")
	fmt.Printf("   errors.Is(wrappedErr, baseErr): %v\n", errors.Is(wrappedErr, baseErr))
	fmt.Println()

	// Распаковка ошибки
	fmt.Println("   Распаковка ошибки:")
	unwrapped := errors.Unwrap(wrappedErr)
	fmt.Printf("   Unwrapped: %v\n", unwrapped)
	fmt.Println()

	// Извлечение конкретной ошибки
	fmt.Println("   Извлечение конкретной ошибки:")
	var target *ValidationError
	customErr := NewValidationError("age", "должно быть больше 0")
	wrappedCustom := fmt.Errorf("валидация пользователя: %w", customErr)
	fmt.Printf("   errors.As(wrappedCustom, &target): %v\n", errors.As(wrappedCustom, &target))
	fmt.Println()
}

// DemoErrorHandling - демонстрация обработки ошибок
func DemoErrorHandling() {
	fmt.Println("4️⃣  Обработка ошибок:")
	fmt.Println()

	// Пример обработки с проверкой типа
	err := NewDatabaseError("INSERT", errors.New("duplicate key"))

	var dbErr *DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("   Обнаружена ошибка БД!\n")
		fmt.Printf("   Операция: %s\n", dbErr.Operation)
		fmt.Printf("   Причина: %v\n", dbErr.Err)
	}
	fmt.Println()

	// Пример с switch по типу
	handleError := func(err error) {
		switch e := err.(type) {
		case *ValidationError:
			fmt.Printf("   Ошибка валидации поля '%s': %s\n", e.Field, e.Message)
		case *DatabaseError:
			fmt.Printf("   Ошибка БД при %s: %v\n", e.Operation, e.Err)
		case *APIError:
			fmt.Printf("   API ошибка %d: %s\n", e.Code, e.Message)
		default:
			fmt.Printf("   Неизвестная ошибка: %v\n", err)
		}
	}

	errs := []error{
		NewValidationError("password", "минимум 8 символов"),
		NewDatabaseError("UPDATE", errors.New("lock timeout")),
		NewAPIError(404, "Not Found", 404),
		errors.New("обычная ошибка"),
	}

	for _, e := range errs {
		handleError(e)
	}
	fmt.Println()
}

// DemoBestPractices - демонстрация лучших практик
func DemoBestPractices() {
	fmt.Println("5️⃣  Лучшие практики:")
	fmt.Println()

	fmt.Println("   ✅ Всегда проверяйте ошибки:")
	fmt.Println("   if err != nil { return err }")
	fmt.Println()

	fmt.Println("   ✅ Обёртывайте ошибки с контекстом:")
	fmt.Println("   fmt.Errorf(\"context: %w\", err)")
	fmt.Println()

	fmt.Println("   ✅ Используйте кастомные типы для ожидаемых ошибок:")
	fmt.Println("   type ValidationError struct {...}")
	fmt.Println()

	fmt.Println("   ✅ Избегайте множественных return err:")
	fmt.Println("   Используйте ранний возврат")
	fmt.Println()

	fmt.Println("   ✅ Логируйте ошибки в точках возникновения:")
	fmt.Println("   log.Printf(\"error: %v\", err)")
	fmt.Println()

	fmt.Println("   ✅ Не игнорируйте ошибки:")
	fmt.Println("   Избегайте _ = someFunc()")
	fmt.Println()
}

// DemoPracticalExample - практический пример
func DemoPracticalExample() {
	fmt.Println("6️⃣  Практический пример - регистрация пользователя:")
	fmt.Println()

	// Имитация регистрации
	type User struct {
		ID    int
		Email string
		Name  string
	}

	registerUser := func(email, name string) (*User, error) {
		// Валидация email
		if email == "" {
			return nil, NewValidationError("email", "не может быть пустым")
		}

		// Валидация имени
		if name == "" {
			return nil, NewValidationError("name", "не может быть пустым")
		}

		// Имитация проверки в БД
		if email == "taken@example.com" {
			dbErr := NewDatabaseError("SELECT", errors.New("user exists"))
			return nil, NewBusinessError("register", "user already exists", dbErr)
		}

		// Успешная регистрация
		return &User{ID: 1, Email: email, Name: name}, nil
	}

	// Тестовые случаи
	testCases := []struct {
		email string
		name  string
	}{
		{"", "Alice"},
		{"alice@example.com", ""},
		{"taken@example.com", "Bob"},
		{"valid@example.com", "Charlie"},
	}

	for i, tc := range testCases {
		fmt.Printf("   Тест %d: email=%q, name=%q\n", i+1, tc.email, tc.name)
		user, err := registerUser(tc.email, tc.name)
		if err != nil {
			fmt.Printf("   ❌ Ошибка: %v\n", err)

			// Детальная обработка
			var ve *ValidationError
			if errors.As(err, &ve) {
				fmt.Printf("      → Это ошибка валидации!\n")
			}
			var be *BusinessError
			if errors.As(err, &be) {
				fmt.Printf("      → Это бизнес-ошибка: %s\n", be.Reason)
			}
		} else {
			fmt.Printf("   ✅ Успешно: %+v\n", user)
		}
		fmt.Println()
	}
}

// DemoErrors - основная демонстрационная функция
func DemoErrors() {
	PrintHeader("❌ Обработка ошибок в Go")

	DemoBasicErrors()
	DemoCustomErrors()
	DemoErrorWrapping()
	DemoErrorHandling()
	DemoBestPractices()
	DemoPracticalExample()

	PrintFooter()
}
