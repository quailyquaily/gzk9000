package loop

import (
	"strings"
)

// SplitMarkdown splits the given markdown text into multiple chunks,
// each no more than maxBytes in size. The splitting is done by lines.
//
//   - We do not break a single line in the middle.
//   - If we are in the middle of a fenced code block (``` or ~~~)
//     and need to start a new chunk, we "close" the code block in the
//     current chunk and "re-open" it in the next.
//
// Returns a slice of chunked Markdown.
func SplitMarkdown(markdown string, maxBytes int) []string {
	lines := strings.Split(markdown, "\n")
	var chunks []string

	var currentChunk strings.Builder
	currentSize := 0

	inCodeBlock := false
	fenceMarker := ""   // e.g. ``` or ~~~
	fenceLanguage := "" // e.g. "js"

	for i, line := range lines {
		// Include the newline in all but the last line
		lineToAdd := line
		if i < len(lines)-1 {
			lineToAdd += "\n"
		}
		lineLen := len(lineToAdd)

		// If adding this line would exceed maxBytes, start a new chunk:
		if currentSize+lineLen > maxBytes && currentSize > 0 {
			// If we are currently in a code block, close it
			if inCodeBlock && fenceMarker != "" {
				currentChunk.WriteString(fenceMarker)
				currentChunk.WriteString("\n")
			}

			chunks = append(chunks, currentChunk.String())

			// Start a new chunk
			currentChunk.Reset()
			currentSize = 0

			// If we were in a code block, reopen it
			if inCodeBlock && fenceMarker != "" {
				// Reopen fence with language if we had one
				currentChunk.WriteString(fenceMarker)
				if fenceLanguage != "" {
					currentChunk.WriteString(fenceLanguage)
				}
				currentChunk.WriteString("\n")
				currentSize += len(fenceMarker) + len(fenceLanguage) + 1
			}
		}

		// Add the current line to this chunk
		currentChunk.WriteString(lineToAdd)
		currentSize += lineLen

		// Check if this line is a fence line (opening or closing)
		if m, lang, isFence := detectFenceLine(line); isFence {
			// If we are not in a code block, we are opening one
			if !inCodeBlock {
				inCodeBlock = true
				fenceMarker = m
				fenceLanguage = lang
			} else {
				// We are closing one
				inCodeBlock = false
				fenceMarker = ""
				fenceLanguage = ""
			}
		}
	}

	// If something remains in currentChunk, finalize it
	if currentChunk.Len() > 0 {
		// If still in a code block, close it
		if inCodeBlock && fenceMarker != "" {
			currentChunk.WriteString(fenceMarker)
			currentChunk.WriteString("\n")
		}
		chunks = append(chunks, currentChunk.String())
	}

	return chunks
}

// detectFenceLine checks if a line is a code fence line, such as:
//
//	```
//	```js
//	~~~
//	~~~python
//
// and returns the fence marker ("```" or "~~~"), the language (e.g. "js", "python"),
// and a bool indicating whether it is indeed a fence line.
//
// This is a simplified function; for complex or unusual fences, you'd expand logic.
func detectFenceLine(line string) (fence string, language string, isFence bool) {
	trimmed := strings.TrimSpace(line)

	// Possible fences to detect
	fences := []string{"```", "~~~"}

	for _, f := range fences {
		if strings.HasPrefix(trimmed, f) {
			// Extract whatever is after the fence as language
			lang := strings.TrimSpace(trimmed[len(f):])
			return f, lang, true
		}
	}
	return "", "", false
}
