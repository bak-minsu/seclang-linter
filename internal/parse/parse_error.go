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

type ParseError struct {
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

func (e *ParseError) Error() string {
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

	builder.WriteString(
		fmt.Sprintf(
			"\t%s\n",
			e.DirectiveContent,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"\t%s%s\n",
			strings.Repeat(" ", e.ColumnStart),
			strings.Repeat("^", e.ColumnEnd-e.ColumnStart),
		),
	)

	return builder.String()
}
