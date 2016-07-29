package misspell

import (
	"testing"
)

func TestIsBinaryFile(t *testing.T) {
	cases := []struct {
		path string
		want bool
	}{
		{"foo.png", true},
		{"foo.PNG", true},
		{"README", false},
		{"foo.txt", false},
	}

	for num, tt := range cases {
		if IsBinaryFile(tt.path) != tt.want {
			t.Errorf("Case %d: %s was not %v", num, tt.path, tt.want)
		}
	}
}
