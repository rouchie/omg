package omg

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Hello() {
	logrus.Infof("Hello: %v", time.Now().String())
}
