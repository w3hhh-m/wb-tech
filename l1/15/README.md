## Task

**Look at the following code and answer these questions:**

* What problems could this code cause?
* How can you fix those problems?

**Also, provide a correct version of the code.**

```go
var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
```

**Question:** What happens to the variable `justString`?

---

## Solution

In Go, strings are immutable and implemented as a structure with a pointer to a byte array and a length. When we slice a string like in this example,
we don't actually copy the first 100 bytes — we create a new string that still references the original underlying memory of the string v. This causes a memory leak.

We need to copy the first 100 characters to avoid storing the whole string in memory.
Also, add a length check to avoid panic if the string is shorter than 100 characters.

```go
var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    
    if len(v) < 100 {
        return
    }
    
    justString = string([]byte(v[:100]))
}

func main() {
    someFunc()
}
```
