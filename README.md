# dess

dess is a Go package for decimal numbers.

## Installation

`go get github.com/kochman/dess`

## Usage

```
package main

import (
    "fmt"
    "github.com/kochman/dess"
)

func main() {
    decimal := dess.NewDecimal(4567, 2)
    fmt.Println(decimal)  // prints 45.67

    d2 := dess.NewDecimal(2345, 3)
    d3 := dess.NewDecimal(19, 1)
    res := decimal.Add(d2).Sub(d3)
    fmt.Println(res)  // prints 46.115
}
```