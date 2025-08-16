## Task

**What will be the output of the program below?**

```go
package main

import "fmt"

func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}

func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}

func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

Explain the order of `defer` execution.

## Solution

### `test` function

1. `x` is defined as a named return
2. `x = 1`
3. Defer executes after the function `return` statement, but before function exits.
4. Function returns, defer increments the variable `x` and function exits with `x = 2`
5. Output: 2

### `anotherTest` function

1. `x` is defined as local variable
2. `x = 1`
3. `return x` **copies** value of variable `x` in this statement. As the result, function returns `x = 1`
4. Defer increments the local variable `x`, but the return value doesn't change
5. Output: 1

Output:\
2\
1
