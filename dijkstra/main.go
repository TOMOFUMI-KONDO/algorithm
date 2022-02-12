package dijkstra

type Node uint
type Cost uint
type Edge struct {
	Node1 Node
	Node2 Node
	Cost
}

const MaxCost = ^Cost(0)

func (e Edge) opposite(n Node) (*Node, bool) {
	if n == e.Node1 {
		return &e.Node2, true
	}
	if n == e.Node2 {
		return &e.Node1, true
	}
	return nil, false
}

func Dijkstra(nodes []Node, edges []Edge, start Node) []Edge {
	p := map[Node]bool{start: true}
	for _, v := range nodes {
		if v != start {
			p[v] = false
		}
	}
	e := make(map[Node]Edge)
	d := map[Node]Cost{start: 0}
	for _, v := range nodes {
		if v != start {
			d[v] = MaxCost
		}
	}
	u := start
	var t []Edge

	for len(p) != len(nodes) {
		for _, v := range edges {
			n, ok := v.opposite(u)
			if ok && !p[*n] {
				if d[u]+v.Cost < d[*n] {
					d[*n] = d[u] + v.Cost
					e[*n] = v
				}
			}
		}
		var minNode Node
		minCost := MaxCost
		for k, v := range d {
			if !p[k] && v < minCost {
				minNode = k
				minCost = v
			}
		}
		p[minNode] = true
		t = append(t, e[minNode])
	}

	return t
}
