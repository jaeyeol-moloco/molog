package main

import (
	"fmt"
	"time"

	"github.com/moloco/molog"
)

func main() {
	molog.SetFormatter(&molog.JSONFormatter{})
	molog.WithFields(molog.Fields{"company": "MOLOCO", "headquarter": "Redwood city"}).Info("Hi")

	sampledLogger := molog.Limited(molog.NewNthSampler(10))
	for i := 0; i < 100; i++ {
		sampledLogger.Errorf("error is logged at 0.1 rate")
	}

	dedupedLogger := molog.Limited(&molog.Deduper{
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

	eventDeduper := &molog.Deduper{
		CoolingTimeSeconds: 1,
		LogKeyGen: func(entry *molog.Entry) string {
			return fmt.Sprint(entry.GetFields()["event"])
		},
	}
	thirdSampler := molog.NewNthSampler(3)
	sampledAndDeduped := molog.Limited(molog.AndLimiters(eventDeduper, thirdSampler))
	for i := 0; i < 30; i++ {
		sampledAndDeduped.WithFields(molog.Fields{"event": "fire", "try": i}).Info("deduped for 1s and sample every 3rd try")
		time.Sleep(100 * time.Millisecond)
	}

	dedupedByCaller := molog.Limited(molog.NewDeduperByCaller(1))
	for i := 0; i < 10; i++ {
		dedupedByCaller.WithFields(molog.Fields{"name": "john", "try": i}).Info("deduped by caller for 1s")
		time.Sleep(100 * time.Millisecond)
	}
}
