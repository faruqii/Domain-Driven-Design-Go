package handlers

import (
	"github.com/faruqii/Domain-Driven-Design-Go/internal/domain/entities"
	"github.com/faruqii/Domain-Driven-Design-Go/internal/domain/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service services.MovieService
}

func (h *MovieHandler) GetMovie(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	movie, err := h.service.GetMovie(string(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) GetMovies(c *gin.Context) {
	movies, err := h.service.GetMovies()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie entities.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	var movie entities.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = string(idInt)
	if err := h.service.UpdateMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if err := h.service.DeleteMovie(string(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
