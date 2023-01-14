package cmd

import (
	"github.com/faruqii/Domain-Driven-Design-Go/infrastructure/repository"
	"github.com/faruqii/Domain-Driven-Design-Go/internal/app/handlers"
	"github.com/faruqii/Domain-Driven-Design-Go/internal/app/routes"
	"github.com/faruqii/Domain-Driven-Design-Go/internal/domain/services"
)

func main() {
	r := gin.Default()

	// Repositories
	movieRepo := repository.MovieRepository()

	// Services
	movieService := services.MovieService(movieRepo)

	// Handlers
	movieHandler := handlers.MovieHandler(movieService)

	// Routes
	routes.SetupMovieRoutes(r, movieHandler)

	// Start the server
	server.Start(r)
}
