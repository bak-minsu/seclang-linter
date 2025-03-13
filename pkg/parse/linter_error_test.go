package parse

import (
	"strings"
	"testing"
)

// constructs a message separated by newline
func joinString(messages ...string) string {
	return strings.Join(messages, "\n")
}

func TestLinterError_Error(t *testing.T) {
	type fields struct {
		Offset     int
		Distance   int
		Message    string
		ParseLevel int
		Content    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "POSITIVE - single column error at the start",
			fields: fields{
				Offset:     0,
				Distance:   1,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content:    "SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 0:",
				"SecRule optionA optionB",
				"^",
				"",
			),
		},
		{
			name: "POSITIVE - single column error with tab start",
			fields: fields{
				Offset:     1,
				Distance:   1,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content:    "\tSecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 1:",
				"	SecRule optionA optionB",
				"	^",
				"",
			),
		},
		{
			name: "POSITIVE - single column error with whitespace start",
			fields: fields{
				Offset:     4,
				Distance:   1,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content:    "    SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 4:",
				"    SecRule optionA optionB",
				"    ^",
				"",
			),
		},
		{
			name: "POSITIVE - multi column error at the start",
			fields: fields{
				Offset:     0,
				Distance:   7,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content:    "SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 0:",
				"SecRule optionA optionB",
				"^^^^^^^",
				"",
			),
		},
		{
			name: "POSITIVE - multi column error with whitespace at the start",
			fields: fields{
				Offset:     2,
				Distance:   7,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content:    "  SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 2:",
				"  SecRule optionA optionB",
				"  ^^^^^^^",
				"",
			),
		},
		{
			name: "POSITIVE - multi line error",
			fields: fields{
				Offset:     16,
				Distance:   25,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content: joinString(
					"SecRule optionA \"this option",
					"    is long\"",
				),
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 16:",
				`SecRule optionA "this option`,
				`                ^^^^^ ^^^^^^`,
				`    is long"`,
				`    ^^ ^^^^^`,
				"",
			),
		},
		{
			name: "POSITIVE - long single-line text error",
			fields: fields{
				Offset:     16,
				Distance:   71,
				Message:    "This column is wrong",
				ParseLevel: ParseLevelError,
				Content: joinString(
					`SecRule optionA "this single option is way too long so it will be split into two lines"`,
				),
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 1, column 16:",
				`SecRule optionA "this single option is way too long so it will be split into two`,
				`                ^^^^^ ^^^^^^ ^^^^^^ ^^ ^^^ ^^^ ^^^^ ^^ ^^ ^^^^ ^^ ^^^^^ ^^^^ ^^^`,
				`     lines"`,
				`     ^^^^^^`,
				"",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &LinterError{
				Offset:     tt.fields.Offset,
				Distance:   tt.fields.Distance,
				Message:    tt.fields.Message,
				ParseLevel: tt.fields.ParseLevel,
				Contents:   tt.fields.Content,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ParseError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
