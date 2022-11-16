package main

import (
	"math/rand"
	"time"

	"github.com/moloco/molog"
)

func main() {
	molog.SetFormatter(&molog.JSONFormatter{})
	molog.WithFields(molog.Fields{"company": "MOLOCO", "headquarter": "Redwood city"}).Info("Hi")
	sampledLogger := molog.Sampled(&molog.RandomSampler{
		Rate: 0.1,
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	})
	for i := 0; i < 100; i++ {
		sampledLogger.Errorf("error is logged at 0.1 rate")
	}
}
