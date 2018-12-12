package main

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/dimfeld/httptreemux"

	"github.com/HectorNM/minha-api-em-go/music"
	"github.com/HectorNM/minha-api-em-go/api"
)

func main() {
	session, err := mgo.Dial("localhost:27017/song-database")
	if err != nil {
		log.Fatal(err)
	}

	repository := music.NewSongRepository(session)


	// creating songs
	song := &music.Song{Id: "1235", Title: "Electrictiy", Artist: "Dua Lipa"}
	err = repository.Create(song)

	if (err != nil) {
		log.Println("Failed to add song: ", err)	
	}
	song2 := &music.Song{Id: "1236", Title: "Piranha", Artist: "Biltre"}
	err = repository.Create(song2)

	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/songs/:id", &api.GetSongHandler{ Repository: repository})

	addr := "127.0.0.1:8084"
	log.Fatal(http.ListenAndServe(addr, router))
}
