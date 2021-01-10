package main

import (
  "fmt"
  "strconv"
)

type personalDetails struct {
  id int
  name string
  add address
}

type address struct {
  area, country string
}

type techDets struct {
  techName string
  exp float32
}

type techDeetColl struct {
  id int
  tech []techDets
}

type contactDetails struct {
  email, phone string
}

type contactDetailsColl struct {
  id int
  contacts contactDetails
}

type user struct {
  name string
  add address
  techDets techDeetColl
  email, phone string
}

func main() {
  countryCode := map[string]string {
    "IND": "+91",
    "UK": "+41",
  }

  personDetailsCollection := []personalDetails {
    {
      id: 1,
      name: "radha",
      add: address {
        area: "bopal",
        country: "IND",
      },
    },
    {
      id: 2,
      name: "aniket",
      add: address {
        area: "maninagar",
        country: "UK",
      },
    },
  }

  techDetailsCollection := []techDeetColl {
    {
      id: 1,
      tech: []techDets{{techName:"React", exp:3}, {techName:"Golang", exp:1.5}},
    },
    {
      id: 2,
      tech: []techDets{{techName:"Vue", exp:0.9}, {techName:"Golang", exp:1.5}},
    },
  }

  contactDetailsCollection := []contactDetailsColl {
    {
      id: 1,
      contacts: contactDetails {
        email: "radha.kotecha@bacancy.com",
        phone: "1234567890",
      },
    },
    {
      id: 2,
      contacts: contactDetails {
        email: "aniket.amin@bacancy.com",
        phone: "0987654321",
      },
    },
  }

  //fmt.Println(countryCode)
  ret := make(map[string]user)
  u := user{}

  for i := 0; i < 2; i++ {
    //fmt.Println(personDetailsCollection[i], "--" ,techDetailsCollection[i],"--", contactDetailsCollection[i])
    if personDetailsCollection[i].id == 1 {
      u.name = personDetailsCollection[i].name
      u.add = personDetailsCollection[i].add
      u.techDets = techDetailsCollection[i]
      u.email = contactDetailsCollection[i].contacts.email
      if personDetailsCollection[i].add.country == "IND" {
        u.phone = countryCode[personDetailsCollection[i].add.country] + contactDetailsCollection[i].contacts.phone
      }
      newId := strconv.Itoa(personDetailsCollection[i].id)
      ret[newId] = u
    }
  }
  fmt.Println(ret)
}
