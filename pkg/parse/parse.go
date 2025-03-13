package parse

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Parses file structure using just the content
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

// Parses file structure using the given path of
// the file and reading its contents
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

// Parses file structure using the content of all
// files that match the glob pattern
func ParseGlob(patterns ...string) ([]*File, error) {
	matches := make([]string, 0)

	for _, pattern := range patterns {
		someMatches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid glob pattern: %w",
				err,
			)
		}

		matches = append(matches, someMatches...)
	}

	errs := make([]error, 0, len(matches))
	files := make([]*File, 0, len(matches))

	for _, match := range matches {
		fmt.Printf("validating file %s", match)

		parsedFile, err := ParseFile(match)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		files = append(files, parsedFile)

		fmt.Println(".........success!")
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf(
			"Linter errors: \n%w",
			errors.Join(errs...),
		)
	}

	return files, nil
}
