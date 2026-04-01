package context_pkg

import (
	"context"
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

// ========== БАЗОВЫЙ CONTEXT ==========

// DemoBasicContext - демонстрация базового контекста с отменой
func DemoBasicContext() {
	fmt.Println("1️⃣  Базовый контекст с отменой:")
	fmt.Println()

	// Создаём контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Обязательно вызываем cancel для очистки ресурсов

	var wg sync.WaitGroup

	// Запускаем горутину, которая работает до отмены контекста
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("     Горутина получила сигнал отмены!")
				return
			default:
				counter++
				fmt.Printf("     Работа... счётчик: %d\n", counter)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Ждём 1 секунду и отменяем контекст
	time.Sleep(1 * time.Second)
	fmt.Println("   Отменяем контекст...")
	cancel()

	// Ждём завершения горутины
	wg.Wait()
	fmt.Println("   Горутина завершена")
	fmt.Println()
}

// ========== CONTEXT С ТАЙМАУТОМ ==========

// DemoContextWithTimeout - демонстрация контекста с таймаутом
func DemoContextWithTimeout() {
	fmt.Println("2️⃣  Контекст с таймаутом:")
	fmt.Println()

	// Контекст автоматически отменится через 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	start := time.Now()

	// Имитация долгой операции
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		
		counter := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("     Операция прервана по таймауту (прошло %v)\n", time.Since(start))
				return
			default:
				counter++
				fmt.Printf("     Обработка... %d\n", counter)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	<-done
	fmt.Println()
}

// ========== CONTEXT С DEADLINE ==========

// DemoContextWithDeadline - демонстрация контекста с дедлайном
func DemoContextWithDeadline() {
	fmt.Println("3️⃣  Контекст с дедлайном:")
	fmt.Println()

	// Устанавливаем конкретное время отмены
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("   Дедлайн установлен на: %v\n", deadline.Format("15:04:05"))
	fmt.Printf("   Текущее время: %v\n", time.Now().Format("15:04:05"))
	fmt.Println()

	start := time.Now()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("     Контекст истёк! (прошло %v)\n", time.Since(start))
			fmt.Printf("     Причина: %v\n", ctx.Err())
			return
		default:
			fmt.Printf("     Работа... %v\n", time.Since(start).Round(time.Millisecond))
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ========== ПЕРЕДАЧА ЗНАЧЕНИЙ В CONTEXT ==========

// DemoContextWithValue - демонстрация передачи значений через контекст
func DemoContextWithValue() {
	fmt.Println("4️⃣  Передача значений через контекст:")
	fmt.Println()

	// Создаём цепочку контекстов с значениями
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", 123)
	ctx = context.WithValue(ctx, "username", "alice")
	ctx = context.WithValue(ctx, "request_id", "req-abc-123")

	// Получаем значения из контекста
	userID := ctx.Value("user_id")
	username := ctx.Value("username")
	requestID := ctx.Value("request_id")

	fmt.Printf("   User ID: %v\n", userID)
	fmt.Printf("   Username: %v\n", username)
	fmt.Printf("   Request ID: %v\n", requestID)
	fmt.Println()

	// Важное замечание: не используйте контекст для передачи данных между горутинами!
	fmt.Println("   ⚠️  Контекст не предназначен для передачи данных между горутинами!")
	fmt.Println("      Используйте для этого каналы или другие механизмы синхронизации.")
	fmt.Println()
}

// ========== ЦЕПОЧКА ОТМЕНЫ CONTEXT ==========

// DemoContextChain - демонстрация цепочки отмены контекстов
func DemoContextChain() {
	fmt.Println("5️⃣  Цепочка отмены контекстов:")
	fmt.Println()

	// Родительский контекст
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	// Дочерние контексты
	child1Ctx, child1Cancel := context.WithCancel(parentCtx)
	defer child1Cancel()

	child2Ctx, child2Cancel := context.WithCancel(parentCtx)
	defer child2Cancel()

	fmt.Println("   Создали родительский и два дочерних контекста")
	fmt.Println()

	var wg sync.WaitGroup

	// Горутина 1 - слушает дочерний контекст 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-child1Ctx.Done():
			fmt.Println("     Горутина 1: дочерний контекст 1 отменён")
		}
	}()

	// Горутина 2 - слушает дочерний контекст 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-child2Ctx.Done():
			fmt.Println("     Горутина 2: дочерний контекст 2 отменён")
		}
	}()

	// Горутина 3 - слушает родительский контекст
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-parentCtx.Done():
			fmt.Println("     Горутина 3: родительский контекст отменён")
		}
	}()

	// Отменяем один из дочерних контекстов
	fmt.Println("   Отменяем дочерний контекст 1...")
	child1Cancel()
	time.Sleep(100 * time.Millisecond)

	// Отменяем родительский контекст - это отменит все оставшиеся
	fmt.Println("   Отменяем родительский контекст...")
	parentCancel()

	wg.Wait()
	fmt.Println("   Все горутины завершены")
	fmt.Println()
}

// ========== ПРАКТИЧЕСКИЙ ПРИМЕР ==========

// Worker - имитация рабочего, который выполняет задачу
type Worker struct {
	id int
}

// DoWork - выполняет работу с поддержкой отмены
func (w *Worker) DoWork(ctx context.Context, duration time.Duration) error {
	fmt.Printf("     Worker %d: начал работу\n", w.id)

	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("     Worker %d: работа завершена успешно\n", w.id)
		return nil
	case <-ctx.Done():
		fmt.Printf("     Worker %d: работа отменена (%v)\n", w.id, ctx.Err())
		return ctx.Err()
	}
}

// DemoPracticalExample - практический пример использования контекста
func DemoPracticalExample() {
	fmt.Println("6️⃣  Практический пример - управление воркерами:")
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	workers := []*Worker{{id: 1}, {id: 2}, {id: 3}}

	for _, worker := range workers {
		wg.Add(1)
		go func(w *Worker) {
			defer wg.Done()
			// Имитируем работу разной длительности
			duration := time.Duration(w.id+1) * time.Second
			err := w.DoWork(ctx, duration)
			if err != nil {
				// Работа была отменена
				return
			}
		}(worker)
	}

	wg.Wait()
	fmt.Println("   Все воркеры завершили работу")
	fmt.Println()
}

// DemoContext - основная демонстрационная функция
func DemoContext() {
	PrintHeader("🔄 Context в Go - управление временем и отменой")

	DemoBasicContext()
	DemoContextWithTimeout()
	DemoContextWithDeadline()
	DemoContextWithValue()
	DemoContextChain()
	DemoPracticalExample()

	PrintFooter()
}
