package test

import (
	"os"
	"strings"
	"testing"

	Gor "github.com/iwhitebird/Gor"
)

func runFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	result := <-Gor.RunFromInput(string(data))
	if result.Error != nil {
		t.Fatalf("%s runtime error: %v", path, result.Error)
	}
	return strings.TrimRight(result.Output, "\n")
}

func TestExampleStack(t *testing.T) {
	out := runFile(t, "../examples/stack.gor")
	if !strings.Contains(out, "Size after 3 pushes:") {
		t.Error("stack example failed")
	}
}

func TestExampleQueue(t *testing.T) {
	out := runFile(t, "../examples/queue.gor")
	if !strings.Contains(out, "Queue size:") {
		t.Error("queue example failed")
	}
}

func TestExampleLinkedList(t *testing.T) {
	out := runFile(t, "../examples/linked_list.gor")
	if !strings.Contains(out, "Sum:") && !strings.Contains(out, "15") {
		t.Error("linked list example failed")
	}
}

func TestExampleBinarySearch(t *testing.T) {
	out := runFile(t, "../examples/binary_search.gor")
	if !strings.Contains(out, "Searching for 23:") {
		t.Error("binary search example failed")
	}
}

func TestExampleBubbleSort(t *testing.T) {
	out := runFile(t, "../examples/bubble_sort.gor")
	if !strings.Contains(out, "After sort:") {
		t.Error("bubble sort example failed")
	}
}

func TestExampleSelectionSort(t *testing.T) {
	out := runFile(t, "../examples/selection_sort.gor")
	if !strings.Contains(out, "After sort:") {
		t.Error("selection sort example failed")
	}
}

func TestExampleHashMap(t *testing.T) {
	out := runFile(t, "../examples/hash_map.gor")
	if !strings.Contains(out, "Alice's score:") {
		t.Error("hash map example failed")
	}
}

func TestExampleFibonacci(t *testing.T) {
	out := runFile(t, "../examples/fibonacci.gor")
	if !strings.Contains(out, "Fibonacci (iterative):") {
		t.Error("fibonacci example failed")
	}
}

func TestExampleClosurePatterns(t *testing.T) {
	out := runFile(t, "../examples/closure_patterns.gor")
	if !strings.Contains(out, "Counter:") {
		t.Error("closure patterns example failed")
	}
}

func TestBenchmarkPrimeSieve(t *testing.T) {
	out := runFile(t, "../benchmark/prime_sieve.gor")
	// There are 1229 primes under 10000
	if !strings.Contains(out, "1229") {
		t.Errorf("prime sieve expected 1229, got: %s", out)
	}
}
