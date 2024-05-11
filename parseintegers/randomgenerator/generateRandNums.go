package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	outfile, err := os.Create("input.txt")
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < n; i++ {
		outfile.WriteString(fmt.Sprintln(rand.Int64()))
	}
	err = outfile.Close()
	if err != nil {
		panic(err)
	}
}
