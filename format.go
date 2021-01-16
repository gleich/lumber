package lumber

import (
	"fmt"
	"strings"
	"time"
)

// If the log should have an extra new line at the bottom
var Padding = true

// Raw formatting for lumber log
func format(stat status, t time.Time, message string) string {
	out := fmt.Sprintf(`%v | %v
%v
`,
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
