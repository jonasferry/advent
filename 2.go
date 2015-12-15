package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := strings.Split(os.Args[1], "x") // in size

	var ls [3]int // length sides

	for i, v := range in {
		ls[i], _ = strconv.Atoi(v) // convert to int
	}

	l, w, h := ls[0], ls[1], ls[2] // length, width, height

	s := []int{2 * l * w, 2 * w * h, 2 * h * l}
	sort.Ints(s)

	wp := s[0] + s[0]/2 + s[1] + s[2] // wrapping paper

	fmt.Println(wp)
}
