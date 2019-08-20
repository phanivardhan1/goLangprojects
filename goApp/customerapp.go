package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

type Customer struct {
	Id        string
	FirstName string
	LastName  string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html", "static/response.html"))
	tmpl.Execute(w, nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fname := r.FormValue("fname") //Reading form Values
	lname := r.FormValue("lname")

	cus1 := Customer{
		Id:        id,
		FirstName: fname, // creaing cusomer object using the form values
		LastName:  lname,
	}

	fmt.Println(cus1)

	db, err := sql.Open("mysql", "hbstudent:hbstudent@tcp(localhost:3306)/goschema")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	stmtIns, err := db.Query("INSERT INTO customer (id, firstname, lastname) VALUES ('" + id + "','" + fname + "','" + lname + "')")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	http.Redirect(w, r, "/response", http.StatusFound)

}

func resHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/response.html"))

	db, err := sql.Open("mysql", "hbstudent:hbstudent@tcp(localhost:3306)/goschema")

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select * from customer")

	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.FirstName, &c.LastName)
		//var customer []Customer
		customers = append(customers, c)
		//fmt.Println(customers)
	}
	cs := &customers
	tmpl.Execute(w, cs)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cusid := vars["id"]

	fmt.Println()

	db, err := sql.Open("mysql", "hbstudent:hbstudent@tcp(localhost:3306)/goschema")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("select * from customer where id=" + cusid)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	var id, fname, lname string
	for rows.Next() {

		rows.Scan(&id, &fname, &lname)

	}
	da := Customer{
		Id:        id,
		FirstName: fname,
		LastName:  lname,
	}

	tmpl := template.Must(template.ParseFiles("static/update.html"))
	data := &da
	tmpl.Execute(w, data)
}

func main() {

	r := mux.NewRouter()
	//f := http.FileServer(http.Dir("./static"))

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", viewHandler)
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/response", resHandler)
	r.HandleFunc("/updatecustomer/{id:[0-9]+}", updateHandler)
	http.ListenAndServe(":8080", r)

}
