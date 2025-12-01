# Control Structures in Go

## Introduction

Control structures are programming constructs that determine the flow of execution in a program. They allow you to make decisions, repeat code blocks, and handle different scenarios based on conditions. Go provides three main types of control structures: **if/else statements**, **switch statements**, and **loops**.

---

## 1. Conditional Statements

### 1.1 **if Statement**

The `if` statement executes a block of code only if a specified condition is true.

**Syntax:**
```go
if condition {
    // Code to execute if condition is true
}
```

**Example:**
```go
i := 10
if i > 10 {
    fmt.Println("Big")
} else {
    fmt.Println("Small")
}
// Output: Small
```

### 1.2 **if-else Statement**

The `if-else` statement executes one block if the condition is true, and another block if the condition is false.

**Syntax:**
```go
if condition {
    // Code if condition is true
} else {
    // Code if condition is false
}
```

**Example:**
```go
n := 5
if n%2 == 0 {
    fmt.Println(n, "is even")
} else {
    fmt.Println(n, "is odd")
}
// Output: 5 is odd
```

### 1.3 **if-else if-else Statement**

Use multiple `else if` statements to check multiple conditions.

**Syntax:**
```go
if condition1 {
    // Code if condition1 is true
} else if condition2 {
    // Code if condition2 is true
} else if condition3 {
    // Code if condition3 is true
} else {
    // Code if all conditions are false
}
```

**Example:**
```go
for n := 1; n <= 20; n++ {
    if n%3 == 0 && n%5 == 0 {
        fmt.Println(n, "FizzBuzz")
    } else if n%3 == 0 {
        fmt.Println(n, "Fizz")
    } else if n%5 == 0 {
        fmt.Println(n, "Buzz")
    } else {
        fmt.Println(n)
    }
}
// Output:
// 1
// 2
// 3 Fizz
// 4
// 5 Buzz
// 6 Fizz
// ... and so on
```

### 1.4 **Logical Operators**

Combine multiple conditions using logical operators:

| Operator | Meaning | Example |
|----------|---------|---------|
| `&&` | AND | `a > 5 && b < 10` |
| `||` | OR | `a > 5 \|\| b < 10` |
| `!` | NOT | `!a` |

**Example:**
```go
if n%3 == 0 && n%5 == 0 {
    fmt.Println("Divisible by both 3 and 5")
}

if n%2 == 0 || n%3 == 0 {
    fmt.Println("Divisible by 2 or 3")
}

if !(n > 10) {
    fmt.Println("n is not greater than 10")
}
```

### 1.5 **Short if Statement with Variable Declaration**

You can declare and initialize a variable within an if statement. The variable is scoped to the if/else block.

**Syntax:**
```go
if variable := expression; condition {
    // Use variable here
}
```

**Example:**
```go
if x := getValue(); x > 10 {
    fmt.Println("x is greater than 10:", x)
} else {
    fmt.Println("x is less than or equal to 10:", x)
}
```

---

## 2. Switch Statements

A `switch` statement is used to execute different code based on different conditions. It's more efficient than multiple `if-else` statements when checking a single variable against many values.

### 2.1 **Basic Switch Statement**

**Syntax:**
```go
switch variable {
case value1:
    // Code if variable == value1
case value2:
    // Code if variable == value2
default:
    // Code if no cases match
}
```

**Example:**
```go
i := 11
switch i {
case 0: 
    fmt.Println("i is zero")
case 1: 
    fmt.Println("i is one")
case 2: 
    fmt.Println("i is two")
default: 
    fmt.Println("i is greater than two")
}
// Output: i is greater than two
```

### 2.2 **Switch with Multiple Values in a Case**

You can have multiple values in a single case, separated by commas.

**Example:**
```go
day := 6
switch day {
case 1, 2, 3, 4, 5:
    fmt.Println("Weekday")
case 6, 7:
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid day")
}
// Output: Weekend
```

### 2.3 **Switch with Expressions**

