package httpclient

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
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

// ========== БАЗОВЫЙ GET ЗАПРОС ==========

// DemoBasicGet - базовый GET запрос
func DemoBasicGet() {
	fmt.Println("1️⃣  Базовый GET запрос:")
	fmt.Println()

	// Используем httpbin.org для тестирования
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Printf("   StatusCode: %d\n", resp.StatusCode)
	fmt.Printf("   Content-Length: %d\n", resp.ContentLength)
	fmt.Printf("   Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// Читаем тело ответа
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n   Тело ответа (первые 200 символов):\n")
	if len(body) > 200 {
		fmt.Printf("   %s...\n", string(body[:200]))
	} else {
		fmt.Printf("   %s\n", string(body))
	}
	fmt.Println()
}

// ========== GET С ПАРАМЕТРАМИ ==========

// DemoGetWithParams - GET с параметрами
func DemoGetWithParams() {
	fmt.Println("2️⃣  GET с параметрами:")
	fmt.Println()

	// Способ 1: вручную в URL
	url1 := "https://httpbin.org/get?name=Alice&age=25&city=Moscow"
	resp1, _ := http.Get(url1)
	defer resp1.Body.Close()
	fmt.Printf("   Способ 1 (вручную): %s\n", url1)

	// Способ 2: через url.Values
	params := url.Values{}
	params.Add("name", "Bob")
	params.Add("age", "30")
	params.Add("hobby", "coding")
	params.Add("hobby", "music")

	url2 := "https://httpbin.org/get?" + params.Encode()
	fmt.Printf("   Способ 2 (url.Values): %s\n", url2)

	resp2, _ := http.Get(url2)
	defer resp2.Body.Close()

	body, _ := io.ReadAll(resp2.Body)
	fmt.Printf("\n   Ответ сервера:\n")
	fmt.Printf("   %s\n", string(body[:300]))
	fmt.Println()
}

// ========== POST ЗАПРОС С JSON ==========

// DemoPostJSON - POST с JSON
func DemoPostJSON() {
	fmt.Println("3️⃣  POST с JSON:")
	fmt.Println()

	// Создаём структуру
	data := map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
		"age":   25,
	}

	// Кодируем в JSON
	jsonData, _ := json.Marshal(data)

	// Создаём запрос
	req, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Отправляем
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Status: %s\n", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n   Ответ сервера:\n")
	fmt.Printf("   %s\n", string(body[:400]))
	fmt.Println()
}

// ========== POST С ФОРМОЙ ==========

// DemoPostForm - POST с формой
func DemoPostForm() {
	fmt.Println("4️⃣  POST с формой (application/x-www-form-urlencoded):")
	fmt.Println()

	// Данные формы
	formData := url.Values{}
	formData.Set("username", "john_doe")
	formData.Set("password", "secret123")
	formData.Set("email", "john@example.com")

	// Отправляем
	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Status: %s\n", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n   Ответ сервера:\n")
	fmt.Printf("   %s\n", string(body[:400]))
	fmt.Println()
}

// ========== КАСТОМНЫЙ HTTP CLIENT ==========

// DemoCustomClient - кастомный клиент
func DemoCustomClient() {
	fmt.Println("5️⃣  Кастомный HTTP Client:")
	fmt.Println()

	// Создаём клиент с настройками
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Делаем запрос
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Клиент с таймаутом 30с\n")
	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Println()
}

// ========== ЗАПРОС С КОНТЕКСТОМ ==========

// DemoWithContext - запрос с контекстом
func DemoWithContext() {
	fmt.Println("6️⃣  Запрос с контекстом (таймаут/отмена):")
	fmt.Println()

	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Создаём запрос с контекстом
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)

	client := &http.Client{}
	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Запрос выполнен за %v\n", time.Since(start))
	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Println()

	// Пример с отменой
	fmt.Println("   Пример отмены запроса:")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel2() // Отменяем через 100мс
	}()

	req2, _ := http.NewRequestWithContext(ctx2, "GET", "https://httpbin.org/delay/5", nil)
	_, err = client.Do(req2)
	if err != nil {
		fmt.Printf("   Запрос отменён: %v\n", err)
	}
	fmt.Println()
}

