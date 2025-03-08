package parse

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	ParseLevelWarning = iota
	ParseLevelError
)

type LinterError struct {
	// error message
	Message string

	// indicator for parse error
	ParseLevel int

	// start index of error
	Offset int

	// distance to the end of the error
	Distance int

	// entire content
	Contents string
}

// returns offset value of the
// end of the content, exclusive
func (e *LinterError) OffsetEnd() int {
	return e.Offset + e.Distance
}

// Implements error interface
func (e *LinterError) Error() string {
	var builder strings.Builder

	// start on a new line
	builder.WriteRune('\n')

	switch e.ParseLevel {
	case ParseLevelError:
		builder.WriteString("Error: ")
	case ParseLevelWarning:
		builder.WriteString("Warning: ")
	}

	builder.WriteString(e.Message)
	builder.WriteRune('\n')

	patternNewline := regexp.MustCompile(`\n`)
	leftNewlines := patternNewline.FindAllStringIndex(e.Contents[:e.Offset], -1)

	column := e.Offset
	if len(leftNewlines) > 0 {
		column = e.Offset - leftNewlines[len(leftNewlines)-1][1]
	}

	builder.WriteString(
		fmt.Sprintf(
			"line %d, column %d:\n",
			len(leftNewlines),
			column,
		),
	)

	builder.WriteString(e.underlined())

	return builder.String()
}

// returns the error lines underlined with carrots
func (e *LinterError) underlined() string {
	var (
		builder              strings.Builder
		patternNewline       = regexp.MustCompile(`\n`)
		patternNotWhiteSpace = regexp.MustCompile(`\S`)
	)

	leftNewlines := patternNewline.FindAllStringIndex(e.Contents[:e.Offset], -1)
	innerNewlines := patternNewline.FindAllStringIndex(e.Contents[e.Offset:e.OffsetEnd()], -1)
	lastNewlines := patternNewline.FindAllStringIndex(e.Contents, len(leftNewlines)+len(innerNewlines)+1)

	lineStartOffset := 0
	if len(leftNewlines) != 0 {
		lineStartOffset = leftNewlines[len(leftNewlines)-1][1]
	}

	lineEndOffset := 0
	if len(lastNewlines) == len(leftNewlines)+len(innerNewlines) {
		lineEndOffset = len(e.Contents)
	} else {
		lineEndOffset = lastNewlines[len(lastNewlines)-1][0]
	}

	spaces := patternNotWhiteSpace.ReplaceAllString(e.Contents[lineStartOffset:e.Offset], " ")
	carrots := patternNotWhiteSpace.ReplaceAllString(e.Contents[e.Offset:e.OffsetEnd()], "^")

	underlineLines := strings.Split(spaces+carrots, "\n")
	contentLines := strings.Split(e.Contents[lineStartOffset:lineEndOffset], "\n")

	for i := range len(contentLines) {
		builder.WriteString(contentLines[i] + "\n")
		builder.WriteString(underlineLines[i] + "\n")
	}

	return builder.String()
}
