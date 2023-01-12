package lmbfmt

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = time.RFC3339
)

type Formatter struct {
	logrus.JSONFormatter
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	serialized, err := f.JSONFormatter.Format(entry)

	timestampFormat := f.TimestampFormat

	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	if _, ok := entry.Data["requestId"]; !ok {
		return []byte(fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			entry.Time.Format(timestampFormat),
			strings.ToUpper(entry.Level.String()),
			entry.Message,
			string(serialized),
		)), err
	}

	return []byte(fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s",
		entry.Time.Format(timestampFormat),
		entry.Data["requestId"],
		strings.ToUpper(entry.Level.String()),
		entry.Message,
		string(serialized),
	)), err
}
