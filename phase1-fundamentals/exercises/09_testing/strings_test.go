package strings

import "testing"

// TODO(phase1-09-test): Write comprehensive tests.
//
// Task:
//   1. Table-driven tests with t.Run subtests.
//   2. Add a benchmark: func BenchmarkReverse(b *testing.B).
//   3. Add an Example function for godoc-style output testing.
//
// Verify: go test ./exercises/09_testing -v -bench=.

func TestReverse(t *testing.T) {
	t.Skip("implement table-driven tests and remove this skip")
}

func BenchmarkReverse(b *testing.B) {
	b.Skip("implement benchmark and remove this skip")
}
