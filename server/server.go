package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "html/template"
  "encoding/json"
)

var empFilePath string = "data/EmployeeMaster.json"

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

// ----------- HELPER FUNCTIONS ---------------

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
// ------------------------------------------



func admin(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    adminTemp := template.Must(template.ParseFiles("static/admin.html"))
    adminTemp.Execute(res, nil)
  }

  if req.Method == "POST" {
    email, pass := req.FormValue("email"), req.FormValue("password")

    adminCreds := Admin{Email: email, Password: pass}
    dashboard := Dashboard{AllEmps: getEmployees(), AdminCreds: adminCreds}

    adminTemp := template.Must(template.ParseFiles("static/dashboard.html"))
    //allEmps := getEmployees()
    adminTemp.Execute(res, dashboard)

    //fmt.Fprintf(res, "Welcome, %s \n", email)

    //disp := template.Must(template.ParseFiles("static/adminHome.html"))
    //disp.Execute(res, allEmps)

    //http.Redirect(res, req, "static/adminHome.html", http.StatusSeeOther)
    /*
    fmt.Fprintf(res, "Welcome, %s \n", email)


    disp := template.Must(template.ParseFiles("static/adminHome.html"))
    //disp, _ := template.ParseFiles("static/admin.html")
    disp.Execute(res, allEmps)
    */
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
    //email, _ := req.FormValue("email"), req.FormValue("password")
    allEmps := getEmployees()

    disp := template.Must(template.ParseFiles("static/adminHome.html"))
    disp.Execute(res, allEmps)
  }
}

func home(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    homeTemp := template.Must(template.ParseFiles("static/home.html"))
    homeTemp.Execute(res, nil)
  }
}

func main() {
  fmt.Println("\033[32mServer is running....")

  http.HandleFunc("/home", home)
  http.HandleFunc("/admin", admin)
  http.HandleFunc("/employee", employee)
  http.HandleFunc("/showEmployees", showEmployees)
  http.ListenAndServe(":8000", nil)
}
