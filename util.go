package lumber

import (
	"fmt"
	"strings"
)

// Join all of the items together with a space
func separateWithSpaces(v ...any) string {
	var joined string
	for _, item := range v {
		joined = fmt.Sprintf("%v %v", joined, item)
	}
	return strings.TrimPrefix(joined, " ")
}
