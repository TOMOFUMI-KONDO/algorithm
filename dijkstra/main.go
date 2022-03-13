package dijkstra

import "fmt"

type Node uint
type Cost uint
type Edge struct {
	Node1 Node
	Node2 Node
	Cost
}

func (e Edge) opposite(n Node) Node {
	if n == e.Node1 {
		return e.Node2
	}
	if n == e.Node2 {
		return e.Node1
	}
	panic(fmt.Errorf("node %d is not included in edge %#v", n, e))
}

const MaxCost = ^Cost(0)

func Dijkstra(nodes []Node, edges []Edge, src Node, dst Node) []Edge {
	visited := make(map[Node]bool, len(nodes))
	for _, v := range nodes {
		visited[v] = false
	}
	visited[src] = true

	cost := make(map[Node]Cost, len(nodes))
	for _, v := range nodes {
		cost[v] = MaxCost
	}
	cost[src] = 0

	neighbor := make(map[Node][]struct {
		node Node
		edge Edge
	}, len(nodes))
	for _, v := range edges {
		neighbor[v.Node1] = append(neighbor[v.Node1], struct {
			node Node
			edge Edge
		}{v.Node2, v})
		neighbor[v.Node2] = append(neighbor[v.Node2], struct {
			node Node
			edge Edge
		}{v.Node1, v})
	}

	edgeToNode := make(map[Node]Edge)
	lastVisited := src
	var tree []Edge

	for !visited[dst] {
		// update neighbor costs
		for _, v := range neighbor[lastVisited] {
			if visited[v.node] {
				continue
			}
			if cost[lastVisited]+v.edge.Cost < cost[v.node] {
				cost[v.node] = cost[lastVisited] + v.edge.Cost
				edgeToNode[v.node] = v.edge
			}
		}

		var minNode Node
		minCost := MaxCost
		for n, c := range cost {
			if !visited[n] && c < minCost {
				minNode = n
				minCost = c
			}
		}
		lastVisited = minNode
		visited[minNode] = true
		tree = append(tree, edgeToNode[minNode])
	}

	var path []Edge
	cur := dst
	for {
		path = append([]Edge{edgeToNode[cur]}, path...)
		oppo := edgeToNode[cur].opposite(cur)
		if oppo == src {
			break
		}
		cur = oppo
	}

	return path
}
