## Task

**What will be the output of the program below?**

```go
package main

import "fmt"

func main() {
  a := [5]int{76, 77, 78, 79, 80}
  var b []int = a[1:4]
  fmt.Println(b)
}
```

Explain the output

## Solution

1. Array `a` is initialized with values [76, 77, 78, 79, 80]
2. Slice `b` is created from elements of array `a` with indexes from 1 to 4: [77, 78, 79]
3. Slice `b` is printed

Output: [77 78 79]