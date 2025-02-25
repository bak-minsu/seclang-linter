package parse

import (
	"fmt"
	"regexp"
	"strings"
)

// all possible directives
const (
	DirectiveEmpty                       = ""
	DirectiveComment                     = "#"
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

// represents the SecLang directive
type Directive struct {
	// Token representing the directive.
	// ex. "Include"
	Token string

	// Entire directive content
	Content string

	// Line number for directive
	Line int
	// Column index for the first character
	// of directive token
	Column int

	// Options for the directive
	Options []*Option
}

// returns all possible directive tokens as a string slice.
func DirectiveTokens() []string {
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

// returns the directive token and the column index it was found on.
// Returns an error if incorrectly formatted
func directiveToken(line int, content string) (string, int, error) {
	compiled := regexp.MustCompile(`^[\s]*(?P<directive>\w+) \S+$`)

	directiveIndex := compiled.SubexpIndex("directive")
	if directiveIndex == -1 {
		panic("unreachable condition - directive subexp should exist")
	}

	submatchIndices := compiled.FindAllStringSubmatchIndex(content, -1)
	if len(submatchIndices) == 0 {
		return "", 0, &LinterError{
			Line:             line,
			ColumnStart:      0,
			ColumnEnd:        len(content),
			Message:          "could not find any directives",
			ParseLevel:       ParseLevelError,
			DirectiveContent: content,
		}
	}

	indices := compiled.FindStringIndex(content)
	if indices == nil {
		return "", -1, &LinterError{
			Line:             line,
			ColumnStart:      0,
			ColumnEnd:        1,
			Message:          "could not find any directives",
			ParseLevel:       ParseLevelError,
			DirectiveContent: content,
		}
	}

	return content[indices[0]:indices[1]], indices[0], nil
}

// returns true if the content represents an empty directive
func isEmptyDirective(content string) bool {
	if len(content) == 0 {
		return true
	}

	compiled := regexp.MustCompile(`^\s+$`)

	return compiled.MatchString(content)
}

// returns true if the content represents a string directive
func isCommentDirective(content string) bool {
	compiled := regexp.MustCompile(`^\s*#.+$`)

	return compiled.MatchString(content)
}

// parses a given directive from string
func ParseDirective(line int, content string) (*Directive, error) {
	switch {
	case isEmptyDirective(content):
		return &Directive{
			Token:   DirectiveEmpty,
			Content: content,
			Line:    line,
			Column:  0,
			Options: nil,
		}, nil
	case isCommentDirective(content):
		return &Directive{
			Token:   DirectiveComment,
			Content: content,
			Line:    line,
			Column:  strings.IndexRune(content, '#'),
			Options: nil,
		}, nil
	default:
		token, index, err := directiveToken(line, content)
		if err != nil {
			return nil, fmt.Errorf(
				"could not get directive token: %w",
				err,
			)
		}

		optionStartColumn := index + len(token)

		options, err := ParseOptions(
			line,
			optionStartColumn,
			// everything after the token
			content,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"could not get options: %w",
				err,
			)
		}

		return &Directive{
			Token:   token,
			Content: content,
			Line:    line,
			Column:  index,
			Options: options,
		}, nil
	}
}

// parses a given directive from byte array
func ParseDirectives(contents []string) ([]*Directive, error) {
	directives := make([]*Directive, len(contents))

	for i, content := range contents {
		directive, err := ParseDirective(i, content)
		if err != nil {
			return nil, fmt.Errorf(
				"could not parse directive: %w",
				err,
			)
		}

		directives = append(directives, directive)
	}

	return directives, nil
}
