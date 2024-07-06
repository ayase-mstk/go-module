package test

import (
    "testing"
    "piscine"
)

// table駆動テスト
func TestAtoiBase(t *testing.T) {
    tests := []struct {
        input    string
        base string
        want int
    }{
        {"", "0123456789", 0},
        {"101", "01", 5},
        {"101", "0123456789", 101},
        {"1A", "0123456789ABCDEF", 26},
        {"Z", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 25},
        {"!@#", "!@#$%^&*()", 12},
        // unicode
        {"𐍈𐍉𐍊", "𐍈𐍉𐍊", 5}, // Gothic script characters
        {"你好", "你好吗", 1}, // Chinese characters
        {"😊😇😈", "😊😇😈", 5}, // Emoji characters
        {"абв", "абвгде", 8}, // Cyrillic characters
        // error
        {"10", "0", 0}, // does not have two char
        {"10", "11", 0}, //not unique
        {"10", "0123+456", 0}, // have +/-
        {"10", "012345-6", 0}, // have +/-
        {"10a", "01", 0}, // input not in base
    }

    for _, test := range tests {
        got := piscine.AtoiBase(test.input, test.base)
        if got != test.want {
            t.Errorf("atoiBase(%q, %q) = %d; want %d", test.input, test.base, got, test.want)
        }
    }
}

