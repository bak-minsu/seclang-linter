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

// returns offset end, exclusive
func (e *LinterError) OffsetEnd() int {
	return e.Offset + e.Distance
}

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

	contentLines := strings.Split(e.Contents, "\n")

	nonWhiteSpace := regexp.MustCompile(`\S`)

	lineStartOffset := 0
	offsetStart := e.Offset
	for lineNumber, subContent := range contentLines {
		// +1 to account for newline character
		lineEndOffset := lineStartOffset + len(subContent)

		if lineEndOffset < e.Offset {
			continue
		}

		lineIndicator := fmt.Sprintf(
			"line %d:",
			lineNumber,
		)

		builder.WriteString(
			fmt.Sprintf(
				"%s %s\n",
				lineIndicator,
				subContent,
			),
		)

		spaceContent := nonWhiteSpace.ReplaceAllString(subContent, " ")
		carrotContent := nonWhiteSpace.ReplaceAllString(subContent, "^")

		localStartOffset := offsetStart - lineStartOffset

		if e.OffsetEnd() <= lineEndOffset {
			underline := spaceContent[:localStartOffset] +
				carrotContent[localStartOffset:e.OffsetEnd()-lineStartOffset]

			builder.WriteString(
				fmt.Sprintf(
					"%s %s\n",
					strings.Repeat(" ", len(lineIndicator)),
					underline,
				),
			)

			break
		}

		underline := spaceContent[:localStartOffset] + carrotContent[localStartOffset:]

		builder.WriteString(
			fmt.Sprintf(
				"%s %s\n",
				strings.Repeat(" ", len(lineIndicator)),
				underline,
			),
		)

		offsetStart = lineEndOffset + 1
		lineStartOffset = lineEndOffset + 1
	}

	return builder.String()
}
