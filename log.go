package lumber

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type logLevel string

const (
	debugLevel logLevel = "DEBUG"
	doneLevel  logLevel = "DONE "
	infoLevel  logLevel = "INFO "
	warnLevel  logLevel = "WARN "
)

// Normal log output
func logNormal(level logLevel, color lipgloss.Style, v ...any) {
	Logger.mutex.RLock()
	defer Logger.mutex.RUnlock()
	out := fmt.Sprintf(
		"%s %s %s",
		time.Now().Format(Logger.timeFormat),
		color.Render(string(level)),
		separateWithSpaces(v...),
	)
	log.New(io.MultiWriter(append(Logger.extraNormalOuts, Logger.normalOut)...), "", 0).Println(out)
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
func Warn(v ...any) {
	logNormal(warnLevel, Logger.colors.WarnStyle, v...)
}
