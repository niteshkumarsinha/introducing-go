# Structs and Interfaces in Go (Golang)

Structs and interfaces are **core building blocks** in Go.

* **Structs** are used to group related data
* **Interfaces** define behavior (what a type can do)

Together, they enable **clean design**, **abstraction**, and **polymorphism**.

---

# PART 1: STRUCTS IN GO

---

## 1. What is a Struct?

A **struct** is a user-defined type that groups **multiple fields** of different data types into a single unit.

### Why Structs?

* Represent real-world entities
* Organize related data
* Create custom data types

---

## 2. Defining a Struct

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

---

## 3. Creating and Using a Struct

### Method 1: Declare and Assign Later

```go
var p Person
p.Name = "Nitesh"
p.Age = 30
p.City = "Delhi"
```

### Method 2: Struct Literal

```go
p := Person{
    Name: "Nitesh",
    Age:  30,
    City: "Delhi",
}
```

### Method 3: Positional Initialization (Not Recommended)

```go
p := Person{"Nitesh", 30, "Delhi"}
```

---

## 4. Accessing Struct Fields

```go
fmt.Println(p.Name)
fmt.Println(p.Age)
```

---

## 5. Nested Structs

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Age     int
    Address Address
}
```

```go
e := Employee{
    Name: "Amit",
    Age:  28,
    Address: Address{
        City:  "Mumbai",
        State: "MH",
    },
}
```

---

## 6. Anonymous Structs

Used for short-lived data structures.

```go
user := struct {
    Username string
    Active   bool
}{
    Username: "admin",
    Active:   true,
}
```

---

## 7. Struct Pointers

Go automatically dereferences struct pointers.

```go
p := Person{Name: "Rahul", Age: 25}
ptr := &p

fmt.Println(ptr.Name) // No need to use (*ptr).Name
```

---

## 8. Methods on Structs

Functions attached to structs are called **methods**.

```go
func (p Person) greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

```go
p.greet()
```

---

## 9. Pointer Receiver Methods

Used when you want to modify struct data.

```go
func (p *Person) birthday() {
    p.Age++
}
```

```go
p.birthday()
fmt.Println(p.Age)
```

---

# PART 2: INTERFACES IN GO

---

## 10. What is an Interface?

An **interface** defines a set of **method signatures**.

> It describes **behavior**, not implementation.

```go
type Shape interface {
    Area() float64
}
```

---

## 11. Implementing an Interface (Implicitly)

Go interfaces are implemented **implicitly**.

```go
type Rectangle struct {
    Width  float64
    Height float64
}
```

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

➡️ `Rectangle` automatically implements `Shape`

---

## 12. Using Interfaces

```go
func printArea(s Shape) {
    fmt.Println(s.Area())
}
```

```go
r := Rectangle{Width: 10, Height: 5}
printArea(r)
```

---

## 13. Multiple Types Implementing the Same Interface

```go
type Circle struct {
    Radius float64
}
```

```go
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
```

```go
shapes := []Shape{
    Rectangle{10, 5},
    Circle{7},
}

for _, s := range shapes {
    fmt.Println(s.Area())
}
```

---

## 14. Interface with Multiple Methods

```go
type Vehicle interface {
    Start()
    Stop()
}
```

```go
type Car struct {}
```

```go
func (c Car) Start() {
    fmt.Println("Car started")
}

func (c Car) Stop() {
    fmt.Println("Car stopped")
}
```

---

## 15. Empty Interface (`interface{}`)

Can hold values of **any type**.

```go
var data interface{}

data = 10
data = "Hello"
data = true
```

---

## 16. Type Assertion

Used to extract concrete value from interface.

```go
value, ok := data.(string)
if ok {
    fmt.Println(value)
}
```

---

## 17. Type Switch

```go
func checkType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

---

## 18. Structs + Interfaces (Real-World Example)

```go
type Payment interface {
    Pay(amount float64)
}
```

```go
type CreditCard struct {}
type UPI struct {}
```

```go
func (c CreditCard) Pay(amount float64) {
    fmt.Println("Paid via Credit Card:", amount)
}

func (u UPI) Pay(amount float64) {
    fmt.Println("Paid via UPI:", amount)
}
```

```go
func processPayment(p Payment, amount float64) {
    p.Pay(amount)
}
```

---

## 19. Best Practices

### Structs

✅ Use meaningful field names
✅ Prefer composition over inheritance
✅ Use pointer receivers for mutations

### Interfaces

✅ Keep interfaces **small** (1–2 methods)
✅ Define interfaces where they are **used**, not implemented
✅ Prefer behavior over concrete types

---

## 20. Summary

| Feature                 | Struct | Interface |
| ----------------------- | ------ | --------- |
| Holds data              | ✅      | ❌         |
| Defines behavior        | ❌      | ✅         |
| Supports methods        | ✅      | ✅         |
| Enables polymorphism    | ❌      | ✅         |
| Implicit implementation | ❌      | ✅         |

---

📌 **Key Takeaway:**

* **Structs define data**
* **Interfaces define behavior**
* Together, they form the foundation of **idiomatic Go design**

---

Happy Coding with Go 🚀
