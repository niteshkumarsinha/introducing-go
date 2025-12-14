# Testing in Go (Golang)

Go provides **first-class support for testing** through its standard library. Testing in Go is **simple, fast, and idiomatic**, encouraging developers to write tests alongside production code.

---

# 1. Why Testing in Go?

Testing helps you:

* Verify correctness of logic
* Prevent regressions
* Improve code quality
* Refactor safely
* Build reliable production systems

Go follows a **code + test together** philosophy.

---

# 2. Go Testing Basics

Go testing is provided by the built-in **`testing`** package.

### Key Features

* No external libraries required
* Convention-based
* Integrated with `go test` command

---

# 3. Test File Naming Convention

| Rule                           | Example                  |
| ------------------------------ | ------------------------ |
| File name ends with `_test.go` | `math_test.go`           |
| Package name same as source    | `package mathutils`      |
| Or black-box testing           | `package mathutils_test` |

---

# 4. Writing Your First Test

### Production Code (`math.go`)

```go
package mathutils

func Add(a int, b int) int {
    return a + b
}
```

### Test Code (`math_test.go`)

```go
package mathutils

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}
```

---

# 5. Running Tests

```bash
go test
```

### Run with Verbose Output

```bash
go test -v
```

---

# 6. Test Function Rules

A test function must:

* Start with `Test`
* Accept `*testing.T`
* Be in a `_test.go` file

```go
func TestSomething(t *testing.T) { }
```

---

# 7. Using `t.Errorf` vs `t.Fatalf`

| Method     | Behavior                   |
| ---------- | -------------------------- |
| `t.Errorf` | Logs error, continues test |
| `t.Fatalf` | Logs error, stops test     |

```go
if result < 0 {
    t.Fatalf("invalid result")
}
```

---

# 8. Table-Driven Tests (Very Important)

Most idiomatic Go testing style.

```go
func TestAdd_TableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"both positive", 2, 3, 5},
        {"with zero", 0, 5, 5},
        {"both negative", -2, -3, -5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("expected %d, got %d", tt.expected, result)
            }
        })
    }
}
```

---

# 9. Subtests (`t.Run`)

Subtests help:

* Isolate failures
* Improve readability
* Run specific cases

```go
t.Run("case name", func(t *testing.T) {
    // test logic
})
```

---

# 10. Testing for Errors

```go
func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

```go
func TestDivide_Error(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Errorf("expected error, got nil")
    }
}
```

---

# 11. Skipping Tests

```go
func TestSkip(t *testing.T) {
    t.Skip("Skipping this test temporarily")
}
```

---

# 12. Running Specific Tests

```bash
go test -run TestAdd
```

```bash
go test -run TestAdd_TableDriven
```

---

# 13. Test Coverage

### Generate Coverage

```bash
go test -cover
```

### Coverage Report

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

# 14. Benchmark Testing

Benchmarks measure performance.

### Benchmark Function Rules

* Start with `Benchmark`
* Use `*testing.B`

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(10, 20)
    }
}
```

### Run Benchmarks

```bash
go test -bench=.
```

---

# 15. Setup and Teardown

Go doesn’t have annotations, but you can simulate setup/teardown.

```go
func TestMain(m *testing.M) {
    // setup
    code := m.Run()
    // teardown
    os.Exit(code)
}
```

---

# 16. Black-Box Testing

Test only exported APIs.

```go
package mathutils_test

import (
    "testing"
    "myapp/mathutils"
)
```

---

# 17. Mocking in Go (Concept)

Go prefers **interfaces** for mocking.

```go
type Store interface {
    Save(data string) error
}
```

```go
type MockStore struct {}

func (m MockStore) Save(data string) error {
    return nil
}
```

---

# 18. Common Testing Commands Cheat Sheet

```bash
go test
go test -v
go test -run TestName
go test -cover
go test -bench=.
```

---

# 19. Best Practices for Testing in Go

✅ Write tests alongside code
✅ Prefer table-driven tests
✅ Keep tests simple and readable
✅ Test behavior, not implementation
✅ Use interfaces for testability

---

# 20. Summary

| Feature                | Supported |
| ---------------------- | --------- |
| Unit testing           | ✅         |
| Table-driven tests     | ✅         |
| Subtests               | ✅         |
| Benchmarks             | ✅         |
| Coverage               | ✅         |
| Mocking via interfaces | ✅         |

---

📌 **Key Takeaway:**
Go’s testing philosophy emphasizes **simplicity**, **clarity**, and **speed**, making testing an integral part of everyday Go development.

---

Happy Testing with Go 🧪🚀
