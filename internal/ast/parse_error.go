package ast

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
	Column int

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
			e.Column,
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
			"\t%s^\n",
			strings.Repeat(" ", e.Column),
		),
	)

	return builder.String()
}
