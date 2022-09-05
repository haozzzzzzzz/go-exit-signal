package main

import (
	"fmt"
	"github.com/haozzzzzzzz/go-exit-signal/usignal"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			success, exitC, exited := usignal.ProcessExitC()
			defer exited()
			fmt.Println(success)

			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-exitC:
					fmt.Println("exit c", i)
					return

				case <-ticker.C:
					fmt.Println("tick", i)
				}
			}
		}(i)
	}

	usignal.WaitTasksExit()
}
