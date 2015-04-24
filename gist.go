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

	fmt.Println(replace_lines)

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

	out := []int{}

	for _, ln := range strings.Split(value, ",") {
		//
		// try to sanitize input 1, 3,  7, 9
		//fmt.Printf("[%s]\n", strings.TrimSpace(ln))
		//
		line, err := strconv.Atoi(ln)
		if err != nil {
			return err
		}
		out = append(out, line)
		// Insertion sort
		fmt.Println("Start:", out)
		for x := 1; x < len(out); x++ {
			value := out[x]
			y := x - 1
			for y >= 0 && out[y] > value {
				fmt.Printf("Value: %d\n out[y]: %d y: %d\n", value, out[y], y)
				out[y+1] = out[y]
				y = y - 1
			}
			out[y+1] = value
		}
		fmt.Printf("End:   %v \n", out)
	}

	fmt.Println("sorted ---------->", out)

	return nil
}

var replace_lines intslice

func main() {

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