The switch value doesn't have to be a simple variable; it can be an expression.

**Example:**
```go
age := 25
switch {
case age < 13:
    fmt.Println("Child")
case age >= 13 && age < 18:
    fmt.Println("Teenager")
case age >= 18 && age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}
// Output: Adult
```

### 2.4 **Fallthrough in Switch**

By default, once a case matches, the execution stops. Use `fallthrough` to continue to the next case.

**Example:**
```go
x := 2
switch x {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
default:
    fmt.Println("Other")
}
// Output:
// Two
// Three
```

---

## 3. Loops

Loops allow you to repeat a block of code multiple times. Go has only one loop type: the `for` loop, but it can be used in different ways.

### 3.1 **for Loop with Three Components**

The traditional for loop consists of an initialization, a condition, and an increment/decrement.

**Syntax:**
```go
for initialization; condition; increment {
    // Code to repeat
}
```

**Example:**
```go
for j := 1; j <= 5; j++ {
    fmt.Println(j)
}
// Output:
// 1
// 2
// 3
// 4
// 5
```

**Breakdown:**
- `j := 1` - Initialize j to 1
- `j <= 5` - Continue while j is less than or equal to 5
- `j++` - Increment j by 1 after each iteration

### 3.2 **for Loop with Condition Only (while loop)**

If you omit the initialization and increment, the for loop behaves like a while loop.

**Syntax:**
```go
for condition {
    // Code to repeat
}
```

**Example:**
```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
// Output:
// 1
// 2
// 3
// 4
// 5
```

### 3.3 **Infinite Loop**

An infinite loop is created by omitting the condition. Use `break` to exit the loop.

**Syntax:**
```go
for {
    // Code repeats forever
    if condition {
        break
    }
}
```

**Example:**
```go
for {
    fmt.Println("Infinite Loop")
    break  // Exit the loop
}
// Output: Infinite Loop
```

### 3.4 **for Loop with break Statement**

The `break` statement exits the loop immediately.

**Example:**
```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break  // Exit when i equals 5
    }
    fmt.Println(i)
}
// Output:
// 1
// 2
// 3
// 4
```

### 3.5 **for Loop with continue Statement**

The `continue` statement skips the current iteration and moves to the next one.

**Example:**
```go
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue  // Skip even numbers
    }
    fmt.Println(i)
}
// Output:
// 1
// 3
// 5
// 7
// 9
```

### 3.6 **for Loop with range Keyword**

The `range` keyword is used to iterate over arrays, slices, maps, and strings. It returns the index and value for each element.

**Syntax:**
```go
for index, value := range collection {
    // Use index and value
}
```

**Example with array:**
```go
numbers := []int{10, 20, 30, 40, 50}
for i, num := range numbers {
    fmt.Println("Index:", i, "Value:", num)
}
// Output:
// Index: 0 Value: 10
// Index: 1 Value: 20
// ... and so on
```

**Example with string:**
```go
str := "Hello"
for i, char := range str {
    fmt.Println("Index:", i, "Character:", string(char))
}
```

---

## 4. Practical Examples

### **Example 1: Check Even or Odd**

```go
func isEven(n int) bool {
    if n%2 == 0 {
        return true
    } 
    return false
}

func isOdd(n int) bool {
    if n%2 != 0 {
        return true
    } 
    return false
}

func printEvenOdd() {
    for i := 1; i <= 10; i++ {
        if isEven(i) {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }
}
```

**Output:**
```
1 is odd
2 is even
3 is odd
4 is even
... and so on
```

### **Example 2: Find Numbers Divisible by N**

```go
func divisibleByN(n int, upToRange int) {
    i := 0
    for i <= upToRange {
        if i%n == 0 {
            fmt.Printf("%v is divisible by %v\n", i, n)
        }
        i++
    }
}
```

**Usage:**
```go
divisibleByN(3, 20)
// Output:
// 0 is divisible by 3
// 3 is divisible by 3
// 6 is divisible by 3
// 9 is divisible by 3
// 12 is divisible by 3
// 15 is divisible by 3
// 18 is divisible by 3
```

