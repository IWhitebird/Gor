package test

import (
	"strings"
	"testing"

	Gor "github.com/iwhitebird/Gor"
)

// run executes Gor source code and returns trimmed stdout output.
func run(t *testing.T, code string) string {
	t.Helper()
	result := <-Gor.RunFromInput(code)
	if result.Error != nil {
		t.Fatalf("runtime error: %v", result.Error)
	}
	return strings.TrimRight(result.Output, "\n")
}

// expectOutput runs code and asserts the output matches expected.
func expectOutput(t *testing.T, code string, expected string) {
	t.Helper()
	got := run(t, code)
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// expectLines runs code and asserts each line of output.
func expectLines(t *testing.T, code string, lines ...string) {
	t.Helper()
	expected := strings.Join(lines, "\n")
	expectOutput(t, code, expected)
}
