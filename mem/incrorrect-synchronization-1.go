package main

import "sync"

var a, b int
//var c = make(chan int, 1)
var l sync.Mutex

func f() {
	a = 1
	b = 2
	//c <- 0
	l.Unlock()
}

func g() {
	print(b)
	print(a)
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	//<-c
	g()
}