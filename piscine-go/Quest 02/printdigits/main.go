package main

import (
	"github.com/01-edu/z01"
)

func main() {
	c := '0'
	for c <= '9' {
		z01.PrintRune(c)
		c++
	}
	z01.PrintRune('\n')
}
