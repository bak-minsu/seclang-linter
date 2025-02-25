package parse

import (
	"reflect"
	"testing"
)

func Test_optionValues(t *testing.T) {
	type args struct {
		line    int
		column  int
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{
			name: "NEGATIVE - single line with no spaces and no quotes",
			args: args{
				line:    0,
				column:  0,
				content: "testOption",
			},
			wantErr: true,
		},
		{
			name: "POSITIVE - single line starting with one space with no quotes",
			args: args{
				line:    0,
				column:  0,
				content: " testOption",
			},
			want: map[string]int{
				"testOption": 1,
			},
		},
		{
			name: "POSITIVE - single line starting with several spaces with no quotes",
			args: args{
				line:    0,
				column:  0,
				content: "    testOption",
			},
			want: map[string]int{
				"testOption": 4,
			},
		},
		{
			name: "NEGATIVE - multiple options in a single line starting with no space with no quotes",
			args: args{
				line:    0,
				column:  0,
				content: "optionA optionB",
			},
			wantErr: true,
		},
		{
			name: "POSITIVE - multiple options in a single line starting with a single space with no quotes",
			args: args{
				line:    0,
				column:  0,
				content: " optionA optionB",
			},
			want: map[string]int{
				"optionA": 1,
				"optionB": 9,
			},
		},
		{
			name: "NEGATIVE - single line starting with a tab with no quotes",
			args: args{
				line:    0,
				column:  0,
				content: "\ttestOption",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := optionValues(tt.args.line, tt.args.column, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("optionValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optionValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
