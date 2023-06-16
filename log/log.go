package log

import (
	"fmt"
	"time"
)

type Logging interface {
	LogDebug(message string)
	LogError(message string)
}

type Logger struct {
	Debug bool
}

func (logger Logger) LogDebug(message string) {
	if !logger.Debug {
		return
	}

	fmt.Printf("%s%s\n", getPrefix(), message)
}

func (logger Logger) LogError(message string) {
	fmt.Printf("%s%s\n", getPrefix(), message)
}

func getPrefix() string {
	currentTime := time.Now()
	return fmt.Sprintf("[%s] ", currentTime.UTC().Format(time.RFC1123Z))
}
