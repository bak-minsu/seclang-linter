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
		Line             int
		ColumnStart      int
		ColumnEnd        int
		Message          string
		ParseLevel       int
		DirectiveContent string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "POSITIVE - single column error at the start",
			fields: fields{
				Line:             0,
				ColumnStart:      0,
				ColumnEnd:        1,
				Message:          "This column is wrong",
				ParseLevel:       ParseLevelError,
				DirectiveContent: "SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Parse Error: This column is wrong",
				"\tline 0, column 0:",
				"\tSecRule optionA optionB",
				"\t^",
				"",
			),
		},
		{
			name: "POSITIVE - single column error with whitespace start",
			fields: fields{
				Line:             0,
				ColumnStart:      4,
				ColumnEnd:        5,
				Message:          "This column is wrong",
				ParseLevel:       ParseLevelError,
				DirectiveContent: "    SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Parse Error: This column is wrong",
				"\tline 0, column 4:",
				"\t    SecRule optionA optionB",
				"\t    ^",
				"",
			),
		},
		{
			name: "POSITIVE - multi column error at the start",
			fields: fields{
				Line:             0,
				ColumnStart:      0,
				ColumnEnd:        7,
				Message:          "This column is wrong",
				ParseLevel:       ParseLevelError,
				DirectiveContent: "SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Parse Error: This column is wrong",
				"\tline 0, column 0:",
				"\tSecRule optionA optionB",
				"\t^^^^^^^",
				"",
			),
		},
		{
			name: "POSITIVE - multi column error with whitespace at the start",
			fields: fields{
				Line:             0,
				ColumnStart:      2,
				ColumnEnd:        9,
				Message:          "This column is wrong",
				ParseLevel:       ParseLevelError,
				DirectiveContent: "  SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Parse Error: This column is wrong",
				"\tline 0, column 2:",
				"\t  SecRule optionA optionB",
				"\t  ^^^^^^^",
				"",
			),
		},
		{
			name: "POSITIVE - multi line error",
			fields: fields{
				Line:        0,
				ColumnStart: 36,
				ColumnEnd:   40,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				DirectiveContent: joinString(
					"SecRule optionA \"this option",
					"    is long\"",
				),
			},
			want: joinString(
				"",
				"Parse Error: This column is wrong",
				"\tline 0, column 36:",
				"\tSecRule optionA \"this option",
				"",
				"\t    is long\"",
				"\t       ^^^^",
				"",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &LinterError{
				Line:             tt.fields.Line,
				ColumnStart:      tt.fields.ColumnStart,
				ColumnEnd:        tt.fields.ColumnEnd,
				Message:          tt.fields.Message,
				ParseLevel:       tt.fields.ParseLevel,
				DirectiveContent: tt.fields.DirectiveContent,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ParseError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
