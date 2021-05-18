package usignal

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var ExitSignals = []os.Signal{
	syscall.SIGTERM,
	syscall.SIGINT,
	syscall.SIGHUP,
}

var taskCount int64
var serviceExiting bool = false

// ProcessExitC Add an task
func ProcessExitC() (
	success bool,
	exitSignalC <-chan os.Signal,
	exitedCallback func(), // 成功退出
) {
	if serviceExiting {
		return
	}

	signalC := make(chan os.Signal, 1)
	exitSignalC = signalC
	signal.Notify(
		signalC,
		ExitSignals...,
	)

	atomic.AddInt64(&taskCount, 1)
	exitedCallback = func() {
		atomic.AddInt64(&taskCount, -1)
	}
	success = true
	return
}

// GetTaskCount Get tasks count
func GetTaskCount() int64 {
	return atomic.LoadInt64(&taskCount)
}

// WaitTasksExit Wait until all tasks exit
func WaitTasksExit() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			curTaskCount := atomic.LoadInt64(&taskCount)
			logrus.Infof("waiting task to exit, left: %d", curTaskCount)
			if curTaskCount <= 0 {
				return
			}
		}
	}
}

func init() {
	go func() {
		success, exitC, exited := ProcessExitC()
		if !success {
			return
		}
		defer exited()

		for {
			select {
			case <-exitC:
				serviceExiting = true
				logrus.Println("service exiting...")
				return
			}
		}
	}()
}
