package lmbfmt_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/elvishp2006/logrus-lambda-formatter/pkg/lmbfmt"
	"github.com/sirupsen/logrus"
)

func Example() {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.DebugLevel,
		Formatter: &lmbfmt.Formatter{},
	}

	entry := logger.WithFields(logrus.Fields{
		"requestId": "4d3d9965-7fde-4c24-b2cc-474ddcd4f862",
	})

	entry.Time = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	entry.Println("Test message")

	// Output:
	// 2000-01-01T00:00:00Z	4d3d9965-7fde-4c24-b2cc-474ddcd4f862	INFO	Test message	{"level":"info","msg":"Test message","requestId":"4d3d9965-7fde-4c24-b2cc-474ddcd4f862","time":"2000-01-01T00:00:00Z"}
}

func TestFormat(t *testing.T) {
	f := &lmbfmt.Formatter{}
	e := logrus.NewEntry(logrus.New())

	e.Message = "Test message"
	e.Level = logrus.InfoLevel
	e.Time = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("WithoutRequestId", func(t *testing.T) {
		b, err := f.Format(e)

		if err != nil {
			t.Errorf("Error formatting: %v", err)
		}

		serialized, _ := f.JSONFormatter.Format(e)

		want := fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			e.Time.Format(time.RFC3339),
			strings.ToUpper(e.Level.String()),
			e.Message,
			string(serialized),
		)

		if string(b) != want {
			t.Errorf("Got %s, want %s", string(b), want)
		}
	})

	t.Run("WithRequestId", func(t *testing.T) {
		e.Data["requestId"] = "1234567890"

		b, err := f.Format(e)

		if err != nil {
			t.Errorf("Error formatting: %v", err)
		}

		serialized, _ := f.JSONFormatter.Format(e)

		want := fmt.Sprintf(
			"%s\t%s\t%s\t%s\t%s",
			e.Time.Format(time.RFC3339),
			e.Data["requestId"],
			strings.ToUpper(e.Level.String()),
			e.Message,
			string(serialized),
		)

		if string(b) != want {
			t.Errorf("Got %s, want %s", string(b), want)
		}
	})
}
