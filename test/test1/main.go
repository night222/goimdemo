package main

import (
	"fmt"
	"sync"
)

type MyMuter struct {
	count int
	sync.Mutex
}

func main() {
	var m1 MyMuter
	m1.Lock()
	m1.count++
	var m2 = m1
	//m1.Unlock()
	//m2.Lock()
	m2.count++
	m2.Unlock()
	fmt.Println(m1.count, m2.count)
}
