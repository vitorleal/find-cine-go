package findcine

import (
  "fmt"
  "testing"
)

func TestFindCineNear(t *testing.T) {
  theaters, err := Near("Miami")

  if err != nil {
    t.Error(err)
  }

  fmt.Println(theaters)
}
