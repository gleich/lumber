package main

import (
	"math/rand"
	"time"

	"github.com/gleich/lumber/v3"
)

func main() {
	lumber.SetTimezone(time.Local)
	lumber.SetTimeFormat("Mon Jan 2 15:04:05 MST 2006")
	lumber.SetFatalExitCode(0)

	randCap := 100
	lumber.Debug(rand.Intn(randCap))
	lumber.Info(rand.Intn(randCap))
	lumber.Done(rand.Intn(randCap))
	lumber.Warning(rand.Intn(randCap))
	lumber.ErrorMsg(rand.Intn(randCap))
	lumber.FatalMsg(rand.Intn(randCap))
}
