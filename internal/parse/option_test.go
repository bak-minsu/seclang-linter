package parse

import (
	"testing"

	"github.com/go-test/deep"
)

func TestParseOptions(t *testing.T) {
	type args struct {
		contents []byte
		offset   int
	}
	tests := []struct {
		name    string
		args    args
		want    []*Option
		wantErr bool
	}{
		{
			name: "NEGATIVE - No options",
			args: args{
				contents: []byte(
					"",
				),
			},
			wantErr: true,
		},
		{
			name: "POSITIVE - Parse single unquoted option",
			args: args{
				contents: []byte(
					"optionA",
				),
			},
			want: []*Option{
				{
					Lexeme: "optionA",
					Offset: 0,
				},
			},
		},
		{
			name: "POSITIVE - Parse single quoted option",
			args: args{
				contents: []byte(
					`"option A"`,
				),
			},
			want: []*Option{
				{
					Lexeme: `"option A"`,
					Offset: 0,
				},
			},
		},
		{
			name: "POSITIVE - Parse single quoted option with escaped quote",
			args: args{
				contents: []byte(
					`"\"option A\""`,
				),
			},
			want: []*Option{
				{
					Lexeme: `"\"option A\""`,
					Offset: 0,
				},
			},
		},
		{
			name: "POSITIVE - Parse single unquoted option with offset",
			args: args{
				contents: []byte(
					"Directive optionA",
				),
				offset: 9,
			},
			want: []*Option{
				{
					Lexeme: "optionA",
					Offset: 10,
				},
			},
		},
		{
			name: "POSITIVE - Parse single quoted option with offset",
			args: args{
				contents: []byte(
					`Directive "option A"`,
				),
				offset: 9,
			},
			want: []*Option{
				{
					Lexeme: `"option A"`,
					Offset: 10,
				},
			},
		},
		{
			name: "POSITIVE - Parse two unquoted options with offset",
			args: args{
				contents: []byte(
					"Directive optionA optionB",
				),
				offset: 9,
			},
			want: []*Option{
				{
					Lexeme: "optionA",
					Offset: 10,
				},
				{
					Lexeme: "optionB",
					Offset: 18,
				},
			},
		},
		{
			name: "POSITIVE - Parse two quoted options",
			args: args{
				contents: []byte(
					`"option A" "option B"`,
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: `"option A"`,
					Offset: 0,
				},
				{
					Lexeme: `"option B"`,
					Offset: 11,
				},
			},
		},
		{
			name: "POSITIVE - Parse two quoted options with escaped quotes",
			args: args{
				contents: []byte(
					`"option \"A\"" "option \"B\""`,
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: `"option \"A\""`,
					Offset: 0,
				},
				{
					Lexeme: `"option \"B\""`,
					Offset: 15,
				},
			},
		},
		{
			name: "POSITIVE - Parse two unquoted options with escaped newline",
			args: args{
				contents: []byte(
					"optionA \\\n" +
						"optionB",
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: "optionA",
					Offset: 0,
				},
				{
					Lexeme: "optionB",
					Offset: 10,
				},
			},
		},
		{
			name: "POSITIVE - Parse two quoted options with escaped newline",
			args: args{
				contents: []byte(
					`"option \"A\""` + "\\\n" +
						`"option \"B\""`,
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: `"option \"A\""`,
					Offset: 0,
				},
				{
					Lexeme: `"option \"B\""`,
					Offset: 16,
				},
			},
		},
		{
			name: "POSITIVE - Parse a quoted option with escaped newline",
			args: args{
				contents: []byte(
					`"option` + "\\\n" +
						`A"`,
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: `"option` + "\\\n" + `A"`,
					Offset: 0,
				},
			},
		},
		{
			name: "POSITIVE - Parse an unquoted option with nonescaped newline",
			args: args{
				contents: []byte(
					"optionA\n" +
						"Directive optionB",
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: "optionA",
					Offset: 0,
				},
			},
		},
		{
			name: "POSITIVE - Parse two quoted options with nonescaped newline",
			args: args{
				contents: []byte(
					`"option \"A\""` + "\n" +
						`Directive "option \"B\""`,
				),
				offset: 0,
			},
			want: []*Option{
				{
					Lexeme: `"option \"A\""`,
					Offset: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOptions(tt.args.contents, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestOption_Content(t *testing.T) {
	type fields struct {
		Lexeme string
		Offset int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "POSITIVE - unquoted option",
			fields: fields{
				Lexeme: "optionA",
			},
			want: "optionA",
		},
		{
			name: "POSITIVE - quoted option",
			fields: fields{
				Lexeme: `"option A"`,
			},
			want: "option A",
		},
		{
			name: "POSITIVE - quoted option with escaped quote",
			fields: fields{
				Lexeme: `"option \"A\""`,
			},
			want: `option "A"`,
		},
		{
			name: "POSITIVE - quoted option with escaped newline",
			fields: fields{
				Lexeme: `"option` + "\\\n" +
					`A"`,
			},
			want: `option A`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Option{
				Lexeme: tt.fields.Lexeme,
				Offset: tt.fields.Offset,
			}
			if got := o.Content(); got != tt.want {
				t.Errorf("Option.Content() = %v, want %v", got, tt.want)
			}
		})
	}
}
