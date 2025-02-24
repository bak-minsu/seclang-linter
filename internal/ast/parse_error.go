package ast

import (
	"fmt"
	"strings"
)

type ParseError struct {
	// line number of error
	Line int

	// column number of error
	Column int

	// error message
	Message string

	// entire directive
	DirectiveContent string
}

func (e *ParseError) Error() string {
	var builder strings.Builder

	builder.WriteString(
		fmt.Sprintf(
			"Parse Error: %s",
			e.Message,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"\tline %d, column %d:",
			e.Line,
			e.Column,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"\t%s",
			e.DirectiveContent,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"\t%s^",
			strings.Repeat(" ", e.Column),
		),
	)

	return builder.String()
}
