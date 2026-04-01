package channels

import (
	"fmt"
	"sync"
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

// ========== БАЗОВЫЕ КАНАЛЫ ==========

// DemoBasicChannels - демонстрация базовых каналов
func DemoBasicChannels() {
	fmt.Println("1️⃣  Базовые каналы:")
	fmt.Println()

	// Создаём небуферизированный канал
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	// Горутина-отправитель
	go func() {
		defer wg.Done()
		fmt.Println("   Отправляем сообщение в канал...")
		ch <- "Привет из горутины!"
		fmt.Println("   Сообщение отправлено")
	}()

	// Получаем сообщение
	fmt.Println("   Получаем сообщение из канала...")
	msg := <-ch
	fmt.Printf("   Получено: %s\n", msg)

	wg.Wait()
	fmt.Println()
}

// ========== БУФЕРИЗИРОВАННЫЕ КАНАЛЫ ==========

// DemoBufferedChannels - демонстрация буферизированных каналов
func DemoBufferedChannels() {
	fmt.Println("2️⃣  Буферизированные каналы:")
	fmt.Println()

	// Создаём канал с буфером на 3 элемента
	ch := make(chan int, 3)

	fmt.Println("   Заполняем буфер...")
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("   Отправлено 3 значения, буфер заполнен\n")

	// Читаем из канала
	fmt.Println("   Читаем из канала:")
	fmt.Printf("   %d\n", <-ch)
	fmt.Printf("   %d\n", <-ch)
	fmt.Printf("   %d\n", <-ch)

	// Попытка отправить в полный канал (закомментировано, чтобы не блокировать)
	// ch <- 4 // Это заблокировало бы выполнение

	fmt.Println()
}

// ========== НАПРАВЛЕНИЕ КАНАЛОВ ==========

// sendValues - функция только отправляет в канал
func sendValues(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
		fmt.Printf("   Отправлено: %d\n", i)
	}
	close(ch)
}

// receiveValues - функция только получает из канала
func receiveValues(ch <-chan int) {
	for val := range ch {
		fmt.Printf("   Получено: %d\n", val)
	}
}

// DemoChannelDirections - демонстрация направления каналов
func DemoChannelDirections() {
	fmt.Println("3️⃣  Направление каналов:")
	fmt.Println()

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sendValues(ch, 3)
	}()

	go func() {
		defer wg.Done()
		receiveValues(ch)
	}()

	wg.Wait()
	fmt.Println()
}

// ========== SELECT ==========

// DemoSelect - демонстрация select
func DemoSelect() {
	fmt.Println("4️⃣  Оператор select:")
	fmt.Println()

	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	// Отправляем в разные каналы с разной задержкой
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Сообщение из канала 1"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Сообщение из канала 2"
	}()

	// Ждём оба сообщения
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("   %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("   %s\n", msg2)
		}
	}

	wg.Wait()
	fmt.Println()
}

// ========== SELECT С DEFAULT ==========

// DemoSelectWithDefault - select с default
func DemoSelectWithDefault() {
	fmt.Println("5️⃣  Select с default (неблокирующая операция):")
	fmt.Println()

	ch := make(chan int)

	// Попытка получить без блокировки
	select {
	case val := <-ch:
		fmt.Printf("   Получено: %d\n", val)
	default:
		fmt.Println("   Канал пуст, выполняем другую работу...")
	}

	// Попытка отправить без блокировки
	select {
	case ch <- 42:
		fmt.Println("   Отправлено в канал")
	default:
		fmt.Println("   Канал полон, не отправляем")
	}

	fmt.Println()
}

// ========== ЗАКРЫТИЕ КАНАЛА ==========

// DemoCloseChannel - демонстрация закрытия канала
func DemoCloseChannel() {
	fmt.Println("6️⃣  Закрытие канала:")
	fmt.Println()

	ch := make(chan int, 5)

	// Отправляем значения
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	// Читаем до закрытия
	fmt.Println("   Чтение из закрытого канала:")
	for val, ok := <-ch; ok; val, ok = <-ch {
		fmt.Printf("   Получено: %d (ok=%v)\n", val, ok)
	}

	// Чтение из закрытого канала возвращяет zero value
	val, ok := <-ch
	fmt.Printf("   После закрытия: val=%d, ok=%v\n", val, ok)
	fmt.Println()
}

// ========== WORKER POOL ==========

// worker - воркер обрабатывает задачи
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("   Worker %d обрабатывает задачу %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

// DemoWorkerPool - демонстрация пула воркеров
func DemoWorkerPool() {
	fmt.Println("7️⃣  Worker Pool (пул воркеров):")
	fmt.Println()

	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Запускаем воркеров
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(w)
	}

	// Отправляем задачи
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Собираем результаты
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("   Результаты:")
	for result := range results {
		fmt.Printf("   Результат: %d\n", result)
	}

	fmt.Println()
}

// ========== PIPELINE ==========

// generator - генерирует числа
func generator(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// doubler - удваивает числа
func doubler(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}

// printer - печатает результаты
func printer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("   %d\n", n)
	}
}

// DemoPipeline - демонстрация конвейера
func DemoPipeline() {
	fmt.Println("8️⃣  Pipeline (конвейер):")
	fmt.Println()

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		generator(ch1)
	}()

	go func() {
		defer wg.Done()
		doubler(ch1, ch2)
	}()

	go func() {
		defer wg.Done()
		printer(ch2)
	}()

	wg.Wait()
	fmt.Println()
}

// ========== RACE CONDITION ==========

// DemoRaceCondition - демонстрация гонки данных
func DemoRaceCondition() {
	fmt.Println("9️⃣  Race Condition и Mutex:")
	fmt.Println()

	// ❌ Плохо: гонка данных
	fmt.Println("   ❌ Гонка данных (без синхронизации):")
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			temp := counter
			temp++
			counter = temp
		}()
	}
	wg.Wait()
	fmt.Printf("   Результат без синхронизации: %d (ожидалось 100)\n", counter)
	fmt.Println()

	// ✅ Хорошо: с Mutex
	fmt.Println("   ✅ С Mutex (безопасно):")
	var mu sync.Mutex
	counterSafe := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			temp := counterSafe
			temp++
			counterSafe = temp
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("   Результат с Mutex: %d (ожидалось 100)\n", counterSafe)
	fmt.Println()
}

// DemoChannels - основная демонстрационная функция
func DemoChannels() {
	PrintHeader("📡 Каналы и Горутины в Go")

	DemoBasicChannels()
	DemoBufferedChannels()
	DemoChannelDirections()
	DemoSelect()
	DemoSelectWithDefault()
	DemoCloseChannel()
	DemoWorkerPool()
	DemoPipeline()
	DemoRaceCondition()

	PrintFooter()
}
