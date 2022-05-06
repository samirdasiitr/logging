package logging

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleLogger(t *testing.T) {
	tests := []struct {
		Name    string
		Level   uint32
		message string
	}{
		{"info", INFO, "testmessgae"},
		{"debug", DEBUG, "testmessgae"},
		{"warning", WARN, "testmessgae"},
		{"error", ERROR, "testmessgae"},
		{"trace", TRACE, "testmessgae"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			Init("TEST", buf, uint32(test.Level))
			logger := NewLogger(uint32(test.Level))

			switch test.Level {
			case INFO:
				logger.Info(test.message)
			case DEBUG:
				logger.Debug(test.message)
			case WARN:
				logger.Warn(test.message)
			case ERROR:
				logger.Error(test.message)
			case TRACE:
				logger.Trace(test.message)
			case FATAL:
				logger.Fatal(test.message)
			case PANIC:
				logger.Panic(test.message)
			}

			logMessage := buf.String()
			t.Log(logMessage)
			assert.True(t, strings.Contains(logMessage, test.message))

			assert.True(t, strings.Contains(logMessage, ":\""+test.Name+"\""))
		})
	}
}

func TestSimpleHook(t *testing.T) {
	buf := new(bytes.Buffer)
	Init("TEST", buf, uint32(INFO))
	logger := NewLogger(uint32(INFO))

	patterns := []string{
		`("password"\s*:\s*")(.*)(")`,
		`("access_key"\s*:\s*")(.*)(")`,
	}

	hook := NewFilterSensitiveInfoHook(patterns)
	RegisterLogHook(hook)

	logger.Info(`"password": "secret", othermsg`)
	logger.Info(`"access_key" : "secret", othermsg`)

	logMessage := buf.String()
	t.Log(logMessage)
	assert.False(t, strings.Contains(logMessage, "secret"))
}

func TestSimpleEvent(t *testing.T) {
	buf := new(bytes.Buffer)
	Init("TEST", buf, uint32(INFO))
	logger := NewLogger(uint32(INFO))

	logger.Eventf("create", "testsubject", "tester", INFO, "simpleevent")

	logMessage := buf.String()
	t.Log(logMessage)
	assert.True(t, strings.Contains(logMessage, "EVENT"))
}

func TestSimpleNotify(t *testing.T) {
	buf := new(bytes.Buffer)
	Init("TEST", buf, uint32(INFO))
	logger := NewLogger(uint32(INFO))

	logger.Notifyf("create", "testsubject", "tester", ERROR, "simpleevent")

	logMessage := buf.String()
	t.Log(logMessage)
	assert.True(t, strings.Contains(logMessage, "NOTIFY"))
}
