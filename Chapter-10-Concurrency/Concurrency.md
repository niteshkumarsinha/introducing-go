# Concurrency in Go (Golang)

Concurrency is one of Go’s **core strengths**. Go makes it easy to build programs that perform **multiple tasks simultaneously**, efficiently utilizing CPU and I/O resources.

> **Concurrency ≠ Parallelism**
>
> * **Concurrency**: dealing with many tasks at once
> * **Parallelism**: executing multiple tasks at the same time

---

# 1. Why Concurrency in Go?

Concurrency helps to:

* Build high-performance applications
* Handle multiple requests efficiently
* Improve responsiveness
* Simplify asynchronous programming

Go provides **language-level primitives** for concurrency.

---

# 2. Goroutines

## What is a Goroutine?

A **goroutine** is a lightweight thread managed by the Go runtime.

* Extremely cheap to create
* Uses small stack
* Managed by Go scheduler

---

## Creating a Goroutine

```go
go functionName()
```

### Example

```go
func sayHello() {
    fmt.Println("Hello from Goroutine")
}

func main() {
    go sayHello()
    time.Sleep(time.Second)
}
```

---

# 3. Goroutines vs Threads

| Feature       | Goroutine                 | OS Thread |
| ------------- | ------------------------- | --------- |
| Creation cost | Very low                  | High      |
| Stack size    | Small (grows dynamically) | Large     |
| Scheduling    | Go runtime                | OS        |

---

# 4. Channels

## What is a Channel?

A **channel** is a typed conduit used for communication between goroutines.

> **Do not communicate by sharing memory; share memory by communicating.**

---

## Creating a Channel

```go
ch := make(chan int)
```

---

## Sending and Receiving Data

```go
ch <- 10   // send
value := <-ch // receive
```

---

## Example: Goroutine + Channel

```go
func worker(ch chan int) {
    ch <- 42
}

func main() {
    ch := make(chan int)
    go worker(ch)

    result := <-ch
    fmt.Println(result)
}
```

---

# 5. Buffered vs Unbuffered Channels

### Unbuffered Channel (Default)

```go
ch := make(chan int)
```

* Send blocks until receive happens

### Buffered Channel

```go
ch := make(chan int, 3)
```

* Allows limited asynchronous sends

---

# 6. Closing Channels

```go
close(ch)
```

### Receiving with `ok`

```go
value, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}
```

---

# 7. Range over Channels

```go
for val := range ch {
    fmt.Println(val)
}
```

Stops automatically when channel is closed.

---

# 8. Select Statement

`select` lets a goroutine wait on **multiple channel operations**.

```go
select {
case msg := <-ch1:
    fmt.Println(msg)
case msg := <-ch2:
    fmt.Println(msg)
default:
    fmt.Println("No data available")
}
```

---

# 9. Synchronization with WaitGroup

Used to wait for multiple goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(2)

go func() {
    defer wg.Done()
    fmt.Println("Task 1")
}()

go func() {
    defer wg.Done()
    fmt.Println("Task 2")
}()

wg.Wait()
```

---

# 10. Mutex (Mutual Exclusion)

Used to protect shared data.

```go
var mu sync.Mutex
count := 0

func increment() {
    mu.Lock()
    count++
    mu.Unlock()
}
```

---

# 11. Race Conditions

Occurs when multiple goroutines access shared data **without synchronization**.

### Detect Race Conditions

```bash
go test -race
```

---

# 12. Atomic Operations

Lock-free synchronization.

```go
var counter int64
atomic.AddInt64(&counter, 1)
```

---

# 13. Context Package

Used for:

* Cancellation
* Timeouts
* Request-scoped values

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

```go
select {
case <-ctx.Done():
    fmt.Println("Cancelled:", ctx.Err())
}
```

---

# 14. Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}
```

```go
jobs := make(chan int, 5)
results := make(chan int, 5)

for i := 1; i <= 3; i++ {
    go worker(i, jobs, results)
}
```

---

# 15. Fan-Out / Fan-In Pattern

* **Fan-out**: distribute work
* **Fan-in**: collect results

```go
func square(n int, out chan<- int) {
    out <- n * n
}
```

---

# 16. Common Concurrency Pitfalls

❌ Forgetting to close channels
❌ Goroutine leaks
❌ Deadlocks
❌ Overusing mutexes

---

# 17. Best Practices

✅ Prefer channels over shared memory
✅ Keep goroutines short-lived
✅ Avoid global shared state
✅ Use context for cancellation
✅ Always handle goroutine lifecycle

---

# 18. Concurrency vs Parallelism in Go

Go achieves parallelism when:

* Multiple CPUs are available
* `GOMAXPROCS > 1`

```go
runtime.GOMAXPROCS(runtime.NumCPU())
```

---

# 19. Common Go Concurrency Tools

| Tool       | Purpose               |
| ---------- | --------------------- |
| Goroutines | Lightweight execution |
| Channels   | Communication         |
| WaitGroup  | Synchronization       |
| Mutex      | Data protection       |
| Context    | Cancellation          |
| Atomic     | Lock-free ops         |

---

# 20. Summary

📌 **Concurrency is built into Go’s DNA**

* Goroutines are cheap
* Channels simplify communication
* Synchronization primitives are powerful
* Patterns make concurrent design scalable

> Mastering concurrency is essential for building **high-performance Go applications**.

---

Happy Concurrent Coding with Go ⚡🚀
