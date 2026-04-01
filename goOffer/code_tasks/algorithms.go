package main

import (
	"fmt"
	"sort"
)

// ============================================
// ЗАДАЧА 1: Two Sum
// ============================================

// TwoSum находит два числа, которые в сумме дают target
// Время: O(n), Память: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		need := target - num
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

// ============================================
// ЗАДАЧА 2: Valid Parentheses
// ============================================

// IsValidParentheses проверяет правильность скобочной последовательности
func IsValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if opening, ok := pairs[ch]; ok {
			// Закрывающая скобка
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// Открывающая скобка
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

// ============================================
// ЗАДАЧА 3: Merge Intervals
// ============================================

// Interval представляет интервал
type Interval struct {
	Start, End int
}

// MergeIntervals объединяет пересекающиеся интервалы
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Сортируем по началу интервала
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	result := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := &result[len(result)-1]
		curr := intervals[i]

		if curr.Start <= last.End {
			// Пересекаются - объединяем
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			// Не пересекаются - добавляем новый
			result = append(result, curr)
		}
	}

	return result
}

// ============================================
// ЗАДАЧА 4: Group Anagrams
// ============================================

// GroupAnagrams группирует анаграммы
func GroupAnagrams(words []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		// Ключ - отсортированные буквы
		key := sortString(word)
		anagrams[key] = append(anagrams[key], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// ============================================
// ЗАДАЧА 5: Longest Substring Without Repeating
// ============================================

// LengthOfLongestSubstring находит длину самой длинной подстроки без повторений
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := make(map[rune]int)
	maxLen := 0
	left := 0

	for right, ch := range s {
		if lastPos, ok := seen[ch]; ok && lastPos >= left {
			left = lastPos + 1
		}
		seen[ch] = right
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ============================================
// ЗАДАЧА 6: Product of Array Except Self
// ============================================

// ProductExceptSelf возвращает массив где каждый элемент - произведение всех остальных
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Левые произведения
	left := 1
	for i := 0; i < n; i++ {
		result[i] = left
		left *= nums[i]
	}

	// Правые произведения
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

// ============================================
// ЗАДАЧА 7: Container With Most Water
// ============================================

// MaxArea находит максимальную площадь контейнера
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		h := min(height[left], height[right])
		w := right - left
		area := h * w
		maxArea = max(maxArea, area)

		// Двигаем меньшую границу
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// ============================================
// ЗАДАЧА 8: 3Sum
// ============================================

// ThreeSum находит все тройки с суммой 0
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// Пропускаем дубликаты
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Пропускаем дубликаты
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// ============================================
// ЗАДАЧА 9: Reverse Linked List (имитация)
// ============================================

// ListNode - узел связного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList разворачивает связный список
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// ============================================
// ЗАДАЧА 10: Binary Search
// ============================================

// BinarySearch ищет элемент в отсортированном массиве
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// ============================================
// ЗАДАЧА 11: Implement Trie (Prefix Tree)
// ============================================

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

// NewTrie создаёт новое дерево
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert вставляет слово в дерево
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search ищет слово в дереве
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith проверяет префикс
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// ============================================
// ЗАДАЧА 12: Sliding Window Maximum
// ============================================

// MaxSlidingWindow находит максимум в скользящем окне
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	result := []int{}
	window := []int{} // Индексы элементов в окне

	for i, num := range nums {
		// Удаляем элементы вне окна
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}

		// Удаляем элементы меньше текущего
		for len(window) > 0 && nums[window[len(window)-1]] < num {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		// Добавляем максимум окна в результат
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}

	return result
}

// ============================================
// ТЕСТИРОВАНИЕ
// ============================================

func testTwoSum() {
	fmt.Println("\n=== Two Sum ===")
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}

	for _, tt := range tests {
		result := TwoSum(tt.nums, tt.target)
		fmt.Printf("TwoSum(%v, %d) = %v\n", tt.nums, tt.target, result)
	}
}

func testParentheses() {
	fmt.Println("\n=== Valid Parentheses ===")
	tests := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}

	for _, tt := range tests {
		result := IsValidParentheses(tt)
		fmt.Printf("IsValidParentheses(%q) = %v\n", tt, result)
	}
}

func testAnagrams() {
	fmt.Println("\n=== Group Anagrams ===")
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := GroupAnagrams(words)
	fmt.Printf("GroupAnagrams(%v) = %v\n", words, result)
}

func testLongestSubstring() {
	fmt.Println("\n=== Longest Substring ===")
	tests := []string{"abcabcbb", "bbbbb", "pwwkew", ""}

	for _, tt := range tests {
		result := LengthOfLongestSubstring(tt)
		fmt.Printf("LengthOfLongestSubstring(%q) = %d\n", tt, result)
	}
}

func testProductExceptSelf() {
	fmt.Println("\n=== Product Except Self ===")
	tests := [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
	}

	for _, tt := range tests {
		result := ProductExceptSelf(tt)
		fmt.Printf("ProductExceptSelf(%v) = %v\n", tt, result)
	}
}

func testBinarySearch() {
	fmt.Println("\n=== Binary Search ===")
	nums := []int{1, 3, 5, 7, 9, 11, 13}
	targets := []int{7, 1, 13, 4, 0}

	for _, target := range targets {
		result := BinarySearch(nums, target)
		fmt.Printf("BinarySearch(%v, %d) = %d\n", nums, target, result)
	}
}

func testTrie() {
	fmt.Println("\n=== Trie ===")
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("application")

	fmt.Printf("Search 'apple': %v\n", trie.Search("apple"))
	fmt.Printf("Search 'app': %v\n", trie.Search("app"))
	fmt.Printf("Search 'apply': %v\n", trie.Search("apply"))
	fmt.Printf("StartsWith 'app': %v\n", trie.StartsWith("app"))
	fmt.Printf("StartsWith 'appl': %v\n", trie.StartsWith("appl"))
}

func main() {
	fmt.Println("📚 Go Algorithm Tasks")
	fmt.Println("======================")

	testTwoSum()
	testParentheses()
	testAnagrams()
	testLongestSubstring()
	testProductExceptSelf()
	testBinarySearch()
	testTrie()

	fmt.Println("\n✅ All tests completed!")
}
