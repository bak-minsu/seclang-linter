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
		OffsetStart int
		Distance    int
		Message     string
		ParseLevel  int
		Content     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "POSITIVE - single column error at the start",
			fields: fields{
				OffsetStart: 0,
				Distance:    1,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				Content:     "SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 0: SecRule optionA optionB",
				"        ^",
				"",
			),
		},
		{
			name: "POSITIVE - single column error with tab start",
			fields: fields{
				OffsetStart: 1,
				Distance:    1,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				Content:     "\tSecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 0: 	SecRule optionA optionB",
				"        	^",
				"",
			),
		},
		{
			name: "POSITIVE - single column error with whitespace start",
			fields: fields{
				OffsetStart: 4,
				Distance:    1,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				Content:     "    SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 0:     SecRule optionA optionB",
				"            ^",
				"",
			),
		},
		{
			name: "POSITIVE - multi column error at the start",
			fields: fields{
				OffsetStart: 0,
				Distance:    7,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				Content:     "SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 0: SecRule optionA optionB",
				"        ^^^^^^^",
				"",
			),
		},
		{
			name: "POSITIVE - multi column error with whitespace at the start",
			fields: fields{
				OffsetStart: 2,
				Distance:    7,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				Content:     "  SecRule optionA optionB",
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				"line 0:   SecRule optionA optionB",
				"          ^^^^^^^",
				"",
			),
		},
		{
			name: "POSITIVE - multi line error",
			fields: fields{
				OffsetStart: 16,
				Distance:    25,
				Message:     "This column is wrong",
				ParseLevel:  ParseLevelError,
				Content: joinString(
					"SecRule optionA \"this option",
					"    is long\"",
				),
			},
			want: joinString(
				"",
				"Error: This column is wrong",
				`line 0: SecRule optionA "this option`,
				`                        ^^^^^ ^^^^^^`,
				`line 1:     is long"`,
				`            ^^ ^^^^^`,
				"",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &LinterError{
				OffsetStart: tt.fields.OffsetStart,
				Distance:    tt.fields.Distance,
				Message:     tt.fields.Message,
				ParseLevel:  tt.fields.ParseLevel,
				Content:     tt.fields.Content,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ParseError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
