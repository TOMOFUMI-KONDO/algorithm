package dijkstra

type Node string
type Cost uint8
type Edge struct {
	Node1 Node
	Node2 Node
	Cost
}

const max Cost = ^Cost(0)

func Dijkstra(nodes []Node, edges []Edge, start Node) []Edge {
	fixed := []Node{start}
	var edge_to_node map[Node]Edge
	costs := map[Node]Cost{start: 0}
	for _, n := range nodes {
		if n != start {
			costs[n] = max
		}
	}

	for len(fixed) != len(nodes) {

	}
}
