package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	out := []int{}
	set := map[int]struct{}{}

	//value := "1,9,5,5,b,4,3,3,a"
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	// convert to int and add to a set to remove dupes
	for _, n := range strings.Split(value, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		if _, ok := set[num]; ok {
			continue
		}
		out = append(out, num)
		set[num] = struct{}{}
	}
	// insertion sort
	for x := 1; x < len(out); x++ {
		value := out[x]
		y := x - 1
		for y >= 0 && out[y] > value {
			out[y+1] = out[y]
			y = y - 1
		}
		out[y+1] = value
	}

	elapsed := time.Since(start)
	fmt.Println(out, elapsed)
}
