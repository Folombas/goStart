package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

// ========== БАЗОВАЯ ЗАПИСЬ В ФАЙЛ ==========

// DemoWriteFile - демонстрация записи в файл
func DemoWriteFile() {
	fmt.Println("1️⃣  Запись в файл (WriteFile):")
	fmt.Println()

	// Простая запись - файл создаётся или перезаписывается
	content := "Привет, Go!\nЭто первая строка.\nЭто вторая строка."
	
	err := os.WriteFile("temp_write.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	fmt.Println("   Файл temp_write.txt создан")
	fmt.Printf("   Содержимое записано: %d байт\n", len(content))
	fmt.Println()

	// Читаем обратно для проверки
	data, _ := os.ReadFile("temp_write.txt")
	fmt.Printf("   Прочитано из файла:\n")
	fmt.Printf("   %s\n", string(data))

	// Удаляем тестовый файл
	os.Remove("temp_write.txt")
	fmt.Println()
}

// ========== БАЗОВОЕ ЧТЕНИЕ ИЗ ФАЙЛА ==========

// DemoReadFile - демонстрация чтения из файла
func DemoReadFile() {
	fmt.Println("2️⃣  Чтение из файла (ReadFile):")
	fmt.Println()

	// Создаём тестовый файл
	testContent := "Строка 1\nСтрока 2\nСтрока 3\nСтрока 4\nСтрока 5"
	os.WriteFile("temp_read.txt", []byte(testContent), 0644)

	// Читаем весь файл
	data, err := os.ReadFile("temp_read.txt")
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}

	fmt.Printf("   Весь файл (%d байт):\n", len(data))
	fmt.Printf("   %s\n", string(data))
	fmt.Println()

	// Читаем по частям
	fmt.Println("   Чтение по частям (по 10 байт):")
	file, _ := os.Open("temp_read.txt")
	buffer := make([]byte, 10)
	
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Printf("   Прочитано %d байт: %q\n", n, string(buffer[:n]))
		}
		if err == io.EOF {
			break
		}
	}
	file.Close()

	os.Remove("temp_read.txt")
	fmt.Println()
}

// ========== ЧТЕНИЕ ПОСТРОЧНО ==========

// DemoReadLineByLine - чтение построчно
func DemoReadLineByLine() {
	fmt.Println("3️⃣  Чтение построчно (bufio.Scanner):")
	fmt.Println()

	// Создаём тестовый файл
	lines := "Первая строка\nВторая строка\nТретья строка\nЧетвёртая строка"
	os.WriteFile("temp_lines.txt", []byte(lines), 0644)

	file, _ := os.Open("temp_lines.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		fmt.Printf("   Строка %d: %s\n", lineNum, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("   Ошибка при чтении: %v\n", err)
	}

	os.Remove("temp_lines.txt")
	fmt.Println()
}

// ========== ЗАПИСЬ ПОСТРОЧНО ==========

