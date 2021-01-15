package main

import (
	"io/ioutil"

	"github.com/Matt-Gleich/lumber"
)

func main() {
	_, err := ioutil.ReadFile("./not-there.txt")
	lumber.Fatal(err, "Failed to read from file")
}
