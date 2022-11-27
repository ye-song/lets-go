//Modify the echo program to print the index and value of each of its
//arguments, one per line.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s1 := strconv.Itoa(i)
		s += sep + s1 + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
