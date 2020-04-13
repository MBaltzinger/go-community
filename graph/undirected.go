package graph

// define undirected graph implements graph
type GraphUnd struct {
	Node []Node

	// UndEdge means it is bi-directionnal
	Edge []UndEdge
}

// Implement general Edge
type UndEdge struct {
	NodesFrom Node
	NodesTo   Node
	W         float64
}

func (e UndEdge) NodesIn() (n Node) {
	return e.NodesFrom
}

func (e UndEdge) NodesOut() (n Node) {
	return e.NodesTo
}

func (e UndEdge) Weight() (i float64) {
	return e.W
}

func (g GraphUnd) sumWeight() (sum float64) {
	for _, e := range g.Edges() {
		sum += e.Weight() / 2
	}
	return sum
}

func (e UndEdge) toEdge() (edge1, edge2 Edge) {
	return Edge{e.NodesIn(), e.NodesOut(), e.Weight()}, Edge{e.NodesOut(), e.NodesIn(), e.Weight()}
}

func (g GraphUnd) Edges() (edges []GeneralEdge) {
	for _, e := range g.Edge {
		e1, e2 := e.toEdge()
		edges = append(edges, e1, e2)
	}
	return edges
}

func (g GraphUnd) Nodes() []Node {
	return g.Node
}

func (g GraphUnd) From(n Node) (l []Node, le []GeneralEdge) {
	for _, e := range g.Edges() {
		if e.NodesIn().Id == n.Id {
			l = append(l, e.NodesOut())
			le = append(le, e)
		}
	}
	return l, le
}

func (g GraphUnd) SubGraph(l []Node) (newg Graph, ext []GeneralEdge) {
	ledge := make([]UndEdge, 0)
	for _, edge := range g.Edge {
		if containsNode(l, edge.NodesIn()) && containsNode(l, edge.NodesOut()) {
			ledge = append(ledge, edge)
		}
		if (containsNode(l, edge.NodesIn()) && !containsNode(l, edge.NodesOut())) || (!containsNode(l, edge.NodesIn()) && containsNode(l, edge.NodesOut())) {
			ext = append(ext, edge)
		}
	}

	return GraphUnd{l, ledge}, ext
}
