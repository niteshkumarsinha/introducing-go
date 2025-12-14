# Packages and Modules in Go (Golang)

Go uses **packages** and **modules** to organize code, manage dependencies, and enable scalable application development.

* **Package** → Code organization unit
* **Module** → Dependency and version management unit

---

# PART 1: PACKAGES IN GO

---

## 1. What is a Package?

A **package** is a collection of Go files in the same directory that:

* Belong to the same namespace
* Are compiled together
* Share exported identifiers

Every Go file starts with a `package` declaration.

```go
package main
```

---

## 2. Why Use Packages?

✅ Code organization
✅ Reusability
✅ Encapsulation
✅ Faster compilation
✅ Clear project structure

---

## 3. Built-in Packages

Go provides many standard packages:

| Package    | Purpose                |
| ---------- | ---------------------- |
| `fmt`      | Input / Output         |
| `math`     | Mathematical functions |
| `os`       | OS-level operations    |
| `time`     | Date & time            |
| `net/http` | HTTP servers & clients |

Example:

```go
import "fmt"

fmt.Println("Hello Go")
```

---

## 4. Creating a Custom Package

### Folder Structure

```
myapp/
 ├── main.go
 └── mathutils/
     └── math.go
```

### math.go

```go
package mathutils

func Add(a int, b int) int {
    return a + b
}
```

> 🔹 Function name starts with **capital letter** → exported

---

## 5. Using a Custom Package

### main.go

```go
package main

import (
    "fmt"
    "myapp/mathutils"
)

func main() {
    result := mathutils.Add(10, 20)
    fmt.Println(result)
}
```

---

## 6. Exported vs Unexported Identifiers

| Name    | Visibility           |
| ------- | -------------------- |
| `Add()` | Exported (Public)    |
| `add()` | Unexported (Private) |

Rule:

> **Capital letter = Public**

---

## 7. Package Initialization (`init` Function)

Each package can have one or more `init()` functions.

```go
func init() {
    fmt.Println("Package initialized")
}
```

✔ Automatically executed
✔ Runs before `main()`
✔ Cannot be called manually

---

## 8. Package Aliases

Used to avoid name conflicts or improve readability.

```go
import m "math"

fmt.Println(m.Sqrt(16))
```

---

## 9. Blank Identifier Import

Used when you only need side effects (e.g., `init()`).

```go
import _ "database/sql/driver"
```

---

## 10. Internal Packages

Restrict access to packages within a module.

```
myapp/
 ├── internal/
 │   └── config/
 │       └── config.go
 └── main.go
```

Only accessible **inside `myapp`**.

---

# PART 2: MODULES IN GO

---

## 11. What is a Go Module?

A **module** is a collection of related packages with:

* A module path
* Versioned dependencies
* A `go.mod` file

Modules were introduced in **Go 1.11**.

---

## 12. Initializing a Module

```bash
go mod init github.com/nitesh/myapp
```

Creates:

```
go.mod
```

---

## 13. go.mod File Explained

```go
module github.com/nitesh/myapp

go 1.22

require (
    github.com/gin-gonic/gin v1.10.0
)
```

| Field     | Meaning      |
| --------- | ------------ |
| `module`  | Module name  |
| `go`      | Go version   |
| `require` | Dependencies |

---

## 14. Adding Dependencies

```bash
go get github.com/gin-gonic/gin
```

Automatically updates:

* `go.mod`
* `go.sum`

---

## 15. go.sum File

* Stores dependency checksums
* Ensures reproducible builds
* Auto-managed by Go

❌ Do NOT edit manually

---

## 16. Using Third-Party Packages

```go
import "github.com/gin-gonic/gin"
```

```go
r := gin.Default()
r.GET("/", func(c *gin.Context) {
    c.String(200, "Hello")
})
r.Run()
```

---

## 17. Updating Dependencies

```bash
go get -u github.com/gin-gonic/gin
```

---

## 18. Tidying Modules

Removes unused dependencies.

```bash
go mod tidy
```

---

## 19. Replacing Modules (Local Development)

```go
replace github.com/example/lib => ../lib
```

Useful for:

* Local testing
* Debugging libraries

---

## 20. Project Structure (Recommended)

```
myapp/
 ├── cmd/
 │   └── server/
 │       └── main.go
 ├── internal/
 ├── pkg/
 ├── go.mod
 └── go.sum
```

---

## 21. Packages vs Modules

| Feature    | Package           | Module                |
| ---------- | ----------------- | --------------------- |
| Purpose    | Code organization | Dependency management |
| Scope      | Directory-level   | Project-level         |
| Versioning | ❌                 | ✅                     |
| go.mod     | ❌                 | ✅                     |

---

## 22. Best Practices

### Packages

✅ Small and focused
✅ Clear naming
✅ Avoid circular dependencies

### Modules

✅ One module per repository
✅ Commit `go.mod` and `go.sum`
✅ Use semantic versioning

---

## 23. Common Commands Cheat Sheet

```bash
go mod init
go mod tidy
go mod download
go get
go list -m all
```

---

## 24. Summary

📦 **Package** → Organizes code
📦 **Module** → Manages dependencies
📦 **go.mod** → Single source of truth

> Mastering packages and modules is essential for writing **scalable**, **maintainable**, and **production-ready Go applications**.

---

Happy Coding with Go 🚀
