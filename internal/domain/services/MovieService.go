package services

import "github.com/faruqii/Domain-Driven-Design-Go/internal/domain/entities"

type MovieService interface {
	GetMovie(id string) (*entities.Movie, error)
	GetMovies() ([]entities.Movie, error)
	CreateMovie(movie *entities.Movie) error
	UpdateMovie(movie *entities.Movie) error
	DeleteMovie(id string) error
}
