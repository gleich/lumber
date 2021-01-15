package lumber

import (
	"testing"
	"time"

	"github.com/tj/assert"
)

func TestFormat(t *testing.T) {
	tt := []struct {
		stat           status
		msg            string
		expectedOutput string
		noPadding      bool
		noColor        bool
	}{
		// Regular
		{
			stat:           fatalStatus,
			msg:            "Hello World!",
			expectedOutput: "\x1b[41;37;1m    FATAL    \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:           errorStatus,
			msg:            "Hello World!",
			expectedOutput: "\x1b[41;37;1m    ERROR    \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:           warningStatus,
			msg:            "Hello World!",
			expectedOutput: "\x1b[43;30;1m   WARNING   \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:           infoStatus,
			msg:            "Hello World!",
			expectedOutput: "\x1b[47;30;1m    INFO     \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:           debugStatus,
			msg:            "Hello World!",
			expectedOutput: "\x1b[44;37;1m    DEBUG    \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:           successStatus,
			msg:            "Hello World!",
			expectedOutput: "\x1b[42;37;1m   SUCCESS   \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		// Padding Off
		{
			stat:           infoStatus,
			msg:            "Hey!",
			expectedOutput: "\x1b[47;30;1m    INFO     \x1b[0m | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			noPadding:      true,
		},
		// Color output off
		{
			stat:           infoStatus,
			msg:            "Hey!",
			expectedOutput: "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!\n",
			noColor:        true,
		},
		// Color output and padding off
		{
			stat:           infoStatus,
			msg:            "Hey!",
			expectedOutput: "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			noColor:        true,
			noPadding:      true,
		},
	}

	for _, test := range tt {
		Padding = !test.noPadding
		ColoredOutput = !test.noColor
		assert.Equal(t, test.expectedOutput, format(test.stat, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), test.msg))
	}
}

func TestSeperateWithSpaces(t *testing.T) {
	tt := []struct {
		items          []interface{}
		expectedOutput string
	}{
		{
			items:          []interface{}{"Hello", 9, false},
			expectedOutput: "Hello 9 false",
		},
		{
			items:          []interface{}{"Hello"},
			expectedOutput: "Hello",
		},
		{
			items:          []interface{}{9},
			expectedOutput: "9",
		},
		{
			items:          []interface{}{" ", 9, " "},
			expectedOutput: "  9  ",
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expectedOutput, separateWithSpaces(test.items...))
	}
}
