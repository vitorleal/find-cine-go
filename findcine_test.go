package findcine

import (
  "testing"
)

func TestFindCineNear(t *testing.T) {
  _, err := Near("Miami")

  if err != nil {
    t.Error(err)
  }
}
