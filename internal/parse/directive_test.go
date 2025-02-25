package parse

import (
	"reflect"
	"testing"
)

func TestParseDirective(t *testing.T) {
	type args struct {
		line    int
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    *Directive
		wantErr bool
	}{
		{
			name: "POSITIVE - empty string",
			args: args{
				line:    0,
				content: "",
			},
			want: &Directive{
				Token:   DirectiveEmpty,
				Content: "",
				Line:    0,
				Column:  0,
				Options: nil,
			},
		},
		{
			name: "POSITIVE - spaces and tabs",
			args: args{
				line:    0,
				content: " \t \t",
			},
			want: &Directive{
				Token:   DirectiveEmpty,
				Content: " \t \t",
				Line:    0,
				Column:  0,
				Options: nil,
			},
		},
		{
			name: "POSITIVE - comment at first column",
			args: args{
				line:    0,
				content: "# comment",
			},
			want: &Directive{
				Token:   DirectiveComment,
				Content: "# comment",
				Line:    0,
				Column:  0,
				Options: nil,
			},
		},
		{
			name: "POSITIVE - comment at some column",
			args: args{
				line:    0,
				content: "\t # comment",
			},
			want: &Directive{
				Token:   DirectiveComment,
				Content: "\t # comment",
				Line:    0,
				Column:  2,
				Options: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDirective(tt.args.line, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDirective() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDirective() = %v, want %v", got, tt.want)
			}
		})
	}
}
