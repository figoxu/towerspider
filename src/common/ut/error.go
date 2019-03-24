package ut

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func Catch(hooks ...func(err ...interface{})) {
	if err := recover(); err != nil {
		logrus.WithFields(logrus.Fields{
			"stack":   string(debug.Stack()),
			"recover": true,
		}).Println("panic")
		for _, hook := range hooks {
			hook(err)
		}
	}
}
