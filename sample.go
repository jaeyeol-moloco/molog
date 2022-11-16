package molog

import "math/rand"

type Sampler interface {
	Sample(entry *Entry) bool
}

type RandomSampler struct {
	Rate float64
	Rand *rand.Rand
}

func (s *RandomSampler) Sample(_ *Entry) bool {
	return s.Rand.Float64() < s.Rate
}
