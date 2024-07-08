package test

import (
	"bytes"
	"os"
	"piscine"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	return buf.String()
}

// table駆動テスト
func TestPrintNbrBase(t *testing.T) {
	tests := []struct {
		input int
		base  string
		want  string
	}{
		{123, "0123456789", "123"},
		{123, "01", "1111011"},
		{123, "01234567", "173"},
		{123, "0123456789ABCDEF", "7B"},
		{-123, "0123456789ABCDEF", "-7B"},
		{0, "0123456789", "0"},
		{123, "!@#$", "@$#$"},
		{-123, "!@#$", "-@$#$"},
		// unicode
		{5, "𐍈𐍉𐍊", "𐍉𐍊"},    // Gothic script characters
		{1, "你好吗", "好"},     // Chinese characters
		{5, "😊😇😈", "😇😈"},    // Emoji characters
		{8, "абвгде", "бв"}, // Cyrillic characters
		// error
		{10, "0", "NV"},        // does not have two char
		{10, "11", "NV"},       //not unique
		{10, "0123+456", "NV"}, // have +/-
		{10, "012345-6", "NV"}, // have +/-
	}

	for _, test := range tests {
		got := captureOutput(func() {
			piscine.PrintNbrBase(test.input, test.base)
		})
		if got != test.want {
			t.Errorf("PrintNbrBase(%d, %q) = %q; want %q", test.input, test.base, got, test.want)
		}
	}
}
