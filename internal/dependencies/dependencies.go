package dependencies

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/services"
)

type Dependencies struct {
	// Handlers
	ItemHandler          *handlers.ItemHandler
	AllArtistsHandler    *handlers.AllArtistsHandler
	ArtistDetailsHandler *handlers.ArtistsDetailsHandler
	// UserHandler *handlers.UserHandler
}

func NewDependencies() *Dependencies {
	// load data in first
	store := repository.New_Store()
	store.LoadData()

	// Instantiate services
	itemService := services.NewItemService()
	allartistservice := services.NewAllArtistsService(store)
	artistDetailsservice := services.NewArtistDetailsService(store)

	// Instantiate handlers
	return &Dependencies{
		// ItemService: itemService,
		// UserService:  userService,
		ItemHandler:          handlers.NewItemHandler(itemService),
		AllArtistsHandler:    handlers.NewAllArtistsHandler(allartistservice),
		ArtistDetailsHandler: handlers.NewArtistDetailsService(artistDetailsservice),
		// UserHandler:  userHandler,
	}
}
