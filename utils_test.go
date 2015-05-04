package findcine

import (
  "testing"
)

func TestMakeUrl(t *testing.T) {
  url := MakeUrl("Miami")
  url2 := MakeUrl("Calle del Pinar")

  if url != "http://www.google.com/movies?near=Miami" {
    t.Error("Error in the MakeUrl")
  }

  if url2 != "http://www.google.com/movies?near=Calle+del+Pinar" {
    t.Error("Error in the MakeUrl")
  }
}
