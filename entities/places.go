package entities

type Place struct {
	Name string
}

type Edge struct {
	To       string
	Distance float64
}

type Graph struct {
	Nodes map[string]*Place
	Edges map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Place),
		Edges: make(map[string][]Edge),
	}
}

func (g *Graph) AddPlace(name string) {
	g.Nodes[name] = &Place{Name: name}
}

func (g *Graph) AddEdge(from, to string, distance float64) {
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Distance: distance})
	g.Edges[to] = append(g.Edges[to], Edge{To: from, Distance: distance})
}
