package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  //"text/template"
  "encoding/json"
)

var empFilePath string = "data/EmployeeMaster.json"

type Employee struct {
  Name string `json:"Name"`
  Email string `json:"Email"`
  Password string `json:"Password"`
  Position string `json:"Position"`
}

type Admin struct {
  Email, Password string
}

func getEmployees() []Employee {
  file, _ := ioutil.ReadFile(empFilePath)
  var allEmps []Employee
  json.Unmarshal([]byte(file), &allEmps)

  return allEmps
}

func handleAdmin(res http.ResponseWriter, req *http.Request) {
  email, _ := req.FormValue("email"), req.FormValue("password")
  fmt.Fprintf(res, "Welcome, %s \n", email)
}

func save(emp Employee) {
  allEmps := getEmployees()
  allEmps = append(allEmps, emp)
  empMarshalInd, _ := json.MarshalIndent(allEmps, "", "")
  _ = ioutil.WriteFile(empFilePath, empMarshalInd, 0644)
}

func register(res http.ResponseWriter, req *http.Request) {
  name := req.FormValue("name")
  email := req.FormValue("email")
  pos := req.FormValue("pos")
  pass := req.FormValue("password")

  e := Employee{Name: name, Email: email, Password: pass, Position: pos}
  save(e)
}

func main() {
  fmt.Println("\033[32mServer is running....")

  http.HandleFunc("/admin", handleAdmin)
  http.HandleFunc("/register", register)
  http.ListenAndServe(":8000", nil)
}
