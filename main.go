package sfw

import (
	"log"
	
	"github.com/go-chi/chi/v5"
	"github.com/syke99/sfw/app/web"
)

func main() {

	mux := chi.NewRouter()

	_, err := web.NewWeb(mux, "")
	if err != nil {
		// TODO: handle err shutdown better
		log.Fatal(err)
	}
}
