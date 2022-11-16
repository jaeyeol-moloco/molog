package molog

import (
	"time"

	ccache "github.com/karlseguin/ccache/v2"
)

var cache = ccache.New(ccache.Configure().MaxSize(1000))

type LogKeyGen func(*Entry) string

type Deduper struct {
	CoolingTimeSeconds uint32
	LogKeyGen          LogKeyGen
}

func (d *Deduper) Suppress(entry *Entry) bool {
	key := d.LogKeyGen(entry)
	item := cache.Get(key)
	if item == nil || item.Expired() {
		cache.Set(key, "", time.Duration(d.CoolingTimeSeconds)*time.Second)
		return false
	}
	return true
}