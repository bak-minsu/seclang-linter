package parse

import (
	"fmt"
	"regexp"
)

// represents a single option of a SecLang directive
type Option struct {
	// string value representing the option.
	// ex.
	//
	//   - "quoted option"
	//
	//   - unquotedOption
	//
	//   - multi-line options \
	//
	//     "with multi-line \
	//
	//     quotes"
	Lexeme string

	// offset number for option
	Offset int
}

func (o *Option) Len() int {
	return len(o.Lexeme)
}

// parses non quoted option content into option object
func ParseOptionNotQuoted(contents []byte, offset int) (*Option, error) {
	if offset >= len(contents) {
		return nil, &LinterError{
			Message:    "EOF - expected unquoted option content",
			ParseLevel: ParseLevelError,
			Offset:     offset,
			Distance:   1,
			Contents:   string(contents),
		}
	}

	offsetContent := contents[offset:]

	patternNotWhitespace := regexp.MustCompile(`^\S+`)

	if matchedIndices := patternNotWhitespace.FindIndex(offsetContent); matchedIndices != nil {
		match := offsetContent[matchedIndices[0]:matchedIndices[1]]

		return &Option{
			Lexeme: string(match),
			Offset: offset + matchedIndices[0],
		}, nil
	}

	return nil, &LinterError{
		Message:    "found unexpected whitepsace while scanning unquoted option syntax",
		ParseLevel: ParseLevelError,
		Offset:     offset,
		Distance:   1,
		Contents:   string(contents),
	}
}

// parses non quoted option content into option object
func ParseOptionQuoted(contents []byte, offset int) (*Option, error) {
	if offset >= len(contents) {
		return nil, &LinterError{
			Message:    "EOF - expected quoted option content",
			ParseLevel: ParseLevelError,
			Offset:     offset,
			Distance:   1,
			Contents:   string(contents),
		}
	}

	offsetContent := contents[offset:]

	patternUntilQuote := regexp.MustCompile(`^"(\\"|[^"])+"`)

	matchedIndices := patternUntilQuote.FindIndex(offsetContent)
	if matchedIndices != nil {
		match := offsetContent[matchedIndices[0]:matchedIndices[1]]

		return &Option{
			Lexeme: string(match),
			Offset: offset + matchedIndices[0],
		}, nil
	}

	return nil, &LinterError{
		Message:    "unexpected sequence while scanning quoted option syntax",
		ParseLevel: ParseLevelError,
		Offset:     offset,
		Distance:   1,
		Contents:   string(contents),
	}
}

// parses content representing multiple options
// declared after a directive
func ParseOptions(contents []byte, offset int) ([]*Option, error) {
	options, err := parseOptions(
		contents,
		offset,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"trouble parsing options: %w",
			err,
		)
	}

	return options, nil
}

func parseOptions(contents []byte, offset int, options []*Option) ([]*Option, error) {
	if offset >= len(contents) {
		return options, nil
	}

	var (
		patternNewline   = regexp.MustCompile(`^\n`)
		patternSkip      = regexp.MustCompile(`^(\\\n| )`)
		patternNotQuoted = regexp.MustCompile(`^[^"]`)
		patternQuoted    = regexp.MustCompile(`^"`)
	)

	offsetContent := contents[offset:]

	if matchIndex := patternSkip.FindIndex(offsetContent); matchIndex != nil {
		return parseOptions(
			contents,
			offset+int(matchIndex[1]),
			options,
		)
	}

	if matchIndex := patternNewline.FindIndex(offsetContent); matchIndex != nil {
		return options, nil
	}

	if matchIndex := patternNotQuoted.FindIndex(offsetContent); matchIndex != nil {
		option, err := ParseOptionNotQuoted(contents, offset)
		if err != nil {
			return nil, fmt.Errorf("could not parse option: %w", err)
		}

		return parseOptions(
			contents,
			offset+option.Len(),
			append(options, option),
		)
	}

	if matchIndex := patternQuoted.FindIndex(offsetContent); matchIndex != nil {
		option, err := ParseOptionQuoted(contents, offset)
		if err != nil {
			return nil, fmt.Errorf("could not parse option: %w", err)
		}

		return parseOptions(
			contents,
			offset+option.Len(),
			append(options, option),
		)
	}

	return nil, &LinterError{
		Message:    "unexpected token",
		Offset:     offset,
		Distance:   1,
		ParseLevel: ParseLevelError,
		Contents:   string(contents),
	}
}
