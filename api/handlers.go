package api

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/dimfeld/httptreemux"
	"github.com/HectorNM/minha-api-em-go/music"
)

type GetSongHandler struct {
	Repository *music.SongRepository
}

func (h *GetSongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	id := params["id"]

	song, err := h.Repository.FindById(id)


	encoder := json.NewEncoder(w)
	err = encoder.Encode(song)

	if err == nil {
		log.Printf("GET /songs/%s: %v\n", id, song)
		//fmt.Fprintf(w, "The song you requested is: %v\n", song)
	} else {
		log.Println("Internal Server Error: ", err)
	}

}