package main

import "github.com/gleich/lumber/v2"

func main() {
	msg := "Hello World!"
	lumber.Debug(msg)
	lumber.Info(msg)
	lumber.Success(msg)
	lumber.Warning(msg)
	lumber.ErrorMsg(msg)
	lumber.FatalMsg(msg)
}
