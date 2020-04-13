package graph

import "math"
import "testing"
import "reflect"
import "os"
import "fmt"

func TestParser(t *testing.T) {
	
	f,_ := os.Open("input.txt")
	graph := ParseFile(f)
	fmt.Println(graph)

}

func TestAjdacency(t *testing.T) {
	node1 := Node{0}
	node2 := Node{1}
	edge := UndEdge{node1, node2, 1}

	g := GraphUnd{[]Node{node1, node2}, []UndEdge{edge}}

	got := Adj_mat(g)
	want := [][]float64{{0, 1},
		{1, 0},
	}

	res := reflect.DeepEqual(want, got)
	if !res {
		t.Errorf("Ajdacency wrong")
	}

}

func TestCommunity(t *testing.T) {
	node1 := Node{0}
	node2 := Node{1}
	node3 := Node{2}
	node4 := Node{3}

	edge1 := UndEdge{node1, node2, 1}
	edge2 := UndEdge{node3, node4, 1}

	g := GraphUnd{[]Node{node1, node2, node3, node4}, []UndEdge{edge1, edge2}}

	part := map[int][]Node{
		1: {node1, node2},
		2: {node4},
		3: {node3},
	}

	community := part[1]
	sub, _ := g.SubGraph(community)
	res := sub.sumWeight()

	if res != 1 {
		t.Errorf("Wrong sum Weight; want 1 - got: %f", res)
	}

	community = part[2]
	sub, _ = g.SubGraph(community)
	res = sub.sumWeight()

	if res != 0 {
		t.Errorf("Wrong sum Weight want 0 - got: %f", res)
	}

}

func TestModularity(t *testing.T) {
	node1 := Node{0}
	node2 := Node{1}
	node3 := Node{2}
	node4 := Node{3}

	edge1 := UndEdge{node1, node2, 10000}
	edge2 := UndEdge{node3, node4, 10000}
	edge3 := UndEdge{node3, node2, 1}

	g := GraphUnd{[]Node{node1, node2, node3, node4}, []UndEdge{edge1, edge2, edge3}}

	part := map[int][]Node{
		1: {node1, node2},
		2: {node4, node3},
		3: {},
	}

	c := PartitionedGraph{g, part}
	res := c.Modularity()
	res = math.Round(res*100) / 100

	if res != 0.5 {
		t.Errorf("Wrong modularity expected 0.5 - got: %f", res)
	}
}

func TestDqAdding(t *testing.T) {
	node1 := Node{0}
	node2 := Node{1}
	node3 := Node{2}
	node4 := Node{3}

	edge1 := UndEdge{node1, node2, 1}
	edge2 := UndEdge{node3, node2, 1}
	edge3 := UndEdge{node3, node4, 1}
	edge4 := UndEdge{node4, node1, 1}
	edge5 := UndEdge{node4, node2, 1}
	edge6 := UndEdge{node3, node1, 1}

	g := GraphUnd{[]Node{node1, node2, node3, node4}, []UndEdge{edge1, edge2, edge3, edge4, edge5, edge6}}

	part := map[int][]Node{
		5: {node1, node4, node2},
		2: {node3},
	}

	part_second := map[int][]Node{
		5: {node1, node4, node2, node3},
	}

	c := PartitionedGraph{g, part}

	c_second := PartitionedGraph{g, part_second}
	modularity := c.Modularity()
	modularity_second := c_second.Modularity()

	dq, _ := c.DeltaQ(node3)

	if math.Round((modularity_second-modularity)*10000/10000) != math.Round(dq*10000/10000) {
		t.Errorf("Wrong dq")
	}

}

func TestDqAddingAndRemobe(t *testing.T) {
	node1 := Node{0}
	node2 := Node{1}
	node3 := Node{2}
	node4 := Node{3}
	node5 := Node{4}
	node6 := Node{5}

	edge1 := UndEdge{node1, node2, 4}
	edge2 := UndEdge{node3, node2, 10}
	edge3 := UndEdge{node3, node1, 6}

	edge4 := UndEdge{node4, node3, 0.02}

	edge5 := UndEdge{node4, node5, 9}
	edge6 := UndEdge{node5, node6, 6}
	edge7 := UndEdge{node6, node4, 34}

	g := GraphUnd{[]Node{node1, node2, node3, node4, node5, node6}, []UndEdge{edge1, edge2, edge3, edge4, edge5, edge6, edge7}}

	part := map[int][]Node{
		1: {node1, node2, node3, node4},
		2: {node5, node6},
	}

	part_second := map[int][]Node{
		1: {node1, node2, node3},
		2: {node5, node6, node4},
	}

	c := PartitionedGraph{g, part}
	c_second := PartitionedGraph{g, part_second}

	modularity := c.Modularity()
	modularity_second := c_second.Modularity()

	dq, _ := c.DeltaQ(node4)

	if math.Round((modularity_second-modularity)*10000/10000) != math.Round(dq*10000/10000) {
		t.Errorf("Wrong dq")
	}

}
