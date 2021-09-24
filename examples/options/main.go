package main

import (
	"math/rand"
	"time"

	"github.com/gleich/lumber/v2"
)

func main() {
	lumber.Timezone = time.Local
	lumber.MultiLine = true
	lumber.TrueColor = false

	randCap := 100
	lumber.Debug(rand.Intn(randCap))
	lumber.Info(rand.Intn(randCap))
	lumber.Success(rand.Intn(randCap))
	lumber.Warning(rand.Intn(randCap))
	lumber.ErrorMsg(rand.Intn(randCap))
	lumber.FatalMsg(rand.Intn(randCap))
}
