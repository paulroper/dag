package logging

import (
	"fmt"
	"time"
)

type Logger interface {
	LogDebug(message string)
	Log(message string)
}

type Log struct {
	Debug bool
}

func (log Log) LogDebug(message string) {
	if !log.Debug {
		return
	}

	fmt.Printf("%s%s\n", getPrefix(), message)
}

func (log Log) Log(message string) {
	fmt.Printf("%s%s\n", getPrefix(), message)
}

func getPrefix() string {
	currentTime := time.Now()
	return fmt.Sprintf("[%s] ", currentTime.UTC().Format(time.RFC1123Z))
}
