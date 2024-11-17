package main

import (
	"github.com/Renan-Parise/graph-ifc/handler"
	"github.com/Renan-Parise/graph-ifc/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	graph := usecases.NewGraph()

	pathFinder := usecases.NewPathFinder(graph)

	handler := handler.NewHandler(pathFinder)

	r := gin.Default()
	r.POST("/findpaths", handler.FindPaths)
	r.Run("127.0.0.1:8181")
}