// DemoWriteLineByLine - запись построчно
func DemoWriteLineByLine() {
	fmt.Println("4️⃣  Запись построчно (bufio.Writer):")
	fmt.Println()

	file, _ := os.Create("temp_write_lines.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)

	lines := []string{
		"Первая строка",
		"Вторая строка",
		"Третья строка",
		"Четвёртая строка",
	}

	for i, line := range lines {
		n, _ := writer.WriteString(fmt.Sprintf("%d: %s\n", i+1, line))
		fmt.Printf("   Записано %d байт\n", n)
	}

	// Обязательно вызываем Flush!
	writer.Flush()

	fmt.Println()
	fmt.Println("   Содержимое файла:")
	data, _ := os.ReadFile("temp_write_lines.txt")
	fmt.Printf("   %s", string(data))

	os.Remove("temp_write_lines.txt")
	fmt.Println()
}

// ========== РАБОТА С ФАЙЛАМИ (CREATE, OPEN, APPEND) ==========

// DemoFileOperations - операции с файлами
func DemoFileOperations() {
	fmt.Println("5️⃣  Операции с файлами (Create, Open, Append):")
	fmt.Println()

	// Create - создаёт или перезаписывает
	fmt.Println("   os.Create:")
	file1, _ := os.Create("temp_create.txt")
	file1.WriteString("Создано через Create\n")
	file1.Close()
	fmt.Println("   Файл создан/перезаписан")

	// OpenFile с разными флагами
	fmt.Println()
	fmt.Println("   os.OpenFile с флагами:")

	// Append - добавление в конец
	file2, _ := os.OpenFile("temp_create.txt", os.O_APPEND|os.O_WRONLY, 0644)
	file2.WriteString("Добавлено через Append\n")
	file2.Close()
	fmt.Println("   Добавлена строка (APPEND)")

	// Читаем результат
	data, _ := os.ReadFile("temp_create.txt")
	fmt.Printf("   Итоговое содержимое:\n   %s", string(data))

	os.Remove("temp_create.txt")
	fmt.Println()
}

// ========== РАБОТА С ПУТЯМИ ==========

// DemoFilePaths - работа с путями
func DemoFilePaths() {
	fmt.Println("6️⃣  Работа с путями (filepath):")
	fmt.Println()

	// Join - соединение путей
	path := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Printf("   filepath.Join: %s\n", path)

	// Base - имя файла
	fmt.Printf("   filepath.Base: %s\n", filepath.Base(path))

	// Dir - директория
	fmt.Printf("   filepath.Dir: %s\n", filepath.Dir(path))

	// Ext - расширение
	fmt.Printf("   filepath.Ext: %s\n", filepath.Ext(path))

	// Abs - абсолютный путь
	abs, _ := filepath.Abs("relative/path/file.txt")
	fmt.Printf("   filepath.Abs: %s\n", abs)

	// Clean - очистка пути
	clean := filepath.Clean("/home/user/../documents/./file.txt")
	fmt.Printf("   filepath.Clean: %s\n", clean)
	fmt.Println()
}

// ========== РАБОТА С ДИРЕКТОРИЯМИ ==========

// DemoDirectories - работа с директориями
func DemoDirectories() {
	fmt.Println("7️⃣  Работа с директориями:")
	fmt.Println()

	// Создаём директорию
	dir := "temp_test_dir"
	os.Mkdir(dir, 0755)
	os.MkdirAll(filepath.Join(dir, "subdir1", "subdir2"), 0755)
	fmt.Printf("   Создана директория: %s\n", dir)

	// Создаём тестовые файлы
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(dir, "file2.txt"), []byte("content2"), 0644)
	os.WriteFile(filepath.Join(dir, "subdir1", "file3.txt"), []byte("content3"), 0644)

	// Читаем директорию
	fmt.Println()
	fmt.Println("   Список файлов в директории:")
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		fileType := "файл"
		if entry.IsDir() {
			fileType = "директория"
		}
		fmt.Printf("   - %s (%s)\n", entry.Name(), fileType)
	}

	// Walk - рекурсивный обход
	fmt.Println()
	fmt.Println("   Рекурсивный обход (filepath.Walk):")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		indent := strings.Repeat("  ", strings.Count(path, string(os.PathSeparator))-1)
		fmt.Printf("   %s%s\n", indent, info.Name())
		return nil
	})

	// Удаляем директорию
	os.RemoveAll(dir)
	fmt.Println()
	fmt.Println("   Директория удалена")
	fmt.Println()
}

// ========== ПРОВЕРКА СУЩЕСТВОВАНИЯ ФАЙЛА ==========

// DemoFileExists - проверка существования файла
func DemoFileExists() {
	fmt.Println("8️⃣  Проверка существования файла:")
	fmt.Println()

	// Создаём тестовый файл
	os.WriteFile("temp_exists.txt", []byte("test"), 0644)

	// Проверка через os.Stat
	checkExists := func(path string) {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			fmt.Printf("   %s: не существует\n", path)
		} else if err != nil {
			fmt.Printf("   %s: ошибка %v\n", path, err)
		} else {
			fmt.Printf("   %s: существует\n", path)
		}
	}

	checkExists("temp_exists.txt")
	checkExists("nonexistent.txt")

	os.Remove("temp_exists.txt")
	fmt.Println()
}

// ========== КОПИРОВАНИЕ ФАЙЛОВ ==========

// DemoCopyFile - копирование файлов
func DemoCopyFile() {
	fmt.Println("9️⃣  Копирование файлов:")
	fmt.Println()

	// Создаём исходный файл
	sourceContent := "Это содержимое будет скопировано в другой файл."
	os.WriteFile("temp_source.txt", []byte(sourceContent), 0644)

	// Копируем файл
	source, _ := os.Open("temp_source.txt")
	defer source.Close()

	dest, _ := os.Create("temp_dest.txt")
	defer dest.Close()

	bytesWritten, _ := io.Copy(dest, source)
	fmt.Printf("   Скопировано %d байт\n", bytesWritten)

	// Проверяем
	destData, _ := os.ReadFile("temp_dest.txt")
	fmt.Printf("   Содержимое копии: %s\n", string(destData))

	os.Remove("temp_source.txt")
	os.Remove("temp_dest.txt")
	fmt.Println()
}

// DemoFileIO - основная демонстрационная функция
func DemoFileIO() {
	PrintHeader("📁 Работа с файлами (File I/O) в Go")

	DemoWriteFile()
	DemoReadFile()
	DemoReadLineByLine()
	DemoWriteLineByLine()
	DemoFileOperations()
	DemoFilePaths()
	DemoDirectories()
	DemoFileExists()
	DemoCopyFile()

	PrintFooter()
}
