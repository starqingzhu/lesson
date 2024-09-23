package mylog

import (
	"fmt"
	"lesson/tools/tools"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	path := "./log/"
	if err := tools.DirCreate(path); err != nil {
		fmt.Printf("test mkdir path failed,err:%v\n", err)
		return
	}

	logger.Out = os.Stdout
}

func TestLogRus() {
	logrus.Info("mylog logrus hello, world.")
	logrus.Info("mylog logrus hello, world.")
	logrus.Info("mylog logrus hello, world.")
	logrus.Info("mylog logrus hello, world.")
	logrus.Info("mylog logrus hello, world.")
}
