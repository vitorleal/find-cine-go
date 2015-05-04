# Find Cine Golang [![Build Status](https://travis-ci.org/vitorleal/find-cine-go.svg?branch=master)](https://travis-ci.org/vitorleal/find-cine-go) [![GoDoc](https://godoc.org/github.com/vitorleal/find-cine-go?status.png)](https://godoc.org/github.com/vitorleal/find-cine-go)

Find move theaters near an address using the [Google Moveis](http://www.google.com/movies) page.


## How to install

```
go get github.com/vitorleal/find-cine-go
```

## Using

How to use:

```go
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

