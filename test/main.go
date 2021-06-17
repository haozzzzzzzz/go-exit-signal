package main

import (
	"fmt"
	"github.com/haozzzzzzzz/go-exit-signal/v1/usignal"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			success, exitC, exited := usignal.ProcessExitC()
			defer exited()
			fmt.Println(success)

			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-exitC:
					fmt.Println("exit c")
					return

				case <-ticker.C:
					fmt.Println("tick")
				}
			}
		}()
	}

	usignal.WaitTasksExit()
}
