package parse

import (
	"fmt"
	"strings"
)

const (
	ParseLevelPass = iota
	ParseLevelWarning
	ParseLevelError
)

type LinterError struct {
	// line number of error
	Line int

	// column start number of error
	ColumnStart int

	// column end number of error
	ColumnEnd int

	// error message
	Message string

	// indicator for parse error
	ParseLevel int

	// entire directive
	DirectiveContent string
}

func (e *LinterError) Error() string {
	var builder strings.Builder

	builder.WriteString(
		fmt.Sprintf(
			"\nParse Error: %s\n",
			e.Message,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"\tline %d, column %d:\n",
			e.Line,
			e.ColumnStart,
		),
	)

	lineStart := 0
	columnStart := e.ColumnStart
	columnEnd := e.ColumnEnd
	for _, line := range strings.Split(e.DirectiveContent, "\n") {
		// add one for missing new-line after split
		lineEnd := lineStart + len(line)

		builder.WriteString(
			fmt.Sprintf("\t%s\n", line),
		)

		switch {
		case columnStart < lineStart || columnStart > lineEnd:
			builder.WriteRune('\n')
		case columnEnd > lineEnd:
			builder.WriteString(
				fmt.Sprintf(
					"\t%s%s\n",
					strings.Repeat(" ", columnStart-lineStart),
					strings.Repeat("^", lineEnd-columnStart),
				),
			)

			columnStart = lineEnd + 1
		default:
			builder.WriteString(
				fmt.Sprintf(
					"\t%s%s\n",
					strings.Repeat(" ", columnStart-lineStart),
					strings.Repeat("^", columnEnd-columnStart),
				),
			)
		}

		lineStart = lineEnd + 1 // +1 for missing newline
	}

	return builder.String()
}
