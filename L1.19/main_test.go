package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverseString_Table(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"ascii", "abc", "cba"},
		{"russian", "главрыба", "абырвалг"},
		{"spaces", "a b", "b a"},
		{"single_emoji", "😊", "😊"},
		{"mixed", "аб😊", "😊ба"},
		{"combining_mark", "e\u0301", "\u0301e"},
		{"flag", "🇷🇺", "🇺🇷"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.in)
			if got != tt.want {
				t.Fatalf("reverseString(%q) = %q; want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestReverseString_Involution(t *testing.T) {
	samples := []string{
		"",
		"abc",
		"главрыба",
		"🙂🙃😉",
		"e\u0301",
		"🇷🇺",
	}

	for _, s := range samples {
		if !utf8.ValidString(s) {
			t.Fatalf("invalid utf-8 in sample: %q", s)
		}
		if reverseString(reverseString(s)) != s {
			t.Fatalf("double reverse failed for %q", s)
		}
	}
}

// go test -fuzz=Fuzz -fuzztime=5s
func FuzzReverseString(f *testing.F) {
	seeds := []string{
		"",
		"abc",
		"главрыба",
		"🙂🙃😉",
		"e\u0301",
		"🇷🇺",
		"aб😊",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}
		if reverseString(reverseString(s)) != s {
			t.Fatalf("double reverse failed: %q", s)
		}
	})
}
