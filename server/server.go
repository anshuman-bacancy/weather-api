package main

import (
  "fmt"
  "net/http"
)

type Employee struct {
  name, email, password, postion string
}

func defaultRoute(res http.ResponseWriter, req *http.Request) {
  name := req.FormValue("name")
  email := req.FormValue("email")
  pos := req.FormValue("pos")
  pass := req.FormValue("password")
  fmt.Fprintf(res, "Name: %s\nEmail: %s\nPosition: %s\nPassword: %s", name, email, pos, pass)
}

func main() {
  http.HandleFunc("/", defaultRoute)
  http.ListenAndServe(":8000", nil)
}
