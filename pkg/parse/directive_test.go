package parse

import (
	"testing"

	"github.com/go-test/deep"
)

func TestParseDirectives(t *testing.T) {
	type args struct {
		contents []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []*Directive
		wantErr bool
	}{
		{
			name: "POSITIVE - No content",
			args: args{
				contents: []byte(
					"",
				),
			},
			want: nil,
		},
		{
			name: "POSITIVE - Newlines with no directives",
			args: args{
				contents: []byte(
					"\n" +
						"\n" +
						"\n",
				),
			},
			want: nil,
		},
		{
			name: "POSITIVE - Comments with no directives",
			args: args{
				contents: []byte(
					"# test comment\n" +
						"# test comment 2\n" +
						"# test comment 3",
				),
			},
			want: nil,
		},
		{
			name: "POSITIVE - Test directive with a single option",
			args: args{
				contents: []byte(
					`Directive optionA "option B"`,
				),
			},
			want: []*Directive{
				{
					Lexeme: "Directive",
					Offset: 0,
					Options: []*Option{
						{
							Lexeme: "optionA",
							Offset: 10,
						},
						{
							Lexeme: `"option B"`,
							Offset: 18,
						},
					},
				},
			},
		},
		{
			name: "POSITIVE - Test one directive with two options",
			args: args{
				contents: []byte(
					`Directive optionA "option B"`,
				),
			},
			want: []*Directive{
				{
					Lexeme: "Directive",
					Offset: 0,
					Options: []*Option{
						{
							Lexeme: "optionA",
							Offset: 10,
						},
						{
							Lexeme: `"option B"`,
							Offset: 18,
						},
					},
				},
			},
		},
		{
			name: "POSITIVE - Test two directives with options",
			args: args{
				contents: []byte(
					`DirectiveA optionA "option B"` + "\n" +
						`DirectiveB optionA "option` + "\\\n" +
						`\"B\""`,
				),
			},
			want: []*Directive{
				{
					Lexeme: "DirectiveA",
					Offset: 0,
					Options: []*Option{
						{
							Lexeme: "optionA",
							Offset: 11,
						},
						{
							Lexeme: `"option B"`,
							Offset: 19,
						},
					},
				},
				{
					Lexeme: "DirectiveB",
					Offset: 30,
					Options: []*Option{
						{
							Lexeme: "optionA",
							Offset: 41,
						},
						{
							Lexeme: `"option` + "\\\n" +
								`\"B\""`,
							Offset: 49,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDirectives(tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDirectives() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
