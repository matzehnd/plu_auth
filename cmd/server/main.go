package main

import (
	"log"
	"net/http"

	"github.com/example/vc-openid-idp/internal/oidc"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/openid-configuration", oidc.HandleMetadata).Methods("GET")
	r.HandleFunc("/authorize", oidc.HandleAuthorize).Methods("GET")
	r.HandleFunc("/presentation-request/{id}", oidc.HandlePresentationRequest).Methods("GET")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
