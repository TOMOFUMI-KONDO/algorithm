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
	ans, err := Kruskal(V, E)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ans)
}

func Kruskal(V []int, E [][3]int) ([][3]int, error) {
	var tree [][3]int
	var comp [][]int
	for _, v := range V {
		comp = append(comp, []int{v})
	}
	sortByWeight(E)
	for len(comp) > 1 {
		//fmt.Printf("%v,%v\n", E, comp)
		e := E[0]
		i0, err := memberIdxHasN(comp, e[0])
		if err != nil {
			return nil, err
		}
		i1, err := memberIdxHasN(comp, e[1])
		if err != nil {
			return nil, err
		}
		//fmt.Printf("%d,%d\n", i0, i1)
		//fmt.Printf("%v,%v\n", comp[i0], comp[i1])
		if !reflect.DeepEqual(comp[i0], comp[i1]) {
			merge(comp, i0, i1)
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

func merge(comp [][]int, i0 int, i1 int) {
	v0 := comp[i0]
	v1 := comp[i1]
	comp = append(comp[:i0], comp[:i0+1]...)
	comp = append(comp[:i1], comp[:i1+1]...)
	comp = append(comp, append(v0, v1...))
}