// ========== РАБОТА С ЗАГОЛОВКАМИ ==========

// DemoWithHeaders - работа с заголовками
func DemoWithHeaders() {
	fmt.Println("7️⃣  Работа с заголовками:")
	fmt.Println()

	req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)

	// Добавляем заголовки
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Add("X-Custom-Header", "value1")
	req.Header.Add("X-Custom-Header", "value2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Заголовки отправлены\n")
	fmt.Printf("   Status: %s\n", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n   Ответ сервера (заголовки):\n")
	fmt.Printf("   %s\n", string(body))
	fmt.Println()
}

// ========== ЧТЕНИЕ БОЛЬШИХ ОТВЕТОВ ==========

// DemoLargeResponse - чтение больших ответов
func DemoLargeResponse() {
	fmt.Println("8️⃣  Чтение больших ответов (потоковое):")
	fmt.Println()

	resp, _ := http.Get("https://httpbin.org/stream/5")
	defer resp.Body.Close()

	fmt.Printf("   Потоковое чтение (stream):\n")

	// Читаем по частям
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("   Ошибка: %v\n", err)
			break
		}
		fmt.Printf("   Получено: %d байт\n", len(line))
	}
	fmt.Println()
}

// ========== ОБРАБОТКА ОШИБОК ==========

// DemoErrorHandling - обработка ошибок
func DemoErrorHandling() {
	fmt.Println("9️⃣  Обработка ошибок:")
	fmt.Println()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Разные типы ошибок
	testCases := []struct {
		name string
		url  string
	}{
		{"Неверный URL", "http://invalid-host-that-does-not-exist.com"},
		{"Таймаут", "https://httpbin.org/delay/10"},
		{"404 Not Found", "https://httpbin.org/status/404"},
		{"500 Server Error", "https://httpbin.org/status/500"},
	}

	for _, tc := range testCases {
		fmt.Printf("   Тест: %s\n", tc.name)
		resp, err := client.Get(tc.url)
		if err != nil {
			fmt.Printf("   ❌ Ошибка запроса: %v\n", err)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode >= 400 {
				fmt.Printf("   ❌ HTTP ошибка: %s\n", resp.Status)
			} else {
				fmt.Printf("   ✅ Успех: %s\n", resp.Status)
			}
		}
	}
	fmt.Println()
}

// ========== ПРАКТИЧЕСКИЙ ПРИМЕР ==========

// GitHubUser - структура для GitHub API
type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	PublicRepos int  `json:"public_repos"`
	Followers int    `json:"followers"`
}

// DemoPracticalExample - практический пример
func DemoPracticalExample() {
	fmt.Println("🔟  Практический пример - GitHub API:")
	fmt.Println()

	// Запрос к GitHub API
	resp, err := http.Get("https://api.github.com/users/torvalds")
	if err != nil {
		fmt.Printf("   Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("   HTTP ошибка: %s\n", resp.Status)
		return
	}

	// Декодируем JSON
	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Printf("   Ошибка декодирования: %v\n", err)
		return
	}

	fmt.Printf("   Пользователь GitHub:\n")
	fmt.Printf("   Login: %s\n", user.Login)
	fmt.Printf("   Name: %s\n", user.Name)
	fmt.Printf("   Public Repos: %d\n", user.PublicRepos)
	fmt.Printf("   Followers: %d\n", user.Followers)
	fmt.Println()
}

// DemoHTTPClient - основная демонстрационная функция
func DemoHTTPClient() {
	PrintHeader("🌐 HTTP Client в Go")

	DemoBasicGet()
	DemoGetWithParams()
	DemoPostJSON()
	DemoPostForm()
	DemoCustomClient()
	DemoWithContext()
	DemoWithHeaders()
	DemoLargeResponse()
	DemoErrorHandling()
	DemoPracticalExample()

	PrintFooter()
}
