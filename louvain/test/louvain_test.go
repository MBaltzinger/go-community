package main 

import "go-community/graph"
import "os"
import "testing"

func TestDPS(t *testing.T) {

	f,_ := os.Open("../examples/karate.txt")
	g := graph.ParseFile(f)

	visited, _ := graph.Startdps(g,g.Nodes()[0])

	res := len(visited) == 34
	if !res {
		t.Errorf("DPS is wrong")
	}
}