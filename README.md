# Find Cine Golang [![GoDoc](https://godoc.org/github.com/vitorleal/find-cine-go?status.png)](https://godoc.org/github.com/vitorleal/find-cine-go)

Find move theaters near and address using the [Google Moveis](http://www.google.com/movies) page.


## How to install

```
go get github.com/vitorleal/find-cine-go
```

## Using

How to use:

```
package main

import (
  "fmt"
  "github.com/vitorleal/find-cine-go"
)

func main() {
  findcine := FineCine{}

  theathers, err := findcine.Near("Miami")

  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println(theathers)
}
```

