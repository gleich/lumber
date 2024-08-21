package main

import (
	"os"

	"github.com/gleich/lumber/v3"
)

func main() {
	fname := "sample.txt"

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		lumber.Warning(fname, "doesn't exist. Creating now")
		err := os.WriteFile(fname, []byte("testing"), 0655)
		if err != nil {
			lumber.Fatal(err, "Failed to write to", fname)
		}
		lumber.Done("Wrote to", fname)
	} else {
		lumber.Info("Reading from file")
		b, err := os.ReadFile(fname)
		if err != nil {
			lumber.Fatal(err, "Failed to read from", fname)
		}
		lumber.Done("Read from", fname, "with content of", string(b))
	}
}
