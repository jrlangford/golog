package golog

import (
	"bytes"
	"github.com/jrlangford/toolbox"
	"testing"
)

func TestLoggers(t *testing.T) {

	var buffer bytes.Buffer

	var log Loggers
	log.Init(&buffer)

	var tests = []struct {
		logger          func(v ...interface{})
		expectedPattern string
	}{
		{log.Info, "INFO: .*"},
		{log.Debug, "DEBUG: .*"},
		{log.Warning, "WARNING: .*"},
		{log.Error, "ERROR: .*"},
	}

	b := make([]byte, 100)

	infoMsg := "test"

	for _, test := range tests {
		t.Run(
			test.expectedPattern,
			func(t *testing.T) {
				test.logger(infoMsg)
				n, err := buffer.Read(b)
				if err != nil {
					t.Fatalf("Error reading buffer")
				}
				receivedString := string(b[:n])
				toolbox.FailOnRegexMismatch(test.expectedPattern, receivedString, t)
			},
		)
	}

}
