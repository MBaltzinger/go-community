package graph 

import "fmt"

type PartitionedGraph struct {
	graph Graph
	Partition map[int][]Node
}

func (g PartitionedGraph) Modularity() (q float64) {

	m := 2*g.graph.sumWeight()
	adj := adj_mat(g.graph)

	for _,listnode := range g.Partition {
		for _,node := range listnode {
			for _,node2 := range listnode {

				k1 := node.degree(adj)
				k2 := node2.degree(adj)
				a := float64(k1*k2)/ (m)
				q += float64(adj[node.Id][node2.Id]) - a
			}
		}
		
	}

	return q/(m)
}

func (g PartitionedGraph) community (n Node) (i int) {
	for k,v := range g.Partition {
		for _, node := range (v) {
			if node.Id == n.Id {
				return k
			}
		}
	}
	return 0
}

func (g PartitionedGraph) NeighbouringCommunity (n Node) (c map[int][]Node) {

	c = make(map[int][]Node, len(g.graph.Nodes()))
	cn := g.community(n)
	c[cn] = append(c[cn],n)
	for _,node := range g.graph.From(n) {
		cn := g.community(node)
		c[cn] = append(c[cn],node)
	}
	return c
}

func (g PartitionedGraph) detalQ (n Node) (dq float64, dst int) {

	m:= g.graph.sumWeight()
	dqtemp := 0.
	dqadd := 0.
	adj := adj_mat(g.graph)
	dqremove := 0.
	for nb, community := range g.NeighbouringCommunity(n) {

		newGraph, ext := g.graph.SubGraph(community)
		SumIn := newGraph.sumWeight()
		SumTot := SumIn
		for _,e := range ext {
			SumTot += e.Weight()
		}

		ki := n.degree(adj)
		kiin := 0.
		for _, node_n := range community {
			// require input node list to be ordered and incremental
			kiin += adj[n.Id][node_n.Id]
		}

		if nb != g.community(n) {

				dqtemp = (kiin - (SumTot*ki)/(2*m))/(m)
				fmt.Println("dqadd", dqtemp)
				if (dqtemp > dqadd) {
					dqadd = dqtemp
					dst = nb
				}

		} else {
			fmt.Println("kiin", kiin)
			fmt.Println("ki", ki)
			fmt.Println("SumTot", SumTot)
			fmt.Println("adj[n.Id][n.Id]", adj[n.Id][n.Id])

			dqremove = (kiin - ki*(SumTot-ki)/(2*m) - adj[n.Id][n.Id])/(m)
			fmt.Println("dqremove", dqremove)
		}
	}

	return dqadd - dqremove , dst
}




