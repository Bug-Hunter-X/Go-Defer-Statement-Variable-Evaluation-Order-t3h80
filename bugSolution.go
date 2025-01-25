```go
package main

import "fmt"

func main() {
    var x int = 10
    fmt.Println(x)
    defer func() {
        fmt.Println(x + 10)
    }()
    x = 5
}
```