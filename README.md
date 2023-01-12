# Logrus Lambda Formatter
Provided formatter allow to easily format [Logrus](https://github.com/sirupsen/logrus) log output into a CloudWatch compatible format.

## Sample Usage
```go
package main

import (
	"os"

	"github.com/elvishp2006/logrus-lambda-formatter/pkg/lmbfmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.DebugLevel,
		Formatter: &lmbfmt.Formatter{},
	}

	logger.WithFields(logrus.Fields{
		"requestId": "4d3d9965-7fde-4c24-b2cc-474ddcd4f862",
	}).Info("Test message")
}
```
Above sample will produce:
```
2000-01-01T00:00:00Z	4d3d9965-7fde-4c24-b2cc-474ddcd4f862	INFO	Test message	{"level":"info","msg":"Test message","requestId":"4d3d9965-7fde-4c24-b2cc-474ddcd4f862","time":"2000-01-01T00:00:00Z"}
```
Note that the requestId is used for compose the message in a format that will be interpreted by CloudWatch

## License
This project is under [MIT License](./LICENSE.md).
