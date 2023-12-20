package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Anime struct {
	Title string
	Year  string
}

func main() {
	fmt.Println("yeahhh boi")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		// key is string, value is map of Anime
		animes := map[string][]Anime{
			"Animes": {
				{Title: "Jujutsu Kaisen", Year: "2020"},
				{Title: "Kaguya-sama: Love is War", Year: "2019"},
				{Title: "Death Note", Year: "2006"},
			},
		}
		tmpl.Execute(w, animes)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		year := r.PostFormValue("year")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "anime-list-element", Anime{Title: title, Year: year})
	}

	http.HandleFunc("/", handler1)
	http.HandleFunc("/add-anime/", handler2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
