// Temporary package to run own community detection algorithm
package graph

type Graph interface {
	// Nodes return all nodes of the graph
	Nodes() []Node

	// Edges retun all edges of the graph
	Edges() []GeneralEdge

	sumWeight() float64

	From(n Node) ([]Node, []GeneralEdge)

	SubGraph(l []Node) (graph Graph, ext []GeneralEdge)
}

type GeneralEdge interface {
	NodesIn() Node
	NodesOut() Node
	Weight() float64
}

type Edge struct {
	NodesFrom Node
	NodesTo   Node
	W         float64
}

func (e Edge) NodesIn() (n Node) {
	return e.NodesFrom
}

func (e Edge) NodesOut() (n Node) {
	return e.NodesTo
}

func (e Edge) Weight() (i float64) {
	return e.W
}

type Node struct {
	Id int
}

func containsNode(s []Node, e Node) bool {
	for _, a := range s {
		if a.Id == e.Id {
			return true
		}
	}
	return false
}

func containsEdge(s []GeneralEdge, e GeneralEdge) bool {
	for _, a := range s {
		if (a.NodesIn().Id == e.NodesIn().Id) && (a.NodesOut().Id == e.NodesOut().Id) {
			return true
		}
	}
	return false
}

func Adj_mat(g Graph) [][]float64 {

	adj := make([][]float64, len(g.Nodes()))

	for _, node := range g.Nodes() {

		adj[node.Id] = make([]float64, len(g.Nodes()))

		// doesn't handle self in edges
		for _, edge := range g.Edges() {

			if node.Id == edge.NodesIn().Id {

				for _, node2 := range g.Nodes() {

					if node2.Id == edge.NodesOut().Id {
						adj[node.Id][node2.Id] = edge.Weight()
					}
				}
			}
		}
	}
	return adj
}

func (n *Node) degree(adj [][]float64) (k float64) {
	for nodej, _ := range adj[n.Id] {
		if adj[n.Id][nodej] != 0 {
			k += adj[n.Id][nodej]
		}
	}
	return k
}
