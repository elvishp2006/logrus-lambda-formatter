package lmbfmt_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/elvishp2006/logrus-lambda-formatter/pkg/lmbfmt"
	"github.com/sirupsen/logrus"
)

func TestFormat(t *testing.T) {
	lf := &lmbfmt.LambdaFormatter{}
	e := logrus.NewEntry(logrus.New())

	e.Message = "Test message"
	e.Level = logrus.InfoLevel
	e.Time = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("WithoutRequestId", func(t *testing.T) {
		b, err := lf.Format(e)

		if err != nil {
			t.Errorf("Error formatting: %v", err)
		}

		serialized, _ := lf.JSONFormatter.Format(e)

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

		b, err := lf.Format(e)

		if err != nil {
			t.Errorf("Error formatting: %v", err)
		}

		serialized, _ := lf.JSONFormatter.Format(e)

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
