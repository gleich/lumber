package lumber

import (
	"fmt"
	"strings"
	"time"
)

// If the log should have an extra new line at the bottom
var Padding = true

// If the log should span multiple lines
var MultiLine = false

// Raw formatting for lumber log
func format(stat string, t time.Time, message string) string {
	template := "%v | %v | %v"
	if MultiLine {
		template = "%v | %v\n%v\n"
	}
	out := fmt.Sprintf(template,
		applyColor(stat, string(stat)),
		t.Format("Mon Jan 2 15:04:05 MST 2006"),
		message)

	if !Padding {
		out = strings.TrimSuffix(out, "\n")
	}

	return out
}

// Join all the items in an interface together with spaces
func separateWithSpaces(items ...interface{}) string {
	var joined string
	for _, item := range items {
		joined = fmt.Sprintf("%v %v", joined, item)
	}
	return strings.TrimPrefix(joined, " ")
}
