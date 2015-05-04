package findcine

import (
  "net/url"
  "fmt"
)

// The Google Movies default url
const (
  URL = "http://www.google.com/movies"
)

// Create the Url using the address as the near parameter
func (findcine *FindCine) MakeUrl(address string) string {

  params := url.Values{}
  params.Add("near", address)

  uri := fmt.Sprintf("%s?%s", URL, params.Encode())

  return uri
}

