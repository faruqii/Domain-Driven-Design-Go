package routes

import "github.com/faruqii/Domain-Driven-Design-Go/internal/app/handlers"

func SetupMovieRoutes(r *gin.Engine, h *handlers.MovieHandler) {
	r.GET("/movies", h.GetMovies)
	r.GET("/movies/:id", h.GetMovie)
	r.POST("/movies", h.CreateMovie)
	r.PUT("/movies/:id", h.UpdateMovie)
	r.DELETE("/movies/:id", h.DeleteMovie)
}
