package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n uint64
	sum := uint64(0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ = strconv.ParseUint(scanner.Text(), 10, 64)
		sum += n
	}
	fmt.Printf("%d\n", sum)
}
