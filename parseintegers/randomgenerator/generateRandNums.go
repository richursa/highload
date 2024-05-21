package main

import (
	"bufio"
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
	writer := bufio.NewWriter(outfile)
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < n; i++ {
		writer.WriteString(fmt.Sprintln(rand.Uint64N(2147483647)))
	}
	err = outfile.Close()
	if err != nil {
		panic(err)
	}
}
