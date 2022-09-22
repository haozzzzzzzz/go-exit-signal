package ugo

import (
	"github.com/haozzzzzzzz/go-exit-signal/usignal"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"runtime/debug"
)

func Go(goFunc func(exitC <-chan os.Signal, args ...interface{}), args ...interface{}) {
	go func() {
		success, exitC, exited := usignal.ProcessExitC()
		if !success {
			return
		}

		defer func() {
			if iRec := recover(); iRec != nil {
				logrus.Errorf("run go func panic : %s, %s", iRec, string(debug.Stack()))
			}
			exited()
		}()

		goFunc(exitC, args...)
	}()
}

func GoRecover(goFunc func(args ...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if iRec := recover(); iRec != nil {
				logrus.Errorf("go func panic . panic: %s, func: %s", iRec, reflect.TypeOf(goFunc))
			}
		}()

		goFunc(args...)
	}()
}
