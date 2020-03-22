package graph

import "testing"

//import "reflect"
import "fmt"

// func TestGraph(t *testing.T) {
// 	node1 := Node{0}
// 	node2 := Node{1}
// 	edge := UndEdge{node1, node2,1}

// 	g := GraphUnd{[]Node{node1, node2}, []UndEdge{edge}}

// 	got := adj_mat(g)
// 	want := [][]int{{0, 1},
// 		{1, 0},
// 	}

// 	fmt.Println(reflect.DeepEqual(want, got))
// 	fmt.Println(node1.degree(got))

// }

// func TestCommunity(t *testing.T) {
// 	node1 := Node{0}
// 	node2 := Node{1}
// 	node3 := Node{2}
// 	node4 := Node{3}

// 	edge1 := UndEdge{node1, node2,1}
// 	edge2 := UndEdge{node3, node4,1}

// 	g := GraphUnd{[]Node{node1, node2,node3,node4}, []UndEdge{edge1,edge2}}

// 	fmt.Println("Adj:",adj_mat(g))

// 	part := map[int][]Node{
// 		1: {node1,node2},
// 		2: {node4},
// 		3: {node3},
// 	}

// 	c := PartitionedGraphUnd{g, part}

// 	fmt.Println("Modularity:",c.Modularity())
// 	fmt.Println("Expected:",0.125)

// 	community := part[1]
// 	sub,_ := SubEdges(g, community)
// 	fmt.Println("Sum weights community 1:",sumWeight(GraphUnd{community,sub}))

// 	community = part[2]
// 	sub,_ = SubEdges(g, community)
// 	fmt.Println("Sum weights community 1:",sumWeight(GraphUnd{community,sub}))

// 	community = part[3]
// 	sub,_ = SubEdges(g, community)
// 	fmt.Println("Sum weights community 1:",sumWeight(GraphUnd{community,sub}))

// }

// func TestCommunity2(t *testing.T) {
// 	node1 := Node{0}
// 	node2 := Node{1}
// 	node3 := Node{2}
// 	node4 := Node{3}

// 	edge1 := UndEdge{node1, node2,10000}
// 	edge2 := UndEdge{node3, node4,10000}
// 	edge3 := UndEdge{node3, node2,1}

// 	g := GraphUnd{[]Node{node1, node2,node3,node4}, []UndEdge{edge1,edge2,edge3}}

// 	fmt.Println("Adj:",adj_mat(g))

// 	part := map[int][]Node{
// 		1: {node1,node2},
// 		2: {node4,node3},
// 		3: {},
// 	}

// 	c := PartitionedGraphUnd{g, part}

// 	fmt.Println("Modularity:",c.Modularity())
// 	fmt.Println("Expected:",0.5)

// 	community := part[1]
// 	sub,_ := SubEdges(g, community)
// 	fmt.Println("Sum weights community 1:",sumWeight(GraphUnd{community,sub}))

// 	community = part[2]
// 	sub,ext := SubEdges(g, community)
// 	fmt.Println("Sum weights community 1:",sumWeight(GraphUnd{community,sub}))
// 	fmt.Println(" links ext:",ext)

// 	community = part[3]
// 	sub,_ = SubEdges(g, community)
// 	fmt.Println("Sum weights community 1:",sumWeight(GraphUnd{community,sub}))

// }

func TestDq(t *testing.T) {
	node1 := Node{0}
	node2 := Node{1}
	node3 := Node{2}
	node4 := Node{3}
	node5 := Node{4}
	node6 := Node{5}

	edge1 := UndEdge{node1, node2, 10000}
	edge2 := UndEdge{node3, node2, 10000}
	edge3 := UndEdge{node3, node4, 10000}
	edge4 := UndEdge{node4, node1, 10000}
	edge5 := UndEdge{node4, node2, 10000}
	edge6 := UndEdge{node5, node4, 1}
	edge7 := UndEdge{node5, node6, 1}

	g := GraphUnd{[]Node{node1, node2, node3, node4, node5,node6}, []UndEdge{edge1, edge2, edge3, edge4, edge5, edge6,edge7}}

	fmt.Println("Adj:", adj_mat(g))

	part := map[int][]Node{
		5: {node1, node4, node2},
		2: {node3},
		3: {node5},
		4: {node6},
	}

	c := PartitionedGraph{g, part}

	fmt.Println("Modularity:", c.Modularity())

	dq, dst := c.detalQ(node3)

	fmt.Println("Gain by moving:", node3," is:", dq, "Destination : ", dst)

}
