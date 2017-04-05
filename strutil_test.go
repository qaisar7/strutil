package strutil

import (
	"testing"
)

func TestXxx(t *testing.T) {
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
