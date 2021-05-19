# go-exit-signal
An light-weight exit signal manage lib. 
Listen task exit signal, and manage task exit.

### Usage
example: test/main.go

- listen sys exit signal for task
- when main process/task exit, it can wait until all task end.
```golang
// RunMyTickTask Run tick every seconds
func RunMyTickTask() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.stop()
	
	// register an signal
    success, exitC, exited := usignal.ProcessExitC()
    if !success { // register signal failed
    	return
    }
    
    defer exited() // callback if task ended
    
    for {
    	select {
            case <-exitC:
            	// exit loop
                fmt.Println("exit")
                return

            case <-ticker.C:
            fmt.Println("tick")	
    	}
    }
}

func main() {
	defer func() {
		usignal.WaitTaskExit() // When main process id exiting, wait until all task end.
    }
	
	// do main process thing
	// like, http listen and serve
	// ...
	
}
```