package graph

type PartitionedGraph struct {
	graph     Graph
	Partition map[int][]Node
}

func (g PartitionedGraph) Modularity() (q float64) {

	m := 2 * g.graph.sumWeight()
	adj := adj_mat(g.graph)

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

func (g PartitionedGraph) community(n Node) (i int) {
	for k, v := range g.Partition {
		for _, node := range v {
			if node.Id == n.Id {
				return k
			}
		}
	}
	return 0
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

	cn := g.community(n)
	nb = append(nb, cn)
	for _, node := range g.graph.From(n) {
		cn := g.community(node)
		if !(contains(nb, cn)) {
			nb = append(nb, cn)
		}
	}
	return nb
}

func (g PartitionedGraph) detalQ(n Node) (dq float64, dst int) {

	m := g.graph.sumWeight()
	dqtemp := 0.
	dqadd := 0.
	adj := adj_mat(g.graph)
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

		if nb != g.community(n) {

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
