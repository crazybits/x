package logger

import (
	"testing"

	logging "github.com/op/go-logging"
)

var logger = logging.MustGetLogger("blockchain")

func Test(t *testing.T) {
	logger.Debug("debug")
	logger.Error("error")

	str := "format"
	logger.Errorf("test %s", str)

}
