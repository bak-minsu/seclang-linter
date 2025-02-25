package parse

import (
	"fmt"
	"regexp"
)

// Represents the root of the parse tree
type Root struct {
	Directives []*Directive
}

func Parse(content []byte) (*Root, error) {
	splitter := regexp.MustCompile(`[^\\](\n|\r\n)`)

	split := splitter.Split(string(content), -1)

	directives, err := ParseDirectives(split)
	if err != nil {
		return nil, fmt.Errorf(
			"could not parse directives: %w",
			err,
		)
	}

	return &Root{
		Directives: directives,
	}, nil
}
