package main

import (
  "fmt"
  "github.com/vitorleal/find-cine-go"
)

func main() {
  theaters, err := findcine.Near("Miami")

  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println(theaters)
}

