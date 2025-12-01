# Chapter 2: Data Types in Go

## Overview

Go is a **statically typed** language, meaning the type of every variable must be known at compile time. Go has several categories of data types:

1. **Basic Types** (Boolean, Numeric, String)
2. **Aggregate Types** (Array, Struct)
3. **Reference Types** (Pointer, Slice, Map, Channel, Interface, Function)

---

## 1. Basic Types

### 1.1 Boolean Type

**Concept:**
- The `bool` type represents truth values: `true` or `false`
- Used in conditional statements and logical operations
- Default value (zero value) is `false`

**Declaration and Usage:**
```go
package main

import "fmt"

func main() {
	var isActive bool = true
	var isCompleted bool = false
	
	fmt.Println("Is active:", isActive)      // Output: Is active: true
	fmt.Println("Is completed:", isCompleted) // Output: Is completed: false
	
	// Logical operations
	result := isActive && isCompleted // AND operator
	fmt.Println("AND result:", result) // Output: AND result: false
	
	result = isActive || isCompleted // OR operator
	fmt.Println("OR result:", result) // Output: OR result: true
	
	result = !isActive // NOT operator
	fmt.Println("NOT result:", result) // Output: NOT result: false
}
```

---

### 1.2 Numeric Types

#### **Integer Types**

**Concept:**
- Integers are whole numbers without a fractional part
- Go provides both signed and unsigned integer types
- Each has different size and range

**Types and Ranges:**

| Type | Size | Range |
|------|------|-------|
| `int8` | 8 bits | -128 to 127 |
| `int16` | 16 bits | -32,768 to 32,767 |
| `int32` (rune) | 32 bits | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64 bits | ±9,223,372,036,854,775,807 |
| `uint8` (byte) | 8 bits | 0 to 255 |
| `uint16` | 16 bits | 0 to 65,535 |
| `uint32` | 32 bits | 0 to 4,294,967,295 |
| `uint64` | 64 bits | 0 to 18,446,744,073,709,551,615 |
| `int` | Platform-dependent (32 or 64 bits) | Platform-dependent |
| `uint` | Platform-dependent (32 or 64 bits) | Platform-dependent |

**Usage Examples:**
```go
package main

import "fmt"

func main() {
	// Signed integers
	var age int8 = 25
	var population int32 = 7800000000
	var distance int64 = 384400
	
	fmt.Println("Age:", age)           // Output: Age: 25
	fmt.Println("Population:", population) // Output: Population: 7800000000
	fmt.Println("Distance:", distance)     // Output: Distance: 384400
	
	// Unsigned integers
	var count uint = 100
	var fileSize uint64 = 1024
	
	fmt.Println("Count:", count)       // Output: Count: 100
	fmt.Println("File size:", fileSize) // Output: File size: 1024
	
	// Using byte (alias for uint8)
	var letterA byte = 65 // ASCII value for 'A'
	fmt.Println("Letter A:", letterA)  // Output: Letter A: 65
	fmt.Println("As char:", string(letterA)) // Output: As char: A
	
	// Arithmetic operations
	sum := 10 + 20
	difference := 30 - 5
	product := 4 * 5
	quotient := 20 / 4
	remainder := 17 % 5
	
	fmt.Println("Sum:", sum)           // Output: Sum: 30
	fmt.Println("Difference:", difference) // Output: Difference: 25
	fmt.Println("Product:", product)   // Output: Product: 20
	fmt.Println("Quotient:", quotient) // Output: Quotient: 5
	fmt.Println("Remainder:", remainder) // Output: Remainder: 2
}
```

#### **Floating-Point Types**

**Concept:**
- Floating-point numbers represent real numbers with decimal points
- Go has two floating-point types: `float32` and `float64`
- `float64` is recommended for most use cases due to better precision

**Types:**

| Type | Size | Precision |
|------|------|-----------|
| `float32` | 32 bits | ~6 decimal digits |
| `float64` | 64 bits | ~15 decimal digits |

