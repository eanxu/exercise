package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := 0; index < 10; index++ {
		fmt.Println(r.Intn(10))
	}
}
