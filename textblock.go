package textblock

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sync"
	"unicode/utf8"
)

type TextBlock struct {
	content      []byte
	width        int
	height       int
	total_lines  int
	current_line int
	mu           sync.RWMutex
}

func Create(
	content []byte,
) *TextBlock {
	return &TextBlock{
		content:     content,
		width:       -1,
		height:      -1,
		total_lines: -1,
	}
}
func (self *TextBlock) GetOriginalContent() []byte {
	return self.content
}
func (self *TextBlock) GetContent() [][]byte {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.getContent()
}
func (self *TextBlock) GetWidth() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.width
}
func (self *TextBlock) GetHeight() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.height
}
func (self *TextBlock) GetLine() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.current_line
}
func (self *TextBlock) GetLines() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.total_lines
}
func (self *TextBlock) getContent() [][]byte {
	if self.width <= 0 ||
		self.height <= 0 {
		return nil
	}
	output := [][]byte{}
	start := 0
	x := 0
	lines := 0
	max_lines := self.current_line + self.height
	pos := 0
	l := len(self.content)
	for pos < l {
		c, size := utf8.DecodeRune(self.content[pos:])
		// we're not going to bother checking for invalid runes
		if c == '\r' ||
			c == '\n' {
			// new line character
			lines++
			if lines > self.current_line {
				if x > 0 {
					// append up to newline
					output = append(output, self.content[start:pos])
				} else {
					// append an empty newline
					output = append(output, nil)
				}
			}
			x = 0
			start = pos + size
		} else {
			// another character
			x++
			if x >= self.width {
				// new line
				lines++
				x = 0
				if lines > self.current_line {
					// append current line
					output = append(output, self.content[start:pos+size])
				}
				start = pos + size
			}
		}
		if lines >= max_lines {
			// no more lines needed
			return output
		}
		pos += size
	}
	if x > 0 {
		// append leftovers
		output = append(output, self.content[start:])
	}
	return output
}
func (self *TextBlock) Reset() {
	self.mu.Lock()
	self.resetMaybe(-1, -1)
	self.mu.Unlock()
}
func (self *TextBlock) ResetMaybe(
	width int,
	height int,
) bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	return self.resetMaybe(width, height)
}
func (self *TextBlock) resetMaybe(
	width int,
	height int,
) bool {
	if width <= 0 ||
		height <= 0 {
		// reset to unusable
		self.width = -1
		self.height = -1
		self.total_lines = -1
		self.current_line = 0
		return true
	}
	if self.width != width ||
		self.height != height {
		// reset
		self.width = width
		self.height = height
		self.current_line = 0
		self.countLines()
		return true
	}
	return false
}
func (self *TextBlock) HasAbove() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.current_line != 0
}
func (self *TextBlock) HasBelow() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	lowest_line := self.total_lines - self.height
	if lowest_line < 0 {
		lowest_line = 0
	}
	return self.current_line < lowest_line
}
func (self *TextBlock) jump() {
	lowest_line := self.total_lines - self.height
	if lowest_line < 0 {
		lowest_line = 0
	}
	if self.current_line > lowest_line {
		self.current_line = lowest_line
	}
	if self.current_line < 0 {
		self.current_line = 0
	}
}
func (self *TextBlock) Top() {
	self.mu.Lock()
	self.current_line = 0
	self.mu.Unlock()
}
func (self *TextBlock) Up() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.width <= 0 ||
		self.height <= 0 {
		return
	}
	self.current_line--
	self.jump()
}
func (self *TextBlock) JumpUp(
	i int,
) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.width <= 0 ||
		self.height <= 0 {
		return
	}
	self.current_line -= i
	self.jump()
}
func (self *TextBlock) JumpDown(
	i int,
) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.width <= 0 ||
		self.height <= 0 {
		return
	}
	self.current_line += i
	self.jump()
}
func (self *TextBlock) Down() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.width <= 0 ||
		self.height <= 0 {
		return
	}
	self.current_line++
	self.jump()
}
func (self *TextBlock) Bottom() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.width <= 0 ||
		self.height <= 0 {
		return
	}
	self.current_line = self.total_lines - self.height
	self.jump()
}
func (self *TextBlock) countLines() {
	self.total_lines = 0
	if self.width <= 0 ||
		self.height <= 0 {
		return
	}
	x := 0
	pos := 0
	l := len(self.content)
	for pos < l {
		c, size := utf8.DecodeRune(self.content[pos:])
		// we're not going to bother checking for invalid runes
		if c == '\r' ||
			c == '\n' {
			// new line character
			self.total_lines++
			x = 0
		} else {
			// another character
			x++
			if x >= self.width {
				// new line
				self.total_lines++
				x = 0
			}
		}
		pos += size
	}
	if x > 0 {
		// leftovers
		// this is a new line
		self.total_lines++
	}
}
