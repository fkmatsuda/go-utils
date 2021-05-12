package fklogger

import "os"

func ConsoleAppender() Appender {
	return initAppender(os.Stdout, os.Stderr)
}
