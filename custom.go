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
	DebugStyle   lipgloss.Style
	InfoStyle    lipgloss.Style
	DoneStyle    lipgloss.Style
	WarningStyle lipgloss.Style
	ErrorStyle   lipgloss.Style
	FatalStyle   lipgloss.Style
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
		WarningStyle: l.normalRenderer.NewStyle().
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
func SetNormalOut(out *os.File) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.normalOut = out
	Logger.normalRenderer = *lipgloss.NewRenderer(out)
	// rerendering colors incase new output doesn't support colors
	Logger.colors.DebugStyle = Logger.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultDebugColor)).
		Bold(true)
	Logger.colors.InfoStyle = Logger.normalRenderer.NewStyle().Bold(true)
	Logger.colors.DoneStyle = Logger.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultDoneColor)).
		Bold(true)
	Logger.colors.WarningStyle = Logger.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultWarnColor)).
		Bold(true)
}

// Set the output or Fatal, FatalMsg, Error, and ErrorMsg.
//
// Default is os.Stderr
func SetErrOut(out *os.File) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.errOut = out
	Logger.errRenderer = *lipgloss.NewRenderer(out)
	// rerendering colors incase new output doesn't support colors
	Logger.colors.FatalStyle = Logger.errRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultFatalColor)).
		Bold(true)
	Logger.colors.ErrorStyle = Logger.errRenderer.NewStyle().
		Foreground(lipgloss.Color(DefaultErrorColor)).
		Bold(true)
}

// Set the extra normal out destinations (e.g. logging to a file)
func SetExtraNormalOuts(outs []io.Writer) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.extraNormalOuts = outs
}

// Set the extra normal out destinations (e.g. logging to a file)
func SetExtraErrOuts(outs []io.Writer) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.extraErrOuts = outs
}

// Set the exit code used by Fatal and FatalMsg.
//
// Default is 1
func SetFatalExitCode(code int) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.fatalExitCode = code
}

// Set if the stack trace should be shown or not when calling Error or Fatal.
//
// Default is true
func SetShowStack(show bool) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.showStack = show
}

// Set the time format
//
// Default is 2006/01/02 15:04:05 MST
func SetTimeFormat(format string) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.timeFormat = format
}

// Set the timezone
//
// Default is time.UTC
func SetTimezone(loc *time.Location) {
	Logger.mutex.Lock()
	defer Logger.mutex.Unlock()
	Logger.timezone = loc
}
