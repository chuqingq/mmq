package tcp

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = logrus.New()

func init() {
	log.SetLevel(logrus.FatalLevel)
}
