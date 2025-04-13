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
	mux.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			find := r.URL.Query().Get("find")
			if id != "" {
				deps.ArtistDetailsHandler.GetArtistDetails(w, r)
			} else if find != "" {
				deps.SearchHandler.Search(w, r)
			} else {
				deps.AllArtistsHandler.GetAllArtists(w, r)
			}
		case http.MethodPost:
			deps.FilterHandler.Filter(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
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
