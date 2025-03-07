package parse

import (
	"fmt"
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
