package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// fist line represents the type of Graph U for undirectionnal edges, D for directionnal edges
// rest of the lines represents the edges
// if edges are composed of two integers, then weight is supposed to be one. If they are composed of three integers, the weight is taken into account
// node supposed to start from 1 to n
func ParseFile(r io.Reader) (G Graph) {
	scanner := bufio.NewScanner(r)
	line := 0
	slice_edge := make([]GeneralEdge, 0)
	slice_node := make([]Node, 0)
	GraphType := "U"
	for scanner.Scan() {
		if line == 0 {
			in := scanner.Text()
			GraphType = in
		} else {
			in := scanner.Text()
			slice := strings.Split(in, " ")
			n1, _ := strconv.ParseInt(slice[0], 10, 32)
			n2, _ := strconv.ParseInt(slice[1], 10, 32)
			w := 1.
			if len(slice) == 3 {
				W, _ := strconv.ParseFloat(slice[2], 64)
				w = W
			}
			node1 := Node{int(n1) - 1}
			node2 := Node{int(n2) - 1}
			var edge GeneralEdge
			if GraphType == "U" {
				edge = UndEdge{node1, node2, w}
			} else {
				edge = Edge{node1, node2, w}
			}

			slice_edge = append(slice_edge, edge)
			if !containsNode(slice_node, node1) {
				slice_node = append(slice_node, node1)
			}
			if !containsNode(slice_node, node2) {
				slice_node = append(slice_node, node2)
			}
		}
		line = 1
	}
	if GraphType == "U" {
		edge_results := make([]UndEdge, 0)
		for _, ed := range slice_edge {
			edge_results = append(edge_results, ed.(UndEdge))
		}
		return GraphUnd{slice_node, edge_results}
	} else {
		edge_results := make([]Edge, 0)
		for _, ed := range slice_edge {
			edge_results = append(edge_results, ed.(Edge))
		}
		return GraphDirect{slice_node, edge_results}
	}
}
