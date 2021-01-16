package lumber

import (
	"testing"
	"time"

	"github.com/tj/assert"
)

func TestFormat(t *testing.T) {
	tt := []struct {
		stat                     status
		msg                      string
		expectedOutputTrueColor  string
		expectedOutputBasicColor string
		noPadding                bool
		noColor                  bool
	}{
		// Regular
		{
			stat:                     fatalStatus,
			msg:                      "Hello World!",
			expectedOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;255;0;0m    FATAL    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedOutputBasicColor: "    FATAL     | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:                     errorStatus,
			msg:                      "Hello World!",
			expectedOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;255;0;0m    ERROR    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedOutputBasicColor: "    ERROR     | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:                     warningStatus,
			msg:                      "Hello World!",
			expectedOutputTrueColor:  "\x1b[38;2;0;0;0m\x1b[48;2;253;255;0m   WARNING   \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedOutputBasicColor: "   WARNING    | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:                     infoStatus,
			msg:                      "Hello World!",
			expectedOutputTrueColor:  "\x1b[38;2;0;0;0m\x1b[48;2;255;255;255m    INFO     \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:                     debugStatus,
			msg:                      "Hello World!",
			expectedOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;0;97;255m    DEBUG    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedOutputBasicColor: "    DEBUG     | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		{
			stat:                     successStatus,
			msg:                      "Hello World!",
			expectedOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;0;176;99m   SUCCESS   \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedOutputBasicColor: "   SUCCESS    | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
		},
		// Padding Off
		{
			stat:                     infoStatus,
			msg:                      "Hey!",
			expectedOutputTrueColor:  "\x1b[38;2;0;0;0m\x1b[48;2;255;255;255m    INFO     \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			expectedOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			noPadding:                true,
		},
		// Color output off
		{
			stat:                     infoStatus,
			msg:                      "Hey!",
			expectedOutputTrueColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!\n",
			expectedOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!\n",
			noColor:                  true,
		},
		// Color output and padding off
		{
			stat:                     infoStatus,
			msg:                      "Hey!",
			expectedOutputTrueColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			expectedOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			noColor:                  true,
			noPadding:                true,
		},
	}

	for _, test := range tt {
		Padding = !test.noPadding
		ColoredOutput = !test.noColor

		TrueColor = true
		assert.Equal(t, test.expectedOutputTrueColor, format(test.stat, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), test.msg))

		TrueColor = false
		assert.Equal(t, test.expectedOutputBasicColor, format(test.stat, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), test.msg))

		TrueColor = true
		ColoredOutput = true
		Padding = true
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
