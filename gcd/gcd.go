package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input a0 a1 > ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	textArray := strings.Fields(text)
	a0, err := strconv.Atoi(textArray[0])
	if err != nil {
		panic(err)
	}
	a1, err := strconv.Atoi(textArray[1])
	if err != nil {
		panic(err)
	}
	if a0 == 0 || a1 == 0 {
		fmt.Println("a0 and a1 must not be 0.")
		return
	}
	if a0 > a1 {
		a0, a1 = a1, a0
	}
	fmt.Printf("gcd of %d and %d is %d.", a0, a1, gcd(a0, a1))
}

func gcd(a0 int, a1 int) int {
	surplus := a0 % a1
	for surplus != 0 {
		a0, a1 = a1, surplus
		surplus = a0 % a1
	}
	return a1
}
