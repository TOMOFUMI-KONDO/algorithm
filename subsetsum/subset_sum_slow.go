package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type YN string

const (
	Y YN = "Yes"
	N YN = "No"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var scanned []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scanned = append(scanned, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var a []int
	fmt.Print("a:")
	for _, v := range strings.Fields(scanned[0]) {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, n)
		fmt.Printf(" %d", n)
	}
	fmt.Println()
	b, err := strconv.Atoi(scanned[1])
	fmt.Printf("b: %d\n", b)
	if err != nil {
		log.Fatal(err)
	}

	yn, subset := SubsetSumSlow(a, b)
	fmt.Printf("YN: %s\n", yn)
	fmt.Printf("Subset: %v", subset)
}

func SubsetSumSlow(a []int, b int) (YN, []bool) {
	lenA := len(a)
	var x []bool
	for i := 0; i < lenA; i++ {
		x = append(x, false)
	}

	for {
		tmp := 0
		for i := 0; i < lenA; i++ {
			if x[i] {
				tmp += a[i]
			}
		}
		if tmp == b {
			return Y, x
		}

		isFull := true
		for i := 0; i < lenA; i++ {
			isFull = x[i]
			if !isFull {
				break
			}
		}
		if isFull {
			return N, nil
		}
		next(x)
	}
}

func next(x []bool) {
	i := len(x) - 1

	for {
		if x[i] {
			x[i] = false
		} else {
			x[i] = true
			break
		}

		i--
	}
}
