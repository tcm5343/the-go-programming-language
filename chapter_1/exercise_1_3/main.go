package main

import (
	"os"
	"strings"
)

func SlowEcho() string {
	var s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
	// fmt.Println(s)
}

func FastEcho() string {
	s := strings.Join(os.Args[1:], " ")
	return s
	// fmt.Println(s)
}

func main() {
	SlowEcho()
	FastEcho()
}
