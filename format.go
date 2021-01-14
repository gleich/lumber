package lumber

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Global options available to a user to turn off.
var (
	Padding       = true
	ColoredOutput = true
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

// Raw formatting for lumber log
func format(stat status, t time.Time, message string) string {
	color.NoColor = !ColoredOutput

	statusColorMap := map[status][]color.Attribute{
		fatalStatus:   {color.BgRed, color.FgWhite, color.Bold},
		successStatus: {color.BgGreen, color.FgWhite, color.Bold},
		errorStatus:   {color.BgRed, color.FgWhite, color.Bold},
		warningStatus: {color.BgYellow, color.FgBlack, color.Bold},
		infoStatus:    {color.BgWhite, color.FgBlack, color.Bold},
		debugStatus:   {color.BgBlue, color.FgWhite, color.Bold},
	}
	col := color.New(statusColorMap[stat]...)

	out := fmt.Sprintf(`

%v %v
%v
`,
		col.Sprintf(string(stat)),
		t.Format("Mon Jan 2 15:04:05 MST 2006"),
		message)

	if !Padding {
		out = strings.Trim(out, "\n")
	}

	return out
}
