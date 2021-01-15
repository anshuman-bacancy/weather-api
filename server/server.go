package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "bytes"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

var empFilePath string = "data/EmployeeMaster.json"

type Employee struct {
  Name string `json:"Name"`
  Email string`json:"Email"`
  Password string`json:"Password"`
  Position string`json:"Position"`
}

type Admin struct {
  Email, Password string
}

func getEmployees() []Employee {
  file, _ := ioutil.ReadFile(empFilePath)
  empList := strings.Split(string(file), "\n")
  var e Employee
  var allEmps []Employee

  for _, emp := range empList {
    if emp == "" { break }
    json.Unmarshal([]byte(emp), &e)
    allEmps = append(allEmps, e)
  }
  return allEmps
}

func handleAdmin(res http.ResponseWriter, req *http.Request) {
  email, _ := req.FormValue("email"), req.FormValue("password")
  fmt.Fprintf(res, "Welcome, %s \n", email)

  employees := getEmployees()
  for _, em := range employees {
    fmt.Println(em.Name)
  }
}

func save(emp Employee) {
  buffer := new(bytes.Buffer)
  encoder := json.NewEncoder(buffer)
  encoder.Encode(emp)

  openFile, openFileErr := os.OpenFile(empFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

  defer openFile.Close()

  if openFileErr != nil {
    log.Fatal(openFileErr)
  }

  io.Copy(openFile, buffer)

  /*

  // marshaling(struct to JSON)
  empMarshalInd, _ := json.MarshalIndent(emp, "", "")
  empMarshalIndStr := string(empMarshalInd)
  _, _ = openFile.WriteString(empMarshalIndStr + "\n")
  */
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
  http.HandleFunc("/admin", handleAdmin)
  http.HandleFunc("/register", register)
  http.ListenAndServe(":8000", nil)
  fmt.Println("\033[32m Server is running....")
}
