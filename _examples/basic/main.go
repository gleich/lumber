package main

import "github.com/gleich/lumber/v3"

func main() {
	msg := "Hello World!"
	lumber.Debug(msg)
	lumber.Done(msg)
	lumber.Info(msg)
	lumber.Warn(msg)
}
