package lumber

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

var (
	// The output file for Debug, Success, Warning, and Info
	NormalOut = os.Stdout
	// The output file for Fatal and Error
	ErrOut = os.Stderr
	// If errors should automatically be checked for a non-nil value
	ErrNilCheck = true

	// Fatal exit code
	ExitStatus = 1

	// Normal logger for Debug, Success, Warning, and Info
	normalLogger = log.New(NormalOut, "", 0)
	// Error logger for Fatal and Error
	errLogger = log.New(ErrOut, "", 0)
)

// Log status
type status string

const (
	successStatus status = "   SUCCESS   "
	fatalStatus   status = "    FATAL    "
	errorStatus   status = "    ERROR    "
	warningStatus status = "   WARNING   "
	infoStatus    status = "    INFO     "
	debugStatus   status = "    DEBUG    "
)

// Log a normal status (Debug, Success, Warning, and Info)
func logNormal(stat status, t time.Time, ctx ...interface{}) {
	out := format(stat, t, separateWithSpaces(ctx...))
	normalLogger.Println(out)
}

// Log a normal status (Debug, Success, Warning, and Info)
func logError(stat status, t time.Time, err error, ctx ...interface{}) {
	var out string

	if err == nil {
		out = format(stat, t, separateWithSpaces(ctx...))
	} else {
		out = format(stat, t, fmt.Sprintf("%v\n\n--- Stack Trace ---\n%+v", separateWithSpaces(ctx...), errors.WithStack(err)))
	}

	errLogger.Println(out)
}

// Output a success log
func Success(ctx ...interface{}) {
	logNormal(successStatus, time.Now(), ctx...)
}

// Output an info log
func Info(ctx ...interface{}) {
	logNormal(infoStatus, time.Now(), ctx...)
}

// Output a debug log
func Debug(ctx ...interface{}) {
	logNormal(debugStatus, time.Now(), ctx...)
}

// Output a warning log
func Warning(ctx ...interface{}) {
	logNormal(warningStatus, time.Now(), ctx...)
}

// Output an error log
func Error(err error, ctx ...interface{}) {
	if !ErrNilCheck || err != nil {
		logError(errorStatus, time.Now(), err, ctx...)
	}
}

// Output an error log with no actual error value
func ErrorMsg(ctx ...interface{}) {
	logError(errorStatus, time.Now(), nil, ctx...)
}

// Output a fatal log
func Fatal(err error, ctx ...interface{}) {
	if !ErrNilCheck || err != nil {
		logError(fatalStatus, time.Now(), err, ctx...)
		os.Exit(ExitStatus)
	}
}

// Output a fatal log with no actual error value
func FatalMsg(ctx ...interface{}) {
	logError(fatalStatus, time.Now(), nil, ctx...)
	os.Exit(ExitStatus)
}
