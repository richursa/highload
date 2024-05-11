package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := uint64(0)
	scanner := bufio.NewScanner(os.Stdin)
	var lookup [10]uint64
	lookup[0] = 0
	lookup[1] = 1
	lookup[2] = 2
	lookup[3] = 3
	lookup[4] = 4
	lookup[5] = 5
	lookup[6] = 6
	lookup[7] = 7
	lookup[8] = 8
	lookup[9] = 9
	for scanner.Scan() {
		sum += stringToint(scanner.Text(), lookup)
	}
	fmt.Printf("%d\n", sum)
}

func stringToint(str string, lookup [10]uint64) uint64 {
	num := uint64(0)
	digit := uint64(1)
	_ = str[0]
	for i := len(str) - 1; i >= 0; i-- {
		num += (lookup[str[i]-'0']) * digit
		digit *= 10
	}
	return num
}
