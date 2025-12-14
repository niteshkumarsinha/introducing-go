# Functions in Go (Golang)

Functions are the building blocks of Go programs. They help you **organize code**, **avoid repetition**, and **improve readability and reusability**.

---

## 1. What is a Function?

A **function** is a named block of code that:

* Performs a specific task
* Can take **input parameters**
* Can return **output values**

In Go, functions are **first-class citizens**, meaning they can be:

* Assigned to variables
* Passed as arguments
* Returned from other functions

---

## 2. Basic Function Syntax

```go
func functionName(parameter1 type, parameter2 type) returnType {
    // function body
    return value
}
```

### Example

```go
func add(a int, b int) int {
    return a + b
}
```

---

## 3. Calling a Function

```go
result := add(10, 20)
fmt.Println(result) // Output: 30
```

---

## 4. Functions with No Parameters

```go
func greet() {
    fmt.Println("Hello, Go!")
}
```

### Usage

```go
greet()
```

---

## 5. Functions with Parameters

```go
func greetUser(name string) {
    fmt.Println("Hello", name)
}
```

```go
greetUser("Nitesh")
```

---

## 6. Functions with Return Values

```go
func square(n int) int {
    return n * n
}
```

```go
fmt.Println(square(5)) // Output: 25
```

---

## 7. Multiple Return Values (Very Important in Go)

Go functions can return **multiple values**, commonly used for **error handling**.

```go
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### Usage

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(result)
}
```

---

## 8. Named Return Values

You can name return variables in the function signature.

```go
func rectangleArea(length int, width int) (area int) {
    area = length * width
    return
}
```

```go
fmt.Println(rectangleArea(10, 5)) // Output: 50
```

---

## 9. Variadic Functions (Variable Number of Arguments)

Use `...` to accept any number of arguments.

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

```go
fmt.Println(sum(1, 2, 3, 4)) // Output: 10
```

---

## 10. Anonymous Functions (Lambda Functions)

Functions without names.

```go
func() {
    fmt.Println("Anonymous function")
}()
```

### Assigned to a Variable

```go
add := func(a int, b int) int {
    return a + b
}

fmt.Println(add(3, 4)) // Output: 7
```

---

## 11. Functions as Parameters (Callback Functions)

```go
func operate(a int, b int, operation func(int, int) int) int {
    return operation(a, b)
}
```

```go
result := operate(5, 3, func(x int, y int) int {
    return x * y
})

fmt.Println(result) // Output: 15
```

---

## 12. Functions Returning Functions

```go
func multiplier(factor int) func(int) int {
    return func(value int) int {
        return value * factor
    }
}
```

```go
double := multiplier(2)
fmt.Println(double(5)) // Output: 10
```

---

## 13. Recursive Functions

A function that calls itself.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

```go
fmt.Println(factorial(5)) // Output: 120
```

---

## 14. Defer with Functions

`defer` delays function execution until the surrounding function returns.

```go
func example() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

### Output

```
Hello
World
```

---

## 15. Function Visibility (Exported vs Unexported)

* **Capital letter** → Exported (public)
* **Small letter** → Unexported (private)

```go
func PrintMessage() { } // Exported
func printMessage() { } // Unexported
```

---

## 16. Best Practices for Functions in Go

✅ Keep functions **small and focused**
✅ Prefer **multiple return values** over global variables
✅ Use **clear and meaningful names**
✅ Handle errors explicitly
✅ Avoid unnecessary side effects

---

## 17. Summary

| Feature                      | Supported |
| ---------------------------- | --------- |
| Multiple return values       | ✅         |
| Anonymous functions          | ✅         |
| Variadic functions           | ✅         |
| Functions as values          | ✅         |
| Recursion                    | ✅         |
| Error handling via functions | ✅         |

---

📌 **Key Takeaway:**
Functions in Go are powerful, simple, and designed to encourage **clean, readable, and maintainable code**.

---

Happy Coding with Go 🚀
