# Arrays, Slices, and Maps in Go

This chapter covers the core collection types in Go: arrays, slices, and maps. You'll learn how each type works, how to use the built-in functions and methods that operate on them, common idioms, and practical examples.

**Quick overview**
- **Array**: fixed-size, value type. Size is part of the type (e.g., `[3]int`).
- **Slice**: dynamic, flexible view over an array. Built-in functions: `len`, `cap`, `append`, `copy`.
- **Map**: unordered key-value store. Built-in functions: `len`, `delete`, `make`.

**Why it matters**: slices are the most commonly used collection in Go — they give dynamic behaviour while still being efficient. Maps are the standard associative container. Understanding the distinction between arrays (value semantics) and slices (reference to an underlying array) is crucial for correct and efficient Go code.

**Contents**
- Arrays
- Slices
- Maps
- Built-in functions and important idioms
- Common pitfalls & tips
- Exercises

**Arrays**

An array is a numbered sequence of elements of a single type with a fixed length that is part of the array type.

Declaration examples:

```go
var a [3]int               // zero-valued array: [0 0 0]
b := [3]int{1, 2, 3}       // array literal
// length inferred
c := [...]int{4, 5, 6, 7} // [4]int
var matrix [2][3]int       // multi-dimensional array
```

Key points:
- Length is part of the type: `[3]int` and `[4]int` are different types.
- Arrays are value types: assigning an array copies all elements.

Example — copying vs referencing:

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a              // copies the array
	b[0] = 99
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [99 2 3]
}
```

If you need reference semantics for arrays, use a slice or a pointer to an array: `p := &a`.

**Slices**

Slices are the idiomatic dynamic array abstraction in Go. A slice describes a segment of an underlying array and contains a pointer to the array, a length, and a capacity.

Declaration:

```go
var s []int           // nil slice, len 0, cap 0
s2 := []int{1, 2, 3}  // slice literal
s3 := make([]int, 5)  // len 5, zero-valued elements
s4 := make([]int, 3, 8) // len 3, cap 8
```

Key built-ins:
- `len(s)` — number of elements in the slice
- `cap(s)` — capacity (size of the underlying array from the first element of the slice)
- `append(s, elems...)` — returns a new slice containing appended elements (may allocate)
- `copy(dst, src)` — copies elements and returns number copied

Slicing operations:

```go
a := []int{1, 2, 3, 4, 5}
sub := a[1:4]    // elements at indices 1,2,3 -> [2 3 4]
head := a[:2]    // [1 2]
tail := a[3:]    // [4 5]
all := a[:]      // full slice view
// Full slice expression includes capacity control: a[low:high:max]
```

append and reallocation:

```go
package main

import "fmt"

func main() {
	s := make([]int, 0, 2)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3) // may allocate new underlying array
	fmt.Println(s)
}
```

Observation: when `append` grows beyond capacity, Go allocates a new underlying array, copies the old contents, and returns a slice pointing to the new array. Any other slices that pointed to the old array will not see appended elements after reallocation.

copy usage:

```go
dst := make([]int, len(src))
n := copy(dst, src) // copies min(len(dst), len(src)) elements
```

Pitfall — shared underlying array:

```go
arr := []int{1, 2, 3, 4}
s1 := arr[:2]   // [1 2], cap possibly 4
s2 := arr[1:3]  // [2 3]
s1[1] = 99      // modifies arr and s2
fmt.Println(arr, s1, s2)
```

Looping and the range variable gotcha (common mistake when taking addresses inside a `range`):

```go
people := []string{"alice", "bob", "charlie"}
ptrs := []*string{}
for _, p := range people {
	// p is reused across iterations — taking &p would append the same address
	// instead, take the address from the slice directly:
	ptrs = append(ptrs, &people[0]) // BAD example; don't do this
}

