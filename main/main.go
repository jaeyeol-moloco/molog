package main

import (
	"fmt"
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

	dedupedLogger := molog.Deduped(&molog.Deduper{
		CoolingTimeSeconds: 1,
		LogKeyGen: func(entry *molog.Entry) string {
			return fmt.Sprint(entry.GetFields()["name"])
		},
	})
	for i := 0; i < 10; i++ {
		dedupedLogger.WithFields(molog.Fields{"name": "john", "try": i}).Info("deduped for 1s")
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
	dedupedLogger.WithFields(molog.Fields{"name": "john", "greeting": "I'm back"}).Info("deduped for 1s")
}
