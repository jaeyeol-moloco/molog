package molog

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

func BenchmarkLogrus(b *testing.B) {
	logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	err := fmt.Errorf("something wrong happened")
	for i := 0; i < b.N; i++ {
		logrus.WithFields(logrus.Fields{"event": "handler-error", "error": err}).Error("let's fix!")
	}
}

func BenchmarkMolog(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFormatter(&logrus.JSONFormatter{})

	err := fmt.Errorf("something wrong happened")
	for i := 0; i < b.N; i++ {
		WithFields(logrus.Fields{"event": "handler-error", "error": err}).Error("let's fix!")
	}
}

func BenchmarkMologNthSampler(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFormatter(&logrus.JSONFormatter{})

	err := fmt.Errorf("something wrong happened")
	deduped := Limited(NewNthSampler(10))
	for i := 0; i < b.N; i++ {
		deduped.WithFields(logrus.Fields{"event": "handler-error", "error": err}).Error("let's fix!")
	}
}

func BenchmarkMologRandomSampler(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFormatter(&logrus.JSONFormatter{})

	err := fmt.Errorf("something wrong happened")
	deduped := Limited(NewRandomSampler(0.1))
	for i := 0; i < b.N; i++ {
		deduped.WithFields(logrus.Fields{"event": "handler-error", "error": err}).Error("let's fix!")
	}
}

func BenchmarkMologDeduperByCaller(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFormatter(&logrus.JSONFormatter{})

	err := fmt.Errorf("something wrong happened")
	deduper := NewDeduperByCaller(60)
	deduped := Limited(deduper)
	for i := 0; i < b.N; i++ {
		deduped.WithFields(logrus.Fields{"event": "handler-error", "error": err}).Error("let's fix!")
	}
}

func BenchmarkMologEventKeyDeduper(b *testing.B) {
	SetOutput(ioutil.Discard)
	SetFormatter(&logrus.JSONFormatter{})

	err := fmt.Errorf("something wrong happened")
	deduper := &Deduper{
		CoolingTimeSeconds: 60,
		LogKeyGen: func(entry *Entry) string {
			return fmt.Sprint(entry.GetFields()["event"])
		},
	}
	deduped := Limited(deduper)
	for i := 0; i < b.N; i++ {
		deduped.WithFields(logrus.Fields{"event": "handler-error", "error": err}).Error("let's fix!")
	}
}
