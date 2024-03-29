package lumber

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	truecolor "github.com/wayneashleyberry/truecolor/pkg/color"
)

// Checking to see if the terminal has true color support
func has256ColorSupport() bool {
	envColor := os.Getenv("TERM")
	if envColor == "xterm-256color" {
		return true
	}

	tputExecPath, err := exec.LookPath("tput")
	if err != nil {
		return false
	}

	support, err := exec.Command(tputExecPath, "colors").Output()
	if err != nil || string(support) != "256\n" {
		return false
	}
	return true
}

// Apply the color for a given status to a string
func applyColor(config Logger, stat string, msg string) string {
	trueColorMap := map[string]*truecolor.Message{
		successStatus: truecolor.White().Background(0, 176, 99),    // Green
		fatalStatus:   truecolor.White().Background(255, 0, 0),     // Red
		errorStatus:   truecolor.White().Background(255, 0, 0),     // Red
		warningStatus: truecolor.Black().Background(253, 255, 0),   // Yellow
		infoStatus:    truecolor.Black().Background(255, 255, 255), // White
		debugStatus:   truecolor.White().Background(0, 97, 255),    // Blue
	}

	basicColorMap := map[string][]color.Attribute{
		successStatus: {color.BgGreen, color.FgWhite},  // Green
		fatalStatus:   {color.BgRed, color.FgWhite},    // Red
		errorStatus:   {color.BgRed, color.FgWhite},    // Red
		warningStatus: {color.BgYellow, color.FgBlack}, // Yellow
		infoStatus:    {color.BgWhite, color.FgBlack},  // White
		debugStatus:   {color.BgBlue, color.FgWhite},   // Blue
	}

	if config.ColoredOutput {
		if config.TrueColor {
			return color.New(color.Bold).Sprint(trueColorMap[stat].Sprint(msg))
		}
		return color.New(basicColorMap[stat]...).Sprint(msg)
	}
	return msg
}
