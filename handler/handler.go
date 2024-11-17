package handler

import (
	"net/http"
	"strings"

	"github.com/Renan-Parise/graph-ifc/usecases"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	PathFinder *usecases.PathFinder
}

func NewHandler(pf *usecases.PathFinder) *Handler {
	return &Handler{PathFinder: pf}
}

func (h *Handler) FindPaths(c *gin.Context) {
	var request struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	from := strings.ToLower(request.From)
	to := strings.ToLower(request.To)

	shortestPath, cost, err := h.PathFinder.FindShortestPath(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	allPaths, costs, err := h.PathFinder.FindAllPaths(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortest_path": gin.H{
			"path": shortestPath,
			"cost": cost,
		},
		"all_paths": gin.H{
			"paths": allPaths,
			"costs": costs,
		},
	})
}
