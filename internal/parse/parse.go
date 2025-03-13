package parse

import (
	"fmt"
	"os"
)

func Parse(content []byte) (*File, error) {
	directives, err := ParseDirectives(content)
	if err != nil {
		return nil, fmt.Errorf(
			"could not parse directives: %w",
			err,
		)
	}

	return &File{
		Directives: directives,
	}, nil
}

func ParseFile(name string) (*File, error) {
	contents, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf(
			"could not read file %q: %w",
			name,
			err,
		)
	}

	parsed, err := Parse(contents)
	if err != nil {
		return nil, fmt.Errorf(
			"could not parse read file contents: %w", err,
		)
	}

	return parsed, nil
}
