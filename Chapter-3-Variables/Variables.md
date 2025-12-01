# Variables in Go

## Introduction to Variables

A variable is a named storage location that holds a value of a specific data type. In Go, variables are used to store data that can be used and modified throughout the program. Go is a **statically-typed language**, which means the data type of a variable must be known at compile time.

---

## Variable Declaration

### 1. **Using `var` Keyword with Explicit Type**

The most common way to declare a variable is using the `var` keyword with an explicit data type.

```go
var variableName dataType = value
```

**Example:**
```go
var x string = "Hello, World!"
var age int = 25
var price float64 = 99.99
var isActive bool = true
```

### 2. **Using `var` with Type Inference**

Go can infer the data type from the assigned value, so you don't always need to specify the type explicitly.

```go
var variableName = value
```

**Example:**
```go
var name = "John"        // Inferred as string
var count = 42           // Inferred as int
var pi = 3.14            // Inferred as float64
```

### 3. **Short Declaration (`:=` Operator)**

Inside functions, you can use the short declaration operator `:=` to declare and initialize variables in one step. This is the most concise way and is commonly used in Go.

```go
variableName := value
```

**Example:**
```go
func main() {
    name := "Alice"          // string
    age := 30                // int
    salary := 50000.50       // float64
    isEmployed := true       // bool
}
```

**Note:** The `:=` operator can only be used inside functions, not at the package level.

### 4. **Using `const` for Constants**

Constants are variables that cannot be changed after initialization. They are declared using the `const` keyword.

```go
const constantName dataType = value
```

**Example:**
```go
const Pi = 3.14159
const MaxUsers int = 100
const DefaultTimeout = "30s"
```

---

## Global vs Local Variables

### **Global Variables**

Global variables are declared at the package level (outside of any function) and can be accessed from any function in the same package.

```go
package main

import "fmt"

var x string = "Global Variable"  // Global variable

func main() {
    fmt.Println(x)  // Can access global variable
}

func f() {
    fmt.Println(x)  // Can also access global variable
}
```

### **Local Variables**

Local variables are declared inside a function and can only be accessed within that function.

```go
func main() {
    var x string = "Local Variable"  // Local variable
    fmt.Println(x)                   // Can access
}

// fmt.Println(x)  // Error: x is not defined here
```

---

## Multiple Variable Declaration

### **Using `var` with Block Declaration**

You can declare multiple variables in a single block using parentheses:

```go
var (
    a = 5
    b = 10
    c = 15
)
```

This is equivalent to:
```go
var a = 5
var b = 10
var c = 15
```

### **Multiple Declaration in One Line**

```go
var x, y, z = 1, 2, 3
var name, age string = "John", "25"  // Both have the same type
```

### **Short Declaration of Multiple Variables**

```go
x, y, z := 10, 20, 30
name, age := "Alice", 25
```

---

## Variable Scope

Variable scope refers to the region of the code where a variable can be accessed.

### **Block Scope**

Variables declared in a block (inside `{}`) can only be accessed within that block and nested blocks.

```go
func main() {
    x := 10
    {
        y := 20
        fmt.Println(x)  // OK
        fmt.Println(y)  // OK
    }
    fmt.Println(x)      // OK
    // fmt.Println(y)   // Error: y is not defined
}
```

### **Function Scope**

Variables declared in a function are accessible throughout the entire function.

```go
func example() {
    if true {
        x := 5
    }
    fmt.Println(x)  // OK: x is still in scope
}
```

### **Package Scope**

Global variables declared at the package level are accessible from all functions in the same package.

```go
package main

var globalVar = "I'm global"

func function1() {
    fmt.Println(globalVar)  // OK
}

func function2() {
    fmt.Println(globalVar)  // OK
}
```

---

## Variable Naming Conventions

Go follows specific conventions for naming variables:

1. **Camel Case (camelCase):** Use camel case for variable names (first letter lowercase, then capitalize the first letter of subsequent words)
   ```go
   var firstName string
   var lastName string
   var totalPrice float64
   ```

2. **Meaningful Names:** Use names that clearly indicate the variable's purpose
   ```go
   var age int              // Good
   var a int                // Poor
   var userAge int          // Good
   ```

