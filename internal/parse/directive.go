package parse

import (
	"fmt"
	"regexp"
	"strings"
)

// all possible directive lexems
const (
	DirectiveInclude                     = "Include"
	DirectiveSecAction                   = "SecAction"
	DirectiveSecArgumentsLimit           = "SecArgumentsLimit"
	DirectiveSecAuditEngine              = "SecAuditEngine"
	DirectiveSecAuditLog                 = "SecAuditLog"
	DirectiveSecAuditLogDir              = "SecAuditLogDir"
	DirectiveSecAuditLogDirMode          = "SecAuditLogDirMode"
	DirectiveSecAuditLogFileMode         = "SecAuditLogFileMode"
	DirectiveSecAuditLogFormat           = "SecAuditLogFormat"
	DirectiveSecAuditLogParts            = "SecAuditLogParts"
	DirectiveSecAuditLogRelevantStatus   = "SecAuditLogRelevantStatus"
	DirectiveSecDebugLog                 = "SecDebugLog"
	DirectiveSecDebugLogLevel            = "SecDebugLogLevel"
	DirectiveSecDefaultAction            = "SecDefaultAction"
	DirectiveSecMarker                   = "SecMarker"
	DirectiveSecRequestBodyAccess        = "SecRequestBodyAccess"
	DirectiveSecRequestBodyInMemoryLimit = "SecRequestBodyInMemoryLimit"
	DirectiveSecRequestBodyLimit         = "SecRequestBodyLimit"
	DirectiveSecRequestBodyLimitAction   = "SecRequestBodyLimitAction"
	DirectiveSecRequestBodyNoFilesLimit  = "SecRequestBodyNoFilesLimit"
	DirectiveSecResponseBodyAccess       = "SecResponseBodyAccess"
	DirectiveSecResponseBodyLimit        = "SecResponseBodyLimit"
	DirectiveSecResponseBodyLimitAction  = "SecResponseBodyLimitAction"
	DirectiveSecRule                     = "SecRule"
	DirectiveSecRuleEngine               = "SecRuleEngine"
	DirectiveSecRuleRemoveByID           = "SecRuleRemoveByID"
	DirectiveSecRuleRemoveByTag          = "SecRuleRemoveByTag"
	DirectiveSecRuleUpdateTargetByID     = "SecRuleUpdateTargetByID"
	DirectiveSecRuleUpdateTargetByTag    = "SecRuleUpdateTargetByTag"
)

// returns all possible directive lexemes as a string slice.
func DirectiveLexemes() []string {
	return []string{
		DirectiveInclude,
		DirectiveSecAction,
		DirectiveSecArgumentsLimit,
		DirectiveSecAuditEngine,
		DirectiveSecAuditLog,
		DirectiveSecAuditLogDir,
		DirectiveSecAuditLogDirMode,
		DirectiveSecAuditLogFileMode,
		DirectiveSecAuditLogFormat,
		DirectiveSecAuditLogParts,
		DirectiveSecAuditLogRelevantStatus,
		DirectiveSecDebugLog,
		DirectiveSecDebugLogLevel,
		DirectiveSecDefaultAction,
		DirectiveSecMarker,
		DirectiveSecRequestBodyAccess,
		DirectiveSecRequestBodyInMemoryLimit,
		DirectiveSecRequestBodyLimit,
		DirectiveSecRequestBodyLimitAction,
		DirectiveSecRequestBodyNoFilesLimit,
		DirectiveSecResponseBodyAccess,
		DirectiveSecResponseBodyLimit,
		DirectiveSecResponseBodyLimitAction,
		DirectiveSecRule,
		DirectiveSecRuleEngine,
		DirectiveSecRuleRemoveByID,
		DirectiveSecRuleRemoveByTag,
		DirectiveSecRuleUpdateTargetByID,
		DirectiveSecRuleUpdateTargetByTag,
	}
}

// represents the SecLang directive
type Directive struct {
	// string representing the directive.
	// ex. "Include"
	Lexeme string

	// offset within the entire file
	Offset int

	// options for the directive
	Options []*Option
}

func (d *Directive) Len() int {
	lastOption := d.Options[len(d.Options)-1]

	directiveEnd := lastOption.Offset + lastOption.Len()

	return directiveEnd - d.Offset
}

// parses a given directive from string
func ParseDirective(contents []byte, offset int) (*Directive, error) {
	patternDirective := regexp.MustCompile(`^[[:alpha:]]+`)

	matchIndex := patternDirective.FindIndex(contents[offset:])
	if matchIndex != nil {
		options, err := ParseOptions(contents, offset+matchIndex[1])
		if err != nil {
			return nil, fmt.Errorf(
				"could not parse options: %w",
				err,
			)
		}

		offsetContent := contents[offset:]

		lexeme := offsetContent[matchIndex[0]:matchIndex[1]]

		return &Directive{
			Lexeme:  string(lexeme),
			Offset:  offset,
			Options: options,
		}, nil
	}

	return nil, &LinterError{
		Offset:     offset,
		Distance:   1,
		Message:    "expected alphabetic characters for directive",
		ParseLevel: ParseLevelError,
		Contents:   string(contents),
	}
}

// parses a given directive from byte array
func ParseDirectives(contents []byte) ([]*Directive, error) {
	if len(contents) == 0 {
		return nil, nil
	}

	// use the number of lines as a guess to how big
	// the directives capacity should be
	lines := strings.Count(string(contents), "\n") + 1

	directives, err := parseDirectives(
		contents,
		0,
		make([]*Directive, 0, lines),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"problems while traversing directive listing: %w",
			err,
		)
	}

	if len(directives) == 0 {
		return nil, nil
	}

	return directives, nil
}

// recursive helper to ParseDirectives.
// Content represents the entire read content,
// offset represents the character index within read content,
// and directives is the array of directives that will be output
// once read is complete.
func parseDirectives(content []byte, offset int, directives []*Directive) ([]*Directive, error) {
	if offset >= len(content) {
		return directives, nil
	}

	var (
		patternWhitespace = regexp.MustCompile(`^\s+`)
		patternComment    = regexp.MustCompile(`^#[^\r\n]*`)
		patternDirective  = regexp.MustCompile(`^[[:alpha:]]+`)
	)

	offsetContents := content[offset:]

	if matchIndices := patternWhitespace.FindIndex(offsetContents); matchIndices != nil {
		return parseDirectives(
			content,
			offset+matchIndices[1],
			directives,
		)
	}

	if matchIndices := patternComment.FindIndex(offsetContents); matchIndices != nil {
		return parseDirectives(
			content,
			offset+matchIndices[1],
			directives,
		)
	}

	if patternDirective.Match(offsetContents) {
		directive, err := ParseDirective(content, offset)
		if err != nil {
			return nil, fmt.Errorf(
				"could not parse directive: %w",
				err,
			)
		}

		return parseDirectives(
			content,
			offset+directive.Len(),
			append(directives, directive),
		)
	}

	return nil, &LinterError{
		Offset:     offset,
		Distance:   1,
		Message:    "unexpected token while attempting to read directive",
		ParseLevel: ParseLevelError,
		Contents:   string(content),
	}
}
