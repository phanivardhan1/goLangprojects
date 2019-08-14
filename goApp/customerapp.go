package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

type Employee struct {
	Name string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "hello this is  a web app")

	s := Employee{
		Name: "phani",
	}
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	fmt.Println(s)

	w.Header().Set("Content-Type", "text/html")
	da := &Employee{Name: "phani"}
	

	tmpl.Execute(w, da)


}

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", viewHandler)
	r.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", r)

}
