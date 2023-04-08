<h2 style="padding-left: 60px"> Usage : </h2>

```go
    package main

    import (
        "euler"
        "fmt"
    )

    func main() {
        fmt.Println(euler.LargestPrimeFactor(2))
        fmt.Println(euler.LargestPrimeFactor(3))
        fmt.Println(euler.LargestPrimeFactor(5))
        fmt.Println(euler.LargestPrimeFactor(8))
    }
```

```bash
    $ go run .
    2
    3
    5
    2
    $
```