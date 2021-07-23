package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
)

func main() {
	V := []int{0, 1, 2, 3, 4}
	E := [][3]int{{0, 1, 7}, {3, 4, 8}, {0, 4, 9}, {1, 4, 12}, {2, 4, 15}, {2, 3, 16}, {1, 2, 16}}
	fmt.Printf("V: %v\n", V)
	fmt.Printf("E: %v\n", E)
	ans, err := Kruskal(V, E)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MST is %v\n", ans)
}

func Kruskal(V []int, E [][3]int) ([][3]int, error) {
	var tree [][3]int
	var comp [][]int
	for _, v := range V {
		comp = append(comp, []int{v})
	}
	sortByWeight(E)
	for len(comp) > 1 {
		e := E[0]
		i0, err := memberIdxHasN(comp, e[0])
		if err != nil {
			return nil, fmt.Errorf("i0:%w", err)
		}
		i1, err := memberIdxHasN(comp, e[1])
		if err != nil {
			return nil, fmt.Errorf("i0:%w", err)
		}
		if !reflect.DeepEqual(comp[i0], comp[i1]) {
			comp = merged(comp, i0, i1)
			tree = append(tree, e)
		}
		E = E[1:]
	}
	return tree, nil
}

func sortByWeight(E [][3]int) {
	sort.SliceStable(E, func(i, j int) bool { return E[i][2] < E[j][2] })
}

func memberIdxHasN(comp [][]int, n int) (int, error) {
	for i, V := range comp {
		for _, v := range V {
			if v == n {
				return i, nil
			}
		}
	}
	return -1, errors.New(fmt.Sprintf("[ERROR] %d not in %v", n, comp))
}

func merged(comp [][]int, i0 int, i1 int) [][]int {
	var newComp [][]int
	for i := range comp {
		if i != i0 && i != i1 {
			newComp = append(newComp, comp[i])
		}
	}
	var c []int
	c = append(c, comp[i0]...)
	c = append(c, comp[i1]...)
	newComp = append(newComp, c)

	return newComp
}
