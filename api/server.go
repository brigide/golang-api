package api

import (
	"fmt"
	"golang-api/api/routes"
	"log"
	"net/http"
)

func listen(p int) {
	fmt.Printf("\n\nListening port %d...", p)
	port := fmt.Sprintf(":%d", p)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, r))
}
