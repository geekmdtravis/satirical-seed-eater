package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SetupInterruptHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rSIGTERM interrupt captured (ctrl + c). Attempting clean exit...")
		os.Exit(0)
	}()
}