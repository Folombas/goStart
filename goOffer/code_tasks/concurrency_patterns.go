package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// ЗАДАЧА 1: Реализация Worker Pool
// ============================================

// Job - задача для выполнения
type Job func() error

// WorkerPool - пул воркеров
type WorkerPool struct {
	workers   int
	jobQueue  chan Job
	resultCh  chan error
	wg        sync.WaitGroup
}

// NewWorkerPool создаёт новый пул воркеров
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:  workers,
		jobQueue: make(chan Job, 100),
		resultCh: make(chan error, 100),
	}
}

// Start запускает воркеров
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go func(id int) {
			defer wp.wg.Done()
			for job := range wp.jobQueue {
				if err := job(); err != nil {
					wp.resultCh <- err
				}
			}
		}(i)
	}
}

// Submit отправляет задачу в пул
func (wp *WorkerPool) Submit(job Job) {
	wp.jobQueue <- job
}

// Stop закрывает пул
func (wp *WorkerPool) Stop() {
	close(wp.jobQueue)
	wp.wg.Wait()
	close(wp.resultCh)
}

// Results возвращает канал с результатами
func (wp *WorkerPool) Results() <-chan error {
	return wp.resultCh
}

// ============================================
// ЗАДАЧА 2: Rate Limiter (Token Bucket)
// ============================================

type RateLimiter struct {
	tokens     int
	maxTokens  int
	refillRate time.Duration
	lastRefill time.Time
	mu         sync.Mutex
}

// NewRateLimiter создаёт ограничитель скорости
func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow проверяет, можно ли выполнить запрос
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)

	// Добавляем токены
	tokensToAdd := int(elapsed / rl.refillRate)
	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		}
		rl.lastRefill = now
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

// ============================================
// ЗАДАЧА 3: Concurrent Cache с TTL
// ============================================

type CacheItem struct {
	value      interface{}
	expiration time.Time
}

type ConcurrentCache struct {
	data map[string]CacheItem
	mu   sync.RWMutex
}

// NewConcurrentCache создаёт новый кэш
func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		data: make(map[string]CacheItem),
	}
}

// Set добавляет значение в кэш с TTL
func (c *ConcurrentCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

// Get получает значение из кэша
func (c *ConcurrentCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.expiration) {
		return nil, false
	}

	return item.value, true
}

// Delete удаляет значение из кэша
func (c *ConcurrentCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// CleanExpired удаляет все истёкшие записи
func (c *ConcurrentCache) CleanExpired() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	count := 0
	now := time.Now()
	for key, item := range c.data {
		if now.After(item.expiration) {
			delete(c.data, key)
			count++
		}
	}
	return count
}

// ============================================
// ЗАДАЧА 4: Pub/Sub система
// ============================================

type Subscriber struct {
	id   int
	ch   chan interface{}
	done chan struct{}
}

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[int]*Subscriber
	nextID      int
}

// NewPubSub создаёт новую Pub/Sub систему
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[int]*Subscriber),
		nextID:      1,
	}
}

// Subscribe подписывается на сообщения
func (ps *PubSub) Subscribe() *Subscriber {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	sub := &Subscriber{
		id:   ps.nextID,
		ch:   make(chan interface{}, 10),
		done: make(chan struct{}),
	}
	ps.subscribers[sub.id] = sub
	ps.nextID++
	return sub
}

// Unsubscribe отписывается
func (ps *PubSub) Unsubscribe(sub *Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	delete(ps.subscribers, sub.id)
	close(sub.done)
	close(sub.ch)
}

// Publish отправляет сообщение всем подписчикам
func (ps *PubSub) Publish(msg interface{}) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, sub := range ps.subscribers {
		select {
		case sub.ch <- msg:
		case <-sub.done:
		default:
			// Буфер полон, пропускаем
		}
	}
}

// Receive получает сообщения
func (s *Subscriber) Receive() <-chan interface{} {
	return s.ch
}

// ============================================
// ЗАДАЧА 5: Semaphore (ограничитель параллелизма)
// ============================================

type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore создаёт семафор
func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, max),
	}
}

// Acquire захватывает разрешение
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release освобождает разрешение
func (s *Semaphore) Release() {
	<-s.ch
}

// TryAcquire пытается захватить без блокировки
func (s *Semaphore) TryAcquire() bool {
	select {
	case s.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

// ============================================
// ЗАДАЧА 6: Circuit Breaker
// ============================================

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	mu           sync.Mutex
	state        CircuitState
	failures     int
	successes    int
	maxFailures  int
	timeout      time.Duration
	lastFailure  time.Time
}

// NewCircuitBreaker создаёт circuit breaker
func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       StateClosed,
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

// Execute выполняет функцию с circuit breaker
func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	// Проверяем состояние
	if cb.state == StateOpen {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = StateHalfOpen
			cb.successes = 0
		} else {
			cb.mu.Unlock()
			return fmt.Errorf("circuit is open")
		}
	}

	cb.mu.Unlock()

	// Выполняем функцию
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
		}
		return err
	}

	// Успех
	if cb.state == StateHalfOpen {
		cb.successes++
		if cb.successes >= 3 {
			cb.state = StateClosed
			cb.failures = 0
		}
	}

	return nil
}

// State возвращает текущее состояние
func (cb *CircuitBreaker) State() CircuitState {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// ============================================
// ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ
// ============================================

func exampleWorkerPool() {
	fmt.Println("=== Worker Pool Example ===")

	pool := NewWorkerPool(3)
	pool.Start()

	// Отправляем задачи
	for i := 1; i <= 10; i++ {
		taskNum := i
		pool.Submit(func() error {
			fmt.Printf("Worker processing task %d\n", taskNum)
			time.Sleep(100 * time.Millisecond)
			return nil
		})
	}

	pool.Stop()
	fmt.Println("All tasks completed")
}

func exampleRateLimiter() {
	fmt.Println("\n=== Rate Limiter Example ===")

	limiter := NewRateLimiter(5, 100*time.Millisecond)

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: allowed\n", i)
		} else {
			fmt.Printf("Request %d: rate limited\n", i)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func exampleCache() {
	fmt.Println("\n=== Concurrent Cache Example ===")

	cache := NewConcurrentCache()

	// Set с TTL
	cache.Set("user:1", "Alice", 2*time.Second)
	cache.Set("user:2", "Bob", 5*time.Second)

	// Get
	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Got user: %v\n", val)
	}

	// Ждём истечения TTL
	time.Sleep(3 * time.Second)

	if _, ok := cache.Get("user:1"); !ok {
		fmt.Println("user:1 expired")
	}

	if val, ok := cache.Get("user:2"); ok {
		fmt.Printf("Got user: %v\n", val)
	}
}

func exampleSemaphore() {
	fmt.Println("\n=== Semaphore Example ===")

	sem := NewSemaphore(3) // Максимум 3 параллельных операции
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(taskNum int) {
			defer wg.Done()

			sem.Acquire()
			defer sem.Release()

			fmt.Printf("Task %d started\n", taskNum)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Task %d finished\n", taskNum)
		}(i)
	}

	wg.Wait()
}

func main() {
	fmt.Println("📚 Go Code Tasks - Solutions")
	fmt.Println("================================")

	exampleWorkerPool()
	exampleRateLimiter()
	exampleCache()
	exampleSemaphore()

	fmt.Println("\n✅ All examples completed!")
}
