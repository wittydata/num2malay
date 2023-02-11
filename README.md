# num2malay
A Golang library - convert numbers to Malay words 


## Installation

With `go get`:

```bash
go get github.com/wittydata/num2malay
```

## Code
```go
package main

import (
	"fmt"
	"github.com/wittydata/num2malay"
)

func main() {
        fmt.Println(num2malay.Convert(300))
	fmt.Println(num2malay.Convert(301))
	fmt.Println(num2malay.Convert(11))
	fmt.Println(num2malay.Convert(150))
	fmt.Println(num2malay.Convert(1501))
	fmt.Println(num2malay.Convert(5621))
	fmt.Println(num2malay.Convert(500202))
	fmt.Println(num2malay.Convert(-1520))
}

```