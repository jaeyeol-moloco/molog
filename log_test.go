package molog

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLevelMatch(t *testing.T) {
	assert.Equal(t, logrus.ErrorLevel, ErrorLevel)
	assert.Equal(t, logrus.WarnLevel, WarnLevel)
	assert.Equal(t, logrus.InfoLevel, InfoLevel)
	assert.Equal(t, logrus.DebugLevel, DebugLevel)
	assert.Equal(t, logrus.TraceLevel, TraceLevel)
}