### **Example 3: FizzBuzz Problem**

The FizzBuzz problem is a classic programming challenge. Print numbers from 1 to 100, but:
- Print "Fizz" if divisible by 3
- Print "Buzz" if divisible by 5
- Print "FizzBuzz" if divisible by both 3 and 5
- Print the number otherwise

```go
func fizzBuzz() {
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println(i, "FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println(i, "Fizz")
        } else if i%5 == 0 {
            fmt.Println(i, "Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
```

**Sample Output (first 20 numbers):**
```
1
2
3 Fizz
4
5 Buzz
6 Fizz
7
8
9 Fizz
10 Buzz
11
12 Fizz
13
14
15 FizzBuzz
16
17
18 Fizz
19
20 Buzz
```

### **Example 4: Combined Control Structures**

```go
func example() {
    var i int = 1
    
    // Loop and conditional
    for i <= 10 {
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
        i += 1
    }
    
    // Switch statement
    switch i {
    case 0:
        fmt.Println("i is zero")
    case 1:
        fmt.Println("i is one")
    case 2:
        fmt.Println("i is two")
    default:
        fmt.Println("i is greater than two")
    }
    
    // Simple conditional
    i = 10
    if i > 10 {
        fmt.Println("Big")
    } else {
        fmt.Println("Small")
    }
}
```

---

## 5. Nested Control Structures

You can nest control structures inside each other to create more complex logic.

**Example:**
```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
// Output:
// i=1, j=1
// i=1, j=2
// i=1, j=3
// i=2, j=1
// i=2, j=2
// i=2, j=3
// i=3, j=1
// i=3, j=2
// i=3, j=3
```

---

## 6. Best Practices

1. **Use if-else for simple two-way decisions:** Keep the code readable and straightforward.
   ```go
   if age >= 18 {
       fmt.Println("Adult")
   } else {
       fmt.Println("Minor")
   }
   ```

2. **Use switch for multiple value checks:** It's more efficient and readable than multiple if-else statements.
   ```go
   switch day {
   case 1, 2, 3, 4, 5:
       fmt.Println("Weekday")
   case 6, 7:
       fmt.Println("Weekend")
   }
   ```

3. **Use for loops for repetitive tasks:** Go only has `for` loops, not `while` loops, so use `for` with a condition when needed.
   ```go
   for i <= 5 {
       fmt.Println(i)
       i++
   }
   ```

4. **Avoid deeply nested structures:** If your code is getting deeply nested, consider refactoring into separate functions.
   ```go
   // Instead of many nested loops/conditionals
   // Extract into helper functions
   func processData(data []int) {
       for _, item := range data {
           if isValid(item) {
               handle(item)
           }
       }
   }
   ```

5. **Use meaningful variable names in loops:** Make it clear what each loop iteration represents.
   ```go
   for index, value := range numbers {  // Good
   for i, v := range numbers {          // Less clear
   ```

6. **Use continue and break judiciously:** They make code clearer by avoiding unnecessary nesting.
   ```go
   for i := 1; i <= 10; i++ {
       if i%2 == 0 {
           continue  // Skip even numbers
       }
       fmt.Println(i)  // Only odd numbers
   }
   ```

---

## Summary

- **if/else statements** are used for conditional execution based on boolean expressions
- **switch statements** are efficient for checking a variable against multiple values
- **for loops** are the primary looping construct in Go, available in multiple forms:
  - Traditional three-component for loop
  - Condition-only for loop (while-like behavior)
  - Infinite for loop with break
  - for-range loop for iterating over collections
- **break** exits a loop immediately
- **continue** skips the current iteration
- **Logical operators** (`&&`, `||`, `!`) combine multiple conditions
- **Nested control structures** allow complex logic but should be refactored if too deeply nested
- Control structures form the foundation of writing dynamic, responsive programs in Go
