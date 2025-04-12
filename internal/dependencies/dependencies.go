package dependencies

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/services"
)

type Dependencies struct {
	// Handlers
	ItemHandler *handlers.ItemHandler
	// UserHandler *handlers.UserHandler
}

func NewDependencies() *Dependencies {
	// Instantiate services
	itemService := services.NewItemService()

	// Instantiate handlers
	return &Dependencies{
		// ItemService: itemService,
		// UserService:  userService,
		ItemHandler: handlers.NewItemHandler(itemService),
		// UserHandler:  userHandler,
	}
}
