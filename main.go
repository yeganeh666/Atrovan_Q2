package main

import (
	log "Atrovan_Q2/log"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	eChar        = 101             // 'e' character ASCII code.
	intervalTime = 5 * time.Second // time interval for 5 sec.
)

var (
	currentChar uint8
)

func init() {
	/*
	** disable input buffering.
	** don't display entered characters on the screen.
	** restore the echoing state when exiting.
	 */
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	// set current character to be empty.
	currentChar = '0'
}

func main() {
	var ch []byte = make([]byte, 1)

	ticker := time.NewTicker(intervalTime)
	quit := make(chan struct{})

	go repetitiveWork(ticker, quit)

	for {
		os.Stdin.Read(ch)
		currentChar = ch[0]
		if currentChar == eChar {
			close(quit)
		}
	}
}

// repetitiveWork do repetitively works due to time passes to it.
func repetitiveWork(ticker *time.Ticker, quit chan struct{}) {
	for {
		select {
		case <-ticker.C:
			// check the input not null
			if currentChar != '0' {
				fmt.Printf(log.Info("Last character is : ", string(currentChar)))
				currentChar = '0'
			}
		case <-quit:
			ticker.Stop()
			fmt.Printf(log.Warning("e character pressed, bye :("))
			os.Exit(0)
		}
	}
}
