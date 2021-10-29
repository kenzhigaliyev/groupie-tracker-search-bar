package student

import (
	"strconv"
	"strings"
)

func Choosen(str string) bool {
	two_str := strings.Split(str, " - ")
	if len(two_str) == 2 {
		SearchFor(two_str[0], two_str[1])
		return true
	}
	return false
}

func Entered(str1 string, str2 string) int {
	num := Options(str1, str2)
	return num
}

func Options(str1 string, str2 string) int {
	switch str2 {
	case "artist/band name":
		SearchFor(str1, str2)
	case "first album date":
		SearchFor(str1, str2)
	case "locations":
		SearchFor(str1, str2)
	case "creation date":
		SearchFor(str1, str2)
	case "members":
		SearchFor(str1, str2)
	}
	return -1
}

func (Searching *SearchData) AddItem(item Artists) {
	Searching.Values = append(Searching.Values, item)
}

func SearchFor(str2 string, str1 string) {
	for index, _ := range ArtistsNew {
		switch str1 {
		case "artist/band name":
			if ArtistsNew[index].Name == str2 {
				Searching.AddItem(ArtistsNew[index])
			}
		case "first album date":
			if ArtistsNew[index].FirstAlbum == str2 {
				Searching.AddItem(ArtistsNew[index])
			}
		case "locations":
			for res := range ArtistsNew[index].DatesLocations {
				if res == str2 {
					Searching.AddItem(ArtistsNew[index])
				}
			}
		case "creation date":
			res, _ := strconv.Atoi(str2)
			if ArtistsNew[index].CreationDate == res {
				Searching.AddItem(ArtistsNew[index])

			}
		case "members":
			for _, res := range ArtistsNew[index].Members {
				if res == str2 {
					Searching.AddItem(ArtistsNew[index])
				}
			}
		}
	}
}
