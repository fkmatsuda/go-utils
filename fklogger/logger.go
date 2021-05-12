package fklogger

var (
	loggerConfig config = config{
		appenders: make([]Appender, 0),
	}
)

// Appender defines interface for logger appender
type Appender interface {
	trace(message string, params ...interface{})

	debug(message string, params ...interface{})

	info(message string, params ...interface{})

	warn(message string, params ...interface{})

	error(message string, params ...interface{})

	fatal(message string, params ...interface{})

	panic(message string, params ...interface{})

	doFatal(message string, params ...interface{})

	doPanic(message string, params ...interface{})
}

// Config defines a log configuration
type config struct {
	appenderIdGen uint64
	appenders     []Appender
}

// DefaultConfig initializes default logger configuration
func DefaultConfig() {
	RegisterAppender(DefaultAppender())
}

// DefaultAppender returns default appender
func DefaultAppender() Appender {
	return ConsoleAppender()
}

// RegisterAppender registers a new appender in the stack of appenders
func RegisterAppender(a Appender) {
	var i interface{} = a
	rawAppender, ck := i.(appender)
	if ck {
		rawAppender.config.ID = loggerConfig.appenderIdGen
		loggerConfig.appenderIdGen++
	}
	newAppenders := make([]Appender, len(loggerConfig.appenders)+1)
	newAppenders[0] = a
	copy(newAppenders[1:], loggerConfig.appenders)
	loggerConfig.appenders = newAppenders
}

func listAppenders() []Appender {
	return loggerConfig.appenders
}

// Trace logs something very low level.
func Trace(message string, params ...interface{}) {
	for _, appender := range listAppenders() {
		appender.trace(message, params...)
	}
}

// Debug logs useful debugging information.
func Debug(message string, params ...interface{}) {
	for _, appender := range listAppenders() {
		appender.debug(message, params...)
	}
}

// Info logs something noteworthy happened!
func Info(message string, params ...interface{}) {
	for _, appender := range listAppenders() {
		appender.info(message, params...)
	}
}

// Warn logs something that you should probably take a look at.
func Warn(message string, params ...interface{}) {
	for _, appender := range listAppenders() {
		appender.warn(message, params...)
	}
}

// Error logs something failed but I'm not quitting.
func Error(message string, params ...interface{}) {
	for _, appender := range listAppenders() {
		appender.error(message, params...)
	}
}

// Fatal logs a fatal error
// Calls os.Exit(1) after logging
func Fatal(message string, params ...interface{}) {
	appenders := listAppenders()
	for _, appender := range appenders {
		appender.fatal(message, params...)
	}
	appenders[0].doFatal(message, params...)
}

// Panic logs a panic error
// Calls panic() after logging
func Panic(message string, params ...interface{}) {
	appenders := listAppenders()
	for _, appender := range appenders {
		appender.panic(message, params...)
	}
	appenders[0].doPanic(message, params...)
}

// Error logs something failed but I'm not quitting.
func ErrorErr(err error) {
	Error(err.Error())
}

// Fatal logs a fatal error
// Calls os.Exit(1) after logging
func FatalErr(err error) {
	Fatal(err.Error())
}

// Panic logs a panic error
// Calls panic() after logging
func PanicErr(err error) {
	Panic(err.Error())
}
