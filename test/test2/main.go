package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

func main() {
	go A()
	time.Sleep(2 * time.Second)
	fmt.Println("断点3")
	mu.Lock()
	fmt.Println("断点4")
	defer mu.RUnlock()
	count++
	fmt.Println(count)
}
func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	fmt.Println("断点1")
	mu.RLock()
	fmt.Println("断点2")
	mu.RUnlock()
}
