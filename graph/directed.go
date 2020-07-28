package graph

// define undirected graph implements graph
type GraphDirect struct {
	Node []Node

	// UndEdge means it is bi-directionnal
	Edge []Edge
}

func (g GraphDirect) sumWeight() (sum float64) {
	for _, e := range g.Edges() {
		sum += e.Weight()
	}
	return sum
}

func (g GraphDirect) Edges() (edges []GeneralEdge) {
	for _, e := range g.Edge {
		edges = append(edges, e)
	}
	return edges
}

func (g GraphDirect) Nodes() []Node {
	return g.Node
}

func (g GraphDirect) From(n Node) (l []Node, le []GeneralEdge) {
	for _, e := range g.Edges() {
		if e.NodesIn().Id == n.Id {
			l = append(l, e.NodesOut())
			le = append(le, e)
		}
	}
	return l, le
}

func (g GraphDirect) SubGraph(l []Node) (newg Graph, ext []GeneralEdge) {
	ledge := make([]Edge, 0)
	for _, edge := range g.Edge {
		if containsNode(l, edge.NodesIn()) && containsNode(l, edge.NodesOut()) {
			ledge = append(ledge, edge)
		}
		if (containsNode(l, edge.NodesIn()) && !containsNode(l, edge.NodesOut())) || (!containsNode(l, edge.NodesIn()) && containsNode(l, edge.NodesOut())) {
			ext = append(ext, edge)
		}
	}

	return GraphDirect{l, ledge}, ext
}
