package tomlseq

import (
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
	var rb []byte
	var sort int
	var mlsTracker bool

	for i, b := range bs {
		if i < 4 {
			rb = append(rb, b)
			continue
		}
		if b == '"' {
			if bs[i-1] == '"' && bs[i-2] == '"' {
				mlsTracker = !mlsTracker
			}
			rb = append(rb, b)
			continue
		}
		if mlsTracker {
			rb = append(rb, b)
			continue
		}

		// Logic for building sequence ids
		if bs[i-1] == '\n' && bs[i-2] == ']' && bs[i-3] == ']' {
			rb = append(rb, []byte(fmt.Sprintf("%s = "+strconv.Itoa(sort)+"\n", id))...)
			rb = append(rb, b)
			sort++
			continue
		}
		rb = append(rb, b)
	}
	return rb
}
