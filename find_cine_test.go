package findcine

import (
  "testing"
)

func TestFindCineNear(t *testing.T) {
  findcine := FindCine{}

  _, err := findcine.Near("Miami")

  if err != nil {
    t.Error(err)
  }
}
