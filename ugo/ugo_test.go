package ugo

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGo2(t *testing.T) {
	Go(func(exitC <-chan os.Signal, args ...interface{}) {
		time.Sleep(10 * time.Second)
		fmt.Println(1)
	})

	Go(func(exitC <-chan os.Signal, args ...interface{}) {
		time.Sleep(4 * time.Second)
		fmt.Println(2)
	})

	<-(make(chan int, 0))
}
