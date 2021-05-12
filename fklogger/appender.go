package fklogger

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
)

type appenderConfig struct {
	ID            uint64
	writer        io.Writer
	errWriter     io.Writer
	traceLogger   *log.Logger
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
}

type appender struct {
	config *appenderConfig
}

func initAppender(writer, errWriter io.Writer) Appender {

	appender := appender{
		&appenderConfig{
			writer:        writer,
			errWriter:     errWriter,
			traceLogger:   log.New(writer, "TRACE: ", log.Ldate|log.Ltime), //|log.Lshortfile
			debugLogger:   log.New(writer, "DEBUG: ", log.Ldate|log.Ltime),
			infoLogger:    log.New(writer, "INFO: ", log.Ldate|log.Ltime),
			warningLogger: log.New(writer, "WARN: ", log.Ldate|log.Ltime),
			errorLogger:   log.New(errWriter, "ERROR: ", log.Ldate|log.Ltime),
			fatalLogger:   log.New(errWriter, "FATAL: ", log.Ldate|log.Ltime),
			panicLogger:   log.New(errWriter, "PANIC: ", log.Ldate|log.Ltime),
		},
	}

	var result Appender = appender

	return result

}

// Trace logs something very low level.
func (a appender) trace(message string, params ...interface{}) {
	a.config.traceLogger.Println(fmt.Sprintf(message, params...))
}

// Debug logs useful debugging information.
func (a appender) debug(message string, params ...interface{}) {
	a.config.debugLogger.Println(fmt.Sprintf(message, params...))
}

// Info logs something noteworthy happened!
func (a appender) info(message string, params ...interface{}) {
	a.config.infoLogger.Println(fmt.Sprintf(message, params...))
}

// Warn logs something that you should probably take a look at.
func (a appender) warn(message string, params ...interface{}) {
	a.config.warningLogger.Println(fmt.Sprintf(message, params...))
}

// Error logs something failed but I'm not quitting.
func (a appender) error(message string, params ...interface{}) {
	a.config.errorLogger.Println(printStack(fmt.Sprintf(message, params...)))
}

// Fatal logs a fatal error
// Calls os.Exit(1) after logging
func (a appender) fatal(message string, params ...interface{}) {
	a.config.fatalLogger.Println(printStack(fmt.Sprintf(message, params...)))
}

// Panic logs a panic error
// Calls panic() after logging
func (a appender) panic(message string, params ...interface{}) {
	a.config.panicLogger.Println(printStack(fmt.Sprintf(message, params...)))
}

func (a appender) doFatal(message string, params ...interface{}) {
	a.config.fatalLogger.Fatalln(fmt.Sprintf(message, params...))
}

func (a appender) doPanic(message string, params ...interface{}) {
	a.config.fatalLogger.Panicln(fmt.Sprintf(message, params...))
}

func printStack(log string) string {
	trace := make([]byte, 1024)
	count := runtime.Stack(trace, true)
	stringTrace := string(trace)

	if count > 0 {
		traceLines := strings.Split(stringTrace, "\n")
		stackLines := traceLines[7:]
		stackText := traceLines[0]
		for _, line := range stackLines {
			stackText = fmt.Sprintf("%s\n%s", stackText, line)
		}

		return fmt.Sprintf("%s\n\tStack trace => %s", log, stackText)
	}

	return log

}
