package lumber

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	successStatus string = "   SUCCESS   "
	fatalStatus   string = "    FATAL    "
	errorStatus   string = "    ERROR    "
	warningStatus string = "   WARNING   "
	infoStatus    string = "    INFO     "
	debugStatus   string = "    DEBUG    "
)

// The default logger used by default logs
var defaultLogger = NewCustomLogger()

// Log a normal status (Debug, Success, Warning, and Info)
func logNormal(config Logger, stat string, t time.Time, ctx ...interface{}) {
	out := format(config, stat, t, separateWithSpaces(ctx...))
	writers := append([]io.Writer{config.NormalOut}, config.ExtraNormalOuts...)
	log.New(io.MultiWriter(writers...), "", 0).Println(out)
}

// Log a normal status (Debug, Success, Warning, and Info)
func logError(config Logger, stat string, t time.Time, err error, ctx ...interface{}) {
	var out string

	if config.ShowStack && err != nil {
		out = format(
			config,
			stat,
			t,
			fmt.Sprintf(
				"%v\n--- Stack Trace ---\n%+v",
				separateWithSpaces(ctx...),
				errors.WithStack(err),
			),
		)
	} else if err == nil {
		out = format(config, stat, t, separateWithSpaces(ctx...))
	} else {
		out = format(config, stat, t, separateWithSpaces(ctx...)+"\n\n"+err.Error())
	}

	writers := append([]io.Writer{config.ErrOut}, config.ExtraErrOuts...)
	log.New(io.MultiWriter(writers...), "", 0).Println(out)
}

// Output a success log
func Success(ctx ...interface{}) {
	logNormal(defaultLogger, successStatus, time.Now(), ctx...)
}

// Output an info log
func Info(ctx ...interface{}) {
	logNormal(defaultLogger, infoStatus, time.Now(), ctx...)
}

// Output a debug log
func Debug(ctx ...interface{}) {
	logNormal(defaultLogger, debugStatus, time.Now(), ctx...)
}

// Output a warning log
func Warning(ctx ...interface{}) {
	logNormal(defaultLogger, warningStatus, time.Now(), ctx...)
}

// Output an error log
func Error(err error, ctx ...interface{}) {
	logError(defaultLogger, errorStatus, time.Now(), err, ctx...)
}

// Output an error log with no actual error value
func ErrorMsg(ctx ...interface{}) {
	logError(defaultLogger, errorStatus, time.Now(), nil, ctx...)
}

// Output a fatal log. Will exit program after logging
func Fatal(err error, ctx ...interface{}) {
	logError(defaultLogger, fatalStatus, time.Now(), err, ctx...)
	os.Exit(defaultLogger.ExitCode)
}

// Output a fatal log with no actual error value. Will exit program after logging
func FatalMsg(ctx ...interface{}) {
	logError(defaultLogger, fatalStatus, time.Now(), nil, ctx...)
	os.Exit(defaultLogger.ExitCode)
}
