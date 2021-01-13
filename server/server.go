package main

import (
  "os"
  "fmt"
  "log"
  "encoding/json"
  "net/http"
)

type Employee struct {
  Name, Email, Password, Position string
}

func read() {
}

func save(emp *Employee) {
  openFile, openFileErr := os.OpenFile("EmployeeMaster.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if openFileErr != nil {
    log.Fatal(openFileErr)
  }
  empData, fileErr := json.Marshal(emp)
  if fileErr != nil {
    fmt.Println(fileErr)
    return
  }
  _, _ = openFile.WriteString(string(empData) + "\n")
}

func addEmp(res http.ResponseWriter, req *http.Request) {
  name := req.FormValue("name")
  email := req.FormValue("email")
  pos := req.FormValue("pos")
  pass := req.FormValue("password")
  //fmt.Fprintf(res, "Name: %s\nEmail: %s\nPosition: %s\nPassword: %s", name, email, pos, pass)

  e := &Employee{Name: name, Email: email, Password: pass, Position: pos}
  save(e)
}

func main() {
  http.HandleFunc("/add", addEmp)
  http.ListenAndServe(":8000", nil)
}
