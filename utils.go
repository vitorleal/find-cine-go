package findcine

import (
  "net/url"
  "fmt"
  "strings"
  "regexp"
)

// Google Movies default url
const (
  URL = "http://www.google.com/movies"
)

// Create the Url using the address as the near parameter
func MakeUrl(address string) string {

  params := url.Values{}
  params.Add("near", address)

  uri := fmt.Sprintf("%s?%s", URL, params.Encode())

  return uri
}

// Parse the theater info to separate values
func GetTheaterInfo(info string) (addr string, phone string) {
    split := strings.Split(info, " - ")

    addr = split[0]
    phone = GetPhoneNumber(split[len(split) - 1])

    return addr, phone
}

// Check phone number
func GetPhoneNumber(phone string) string {
  // regexp to match phone numbers like:
  //   (999) 999-9999
  //   999-999-9999
  //   (99) 9999-9999
  //   (99) 99999-9999
  //   99 99999-9999
  //   99 9999-9999
  //   999 99 99 99
  var rgx, _ = regexp.Compile(`^(\(\d{2,3}\)\s?|\d{2,5}(-|\s))\d{2,5}(-|\s)(\d{2}\s?)?\d{2,4}$`)

  // test if the string is a phone number
  if rgx.Match([]byte(phone)) {
    return phone
  }

  return ""
}

