package main

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

func main() {

}

func Kruskal(V []int, E [][3]int) error {
	var tree [][3]int
	var comp [][]int
	sortByWeight(E)
	for len(comp) > 1 {
		e := E[0]
		v1, err := memberHasN(comp, e[0])
		if err != nil {
			return err
		}
		v2, err := memberHasN(comp, e[1])
		if err != nil {
			return err
		}
		if reflect.DeepEqual(v1, v2) {

		}
	}
}

func sortByWeight(E [][3]int) {
	sort.SliceStable(E, func(i, j int) bool { return E[i][2] < E[j][2] })
}

func memberHasN(comp [][]int, n int) ([]int, error) {
	for _, V := range comp {
		for _, v := range V {
			if v == n {
				return V, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("%d not in %v", n, comp))
}

func merge(comp [][]int, v1 []int, v2 []int) error {
	v1I :=
}

func indexOf()
