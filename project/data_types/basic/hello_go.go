package basic

import "fmt"

// HelloGo - первая программа на Go
// Демонстрирует базовый синтаксис и вывод на экран
func HelloGo() {
	// Вывод текста на экран
	fmt.Println("👋 Привет, Go!")
	fmt.Println("🎉 Добро пожаловать в мир языка Go!")
	
	// Пустая строка для разделения
	fmt.Println()
	
	// Вывод с форматированием
	name := "Go-Разработчик"
	level := 1
	fmt.Printf("📛 Имя: %s\n", name)
	fmt.Printf("🏆 Уровень: %d\n", level)
	fmt.Println()
	
	// Базовые операции
	fmt.Println("📚 Основы Go:")
	fmt.Println("  • Переменные объявляются через var или :=")
	fmt.Println("  • Типы данных: int, string, bool, float64")
	fmt.Println("  • Функции начинаются с ключевого слова func")
	fmt.Println("  • Импорт пакетов через import")
	fmt.Println()
	
	// Пример с переменными
	var goVersion string = "1.24"
	year := 2026
	isLearning := true
	rating := 95.5
	
	fmt.Println("💡 Информация о языке:")
	fmt.Printf("  Версия Go: %s\n", goVersion)
	fmt.Printf("  Год: %d\n", year)
	fmt.Printf("  Изучаю Go: %v\n", isLearning)
	fmt.Printf("  Рейтинг языка: %.1f/100\n", rating)
	fmt.Println()
	
	// Мотивация
	fmt.Println("🚀 Мотивация на сегодня:")
	fmt.Println(`  "Каждая строка кода на Go — кирпичик в фундаменте твоей карьеры!"`)
	fmt.Println()
	
	fmt.Println("✅ Первая программа на Go выполнена успешно!")
}

// PrintSeparator выводит разделитель
func PrintSeparator() {
	fmt.Println("========================================")
}

// PrintHeader выводит заголовок программы
func PrintHeader(title string) {
	fmt.Println()
	PrintSeparator()
	fmt.Printf("🚀 %s\n", title)
	PrintSeparator()
}

// PrintFooter выводит подвал программы
func PrintFooter() {
	PrintSeparator()
	fmt.Println("🎉 Программа успешно завершена!")
	fmt.Println()
}
