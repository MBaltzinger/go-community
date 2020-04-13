package graph

import (
	"bufio"
	"strings"
	"io"
	"strconv"
)

// fist line represents the type of Graph U for undirectionnal edges, D for directionnal edges
// rest of the lines represents the edges
// if edges are composed of two integers, then weight is supposed to be one. If they are composed of three integers, the weight is taken into account
// node supposed to start from 1 to n
func ParseFile(r io.Reader) (G Graph){
	scanner := bufio.NewScanner(r)
	line :=0
	slice_edge := make([]UndEdge,0)
	slice_node := make([]Node,0)
	GraphType := 'U'
	for scanner.Scan() {
		if line == 0 {
			in := scanner.Text()
			GraphType := in
			if GraphType == "U" {
				slice_edge = make([]UndEdge,0)
			} else {
				slice_edge = make([]UndEdge,0)
			}
		} else {
			in := scanner.Text()
			slice := strings.Split(in, " ")
			n1,_ := strconv.ParseInt(slice[0],10,32)
			n2,_ := strconv.ParseInt(slice[1],10,32)
			w :=1.
			if len(slice) == 3 {
				W,_ := strconv.ParseFloat(slice[2], 64)
				w = W
			}
			node1 := Node{int(n1)-1}
			node2 := Node{int(n2)-1}
			edge := UndEdge{node1, node2, w}
			slice_edge = append(slice_edge,edge)
			if !containsNode(slice_node,node1){
				slice_node = append(slice_node,node1)
			}
			if !containsNode(slice_node,node2){
				slice_node = append(slice_node,node2)
			}
		}
		line +=1
	}
	if GraphType == 'U' {
		return GraphUnd{slice_node,slice_edge}
	} else {
		return GraphUnd{slice_node,slice_edge}
	}
}