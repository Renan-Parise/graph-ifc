package usecases

import "github.com/Renan-Parise/graph-ifc/entities"

func NewGraph() *entities.Graph {
	graph := entities.NewGraph()

	places := []string{
		"guarita",
		"administrativo",
		"extensão",
		"almoxarifado",
		"salas de aula",
		"ginásio",
		"refeitório",
		"academia",
		"auditório",
		"biblioteca",
	}

	for _, place := range places {
		graph.AddPlace(place)
	}

	graph.AddEdge("guarita", "administrativo", 1)
	graph.AddEdge("guarita", "extensão", 2.5)
	graph.AddEdge("guarita", "salas de aula", 3)
	graph.AddEdge("administrativo", "salas de aula", 2)
	graph.AddEdge("extensão", "salas de aula", 0.5)
	graph.AddEdge("salas de aula", "biblioteca", 2)
	graph.AddEdge("salas de aula", "ginásio", 4)
	graph.AddEdge("biblioteca", "auditório", 1.5)
	graph.AddEdge("auditório", "refeitório", 3)
	graph.AddEdge("ginásio", "academia", 0.5)
	graph.AddEdge("refeitório", "academia", 2)
	graph.AddEdge("academia", "almoxarifado", 5)
	graph.AddEdge("almoxarifado", "ginásio", 5.5)

	return graph
}
