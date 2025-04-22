package api

import (
	"net/http"

	"go-rest-api/internal/dependencies"
)

func NewRouter(deps *dependencies.Dependencies) http.Handler {
	mux := http.NewServeMux()

	// Register routes by domain
	registerArtistsRoutes(mux, deps)
	// registerUserRoutes(mux, deps)

	return mux
}

func registerArtistsRoutes(mux *http.ServeMux, deps *dependencies.Dependencies) {
	mux.HandleFunc("/api/artists",deps.AllArtistsHandler.GetAllArtists)
	mux.HandleFunc("/api/artists/details",deps.ArtistDetailsHandler.GetArtistDetails)	
}

// func registerUserRoutes(mux *http.ServeMux, deps *dependencies.Dependencies) {
// 	mux.Handle("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			deps.UserHandler.GetAll(w, r)
// 		case http.MethodPost:
// 			deps.UserHandler.Create(w, r)
// 		default:
// 			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	}))
// }
