package student

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Err(Str string, Status int, w http.ResponseWriter, r *http.Request) {

	Info := Error{Str, Status}
	val, err := template.ParseFiles("templates/err.html")

	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
	w.WriteHeader(Status)
	err = val.Execute(w, Info)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
}

func Media(w http.ResponseWriter, r *http.Request) {

	Searching.Values = nil

	if !Result {

		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	if r.URL.Path != "/" && r.Method == "GET" {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}

	if r.URL.Path == "/" && r.Method != "GET" {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}

	val, err := template.ParseFiles("templates/groupie.html")
	if err != nil {

		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	err = val.Execute(w, ArtistsNew)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

}

func Album(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/artists/" && r.Method == "POST" {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}

	if r.Method != "POST" && r.URL.Path == "/artists/" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w, r)
		return
	}

	val, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
	name := strings.TrimPrefix(r.URL.Path, "/artists/")
	id, smt := strconv.Atoi(name)
	if smt != nil {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}
	if (len(RelationNew.Index) < id) || (id < 1) {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}

	err = val.Execute(w, ArtistsNew[id-1])
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/search/" && r.Method == "POST" {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}

	if r.Method != "POST" && r.URL.Path == "/search/" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w, r)
		return
	}

	val, err := template.ParseFiles("templates/search.html")
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	str := r.FormValue("myData")
	if len(str) == 0 {
		Err("400 Bad Request", http.StatusBadRequest, w, r)
		return
	}

	res := Choosen(str)
	if !res {
		option := r.FormValue("options")
		SearchFor(str, option)
	}

	if len(Searching.Values) == 0 {
		Err("Not found :)", http.StatusOK, w, r)
		return
	}

	err = val.Execute(w, Searching.Values)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
}

func MainFunc() {
	val := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style", val))
	Func()
	http.HandleFunc("/", Media)
	http.HandleFunc("/artists/", Album)
	http.HandleFunc("/search/", Search)
	http.ListenAndServe(":7770", nil)
}
