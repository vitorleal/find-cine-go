// Package find-cine to search for movie theaters close to an address.
//
// Example:
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/vitorleal/find-cine-go"
//  )
//
//  func main() {
//    theaters, err := findcine.Near("Miami")
//
//    if err != nil {
//      fmt.Println(err.Error())
//    }
//
//    fmt.Println(theaters)
//  }
//

package findcine

