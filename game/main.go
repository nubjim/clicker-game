package main

import (
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"html/template"
	"./models"
)

//var balance = 0


var templates *template.Template
func main() {
	r := mux.NewRouter()
	templates = template.Must(template.ParseGlob("pages/*.html"))
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("POST")
	r.HandleFunc("/upgrade-click/", models.cUpgradeHandler).Methods("POST")
	r.HandleFunc("/add/", models.balanceHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
	err := r.ParseForm()
	
	if err != nil {
		w.Write([]byte("Internal server error"))
		return
	}

}

