package strutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "comparing Hey",
			input: "Hey",
			want:  "yeH",
		},
		{
			name:    "fail, empty string",
			input:   "",
			wantErr: true,
		},
	}
	for _, test := range tests {
		got, err := Reverse(test.input)
		if (err != nil) != test.wantErr {
			t.Errorf("want error:%v, got:%v", test.wantErr, err)
		}
		if err != nil {
			continue
		}
		if got != test.want {
			t.Errorf("Test(%s): got:%s, want:%s", test.name, got, test.want)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "hey mate",
			input: "hey mate",
			want:  "heyMate",
		},
		{
			name:  "extra characters",
			input: "what's up mate?",
			want:  "what'sUpMate?",
		},
		{
			name:  "simple",
			input: "wow",
			want:  "wow",
		},
	}
	for _, test := range tests {
		got := CamelCase(test.input)
		if got != test.want {
			t.Errorf("Test(%s): got:%s., want:%s. %v %d", test.name, got, test.want)
		}
	}
}
