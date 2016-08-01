package textblock

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"sabey.co/unittest"
	"testing"
)

func TestTextBlock(t *testing.T) {
	fmt.Println("TestTextBlock")

	for _, bt := range []*blocktest{
		&blocktest{
			width:  1,
			height: 1,
		},
		&blocktest{
			content:     "a",
			width:       1,
			height:      1,
			total_lines: 1,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "a",
			width:       2,
			height:      2,
			total_lines: 1,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "ab",
			width:       1,
			height:      1,
			total_lines: 2,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "ab",
			width:       2,
			height:      2,
			total_lines: 1,
			output: []string{
				"ab",
			},
		},
		&blocktest{
			content:     "abc",
			width:       1,
			height:      1,
			total_lines: 3,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "abc",
			width:       2,
			height:      2,
			total_lines: 2,
			output: []string{
				"ab",
				"c",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       1,
			height:      1,
			total_lines: 4,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       1,
			height:      1,
			total_lines: 4,
			top:         true,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       1,
			height:      1,
			total_lines: 4,
			up:          1,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       1,
			height:      1,
			total_lines: 4,
			up:          2,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:      "abcd",
			width:        1,
			height:       1,
			total_lines:  4,
			down:         1,
			current_line: 1,
			has_above:    true,
			has_below:    true,
			output: []string{
				"b",
			},
		},
		&blocktest{
			content:      "abcd",
			width:        1,
			height:       1,
			total_lines:  4,
			down:         2,
			current_line: 2,
			has_above:    true,
			has_below:    true,
			output: []string{
				"c",
			},
		},
		&blocktest{
			content:      "abcd",
			width:        1,
			height:       1,
			total_lines:  4,
			down:         3,
			current_line: 3,
			has_above:    true,
			output: []string{
				"d",
			},
		},
		&blocktest{
			content:      "abcd",
			width:        1,
			height:       1,
			total_lines:  4,
			down:         4,
			current_line: 3,
			has_above:    true,
			output: []string{
				"d",
			},
		},
		&blocktest{
			content:      "abcd",
			width:        1,
			height:       1,
			total_lines:  4,
			bottom:       true,
			current_line: 3,
			has_above:    true,
			output: []string{
				"d",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       2,
			height:      2,
			total_lines: 2,
			output: []string{
				"ab",
				"cd",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       2,
			height:      2,
			total_lines: 2,
			down:        1,
			output: []string{
				"ab",
				"cd",
			},
		},
		&blocktest{
			content:     "abcd",
			width:       2,
			height:      2,
			total_lines: 2,
			down:        2,
			output: []string{
				"ab",
				"cd",
			},
		},
		&blocktest{
			content:     "abcde",
			width:       1,
			height:      1,
			total_lines: 5,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "abcde",
			width:       2,
			height:      2,
			total_lines: 3,
			has_below:   true,
			output: []string{
				"ab",
				"cd",
			},
		},
		&blocktest{
			content:     "\r",
			width:       1,
			height:      1,
			total_lines: 1,
			output: []string{
				"",
			},
		},
		&blocktest{
			content:     "\n",
			width:       1,
			height:      1,
			total_lines: 1,
			output: []string{
				"",
			},
		},
		&blocktest{
			content:     "\n\n",
			width:       1,
			height:      1,
			total_lines: 2,
			has_below:   true,
			output: []string{
				"",
			},
		},
		&blocktest{
			content:     "\n\n\n",
			width:       1,
			height:      1,
			total_lines: 3,
			has_below:   true,
			output: []string{
				"",
			},
		},
		&blocktest{
			content:     "a\n",
			width:       1,
			height:      1,
			total_lines: 2,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "a\nbc",
			width:       1,
			height:      1,
			total_lines: 4,
			has_below:   true,
			output: []string{
				"a",
			},
		},
		&blocktest{
			content:     "a\nbc",
			width:       2,
			height:      2,
			total_lines: 2,
			output: []string{
				"a",
				"bc",
			},
		},
		&blocktest{
			content:     "a\nbcd",
			width:       2,
			height:      2,
			total_lines: 3,
			has_below:   true,
			output: []string{
				"a",
				"bc",
			},
		},
		&blocktest{
			content:      "a\nbcd",
			width:        2,
			height:       2,
			total_lines:  3,
			down:         1,
			current_line: 1,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
		&blocktest{
			content:      "a\nbcd",
			width:        2,
			height:       2,
			total_lines:  3,
			down:         2,
			current_line: 1,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
		&blocktest{
			content:      "a\nbcd",
			width:        2,
			height:       2,
			total_lines:  3,
			down:         3,
			current_line: 1,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
		&blocktest{
			content:      "a\nbcde",
			width:        2,
			height:       2,
			total_lines:  3,
			down:         1,
			current_line: 1,
			has_above:    true,
			output: []string{
				"bc",
				"de",
			},
		},
		&blocktest{
			content:      "a\nbcde",
			width:        2,
			height:       2,
			total_lines:  3,
			down:         2,
			current_line: 1,
			has_above:    true,
			output: []string{
				"bc",
				"de",
			},
		},
		&blocktest{
			content:      "a\nbcdef",
			width:        2,
			height:       2,
			total_lines:  4,
			down:         1,
			current_line: 1,
			has_above:    true,
			has_below:    true,
			output: []string{
				"bc",
				"de",
			},
		},
		&blocktest{
			content:      "a\nbcdef",
			width:        2,
			height:       2,
			total_lines:  4,
			down:         2,
			current_line: 2,
			has_above:    true,
			output: []string{
				"de",
				"f",
			},
		},
		&blocktest{
			content:     "a\n\nbcd",
			width:       2,
			height:      2,
			total_lines: 4,
			has_below:   true,
			output: []string{
				"a",
				"",
			},
		},
		&blocktest{
			content:      "a\n\nbcd",
			width:        2,
			height:       2,
			total_lines:  4,
			down:         1,
			current_line: 1,
			has_above:    true,
			has_below:    true,
			output: []string{
				"",
				"bc",
			},
		},
		&blocktest{
			content:      "a\n\nbcd",
			width:        2,
			height:       2,
			total_lines:  4,
			down:         2,
			current_line: 2,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
		&blocktest{
			content:      "a\n\nbcd",
			width:        2,
			height:       2,
			total_lines:  4,
			down:         3,
			current_line: 2,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
		// jump
		&blocktest{
			content:      "a\n\nbcd",
			width:        2,
			height:       2,
			total_lines:  4,
			jumpdown:     1,
			current_line: 1,
			has_above:    true,
			has_below:    true,
			output: []string{
				"",
				"bc",
			},
		},
		&blocktest{
			content:      "a\n\nbcd",
			width:        2,
			height:       2,
			total_lines:  4,
			jumpdown:     2,
			current_line: 2,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
		&blocktest{
			content:      "a\n\nbcd",
			width:        2,
			height:       2,
			total_lines:  4,
			jumpdown:     3,
			current_line: 2,
			has_above:    true,
			output: []string{
				"bc",
				"d",
			},
		},
	} {
		tb := Create([]byte(bt.content))
		//fmt.Printf("\nwidth: %d height: %d top: %t bottom: %t jumpup: %d jumpdown: %d up: %d down: %d content: \"%s\"\n", bt.width, bt.height, bt.top, bt.bottom, bt.jumpup, bt.jumpdown, bt.up, bt.down, tb.GetOriginalContent())
		tb.ResetMaybe(bt.width, bt.height) // init
		// move after we set width/height
		if bt.bottom {
			tb.Bottom()
		}
		if bt.jumpdown > 0 {
			tb.JumpDown(bt.jumpdown)
		}
		if bt.jumpup > 0 {
			tb.JumpUp(bt.jumpup)
		}
		if bt.down > 0 {
			for i := 0; i < bt.down; i++ {
				tb.Down()
			}
		}
		if bt.up > 0 {
			for i := 0; i < bt.up; i++ {
				tb.Up()
			}
		}
		if bt.top {
			tb.Top()
		}
		// get output after movement
		output := tb.GetContent()
		//fmt.Printf("output: \"%s\"\n", output)
		unittest.Equals(t, tb.GetWidth(), bt.width)
		unittest.Equals(t, tb.GetHeight(), bt.height)
		unittest.Equals(t, tb.GetLines(), bt.total_lines)
		unittest.Equals(t, tb.GetLine(), bt.current_line)
		unittest.Equals(t, len(output), len(bt.output))
		unittest.Equals(t, tb.HasAbove(), bt.has_above)
		unittest.Equals(t, tb.HasBelow(), bt.has_below)
		for i, _ := range bt.output {
			unittest.Equals(t, string(output[i]), bt.output[i])
		}
	}
}

type blocktest struct {
	content      string
	width        int
	height       int
	total_lines  int
	bottom       bool
	jumpdown     int
	jumpup       int
	down         int
	up           int
	top          bool
	current_line int
	has_above    bool
	has_below    bool
	output       []string
}
