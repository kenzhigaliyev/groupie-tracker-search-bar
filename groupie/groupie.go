package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Error struct {
	Str  string
	Type int
}

type Groupie struct {
	Artists  string `json:"artists"`
	Relation string `json:"relation"`
}

type Relation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Artists struct {
	ID             int64    `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}

var GroupieNew = Groupie{}
var RelationNew = Relation{}
var ArtistsNew []Artists
var Result bool

var Searching = SearchData{}

type SearchData struct {
	Values []Artists
}

func Dates() {
	for index := range ArtistsNew {
		ArtistsNew[index].DatesLocations = RelationNew.Index[index].DatesLocations
	}
}

func Func() {
	var Url = "https://groupietrackers.herokuapp.com/api"
	Data(Url, &GroupieNew)
	Data(GroupieNew.Relation, &RelationNew)
	Data(GroupieNew.Artists, &ArtistsNew)
	Dates()
}

func Data(url string, val interface{}) {
	res, err := http.Get(url)
	if err != nil {
		Result = false
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Result = false
		return
	}
	Result = true
	json.Unmarshal(body, &val)
	// fmt.Println(string(body))
}
