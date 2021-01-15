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

// Output an error log with a given time
func ErrorWithTime(ctx ...interface{}) {
	logError(errorStatus, time.Now(), ctx...)
}

// Output a fatal log with a given time
func FatalWithTime(ctx ...interface{}) {
	logError(fatalStatus, time.Now(), ctx...)
	os.Exit(ExitStatus)
}
