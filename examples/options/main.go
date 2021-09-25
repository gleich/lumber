package main

import (
	"math/rand"
	"time"

	"github.com/gleich/lumber/v2"
)

func main() {
	log := lumber.NewCustomLogger()
	log.Timezone = time.Local
	log.Multiline = true
	log.TrueColor = true

	randCap := 100
	log.Debug(rand.Intn(randCap))
	log.Info(rand.Intn(randCap))
	log.Success(rand.Intn(randCap))
	log.Warning(rand.Intn(randCap))
	log.ErrorMsg(rand.Intn(randCap))
	log.FatalMsg(rand.Intn(randCap))
}
