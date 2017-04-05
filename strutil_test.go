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
	}
	for _, test := range tests {
		got := Reverse(test.input)
		if got != test.want {
			t.Errorf("Test(%s): got:%s, want:%s", test.name, got, test.want)
		}
	}
}
