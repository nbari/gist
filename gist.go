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

const NOT_SPACE = "[^\\s]"

var buffer []string

func readLine(scanner *bufio.Scanner, replace_lines IntSet, replace_strings StrSlice) {

	scanner.Split(bufio.ScanLines)

	line := 1

	ra, _ := regexp.Compile(NOT_SPACE)

	for scanner.Scan() {
		_, s := replace_lines[line]
		if s {
			buffer = append(buffer, ra.ReplaceAllString(scanner.Text(), "-"))
		} else {
			if len(replace_strings) > 0 {
				str := scanner.Text()
				for _, rs := range replace_strings {
					str = strings.Replace(str, rs, strings.Repeat("-", len(rs)), -1)
				}
				buffer = append(buffer, str)
			} else {
				buffer = append(buffer, scanner.Text())
			}
		}
		line += 1
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

// set of ints
type IntSet map[int]struct{}

// set of strings
type StrSlice []string

func (i *IntSet) String() string {
	return fmt.Sprint(*i)
}

func (i *IntSet) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("line flag already set")
	}
	for _, n := range strings.Split(value, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		if _, found := (*i)[num]; found {
			continue
		}
		(*i)[num] = struct{}{}
	}
	return nil
}

func (s *StrSlice) String() string {
	return fmt.Sprint(*s)
}

func (s *StrSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {

	var replace_lines = IntSet{}
	var replace_strings StrSlice

	flag.Var(&replace_lines, "l", ">>>>>>>>>>>>>>>>> l")
	flag.Var(&replace_strings, "r", ">>>>>>>>>>>>>>> r")
	push := flag.Bool("p", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-lrp] file\n\n", os.Args[0])
		fmt.Printf("  example: %s -l 3,7 -r secret -r 'my passphrase' file.conf\n\n", os.Args[0])
		fmt.Println("  -l: Number of the line(s) to be replaced, comma separated.")
		fmt.Println("  -r: Word to be replaced, can be used multiple times.")
		fmt.Println("  -p: Push the gist.")
	}

	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		readLine(bufio.NewScanner(os.Stdin), replace_lines, replace_strings)
	} else {
		if flag.NArg() != 1 {
			fmt.Fprintf(os.Stderr, "Usage: %s [-lrp] file, use -h for more info\n\n", os.Args[0])
			os.Exit(1)
		}
		f := flag.Arg(0)
		if Exists(f) {
			file, err := os.Open(f)
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			readLine(bufio.NewScanner(file), replace_lines, replace_strings)
		} else {
			fmt.Printf("Cannot read file: %s\n", f)
		}
	}

	// print preview
	pad := len(fmt.Sprint(len(buffer)))
	for k, v := range buffer {
		fmt.Printf("%*d  %s\n", pad, k+1, v)
	}

	if *push {
		fmt.Println("->>>>>>>>>>>>>>>> push")
	}

}