**Usage Examples:**
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// Float32
	var pi32 float32 = 3.14159
	fmt.Println("Pi (float32):", pi32) // Output: Pi (float32): 3.14159
	
	// Float64 (default for floating-point literals)
	var pi64 float64 = 3.14159265359
	fmt.Println("Pi (float64):", pi64) // Output: Pi (float64): 3.14159265359
	
	// Arithmetic operations
	radius := 5.0
	area := math.Pi * radius * radius
	fmt.Printf("Area of circle: %.2f\n", area) // Output: Area of circle: 78.54
	
	// Scientific notation
	scientific := 1.5e-3 // 0.0015
	fmt.Println("Scientific:", scientific) // Output: Scientific: 0.0015
	
	// Special values
	fmt.Println("Infinity:", math.Inf(1))  // Output: Infinity: +Inf
	fmt.Println("NaN:", math.NaN())        // Output: NaN: NaN
	
	// Trigonometric functions
	angle := math.Pi / 4
	sine := math.Sin(angle)
	cosine := math.Cos(angle)
	fmt.Printf("Sin(π/4): %.4f, Cos(π/4): %.4f\n", sine, cosine)
}
```

#### **Complex Numbers**

**Concept:**
- Go natively supports complex numbers with `complex64` and `complex128`
- Useful for mathematical and scientific calculations

**Types:**

| Type | Real Part | Imaginary Part |
|------|-----------|----------------|
| `complex64` | `float32` | `float32` |
| `complex128` | `float64` | `float64` |

**Usage Examples:**
```go
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Complex number literals
	var c1 complex128 = 3 + 4i
	var c2 complex64 = 1 + 2i
	
	fmt.Println("Complex 1:", c1) // Output: Complex 1: (3+4i)
	fmt.Println("Complex 2:", c2) // Output: Complex 2: (1+2i)
	
	// Arithmetic on complex numbers
	c3 := c1 + c2
	c4 := c1 * c2
	
	fmt.Println("Addition:", c3)      // Output: Addition: (4+6i)
	fmt.Println("Multiplication:", c4) // Output: Multiplication: (-5+10i)
	
	// Real and imaginary parts
	fmt.Println("Real part:", real(c1))      // Output: Real part: 3
	fmt.Println("Imaginary part:", imag(c1)) // Output: Imaginary part: 4
	
	// Complex functions
	fmt.Println("Absolute value:", cmplx.Abs(c1)) // Output: Absolute value: 5
	fmt.Println("Conjugate:", cmplx.Conj(c1))     // Output: Conjugate: (3-4i)
}
```

---

### 1.3 String Type

**Concept:**
- A `string` represents a sequence of UTF-8 encoded characters
- Strings are **immutable** - once created, they cannot be changed
- Default value (zero value) is an empty string `""`

**Declaration and Usage:**
```go
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// String literals
	var greeting string = "Hello, Go!"
	name := "Alice"
	
	fmt.Println(greeting)                    // Output: Hello, Go!
	fmt.Println(name)                        // Output: Alice
	
	// String concatenation
	message := greeting + " My name is " + name
	fmt.Println(message)                     // Output: Hello, Go! My name is Alice
	
	// String functions
	fmt.Println("Length:", len(greeting))    // Output: Length: 11
	fmt.Println("Uppercase:", strings.ToUpper(greeting))
	fmt.Println("Lowercase:", strings.ToLower(greeting))
	fmt.Println("Contains:", strings.Contains(greeting, "Go"))  // Output: Contains: true
	fmt.Println("Replace:", strings.Replace(greeting, "Go", "Rust", 1))
	
	// String indexing and slicing
	fmt.Println("First char:", string(greeting[0])) // Output: First char: H
	fmt.Println("Substring:", greeting[0:5])        // Output: Substring: Hello
	
	// Raw strings (backticks)
	rawString := `Line 1
Line 2
Line 3`
	fmt.Println("Raw string:\n" + rawString)
	
	// UTF-8 support
	unicode := "Hello 世界 🌍"
	fmt.Println("String:", unicode)
	fmt.Println("Length in bytes:", len(unicode))
	fmt.Println("Length in runes:", utf8.RuneCountInString(unicode))
	
	// String formatting
	age := 25
	formattedString := fmt.Sprintf("I am %d years old", age)
	fmt.Println(formattedString)             // Output: I am 25 years old
}
```

---

## 2. Aggregate Types

### 2.1 Arrays

**Concept:**
- An array is a fixed-length collection of elements of the same type
- Array length is part of its type definition and cannot be changed
- Arrays are indexed starting from 0

**Declaration and Usage:**
```go
package main

