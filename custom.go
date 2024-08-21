package lumber

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var Logger logger = *DefaultLogger()

type logger struct {
	mutex           *sync.RWMutex
	normalOut       io.Writer // Output for Debug, Done, Warning, and Info. Default is os.Stdout
	errOut          io.Writer // Output for Fatal, FatalMsg, Error, and ErrorMsg. Default is os.Stderr
	normalRenderer  lipgloss.Renderer
	errRenderer     lipgloss.Renderer
	extraNormalOuts []io.Writer    // Extra normal output destinations (e.g. logging to a file)
	extraErrOuts    []io.Writer    // Extra error output destinations (e.g. logging to a file)
	fatalExitCode   int            // Fatal exit code. Default is 1
	showStack       bool           // If stack traces should be logged. Default is true
	timeFormat      string         // Time format of the logs. Defaults to 2006/01/02 15:04:05 MST
	timezone        *time.Location // Timezone for logging. Default is UTC
	colors          Colors         // Colors for the outputs
}

type Colors struct {
	DebugStyle lipgloss.Style
	InfoStyle  lipgloss.Style
	DoneStyle  lipgloss.Style
	WarnStyle  lipgloss.Style
	ErrorStyle lipgloss.Style
	FatalStyle lipgloss.Style
}

const (
	DefaultDebugColor = "#2B95FF"
	DefaultDoneColor  = "#30CE75"
	DefaultWarnColor  = "#E1DC3F"
	DefaultFatalColor = "#FF4747"
	DefaultErrorColor = "#FF4747"
)

// Initialize the default logger with the default values
func DefaultLogger() *logger {
	l := logger{
		mutex:           &sync.RWMutex{},
		normalOut:       os.Stdout,
		errOut:          os.Stderr,
		normalRenderer:  *lipgloss.NewRenderer(os.Stdout),
		errRenderer:     *lipgloss.NewRenderer(os.Stderr),
		extraNormalOuts: []io.Writer{},
		extraErrOuts:    []io.Writer{},
		fatalExitCode:   1,
		showStack:       true,
		timeFormat:      "01/02/2006 15:04:05 MST",
		timezone:        time.UTC,
	}
	l.colors = Colors{
		DebugStyle: l.normalRenderer.NewStyle().
			Foreground(lipgloss.Color(DefaultDebugColor)).
			Bold(true),
		InfoStyle: l.normalRenderer.NewStyle().Bold(true),
		WarnStyle: l.normalRenderer.NewStyle().
			Foreground(lipgloss.Color(DefaultWarnColor)).
			Bold(true),
		DoneStyle: l.normalRenderer.NewStyle().
			Foreground(lipgloss.Color(DefaultDoneColor)).
			Bold(true),
		FatalStyle: l.errRenderer.NewStyle().
			Foreground(lipgloss.Color(DefaultFatalColor)).
			Bold(true),
		ErrorStyle: l.errRenderer.NewStyle().
			Foreground(lipgloss.Color(DefaultErrorColor)).
			Bold(true),
	}
	return &l
}

// Set the output or Debug, Done, Warning, and Info.
//
// Default is os.Stdout
func (l *logger) SetNormalOut(out *os.File) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.normalOut = out
	l.normalRenderer = *lipgloss.NewRenderer(out)
	// rerendering colors incase new output doesn't support colors
	l.colors.DebugStyle = l.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultDebugColor)).
		Bold(true)
	l.colors.InfoStyle = l.normalRenderer.NewStyle().Bold(true)
	l.colors.DoneStyle = l.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultDoneColor)).
		Bold(true)
	l.colors.WarnStyle = l.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultWarnColor)).
		Bold(true)
}

// Set the output or Fatal, FatalMsg, Error, and ErrorMsg.
//
// Default is os.Stderr
func (l *logger) SetErrOut(out *os.File) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.errOut = out
	l.errRenderer = *lipgloss.NewRenderer(out)
	// rerendering colors incase new output doesn't support colors
	l.colors.FatalStyle = l.errRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultFatalColor)).
		Bold(true)
	l.colors.ErrorStyle = l.errRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultErrorColor)).
		Bold(true)
}

// Set the extra normal out destinations (e.g. logging to a file)
func (l *logger) SetExtraNormalOuts(outs []io.Writer) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.extraNormalOuts = outs
}

// Set the extra normal out destinations (e.g. logging to a file)
func (l *logger) SetExtraErrOuts(outs []io.Writer) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.extraErrOuts = outs
}

// Set the exit code used by Fatal and FatalMsg.
//
// Default is 1
func (l *logger) SetFatalExitCode(code int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.fatalExitCode = code
}

// Set if the stack trace should be shown or not when calling Error or Fatal.
//
// Default is true
func (l *logger) SetShowStack(show bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.showStack = show
}

// Set the time format
//
// Default is 2006/01/02 15:04:05 MST
func (l *logger) SetTimeFormat(format string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.timeFormat = format
}

// Set the timezone
//
// Default is time.UTC
func (l *logger) SetTimezone(loc *time.Location) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.timezone = loc
}
