package lumber

import (
	"fmt"
	"strings"
	"time"
)

// Raw formatting for lumber log
func format(config Logger, stat string, t time.Time, message string) string {
	template := "%v | %v | %v"
	if config.Multiline {
		template = "%v | %v\n%v\n"
	}
	timeInLoc := t.In(config.Timezone)
	out := fmt.Sprintf(template,
		applyColor(config, stat, string(stat)),
		timeInLoc.Format("Mon Jan 2 15:04:05 MST 2006"),
		message)

	if !config.Padding {
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