import "fmt"

func main() {
	// Declare an array with explicit length
	var numbers [5]int
	fmt.Println("Empty array:", numbers) // Output: Empty array: [0 0 0 0 0]
	
	// Initialize with values
	var colors [3]string = [3]string{"Red", "Green", "Blue"}
	fmt.Println("Colors:", colors)       // Output: Colors: [Red Green Blue]
	
	// Short declaration with initialization
	scores := [4]int{85, 90, 78, 92}
	fmt.Println("Scores:", scores)       // Output: Scores: [85 90 78 92]
	
	// Using ellipsis for implicit length
	fruits := [...]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Length:", len(fruits))  // Output: Length: 3
	
	// Accessing elements
	fmt.Println("First color:", colors[0])     // Output: First color: Red
	fmt.Println("Last score:", scores[len(scores)-1]) // Output: Last score: 92
	
	// Modifying elements
	colors[1] = "Yellow"
	fmt.Println("Modified colors:", colors)    // Output: Modified colors: [Red Yellow Blue]
	
	// Iterating through array
	fmt.Println("Iterating with for loop:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("scores[%d] = %d\n", i, scores[i])
	}
	
	// Iterating with range
	fmt.Println("Iterating with range:")
	for index, value := range colors {
		fmt.Printf("colors[%d] = %s\n", index, value)
	}
	
	// 2D Array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Element at [1][2]:", matrix[1][2]) // Output: Element at [1][2]: 6
}
```

---

### 2.2 Structs

**Concept:**
- A struct is a collection of fields with different types grouped together
- Structs allow you to create custom data types
- Fields can be exported (public) or unexported (private) based on capitalization

**Declaration and Usage:**
```go
package main

import "fmt"

// Define a struct
type Person struct {
	Name  string
	Age   int
	Email string
	City  string
}

// Another struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Method on struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Create struct instance
	person := Person{
		Name:  "John",
		Age:   30,
		Email: "john@example.com",
		City:  "New York",
	}
	
	fmt.Println("Person:", person)
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	
	// Modify struct fields
	person.Age = 31
	fmt.Println("Updated age:", person.Age)
	
	// Create struct with default values
	person2 := Person{Name: "Alice"}
	fmt.Println("Person2:", person2)
	
	// Struct with positional values
	person3 := Person{"Bob", 25, "bob@example.com", "London"}
	fmt.Println("Person3:", person3)
	
	// Rectangle struct
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	
	// Nested structs
	type Address struct {
		Street string
		City   string
		ZipCode string
	}
	
	type Employee struct {
		Name    string
		Address Address
	}
	
	emp := Employee{
		Name: "Charlie",
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			ZipCode: "02101",
		},
	}
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Println("City:", emp.Address.City)
}
```

---

## 3. Reference Types

### 3.1 Pointers

**Concept:**
- A pointer holds the memory address of a value
- Use `&` operator to get the address of a variable
- Use `*` operator to dereference a pointer (access the value it points to)
- Pointers are useful for passing variables by reference

**Declaration and Usage:**
```go
package main

import "fmt"

func modifyByPointer(x *int) {
	*x = *x * 2 // Dereference and modify
}

func main() {
	// Basic pointer operations
	var value int = 42
	var ptr *int = &value // Pointer to int
	
	fmt.Println("Value:", value)        // Output: Value: 42
	fmt.Println("Address:", &value)     // Output: Address: 0x...
	fmt.Println("Pointer:", ptr)        // Output: Pointer: 0x...
	fmt.Println("Dereference:", *ptr)   // Output: Dereference: 42
	
	// Modify value through pointer
	*ptr = 100
	fmt.Println("Modified value:", value) // Output: Modified value: 100
	
	// Pointer to struct
	person := Person{Name: "John", Age: 30}
	personPtr := &person
	
	fmt.Println("Person via pointer:", personPtr.Name) // Output: Person via pointer: John
	
	// Modifying struct through pointer
	personPtr.Age = 31
	fmt.Println("Updated age:", person.Age)  // Output: Updated age: 31
	
	// Function call with pointer
	x := 5
	modifyByPointer(&x)
	fmt.Println("After modification:", x)    // Output: After modification: 10
	
	// Nil pointer
	var nilPtr *int
	fmt.Println("Nil pointer:", nilPtr)      // Output: Nil pointer: <nil>
	
	// Check if pointer is nil
	if nilPtr == nil {
		fmt.Println("Pointer is nil")        // Output: Pointer is nil
	}
}
```

---

### 3.2 Slices

**Concept:**
- A slice is a **dynamic array** with a flexible length
- Slices are built on top of arrays
- Unlike arrays, slices can grow and shrink dynamically
- Slices are reference types and pass by reference

**Declaration and Usage:**
```go
package main

