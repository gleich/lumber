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

	// Fatal exit code
	ExitCode = 1
	// If stack traces should be included
	ShowStack = true
	// Timezone for the time to be outputted in
	Timezone = time.UTC

	// Normal logger for Debug, Success, Warning, and Info
	normalLogger = log.New(NormalOut, "", 0)
	// Error logger for Fatal and Error
	errLogger = log.New(ErrOut, "", 0)
)

const (
	successStatus string = "   SUCCESS   "
	fatalStatus   string = "    FATAL    "
	errorStatus   string = "    ERROR    "
	warningStatus string = "   WARNING   "
	infoStatus    string = "    INFO     "
	debugStatus   string = "    DEBUG    "
)

// Log a normal status (Debug, Success, Warning, and Info)
func logNormal(stat string, t time.Time, ctx ...interface{}) {
	out := format(stat, t, separateWithSpaces(ctx...))
	normalLogger.Println(out)
}

// Log a normal status (Debug, Success, Warning, and Info)
func logError(stat string, t time.Time, err error, ctx ...interface{}) {
	var out string

	if ShowStack && err != nil {
		out = format(
			stat,
			t,
			fmt.Sprintf(
				"%v\n--- Stack Trace ---\n%+v",
				separateWithSpaces(ctx...),
				errors.WithStack(err),
			),
		)
	} else if err == nil {
		out = format(stat, t, separateWithSpaces(ctx...))
	} else {
		out = format(stat, t, separateWithSpaces(ctx...)+"\n\n"+err.Error())
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
	logError(errorStatus, time.Now(), err, ctx...)
}

// Output an error log with no actual error value
func ErrorMsg(ctx ...interface{}) {
	logError(errorStatus, time.Now(), nil, ctx...)
}

// Output a fatal log
func Fatal(err error, ctx ...interface{}) {
	logError(fatalStatus, time.Now(), err, ctx...)
	os.Exit(ExitCode)
}

// Output a fatal log with no actual error value
func FatalMsg(ctx ...interface{}) {
	logError(fatalStatus, time.Now(), nil, ctx...)
	os.Exit(ExitCode)
}
