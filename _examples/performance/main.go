package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gleich/lumber/v3"
)

func main() {
	lumberResult := timeFunc(lumber.Done)
	logResult := timeFunc(log.Println)

	fmt.Println()
	fmt.Println("lumber:", lumberResult)
	fmt.Println("log:", logResult)
}

func timeFunc(f func(v ...any)) string {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		f("foo", "bar")
	}
	return time.Since(start).String()
}
