package lumber

import (
	"log"
	"os"
	"time"
)

// Global options available to a user to change.
var (
	// The output file for Debug, Success, Warning, and Info
	NormalOut = os.Stdout
	// The output file for Fatal and Error
	ErrOut = os.Stderr

	// Fatal exit status
	ExitStatus = 1

	// Loggers

	// Normal logger for Debug, Success, Warning, and Info
	normalLogger = log.New(NormalOut, "", 0)
	// Error logger for Fatal and Error
	errLogger = log.New(NormalOut, "", 0)
)

// Log a normal status (Debug, Success, Warning, and Info)
func logNormal(stat status, t time.Time, ctx ...interface{}) {
	out := format(stat, t, separateWithSpaces(ctx...))
	normalLogger.Println(out)
}

// Log a normal status (Debug, Success, Warning, and Info)
func logError(stat status, t time.Time, ctx ...interface{}) {
	out := format(stat, t, separateWithSpaces(ctx...))
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
func Error(ctx ...interface{}) {
	logError(errorStatus, time.Now(), ctx...)
}

// Output a fatal log
func Fatal(ctx ...interface{}) {
	logError(fatalStatus, time.Now(), ctx...)
	os.Exit(ExitStatus)
}
