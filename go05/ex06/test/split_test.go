package test

import (
    "testing"
    "reflect"
    "strings"
    "piscine"
)

// table駆動テスト
func TestSplit(t *testing.T) {
    tests := []struct {
        s string
        sep string
    }{
        {"a,b,c", ","},
        {"abc", ","},
        {"a,b,c", ",,"},
        {"", ","},
        {",,a,,b,,,,", ","},
        // non ascii case
        {"ううあいうえおう", "う"},
        {"$a$$$b$$$c", "$$$"},

        {"🍎🍐🍊🍋🍌", "🍊"},
        {"😄😄😄😄😄", "😄"},
        {"🚗💨🚙💨🚕", "💨"},

        {"λμνξοπρστ", "ξ"},
        {"x→→y→→→z→", "→"},
        {"♫♫1♫♫♫2♫♫3♫4♫♫♫", "♫"},
    }

    for _, test := range tests {
        t.Run(test.s, func(t *testing.T) {
            got := piscine.Split(test.s, test.sep)
            want := strings.Split(test.s, test.sep)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("Split(%q, %q) = %v, want %v", test.s, test.sep, got, want)
            }
        })
    }
}
