package graph

type PartitionedGraph struct {
	Graph     Graph
	Partition map[int][]Node
}

type Hierachical struct {
	Graph   PartitionedGraph
	PartMap map[Node][]Node
	Parent  *Hierachical
}

func (g Hierachical) CreateHierachical() Hierachical {

	Nodes := make([]Node, 0)
	PartMap := make(map[Node][]Node)

	ReverseComm := make(map[Node]Node)
	i := 0
	newPartition := make(map[int][]Node)
	for _, list_node := range g.Graph.Partition {

		PartMap[Node{i}] = list_node
		Nodes = append(Nodes, Node{i})
		newPartition[i] = []Node{Node{i}}
		for _, node := range list_node {
			ReverseComm[node] = Node{i}
		}
		i += 1
	}
	var gt Graph

	switch g.Graph.Graph.(type) {
	default:
		Edges := make([]UndEdge, 0)
		for _, list_node := range g.Graph.Partition {
			for _, node := range list_node {
				_, list_edges := g.Graph.Graph.From(node)
				for _, edge := range list_edges {

					EdgeTemp := UndEdge{
						NodesFrom: ReverseComm[node],
						NodesTo:   ReverseComm[edge.NodesOut()],
						W:         edge.Weight() / 2,
					}

					update := false
					for i, NewEdge := range Edges {
						if ((NewEdge.NodesIn().Id == ReverseComm[node].Id) && (ReverseComm[edge.NodesOut()].Id == NewEdge.NodesOut().Id)) || ((NewEdge.NodesOut().Id == ReverseComm[node].Id) && (ReverseComm[edge.NodesOut()].Id == NewEdge.NodesIn().Id)) {
							NewEdge.W += EdgeTemp.W
							Edges[i].W = NewEdge.W
							update = true
						}
					}

					if !update {
						Edges = append(Edges, EdgeTemp)
					}
				}
			}
		}

		gt = GraphUnd{Nodes, Edges}

	case GraphDirect:
		Edges := make([]Edge, 0)
		for _, list_node := range g.Graph.Partition {
			for _, node := range list_node {
				_, list_edges := g.Graph.Graph.From(node)
				for _, edge := range list_edges {

					EdgeTemp := Edge{
						NodesFrom: ReverseComm[node],
						NodesTo:   ReverseComm[edge.NodesOut()],
						W:         edge.Weight(),
					}

					update := false
					for i, NewEdge := range Edges {
						if (NewEdge.NodesIn().Id == ReverseComm[node].Id) && (ReverseComm[edge.NodesOut()].Id == NewEdge.NodesOut().Id) {
							NewEdge.W += EdgeTemp.W
							Edges[i].W = NewEdge.W
							update = true
						}
					}

					if !update {
						Edges = append(Edges, EdgeTemp)
					}
				}
			}
		}

		gt = GraphDirect{Nodes, Edges}
	}

	return Hierachical{PartitionedGraph{gt, newPartition}, PartMap, &g}
}

func (g PartitionedGraph) Modularity() (q float64) {

	m := 2 * g.Graph.sumWeight()
	adj := Adj_mat(g.Graph)

	for _, listnode := range g.Partition {
		for _, node := range listnode {
			for _, node2 := range listnode {

				k1 := node.degree(adj)
				k2 := node2.degree(adj)
				a := float64(k1*k2) / (m)
				q += float64(adj[node.Id][node2.Id]) - a
			}
		}

	}

	return q / (m)
}

func (g PartitionedGraph) Community(n Node) (i int) {
	for k, v := range g.Partition {
		for _, node := range v {
			if node.Id == n.Id {
				return k
			}
		}
	}
	// shouldn't happen
	return 100000
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (g PartitionedGraph) NeighbouringCommunity(n Node) (nb []int) {

	cn := g.Community(n)
	nb = append(nb, cn)
	list_node, _ := g.Graph.From(n)
	for _, node := range list_node {
		cn := g.Community(node)
		if !(contains(nb, cn)) {
			nb = append(nb, cn)
		}
	}
	return nb
}

func (g PartitionedGraph) DeltaQ(n Node) (dq float64, dst int) {

	m := g.Graph.sumWeight()
	dqtemp := 0.
	dqadd := 0.
	adj := Adj_mat(g.Graph)
	dqremove := 0.
	for _, nb := range g.NeighbouringCommunity(n) {

		community := g.Partition[nb]
		SumTot := 0.
		ki := n.degree(adj)
		kiin := 0.

		for _, node_n := range community {
			SumTot += node_n.degree(adj)
			kiin += adj[n.Id][node_n.Id]
		}

		if nb != g.Community(n) {

			dqtemp = (kiin - (SumTot*ki)/(2*m)) / (m)
			if dqtemp > dqadd {
				dqadd = dqtemp
				dst = nb
			}

		} else {
			dqremove = (kiin - ki*(SumTot-ki)/(2*m) - adj[n.Id][n.Id]) / (m)

		}
	}

	return dqadd - dqremove, dst
}