import "fmt"

func main() {
	// Declare a slice (no length specified)
	var numbers []int
	fmt.Println("Empty slice:", numbers)     // Output: Empty slice: []
	fmt.Println("Length:", len(numbers))     // Output: Length: 0
	fmt.Println("Capacity:", cap(numbers))   // Output: Capacity: 0
	
	// Initialize slice from array
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4] // Elements at indices 1, 2, 3
	fmt.Println("Slice:", slice)             // Output: Slice: [2 3 4]
	
	// Make a slice with specific length and capacity
	slice2 := make([]int, 3, 5)
	fmt.Println("Slice2:", slice2)           // Output: Slice2: [0 0 0]
	fmt.Println("Length:", len(slice2))      // Output: Length: 3
	fmt.Println("Capacity:", cap(slice2))    // Output: Capacity: 5
	
	// Slice literal
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println("Colors:", colors)           // Output: Colors: [Red Green Blue]
	
	// Append to slice
	colors = append(colors, "Yellow")
	fmt.Println("After append:", colors)     // Output: After append: [Red Green Blue Yellow]
	
	// Append multiple values
	colors = append(colors, "Orange", "Purple")
	fmt.Println("After multiple append:", colors)
	
	// Slice operations
	fmt.Println("First element:", colors[0])
	fmt.Println("Last element:", colors[len(colors)-1])
	fmt.Println("Subslice:", colors[1:3])    // Output: Subslice: [Green Blue]
	
	// Copy slice
	colorsCopy := make([]string, len(colors))
	copy(colorsCopy, colors)
	fmt.Println("Copied slice:", colorsCopy)
	
	// Modifying through subslice affects original
	slice3 := []int{1, 2, 3, 4, 5}
	subslice := slice3[1:3]
	subslice[0] = 20
	fmt.Println("Modified slice:", slice3)   // Output: Modified slice: [1 20 3 4 5]
	
	// Iterating through slice
	fmt.Println("Iterating:")
	for i, value := range slice3 {
		fmt.Printf("slice3[%d] = %d\n", i, value)
	}
}
```

---

### 3.3 Maps

**Concept:**
- A map is an **unordered collection** of key-value pairs
- Each key appears only once in a map
- Maps are created using `make` function
- Maps are reference types

**Declaration and Usage:**
```go
package main

import "fmt"

