package pad

import (
	"fmt"
	"strconv"
)

func A(pad int) int {
	return len(fmt.Sprint(pad))
}

func B(pad int) int {
	return len(strconv.Itoa(pad))
}
