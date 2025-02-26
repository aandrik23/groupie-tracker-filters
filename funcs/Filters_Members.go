package groupie_tracker_search

import "fmt"

func Members(artists []Artist) {
	for _, artist := range artists {
		artist.NumMembers = len(artist.Members)
		fmt.Println(artist.NumMembers)
	}
}