func main() {
	// Declare a map
	var age map[string]int
	fmt.Println("Uninitialized map:", age) // Output: Uninitialized map: map[]
	
	// Initialize using make
	age = make(map[string]int)
	age["Alice"] = 25
	age["Bob"] = 30
	age["Charlie"] = 28
	
	fmt.Println("Age map:", age)
	
	// Map literal
	country := map[string]string{
		"US":   "United States",
		"UK":   "United Kingdom",
		"FR":   "France",
		"DE":   "Germany",
	}
	fmt.Println("Country map:", country)
	
	// Accessing values
	fmt.Println("Alice's age:", age["Alice"]) // Output: Alice's age: 25
	
	// Check if key exists
	value, exists := age["Alice"]
	if exists {
		fmt.Println("Alice exists with age:", value) // Output: Alice exists with age: 25
	}
	
	// Non-existent key
	value2, exists2 := age["Diana"]
	fmt.Println("Diana exists:", exists2)  // Output: Diana exists: false
	fmt.Println("Diana's age:", value2)    // Output: Diana's age: 0 (zero value)
	
	// Add new key-value pair
	age["Diana"] = 35
	fmt.Println("Updated age map:", age)
	
	// Delete a key
	delete(age, "Bob")
	fmt.Println("After deletion:", age)
	
	// Map length
	fmt.Println("Map length:", len(age))   // Output: Map length: 3
	
	// Iterating through map
	fmt.Println("Iterating through age map:")
	for name, ageValue := range age {
		fmt.Printf("%s is %d years old\n", name, ageValue)
	}
	
	// Nested maps
	nestedMap := make(map[string]map[string]int)
	nestedMap["scores"] = make(map[string]int)
	nestedMap["scores"]["Alice"] = 95
	nestedMap["scores"]["Bob"] = 87
	
	fmt.Println("Nested map:", nestedMap)
	fmt.Println("Alice's score:", nestedMap["scores"]["Alice"])
}
```

---

## 4. Type Conversion

**Concept:**
- Go requires explicit type conversion between different types
- Type conversion syntax: `TargetType(value)`
- Cannot convert between incompatible types

**Examples:**
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer conversions
	var x int = 42
	var y float64 = float64(x)
	fmt.Printf("int to float64: %v (type %T)\n", y, y) // Output: int to float64: 42 (type float64)
	
	// Float to integer (truncates decimal part)
	var pi float64 = 3.14159
	var rounded int = int(pi)
	fmt.Printf("float64 to int: %v (type %T)\n", rounded, rounded) // Output: float64 to int: 3 (type int)
	
	// String to integer
	str := "123"
	num, _ := strconv.Atoi(str) // Returns (int, error)
	fmt.Println("String to int:", num)
	
	// Integer to string
	numStr := strconv.Itoa(456)
	fmt.Println("Int to string:", numStr)
	
	// String to float64
	floatStr := "3.14"
	floatVal, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println("String to float64:", floatVal)
	
	// Byte to string
	bytes := []byte{'H', 'e', 'l', 'l', 'o'}
	text := string(bytes)
	fmt.Println("Bytes to string:", text)
	
	// String to bytes
	text2 := "Hello"
	bytes2 := []byte(text2)
	fmt.Println("String to bytes:", bytes2)
	
	// Rune (int32) to string
	var r rune = 65 // ASCII 'A'
	char := string(r)
	fmt.Println("Rune to string:", char)
}
```

---

## 5. Zero Values

**Concept:**
- Every variable in Go is initialized with a zero value if not explicitly initialized
- Zero values depend on the data type

**Zero Values by Type:**

| Type | Zero Value |
|------|-----------|
| Numeric types (int, float, etc.) | 0 |
| Boolean | false |
| String | "" (empty string) |
| Pointers | nil |
| Slices | nil |
| Maps | nil |
| Channels | nil |
| Interfaces | nil |
| Functions | nil |

**Example:**
```go
package main

import "fmt"

func main() {
	var num int
	var text string
	var flag bool
	var ptr *int
	var slice []int
	var myMap map[string]int
	
	fmt.Println("int:", num)        // Output: int: 0
	fmt.Println("string:", text)    // Output: string: 
	fmt.Println("bool:", flag)      // Output: bool: false
	fmt.Println("pointer:", ptr)    // Output: pointer: <nil>
	fmt.Println("slice:", slice)    // Output: slice: []
	fmt.Println("map:", myMap)      // Output: map: map[]
}
```

---

## Summary Table

| Category | Type | Use Case |
|----------|------|----------|
| **Boolean** | bool | True/false conditions |
| **Integers** | int8, int16, int32, int64, uint8, uint16, uint32, uint64 | Counting, indexing |
| **Floating-Point** | float32, float64 | Decimal values, calculations |
| **Complex** | complex64, complex128 | Mathematical operations |
| **String** | string | Text data |
| **Array** | [n]Type | Fixed-size collections |
| **Struct** | struct | Custom data structures |
| **Pointer** | *Type | Memory addresses |
| **Slice** | []Type | Dynamic arrays |
| **Map** | map[K]V | Key-value storage |

---

## Best Practices

1. **Use `int` for most integer operations** - Go will adjust based on platform
2. **Prefer `float64` over `float32`** - Better precision for most applications
3. **Use slices instead of arrays** - More flexible for collections
4. **Use maps for fast lookups** - O(1) average access time
5. **Be explicit with type conversions** - Go requires explicit conversion
6. **Check for nil values** - Especially with pointers and maps
7. **Use meaningful struct field names** - Improves code readability
8. **Leverage zero values** - Simplifies initialization