// Correct pattern when you need a pointer per element:
for i := range people {
	ptrs = append(ptrs, &people[i])
}
```

Multi-dimensional slices (slice of slices):

```go
grid := make([][]int, rows)
for i := range grid {
	grid[i] = make([]int, cols)
}
```

**Maps**

Maps are built-in unordered key-value stores.

Declaration and creation:

```go
var m map[string]int        // nil map: cannot assign to m["k"] until made
m2 := map[string]int{"a": 1, "b": 2} // map literal
m3 := make(map[string]int) // ready to use
```

Basic operations:

- Set: `m[k] = v`
- Get: `v := m[k]` (returns zero value if key absent)
- Comma-ok idiom: `v, ok := m[k]` — `ok` tells if key present
- Delete: `delete(m, k)`
- Length: `len(m)`
- Iterate: `for k, v := range m { ... }` (iteration order is random)

Example:

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["alice"] = 30
	m["bob"] = 25

	if age, ok := m["carol"]; ok {
		fmt.Println("Carol is", age)
	} else {
		fmt.Println("Carol not found")
	}

	delete(m, "bob")
	fmt.Println("len:", len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

Nil vs empty map:

```go
var m map[string]int   // nil map, cannot assign: m["k"] = 1 -> panic
m := make(map[string]int) // empty map, can assign
```

Check presence with comma-ok before using value semantics when zero values ambiguous:

```go
v, ok := m["key"]
if !ok {
	// key not present
}
```

Concurrency note: A plain `map` is not safe for concurrent writes. Use `sync.RWMutex` or `sync.Map` (for certain patterns) to avoid races.

```go
// Example: using sync.Map for concurrent use cases
import "sync"

var sm sync.Map
sm.Store("k", 42)
v, ok := sm.Load("k")
fmt.Println(v, ok)
```

**Built-in functions commonly used**

- `len(x)` — arrays, slices, maps, strings, channels
- `cap(s)` — slices, channels
- `append(slice, elems...)` — append elements, may allocate
- `copy(dst, src)` — copy between slices
- `delete(map, key)` — remove key from map
- `make(type, ...)` — allocate and initialize slices, maps, channels

**Performance notes**

- Slices are lightweight views: passing a slice to a function does not copy the underlying array (only the slice header is copied).
- When `append` causes growth beyond capacity, a new backing array is allocated and the contents copied. For performance-critical code, pre-allocate with `make` and a sensible capacity.
- Map lookups are roughly O(1) on average but can allocate; frequent small allocations can add GC pressure.

**Common pitfalls & tips**

- Remember arrays are copied on assignment; use slices or pointers when you want reference semantics.
- Distinguish `nil` vs empty slices/maps: `var s []int` is nil; `s := []int{}` is empty but non-nil. `len` works the same, but assignment/append behaviour differs for maps.
- Avoid taking the address of the `range` loop variable — it is reused each iteration.
- Map iteration order is random — do not rely on it.
- Use `copy` to clone slice contents when you need independent storage.
- Control slice capacity using the full slice expression `a[low:high:max]` when you want smaller capacity to prevent accidental sharing.

**Practical examples**

1) Safe append patterns:

```go
func buildNumbers(n int) []int {
	s := make([]int, 0, n) // pre-allocate capacity
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}
```

2) Remove element from slice (order-preserving):

```go
func removeAt(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
```

3) Fast delete when order doesn't matter:

```go
func removeUnordered(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
```

4) Counting with a map:

```go
func countWords(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	return m
}
```

**Exercises**

1. Implement `func uniqueInts(src []int) []int` that returns a slice with duplicates removed (order preserved). Use a map.
2. Implement `func rotateLeft(s []int, k int) []int` that rotates a slice left by `k` positions without allocating a new underlying array (in-place if possible).
3. Write a function to merge two maps `map[string]int` summing values for duplicate keys.
4. Benchmark different ways to build a slice of `n` ints: repeated `append` vs preallocated `make` + index assignment.
5. Demonstrate the `range` address-of-variable gotcha and fix it.

**Further reading**

- Official Go spec: https://golang.org/ref/spec
- Effective Go: https://golang.org/doc/effective_go.html (search for slices, maps)
- Go blog posts: "Go Slices: usage and internals" (blog.golang.org)

---

This file provides both the conceptual model and practical examples to help you use arrays, slices, and maps effectively in Go. Try the exercises and experiment with `append`, `copy`, and `make` so you get a feel for how memory and capacity behave.

