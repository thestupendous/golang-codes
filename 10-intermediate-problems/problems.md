A set of **10 intermediate-level Go (Golang) problems** that you can try to solve. These involve concurrency, error handling, data structures, and idiomatic Go features:

---

### 1. **Word Frequency Counter**

* Write a function that takes a string and returns a map of words with their frequency counts.
* Handle punctuation and case sensitivity.
* Example: `"Go go Gophers!" → {"go": 2, "gophers": 1}`.

---

### 2. **Concurrent Prime Finder**

* Use goroutines to find all prime numbers up to `N`.
* Distribute the workload among multiple goroutines and combine results using channels.

---

### 3. **JSON to Struct Converter**

* Parse a given JSON string into a struct and print it nicely.
* Example JSON:

  ```json
  {"id": 1, "name": "Alice", "skills": ["Go", "Docker"]}
  ```

---

### 4. **Worker Pool**

* Implement a worker pool with goroutines and channels.
* Given a list of integers, compute their squares in parallel.

---

### 5. **LRU Cache**

* Implement an LRU (Least Recently Used) cache in Go using a map and a doubly-linked list.
* Support operations: `Get(key)` and `Put(key, value)`.

---

### 6. **HTTP Server with Routing**

* Build a simple HTTP server in Go.
* Add support for multiple routes (e.g., `/hello`, `/time`) using `net/http`.

---

### 7. **Pipeline Processing**

* Create a pipeline of goroutines where:

  1. First goroutine generates numbers from 1–100.
  2. Second goroutine filters even numbers.
  3. Third goroutine squares them.
  4. Final goroutine prints them.

---

### 8. **CSV Reader & Aggregator**

* Read a CSV file containing `Name, Score`.
* Compute the average score per person and print results.

---

### 9. **Custom Error Handling**

* Create a function that divides two numbers.
* If denominator is `0`, return a custom error type (`DivideByZeroError`).

---

### 10. **Chat Application (CLI-based)**

* Build a simple command-line chat app using goroutines and channels.
* Multiple users should be able to send messages (simulated by goroutines).
* Messages are broadcast to all "connected" goroutines.

---
