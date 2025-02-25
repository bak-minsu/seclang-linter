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
	Value string

	// line number for option
	Line int

	// column index for the first character
	// of the option
	Column int
}

// returns true if option content is valid
func optionsContentValid(content string) bool {
	optionsExpression := `^` + // start of expression
		`(?:` + // group representing a full option string
		`(?:\\\n|\\\r\n|[ ]+)` + // ignore space or escaped newspace before option
		`(?:` + // start of OR relation
		`"(?:\\"|[^"])+"` + // captures quotes syntax option
		`|` + // OR
		`[^\s]+` + // captures a non-quote syntax option
		`)` + // end of OR relation
		`)+` + // end of full option string
		`$` // end of expression

	return regexp.MustCompile(optionsExpression).MatchString(content)
}

// returns the option tokens as the key and
// the column indices they were found on as the value
// within a map
func optionValues(line, column int, content string) (map[string]int, error) {
	if !optionsContentValid(content) {
		return nil, &ParseError{
			Line:             line,
			ColumnStart:      column,
			ColumnEnd:        len(content),
			Message:          "invalid characters in option syntax",
			ParseLevel:       ParseLevelError,
			DirectiveContent: content,
		}
	}

	optionsExpression := `(?:\\\n|[ ]+)` + // ignore space or escaped newspace before option
		`(?P<options>` + // start of 'options' capture group
		`"(?:\\"|[^"])+"` + // captures quotes syntax option
		`|` + // OR
		`[^\s]+` + // captures a non-quote syntax option
		`)` // end of 'options' capture group

	compiled := regexp.MustCompile(optionsExpression)

	subMatches := compiled.FindAllStringSubmatchIndex(content, -1)
	if len(subMatches) == 0 {
		return nil, &ParseError{
			Line:             line,
			ColumnStart:      column,
			ColumnEnd:        len(content),
			ParseLevel:       ParseLevelError,
			Message:          "no options found for this directive",
			DirectiveContent: content,
		}
	}

	subMatchIndex := compiled.SubexpIndex("options")
	if subMatchIndex == -1 {
		panic("unreachable condition - options sub-index was not found")
	}

	subMatchIndex *= 2 // multiply since there are two indices per sub match

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
