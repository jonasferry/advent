package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func first(bxs []string) {
	allwp := 0 // total wrapping paper

	for _, bx := range bxs {

		in := strings.Split(string(bx), "x") // in size

		var ls [3]int // length sides

		for i, v := range in {
			ls[i], _ = strconv.Atoi(v) // convert to int
		}

		l, w, h := ls[0], ls[1], ls[2] // length, width, height

		s := []int{2 * l * w, 2 * w * h, 2 * h * l}
		sort.Ints(s)

		wp := s[0] + s[0]/2 + s[1] + s[2] // wrapping paper

		allwp += wp
	}
	fmt.Println("Total wrapping paper: " + strconv.Itoa(allwp))
}

func second(d []string) {
}

func main() {
	var bxs []string
	part, e := strconv.Atoi(os.Args[1]) //part 1 or 2
	if e != nil {
		panic(e)
	}

	f, e := os.Open(os.Args[2]) //open file
	if e != nil {
		panic(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bxs = append(bxs, scanner.Text())
	}

	if e := scanner.Err(); e != nil {
		panic(e)
	}

	if part == 1 {
		first(bxs)
	} else if part == 2 {
		second(bxs)
	} else {
		fmt.Println("Error: Select part 1 or 2")
	}
}
