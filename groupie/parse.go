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
	ArtistsNew[id-1].DatesLocations = RelationNew.Index[id-1].DatesLocations
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
	num := Choosen(str)
	if num == -1 {
		option := r.FormValue("options")
		num = Entered(str, option)
	}

	if num == -1 {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
	err = val.Execute(w, ArtistsNew[num])
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}
}

func Choosen(str string) int {
	two_str := strings.Split(str, " - ")
	if len(two_str) == 2 {
		num := SearchFor(two_str[0], two_str[1])
		return num
	}
	return -1
}

func Entered(str1 string, str2 string) int {
	num := Options(str1, str2)
	return num
}

func Options(str1 string, str2 string) int {
	switch str2 {
	case "artist/band name":
		return SearchFor(str1, str2)
	case "first album date":
		return SearchFor(str1, str2)
	case "locations":
		return SearchFor(str1, str2)
	case "creation date":
		return SearchFor(str1, str2)
	case "members":
		return SearchFor(str1, str2)
	}
	return -1
}

func SearchFor(str2 string, str1 string) int {
	for index, _ := range ArtistsNew {
		switch str1 {
		case "artist/band name":
			if ArtistsNew[index].Name == str2 {
				return index
			}
		case "first album date":
			if ArtistsNew[index].FirstAlbum == str2 {
				return index
			}
		case "locations":
			for res, _ := range ArtistsNew[index].DatesLocations {
				if res == str2 {
					return index
				}
			}
		case "creation date":
			res, _ := strconv.Atoi(str2)
			if ArtistsNew[index].CreationDate == res {
				return index
			}
		case "members":
			for _, res := range ArtistsNew[index].Members {
				if res == str2 {
					return index
				}
			}
		}
	}
	return -1
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
