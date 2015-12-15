package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func first(d []byte) {
	lvl := 0

	for i := 0; i < len(d); i++ {
		switch string(d[i]) {
		case "(":
			lvl++
		case ")":
			lvl--
		default:
			panic("Unexpected instruction: " + string(d[i]))
		}
	}
	fmt.Println("Santa is on level: " + strconv.Itoa(lvl))
}

func second(d []byte) {
	lvl := 0

	for i := 0; i < len(d); i++ {
		switch string(d[i]) {
		case "(":
			lvl++
		case ")":
			lvl--
		default:
			panic("Unexpected instruction: " + string(d[i]))
		}

		if lvl == -1 {
			fmt.Println("Santa reached level -1 on instruction: " + strconv.Itoa(i+1))
			return
		}
	}
	fmt.Println("Error: Santa didn't reach level -1!")
}

func main() {
	part, e := strconv.Atoi(os.Args[1]) //part 1 or 2
	f := os.Args[2]                     //input file
	d, e := ioutil.ReadFile(f)          //get data
	if e != nil {
		panic(e)
	}

	if part == 1 {
		first(d)
	} else if part == 2 {
		second(d)
	} else {
		fmt.Println("Error: Select part 1 or 2")
	}
}