3. **Exported vs Unexported:** Variables starting with uppercase letters are exported (accessible outside the package), while lowercase letters are unexported (private to the package)
   ```go
   var PublicVar = 10       // Exported
   var privateVar = 20      // Unexported
   ```

---

## Practical Examples

### **Example 1: Basic Variable Usage**

```go
package main

import "fmt"

func main() {
    var x string = "Hello, World!"
    x = "First Go Program"
    fmt.Println(x)  // Output: First Go Program
    
    x = "Hello, Go!"
    fmt.Println(x)  // Output: Hello, Go!
}
```

### **Example 2: Type Inference with Short Declaration**

```go
func main() {
    y := 5
    y += 1
    fmt.Println(y)  // Output: 6
}
```

### **Example 3: Working with Functions**

```go
func fahrenheitToCelsius(f float64) float64 {
    return (f - 32) * 5 / 9
}

func celsiusToFahrenheit(c float64) float64 {
    return (c * 9 / 5) + 32
}

func main() {
    temp := 98.6
    celsius := fahrenheitToCelsius(temp)
    fmt.Println(celsius)  // Output: 37
    
    temp2 := 37.0
    fahrenheit := celsiusToFahrenheit(temp2)
    fmt.Println(fahrenheit)  // Output: 98.6
}
```

### **Example 4: Unit Conversion Functions**

```go
func feetIntoMeters(feet float64) float64 {
    return feet * 0.3048
}

func metersIntoFeet(meters float64) float64 {
    return meters / 0.3048
}

func main() {
    height := 5.5
    heightInMeters := feetIntoMeters(height)
    fmt.Println(heightInMeters)  // Output: 1.6764
    
    length := 1.6764
    lengthInFeet := metersIntoFeet(length)
    fmt.Println(lengthInFeet)  // Output: 5.5
}
```

### **Example 5: Reading User Input**

```go
func getUserInput() {
    fmt.Println("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println("Output:", output)
}

func main() {
    getUserInput()
}
```

---

## Variable Assignment and Modification

### **Simple Assignment**

```go
x := 10
x = 20      // Reassign a new value
fmt.Println(x)  // Output: 20
```

### **Compound Assignment Operators**

Go supports various compound assignment operators that combine an operation with assignment:

| Operator | Example | Equivalent |
|----------|---------|-----------|
| `+=` | `x += 5` | `x = x + 5` |
| `-=` | `x -= 3` | `x = x - 3` |
| `*=` | `x *= 2` | `x = x * 2` |
| `/=` | `x /= 4` | `x = x / 4` |
| `%=` | `x %= 3` | `x = x % 3` |

**Example:**
```go
y := 5
y += 1
fmt.Println(y)  // Output: 6
```

---

## Blank Identifier `_`

The blank identifier `_` is used when you want to ignore a returned value. It's a special variable that accepts any value without storing it.

```go
result, _ := strconv.Atoi("123")  // Ignore the error
fmt.Println(result)  // Output: 123
```

---

## Best Practices

1. **Use short variable names in small scopes:** In loops and short functions, single-letter variables are acceptable.
   ```go
   for i := 0; i < 10; i++ {  // i is acceptable here
       fmt.Println(i)
   }
   ```

2. **Use meaningful names in larger scopes:** For global variables and function parameters, use descriptive names.
   ```go
   var maxConnectionTimeout = 30  // Clear and meaningful
   ```

3. **Prefer short declaration operator (`:=`) inside functions:** It's more concise and idiomatic.
   ```go
   // Inside a function
   name := "John"  // Preferred
   
   // At package level
   var name = "John"  // Must use var
   ```

4. **Use constants for fixed values:** If a value shouldn't change, declare it as a constant.
   ```go
   const MaxRetries = 3
   const DefaultTimeout = 30
   ```

5. **Avoid unused variables:** Go compiler will raise an error if you declare but don't use a variable. Use the blank identifier `_` if you need to ignore a value.

---

## Summary

- Variables in Go are declared using `var`, with optional type inference, or using `:=` for short declarations
- Go supports both global and local variable scopes
- Multiple variables can be declared in a single statement
- Variable names should follow camelCase convention and be meaningful
- Use constants for fixed values that shouldn't change
- The short declaration operator `:=` is the idiomatic way to declare variables inside functions
- Go's strict type system means variables must have a defined type at compile time
