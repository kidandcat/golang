package main

import (
  "fmt"
  "net/http"
  "html/template"
  "strings"
  "io/ioutil"
)

func main(){
  http.HandleFunc("/", serve) //set router
  fmt.Println("Server listening on port 9090")
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    fmt.Println("Error opening server on port 9090")
  }
}

//req.ParseForm() //Parse arguments
//fmt.Println("Form:", req.Form) //print form information
//fmt.Println("Url", req.URL.Path)

func serve(res http.ResponseWriter, req * http.Request){
  fmt.Println("Url", req.URL.Path)
  target := req.URL.Path
  a := strings.Split(target, "/")
  file := a[len(a)-1]
  path := strings.Join(a[0:len(a)-1], "/")
  found := false
  files, _ := ioutil.ReadDir("./static" + path)
  for _, f := range files {
    if file == f.Name() && !f.Mode().IsDir() {
      final := "./static" + path + "/" + f.Name()
      found = true
      render(res, req, final, TemplateVars{})
    }
  }
  if !found {
    render(res, req, "./static/error.html", TemplateVars{})
  }
}

func render(res http.ResponseWriter, req * http.Request, file string, p TemplateVars){
  html, err := template.ParseFiles(file)
  if err != nil {
    fmt.Println("Error reading file ", file)
  }
  html.Execute(res, p)
}

type TemplateVars map[string]string
