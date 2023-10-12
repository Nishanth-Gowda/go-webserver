package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Formula1 struct{
	TeamName string
	TeamPrincipal string
	ConstChampionships string
}

func main() {
	fmt.Println("Server Running on port 5000 ..../")

	h1 := func (w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		formula1Teams := map[string][]Formula1{
			"Teams": {
				{TeamName: "REDBULL RACING", TeamPrincipal: "Christian Horner", ConstChampionships: "5"},
				{TeamName: "Mercedes", TeamPrincipal: "Toto", ConstChampionships: "7"},
			},

		}
		templ.Execute(w, formula1Teams)
	}

	h2:= func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		teamName := r.PostFormValue("teamName")
		teamPrincipal := r.PostFormValue("teamPrincipal")
		constChampionships := r.PostFormValue("constChampionships")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "team-list-element", Formula1{TeamName: teamName, TeamPrincipal: teamPrincipal, ConstChampionships: constChampionships})

	}

	health := func (w http.ResponseWriter, r *http.Request) {
		
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-team/", h2)
	http.HandleFunc("/health/", health)

	log.Fatal(http.ListenAndServe(":5000", nil))
}