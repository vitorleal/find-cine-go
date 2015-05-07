# Find Cine Golang [![Build Status](https://travis-ci.org/vitorleal/find-cine-go.svg?branch=master)](https://travis-ci.org/vitorleal/find-cine-go) [![GoDoc](https://godoc.org/github.com/vitorleal/find-cine-go?status.png)](https://godoc.org/github.com/vitorleal/find-cine-go)

Find move theaters near an address using the [Google Moveis](http://www.google.com/movies) page.


## Install

First install is using ```go get```:

```
go get github.com/vitorleal/find-cine-go
```


## Example

```go
package main

import (
  "fmt"
  "github.com/vitorleal/find-cine-go"
)

func main() {
  // pass an address to the method Near in findcine to get the movie theaters
  theaters, err := findcine.Near("Miami")

  // if return any error 
  if err != nil {
    fmt.Println(err.Error())
  }
  
  // print a list of theaters and movies
  fmt.Println(theaters)
}
```

#### Theater ```struct```:

```go
type Theater struct {
    Id      string
    Name    string
    Address string
    Phone   string
    Movies  []Movie
}
```

#### Movie ```struct```:

```go
type Movie struct {
    Title    string
    Duration string
    Times    []string
}
```


## Docs

You can find the docs in [GoDoc.org](https://godoc.org/github.com/vitorleal/find-cine-go)


## Tests

To test it:

```
go test
```

If you want to test with code coverage:

```
go test -cover
```


## Author

| [![twitter/vitorleal](http://gravatar.com/avatar/e133221d7fbc0dee159dca127d2f6f00?s=80)](http://twitter.com/vitorleal "Follow @vitorleal on Twitter") |
|---|
| [Vitor Leal](http://vitorleal.com) |
