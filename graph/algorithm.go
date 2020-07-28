package graph

func alreadyVisited(visited []Node, node Node) bool {
	for _, nodeTemp := range visited {
		if nodeTemp.Id == node.Id {
			return true
		}
	}
	return false
}

func dps(g Graph, stack *[]Node, visited *[]Node, parent *map[Node]Node) {
	origin := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	list_node, _ := g.From(origin)

	for _, node := range list_node {
		if !(alreadyVisited(*visited, node)) {

			(*visited) = append((*visited), node)
			(*stack) = append((*stack), node)
			(*parent)[node] = origin
			dps(g, stack, visited, parent)

		}
	}

}

func Startdps(g Graph, origin Node) (visited []Node, parent map[Node]Node) {
	stack := []Node{origin}
	visited = append(visited, origin)
	parent = make(map[Node]Node)
	parent[origin] = Node{-1}

	dps(g, &stack, &visited, &parent)

	return visited, parent
}
