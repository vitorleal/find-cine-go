package findcine

import (
  "strings"
  "github.com/PuerkitoBio/goquery"
)

// Return the current implementation version
func Version() string {
  return "0.1.0"
}

// Theater struct
type Theater struct {
  Id string
  Name string
  Address string
  Phone string
  Movies []Movie
}

// Movie struct
type Movie struct {
  Name string
  Duration string
  Time []string
}


// Get the near theaters and movies based on address
func Near(address string) (theaters []Theater, err error) {

  doc, err := goquery.NewDocument(MakeUrl(address))

  if err != nil {
    return nil, err
  }

  doc.Find(".theater").Each(func (i int, s *goquery.Selection) {
    desc := s.Find(".desc")

    name := desc.Find(".name a")
    info := desc.Find(".info").Text()
    addr := strings.Split(info," - ")[0]
    phone := strings.Split(info, " - ")[1]

    theater := Theater{Name: name.Text(), Address: addr, Phone: phone}
    theaters = append(theaters, theater)
  })

  return theaters, nil
}

