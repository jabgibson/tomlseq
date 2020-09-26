package tomlseq

import (
	"bytes"
	"fmt"
	"strconv"
)

// Process iterates over each line starting with "[[" (the beginning of a toml table) and writes the current iteration
// index to the next line starting with configured identifier. This method also attempts to ignore comments and multiline
// strings, focusing on tables only.
func Process(identifier string, data []byte) []byte {
	return process(identifier, data)
}

func process(id string, bs []byte) []byte {
	bs = bytes.Replace(bs, []byte("\r\n"), []byte("\n"), -1)
	var rb []byte
	var tableCount int
	var mlsTracker bool
	var commentTracker bool

	for i, b := range bs {
		if i < 4 {
			rb = append(rb, b)
			continue
		}
		// if starting or ending a multiline string, set mlsTracker
		if b == '"' {
			if bs[i-1] == '"' && bs[i-2] == '"' {
				mlsTracker = !mlsTracker
			}
			rb = append(rb, b)
			continue
		}

		// if within a multiline string or a line comment, continue without logic
		if mlsTracker || (commentTracker && b != '\n') {
			rb = append(rb, b)
			continue
		}

		// flip commentTracker to false if line comment is ending
		if b == '\n' && commentTracker {
			commentTracker = false
			rb = append(rb, b)
			continue
		}

		// flip commentTracker to true if a line comment is starting
		if b == '#' && bs[i-1] == '\n' {
			commentTracker = true
			rb = append(rb, b)
			continue
		}

		// Logic for building sequence ids
		if bs[i-1] == '\n' && bs[i-2] == ']' && bs[i-3] == ']' {
			rb = append(rb, []byte(fmt.Sprintf("%s = "+strconv.Itoa(tableCount)+"\n", id))...)
			rb = append(rb, b)
			tableCount++
			continue
		}
		rb = append(rb, b)
	}
	return rb
}
