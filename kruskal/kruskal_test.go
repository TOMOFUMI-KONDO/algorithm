package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKruskal(t *testing.T) {
	V := []int{0, 1, 2, 3, 4}
	E := [][3]int{{0, 1, 7}, {3, 4, 8}, {0, 4, 9}, {1, 4, 12}, {2, 4, 15}, {2, 3, 16}, {1, 2, 16}}

	ans, err := Kruskal(V, E)
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}
	fmt.Println(ans)
	if !reflect.DeepEqual(ans, [][3]int{{0, 1, 7}, {3, 4, 8}, {0, 4, 9}, {2, 4, 15}}) {
		t.Errorf("ans = %v; want [[0 1 7] [3 4 8] [0 4 9] [2 4 15]]", ans)
	}
}
