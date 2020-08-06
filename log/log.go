package log

import (
	"fmt"
	"time"
)

var (
	// Info for log all info state with teal color
	Info = teal
	// Warning for log all warn state with yellow color
	Warning = yellow
)

var (
	yellow = log("\n\033[1;33m%s\033[0m")
	teal   = log("\n\033[1;36m%s\033[0m") // blue-green color.
)

// log makes output colorful & add time.
func log(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		return fmt.Sprintf(colorString, fmt.Sprint(currentTime, " | ")) + fmt.Sprintln(args...)
	}
	return sprint
}
