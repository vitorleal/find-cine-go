package findcine

import (
  "testing"
  "strings"
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

func TestGetTheaterInfo(t *testing.T) {
  info := "Av. Dr.Chucri Zaidan, 920, São Paulo - SP, Brazil - (11) 3048-7405"
  info2 := "Av. Dr.Chucri Zaidan, 920, São Paulo"
  info3 := "1508 SW 8th Street, Miami, FL, United States - (305) 643-8706"

  addr, phone := GetTheaterInfo(info)
  addr2, phone2 := GetTheaterInfo(info2)
  addr3, phone3 := GetTheaterInfo(info3)

  if strings.TrimSpace(addr) != "Av. Dr.Chucri Zaidan, 920, São Paulo" {
    t.Error("Error in Brazilian Address. Got:", addr)
  }
  if strings.TrimSpace(phone) != "(11) 3048-7405" {
    t.Error("Error in Brazilian Phone. Got:", phone)
  }

  if strings.TrimSpace(addr2) != "Av. Dr.Chucri Zaidan, 920, São Paulo" {
    t.Error("Error in Brazilian Address. Got:", addr2)
  }
  if strings.TrimSpace(phone2) != "" {
    t.Error("Error in Brazilian Phone. Got:", phone2)
  }

  if strings.TrimSpace(addr3) != "1508 SW 8th Street, Miami, FL, United States" {
    t.Error("Error in US Address. Got:", addr3)
  }
  if strings.TrimSpace(phone3) != "(305) 643-8706" {
    t.Error("Erro in US Phone. Got:", phone3)
  }
}

func TestGetPhoneNumber(t *testing.T) {
  phones := []string{
    "(305) 643-8706",
    "305-643-8706",
    "305 643-8706",
    "(11) 3048-7405",
    "(11) 93048-7405",
    "11 93048-7405",
    "11 3048-7405",
    "902 33 32 31",
    "902 33 32 31",
    "902 33 3231",
  }
  notPhones := []string{"Lorem isoum", " "}

  for _, val := range phones {
    var check = GetPhoneNumber(val)

    if check == "" {
      t.Error("Error in the phone format:", val)
    }
  }

  for _, val := range notPhones {
    var check = GetPhoneNumber(val)

    if check != "" {
      t.Error("Error in the phone match", val)
    }
  }
}

func TestGetMovieInfo(t *testing.T) {
  info := "2hr 3min - Rated R - Drama - Trailer - IMDb"
  duration := GetMovieInfo(info)

  if duration != "2hr 3min" {
    t.Error("Error geting the description")
  }
}
