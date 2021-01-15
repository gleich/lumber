package lumber

import (
	"os"
	"time"
)

// Output a success log with a given time
func SuccessWithTime(t time.Time, ctx ...interface{}) {
	logNormal(successStatus, t, ctx...)
}

// Output an info log with a given time
func InfoWithTime(t time.Time, ctx ...interface{}) {
	logNormal(infoStatus, t, ctx...)
}

// Output a debug log with a given time
func DebugWithTime(t time.Time, ctx ...interface{}) {
	logNormal(debugStatus, t, ctx...)
}

// Output a warning log with a given time
func WarningWithTime(t time.Time, ctx ...interface{}) {
	logNormal(warningStatus, t, ctx...)
}

// Output an error log
func ErrorWithTime(err error, t time.Time, ctx ...interface{}) {
	if err != nil {
		logError(errorStatus, t, err, ctx...)
	}
}

// Output an error log and run a given function before
func ErrorHookWithTime(hook func(), err error, t time.Time, ctx ...interface{}) {
	if err != nil {
		hook()
		logError(errorStatus, t, err, ctx...)
	}
}

// Output an error log with no actual error value
func ErrorMsgWithTime(t time.Time, ctx ...interface{}) {
	logError(errorStatus, t, nil, ctx...)
}

// Output a fatal log
func FatalWithTime(err error, t time.Time, ctx ...interface{}) {
	if err != nil {
		logError(fatalStatus, t, err, ctx...)
		os.Exit(ExitStatus)
	}
}