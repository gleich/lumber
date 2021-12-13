package lumber

import (
	"io"
	"os"
	"time"
)

// Custom logger for lumber to use
type Logger struct {
	NormalOut       *os.File       // The output file for Debug, Success, Warning, and Info. Default is os.Stdout
	ErrOut          *os.File       // The output file for Fatal, FatalMsg, Error, and ErrorMsg. Default is os.Stderr
	ExtraNormalOuts []io.Writer    // Extra normal output destinations (e.g. outputting to a file as well)
	ExtraErrOuts    []io.Writer    // Extra error output destinations (e.g. outputting to a file as well)
	ExitCode        int            // Fatal exit code. Default is 1
	ShowStack       bool           // If stack trades should be included. Default is true
	Timezone        *time.Location // Timezone for the time to be outputted in. Default is time.UTC
	Padding         bool           // If the log should have an extra new line at the bottom. Default is true
	Multiline       bool           // If the log should span multiple lines. Default is false
	ColoredOutput   bool           // If the output should have color. Default is true
	TrueColor       bool           // If the output should be true color or basic colors. Default is true if the terminal supports it
}

// Default value for true color
var defaultTrueColor = has256ColorSupport()

func NewCustomLogger() Logger {
	config := Logger{}
	config.NormalOut = os.Stdout
	config.ErrOut = os.Stderr
	config.ExtraNormalOuts = []io.Writer{}
	config.ExtraErrOuts = []io.Writer{}
	config.ExitCode = 1
	config.ShowStack = true
	config.Timezone = time.UTC
	config.Padding = true
	config.Multiline = false
	config.ColoredOutput = true
	config.TrueColor = defaultTrueColor
	return config
}

// Output a success log using a custom logger
func (config Logger) Success(ctx ...interface{}) {
	logNormal(config, successStatus, time.Now(), ctx...)
}

// Output an info log using a custom logger
func (config Logger) Info(ctx ...interface{}) {
	logNormal(config, infoStatus, time.Now(), ctx...)
}

// Output a debug log using a custom logger
func (config Logger) Debug(ctx ...interface{}) {
	logNormal(config, debugStatus, time.Now(), ctx...)
}

// Output a warning log using a custom logger
func (config Logger) Warning(ctx ...interface{}) {
	logNormal(config, warningStatus, time.Now(), ctx...)
}

// Output an error log using a custom logger
func (config Logger) Error(err error, ctx ...interface{}) {
	logError(config, errorStatus, time.Now(), err, ctx...)
}

// Output an error log with no actual error value using a custom logger
func (config Logger) ErrorMsg(ctx ...interface{}) {
	logError(config, errorStatus, time.Now(), nil, ctx...)
}

// Output a fatal log using a custom logger
func (config Logger) Fatal(err error, ctx ...interface{}) {
	logError(config, fatalStatus, time.Now(), err, ctx...)
	os.Exit(config.ExitCode)
}

// Output a fatal log with no actual error value using a custom logger
func (config Logger) FatalMsg(ctx ...interface{}) {
	logError(config, fatalStatus, time.Now(), nil, ctx...)
	os.Exit(config.ExitCode)
}
