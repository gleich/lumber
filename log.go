package lumber

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/errors"
)

type logLevel string

const (
	debugLevel   logLevel = "DEBUG"
	doneLevel    logLevel = "DONE "
	infoLevel    logLevel = "INFO "
	warningLevel logLevel = "WARN "
	errorLevel   logLevel = "ERROR"
	fatalLevel   logLevel = "FATAL"
)

func format(level logLevel, color lipgloss.Style, v ...any) string {
	var joined string
	for _, item := range v {
		joined = fmt.Sprintf("%v %v", joined, item)
	}
	return fmt.Sprintf(
		"%s %s %s",
		time.Now().In(Logger.timezone).Format(Logger.timeFormat),
		color.Render(string(level)),
		strings.TrimPrefix(joined, " "),
	)
}

// Normal log output
func logNormal(level logLevel, color lipgloss.Style, v ...any) {
	Logger.mutex.RLock()
	defer Logger.mutex.RUnlock()
	out := format(level, color, v...)
	log.New(io.MultiWriter(append(Logger.extraNormalOuts, Logger.normalOut)...), "", 0).Println(out)
}

func logError(err error, level logLevel, color lipgloss.Style, v ...any) {
	Logger.mutex.RLock()
	defer Logger.mutex.RUnlock()
	out := format(level, color, v...)
	if err != nil && Logger.showStack {
		out += fmt.Sprintf("\n%+v", errors.WithStack(err))
	} else if err != nil {

		out += fmt.Sprintf("\n%s", err)
	}
	log.New(io.MultiWriter(append(Logger.extraErrOuts, Logger.errOut)...), "", 0).Println(out)
}

// Output a INFO log message
func Debug(v ...any) {
	logNormal(debugLevel, Logger.colors.DebugStyle, v...)
}

// Output a DONE log message
func Done(v ...any) {
	logNormal(doneLevel, Logger.colors.DoneStyle, v...)
}

// Output a INFO log message
func Info(v ...any) {
	logNormal(infoLevel, Logger.colors.InfoStyle, v...)
}

// Output a WARN log message
func Warning(v ...any) {
	logNormal(warningLevel, Logger.colors.WarningStyle, v...)
}

// Output a ERROR log message with information about the error
func Error(err error, v ...any) {
	logError(err, errorLevel, Logger.colors.ErrorStyle, v...)
}

// Output a ERROR log message
func ErrorMsg(v ...any) {
	logError(nil, errorLevel, Logger.colors.ErrorStyle, v...)
}

// Output a FATAL log message with information about the error
func Fatal(err error, v ...any) {
	logError(err, fatalLevel, Logger.colors.FatalStyle, v...)
	os.Exit(Logger.fatalExitCode)
}

// Output a FATAL log message
func FatalMsg(v ...any) {
	logError(nil, fatalLevel, Logger.colors.FatalStyle, v...)
	os.Exit(Logger.fatalExitCode)
}
