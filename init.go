package omg

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			file = fmt.Sprintf("%v:%v", path.Base(frame.File), frame.Line)
			return frame.Function, file
		},
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
}
