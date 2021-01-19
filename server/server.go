package main

import (
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "html/template"
  "encoding/json"
)

//-------- GLOBAL VARIABLES --------------
var adminGlobalEmail string
var adminGlobalPass string
var empFilePath string = "data/EmployeeMaster.json"
//----------------------------------------

//-------- TEMPLATES ---------------------
type UpdateData struct {
  Updater string
  Updatee Employee
}

type Dashboard struct {
  AllEmps []Employee
  AdminCreds Admin
}

type Employee struct {
  Name string `json:"Name"`
  Email string `json:"Email"`
  Password string `json:"Password"`
  Position string `json:"Position"`
}

type Admin struct {
  Email, Password string
}
//----------------------------------------

// ----------- HELPER FUNCTIONS ----------
func addOrRemove(email string) {
  allEmps := getEmployees()
  empty()
  for i := len(allEmps)-1; i >= 0; i-- {
    emp := allEmps[i]
    if emp.Email == email {
      allEmps = append(allEmps[:i], allEmps[i+1:]...)
    } else { save(emp) }
  }
}

func empty() {
  os.Remove(empFilePath)
  os.Create(empFilePath)
}

func getEmployees() []Employee {
  file, _ := ioutil.ReadFile(empFilePath)
  var allEmps []Employee
  json.Unmarshal([]byte(file), &allEmps)
  return allEmps
}

func save(emp Employee) {
  allEmps := getEmployees()
  allEmps = append(allEmps, emp)
  empMarshalInd, _ := json.MarshalIndent(allEmps, "", "")
  _ = ioutil.WriteFile(empFilePath, empMarshalInd, 0644)
}
// ---------------------------------------


// ----------- ROUTE HANDLERS ----------
func admin(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    adminTemp := template.Must(template.ParseFiles("static/admin.html"))
    adminTemp.Execute(res, nil)
  }

  if req.Method == "POST" {
    adminGlobalEmail, adminGlobalPass = req.FormValue("email"), req.FormValue("password")

    adminCreds := Admin{Email: adminGlobalEmail, Password: adminGlobalPass}
    dashboard := Dashboard{AllEmps: getEmployees(), AdminCreds: adminCreds}

    adminTemp := template.Must(template.ParseFiles("static/dashboard.html"))
    adminTemp.Execute(res, dashboard)
  }
}

func employee(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET"  {
    empTemp := template.Must(template.ParseFiles("static/reg.html"))
    empTemp.Execute(res, nil)
  }
  if req.Method == "POST" {
    name := req.FormValue("name")
    email := req.FormValue("email")
    pos := req.FormValue("pos")
    pass := req.FormValue("password")

    e := Employee{Name: name, Email: email, Password: pass, Position: pos}
    save(e)

    empTemp := template.Must(template.ParseFiles("static/reg.html"))
    empTemp.Execute(res, nil)
  }
}

func showEmployees(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    allEmps := getEmployees()
    disp := template.Must(template.ParseFiles("static/adminHome.html"))
    disp.Execute(res, allEmps)
  }
}

func update(res http.ResponseWriter, req *http.Request) {
  email := req.URL.Query().Get("email")
  updateTemp := template.Must(template.ParseFiles("static/update.html"))

  var empToUpdate Employee
  allEmps := getEmployees()

  for _, emp := range allEmps {
    if emp.Email == email {
      empToUpdate = emp
      break
    }
  }
  addOrRemove(email)

  updateInfo := UpdateData{Updater: adminGlobalEmail, Updatee: empToUpdate}
  updateTemp.Execute(res, updateInfo)
}

func saveEmp(res http.ResponseWriter, req *http.Request) {
  if req.Method == "POST" {
    newName := req.FormValue("name")
    newPass := req.FormValue("pass")
    newEmail := req.FormValue("email")
    newPosition := req.FormValue("pos")

    updatedEmp := Employee{Name:newName, Email:newEmail, Password:newPass, Position:newPosition}
    save(updatedEmp)

    adminCreds := Admin{Email: adminGlobalEmail, Password: adminGlobalPass}
    dashboard := Dashboard{AllEmps: getEmployees(), AdminCreds: adminCreds}

    t := template.Must(template.ParseFiles("static/dashboard.html"))
    t.Execute(res, dashboard)
  }
}

func remove(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    email := req.FormValue("email")
    addOrRemove(email)

    adminCreds := Admin{Email: adminGlobalEmail, Password: adminGlobalPass}
    dashboard := Dashboard{AllEmps: getEmployees(), AdminCreds: adminCreds}

    t := template.Must(template.ParseFiles("static/dashboard.html"))
    t.Execute(res, dashboard)
  }
}

func home(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    homeTemp := template.Must(template.ParseFiles("static/home.html"))
    homeTemp.Execute(res, nil)
  }
}
// ---------------------------------------

func main() {
  fmt.Println("\033[32mServer is running....")

  http.HandleFunc("/", home)
  http.HandleFunc("/home", home)
  http.HandleFunc("/admin", admin)
  http.HandleFunc("/employee", employee)
  http.HandleFunc("/showEmployees", showEmployees)
  http.HandleFunc("/admin/remove/", remove)
  http.HandleFunc("/admin/update/", update)
  http.HandleFunc("/admin/save", saveEmp)
  http.ListenAndServe(":8000", nil)
}
