package main

import (
	"bufio"
	//	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func readLine(scanner *bufio.Scanner) {

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

type replace_lines []int

func main() {

	//	lines := flag.String()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		readLine(bufio.NewScanner(os.Stdin))
	} else {
		if len(os.Args) > 1 {

			f := os.Args[1]

			fmt.Println(os.Args)

			if Exists(f) {
				file, err := os.Open(f)
				defer file.Close()

				if err != nil {
					log.Fatal(err)
				}

				readLine(bufio.NewScanner(file))
			}
		}
	}
}
