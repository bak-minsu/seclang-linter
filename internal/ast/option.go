package ast

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
	Value string

	// line number for option
	Line int

	// column index for the first character
	// of the option
	Column int
}

// returns the option tokens as the key and
// the column indices they were found on as the value
// within a map
func optionValues(line, column int, content string) (map[string]int, error) {
	optionsExpression := `(?:\\\n|\s)+` + // ignore space or escaped newspace before option
		`(?P<options>` + // start of 'options' capture group
		`"(?:\\"|[^"])+"` + // captures quotes syntax option
		`|` + // OR
		`[^ \n]+` + // captures a non-quote syntax option
		`)` // end of 'options' capture gorup

	compiled := regexp.MustCompile(optionsExpression)

	subMatches := compiled.FindAllStringSubmatchIndex(content, -1)
	if len(subMatches) == 0 {
		return nil, &ParseError{
			Line:             line,
			Column:           column,
			Message:          "no options found for this directive",
			DirectiveContent: content,
		}
	}

	subMatchIndex := compiled.SubexpIndex("options")
	if subMatchIndex == -1 {
		panic("unreachable condition - options subindex was not found")
	}

	subMatchIndex *= 2

	options := make(map[string]int, len(subMatches))

	for _, subMatch := range subMatches {
		left, right := subMatch[subMatchIndex], subMatch[subMatchIndex+1]

		match := content[left:right]

		options[match] = left
	}

	return options, nil
}

// parses option content into option.
// Examples:
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
func ParseOption(line, column int, optionContent string) *Option {
	return &Option{
		Value:  optionContent,
		Line:   line,
		Column: column,
	}
}

// parses content representing multiple options
// declared after a directive
func ParseOptions(line, startColumn int, content string) ([]*Option, error) {
	values, err := optionValues(line, startColumn, content)
	if err != nil {
		return nil, fmt.Errorf("could not parse options: %w", err)
	}

	options := make([]*Option, 0, len(values))
	for value, optionColumn := range values {
		options = append(options, ParseOption(
			line,
			optionColumn+startColumn,
			value,
		))
	}

	return options, nil
}
