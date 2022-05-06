package logging

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

type FilterSensitiveInfoHook struct {
	Pattern *regexp.Regexp
}

//var sensitiveInfoMask = regexp.MustCompile(`("password"\s*:\s")(.*)(")`)

// Fire called when a log entry is about to be written,
func (hook *FilterSensitiveInfoHook) Fire(entry *logrus.Entry) error {
	entry.Message = string(hook.Pattern.ReplaceAllString(entry.Message, "$1***$3"))
	return nil
}

// Levels at which masking to be done.
func (hook *FilterSensitiveInfoHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

// FilterSensitiveInfoHook creates new sensitive info log filter
func NewFilterSensitiveInfoHook(patterns []string) *FilterSensitiveInfoHook {
	obj := new(FilterSensitiveInfoHook)
	obj.Pattern = regexp.MustCompile(strings.Join(patterns, "|"))
	return obj
}
