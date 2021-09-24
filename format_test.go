package lumber

import (
	"testing"
	"time"

	"github.com/tj/assert"
)

func TestFormat(t *testing.T) {
	tt := []struct {
		stat                               string
		msg                                string
		expectedMultilineOutputTrueColor   string
		expectedMultilineOutputBasicColor  string
		expectedSinglelineOutputTrueColor  string
		expectedSinglelineOutputBasicColor string
		noPadding                          bool
		noColor                            bool
	}{
		// Regular
		{
			stat:                               fatalStatus,
			msg:                                "Hello World!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;255;255;255m\x1b[48;2;255;0;0m    FATAL    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedMultilineOutputBasicColor:  "    FATAL     | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;255;0;0m    FATAL    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
			expectedSinglelineOutputBasicColor: "    FATAL     | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
		},
		{
			stat:                               errorStatus,
			msg:                                "Hello World!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;255;255;255m\x1b[48;2;255;0;0m    ERROR    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedMultilineOutputBasicColor:  "    ERROR     | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;255;0;0m    ERROR    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
			expectedSinglelineOutputBasicColor: "    ERROR     | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
		},
		{
			stat:                               warningStatus,
			msg:                                "Hello World!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;0;0;0m\x1b[48;2;253;255;0m   WARNING   \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedMultilineOutputBasicColor:  "   WARNING    | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;0;0;0m\x1b[48;2;253;255;0m   WARNING   \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
			expectedSinglelineOutputBasicColor: "   WARNING    | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
		},
		{
			stat:                               infoStatus,
			msg:                                "Hello World!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;0;0;0m\x1b[48;2;255;255;255m    INFO     \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedMultilineOutputBasicColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;0;0;0m\x1b[48;2;255;255;255m    INFO     \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
			expectedSinglelineOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
		},
		{
			stat:                               debugStatus,
			msg:                                "Hello World!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;255;255;255m\x1b[48;2;0;97;255m    DEBUG    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedMultilineOutputBasicColor:  "    DEBUG     | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;0;97;255m    DEBUG    \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
			expectedSinglelineOutputBasicColor: "    DEBUG     | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
		},
		{
			stat:                               successStatus,
			msg:                                "Hello World!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;255;255;255m\x1b[48;2;0;176;99m   SUCCESS   \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedMultilineOutputBasicColor:  "   SUCCESS    | Tue Nov 30 00:00:00 UTC -0001\nHello World!\n",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;255;255;255m\x1b[48;2;0;176;99m   SUCCESS   \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
			expectedSinglelineOutputBasicColor: "   SUCCESS    | Tue Nov 30 00:00:00 UTC -0001 | Hello World!",
		},
		// Padding Off
		{
			stat:                               infoStatus,
			msg:                                "Hey!",
			expectedMultilineOutputTrueColor:   "\x1b[38;2;0;0;0m\x1b[48;2;255;255;255m    INFO     \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			expectedMultilineOutputBasicColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			expectedSinglelineOutputTrueColor:  "\x1b[38;2;0;0;0m\x1b[48;2;255;255;255m    INFO     \x1b[39m\x1b[49m | Tue Nov 30 00:00:00 UTC -0001 | Hey!",
			expectedSinglelineOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001 | Hey!",
			noPadding:                          true,
		},
		// Color output off
		{
			stat:                               infoStatus,
			msg:                                "Hey!",
			expectedMultilineOutputTrueColor:   "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!\n",
			expectedMultilineOutputBasicColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!\n",
			expectedSinglelineOutputTrueColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001 | Hey!",
			expectedSinglelineOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001 | Hey!",
			noColor:                            true,
		},
		// Color output and padding off
		{
			stat:                               infoStatus,
			msg:                                "Hey!",
			expectedMultilineOutputTrueColor:   "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			expectedMultilineOutputBasicColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001\nHey!",
			expectedSinglelineOutputTrueColor:  "    INFO      | Tue Nov 30 00:00:00 UTC -0001 | Hey!",
			expectedSinglelineOutputBasicColor: "    INFO      | Tue Nov 30 00:00:00 UTC -0001 | Hey!",
			noColor:                            true,
			noPadding:                          true,
		},
	}

	testTime := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for _, test := range tt {
		Padding = !test.noPadding
		ColoredOutput = !test.noColor

		MultiLine = false

		TrueColor = true
		assert.Equal(
			t,
			test.expectedSinglelineOutputTrueColor,
			format(test.stat, testTime, test.msg),
		)

		TrueColor = false
		assert.Equal(
			t,
			test.expectedSinglelineOutputBasicColor,
			format(test.stat, testTime, test.msg),
		)

		MultiLine = true

		TrueColor = true
		assert.Equal(
			t,
			test.expectedMultilineOutputTrueColor,
			format(test.stat, testTime, test.msg),
		)

		TrueColor = false
		assert.Equal(
			t,
			test.expectedMultilineOutputBasicColor,
			format(test.stat, testTime, test.msg),
		)

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
