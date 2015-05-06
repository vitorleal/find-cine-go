package findcine

import (
  "fmt"
  "testing"
)

func TestFindCineNear(t *testing.T) {
  theaters, err := Near("Rua Guaraiuva")
  theaters2, err2 := Near("Miami")

  if err != nil {
    t.Error(err)
  }

  if err2 != nil {
    t.Error(err2)
  }

  for _, theater := range theaters {
    fmt.Printf("\n• %s | %s | %s\n", theater.Name, theater.Address, theater.Phone)

    for _, movie := range theater.Movies {
      fmt.Printf("    - %s\n", movie.Title)
      fmt.Printf("      Duração %s\n\n", movie.Duration)
    }
  }

  for _, theater := range theaters2 {
    fmt.Printf("\n• %s | %s | %s\n", theater.Name, theater.Address, theater.Phone)

    for _, movie := range theater.Movies {
      fmt.Printf("    - %s\n", movie.Title)
      fmt.Printf("      Duração %s\n\n", movie.Duration)
    }
  }
}

