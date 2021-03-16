package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println(time.Now()).UnixNano())
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("I have %d dogs.\n", rand.Intn(100))
}
