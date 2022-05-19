package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go DatabaseConnection()
	go apiConnection()

	//go routine olduğundan hemen kapanmaması için
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigs
		_ = sig
		done <- true
	}()
	<-done
	fmt.Println("Program Closed")
	//go routine olduğundan hemen kapanmaması içinS
}
