package main

import "go-community/graph"
import "fmt"
import "os"

// This implements the louvain algorithm, checkout : https://perso.uclouvain.be/vincent.blondel/research/louvain.html
func remove(s []graph.Node, i int) []graph.Node {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RunLouvain (g *graph.Hierachical) (*graph.Hierachical) {
	if movingNodes(&g.Graph) {
		h := g.CreateHierachical()
		return RunLouvain(&h)
	} else {
		return g
	}
}

func UnpackLouvain(h *graph.Hierachical) (map[int][]graph.Node) {

	for {
		mapping := h.PartMap
		if (mapping != nil) {
			
			partNew := make(map[int][]graph.Node)
			
			for nb,list_node := range h.Graph.Partition {
				for _, node := range list_node {
					for _, node2 := range h.PartMap[node] {
						partNew[nb] = append(partNew[nb], node2)
					}
				}
			}

			h.Parent.Graph.Partition = partNew
			hold := graph.Hierachical{
				h.Parent.Graph,
				h.Parent.PartMap,
				h.Parent.Parent,
			}

			return UnpackLouvain(&hold)
		} else {
			return h.Graph.Partition
		}
	}

}

func movingNodes(g *graph.PartitionedGraph) bool {

	moved := false

	for _, n := range g.Graph.Nodes() {
		dq, dst := g.DeltaQ(n)
		if dq <= 0 {
			continue
		}

		list_node := g.Partition[g.Community(n)]
		l_temp := make([]graph.Node, len(list_node))
		copy(l_temp, list_node)
		for i, node := range l_temp {
			if node.Id == n.Id {
				l_temp = remove(l_temp, i)
				break
			}
		}

		g.Partition[g.Community(n)] = l_temp
		g.Partition[dst] = append(g.Partition[dst], n)

		moved = true
	}

	final_partition := map[int][]graph.Node{}
	nb_comm := 0

	for _,value := range (g.Partition) {
		if len(value) > 0 {
			nb_comm +=1
			final_partition[nb_comm] = value
		  }
	}

	g.Partition = final_partition

	return moved

}

func main() {

	argsWithoutProg := os.Args[1:]

	f,_ := os.Open(argsWithoutProg[0])
	g := graph.ParseFile(f)

	part_test := make(map[int][]graph.Node)
	nb_comm := 0
	for _,n := range g.Nodes() {
		nb_comm +=1
		part_test[nb_comm] = []graph.Node{n}
	}

	c_test := graph.PartitionedGraph{
		Graph:     g,
		Partition: part_test}

	h := graph.Hierachical{c_test,nil,nil}

	fH := RunLouvain(&h)

	partition := UnpackLouvain(fH)

	fmt.Println(partition)
}