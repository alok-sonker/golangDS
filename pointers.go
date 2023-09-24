package main

import "fmt"

func poiter() {
	var p *int32
	var x int32 = 100
	p = &x
	fmt.Println(*p)
}
