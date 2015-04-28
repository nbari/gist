package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLine(scanner *bufio.Scanner, replace_lines intslice) {

	scanner.Split(bufio.ScanLines)

	line := 0

	ra, _ := regexp.Compile("[^\\s]")

	for scanner.Scan() {
		line += 1
		if line == 4 {
			fmt.Printf("%d: %v\n", line, ra.ReplaceAllString(scanner.Text(), "-"))
		} else {
			fmt.Printf("%d: %v\n", line, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// slice of ints
type intslice []int

// String is the method to format the flag's value.
func (i *intslice) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *intslice) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("line flag already set")
	}

	fmt.Println(*i)

	//out := []int{}
	set := map[int]struct{}{}

	for _, n := range strings.Split(value, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		if _, ok := set[num]; ok {
			continue
		}
		*i = append(*i, num)
		set[num] = struct{}{}

		// Insertion sort
		for x := 1; x < len(*i); x++ {
			value := (*i)[x]
			y := x - 1
			for y >= 0 && (*i)[y] > value {
				(*i)[y+1] = (*i)[y]
				y = y - 1
			}
			(*i)[y+1] = value
		}
	}

	return nil
}

func main() {

	var replace_lines intslice

	flag.Var(&replace_lines, "l", "Number of the line(s) to be replaced")

	flag.Parse()

	//	if len(flag.Args()) == 0 {
	fmt.Println("--:", flag.NArg(), flag.Args(), replace_lines)
	//	}

	//if flag.NFlag() > 0 {
	//fmt.Println("lines to be replaced:")
	//for i := 0; i < len(replace_lines); i++ {
	//fmt.Printf("%d ", replace_lines[i])
	//}
	//}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		readLine(bufio.NewScanner(os.Stdin), replace_lines)
	} else {
		if flag.NArg() > 0 {

			f := flag.Arg(0)

			if Exists(f) {
				file, err := os.Open(f)
				defer file.Close()

				if err != nil {
					log.Fatal(err)
				}

				readLine(bufio.NewScanner(file), replace_lines)
			}
		}
	}
}
