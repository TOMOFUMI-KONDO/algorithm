package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	tests := []struct {
		nodes []Node
		edges []Edge
		start Node
	}{
		{[]Node{"n1", "n2"}, []Edge{{"n1", "n2", 1}}, "n1"},
	}

	for _, test := range tests {
		edges := Dijkstra(test.nodes, test.edges, test.start)
		expected := []Edge{{"n1", "n2", 1}}
		if equal(edges, expected) {
			t.Errorf("edges = %#v; want %#v", edges, test.edges)
		}
	}
}

func equal(a, b []Edge) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
