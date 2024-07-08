package test

import (
    "testing"
    "piscine"
)

// table駆動テスト
func TestConvertBase(t *testing.T) {
    tests := []struct {
        nbr string
        baseFrom string
        baseTo string
        want string
    }{
        {"1010",
         "01",
         "0123456789ABCDEF",
         "A"},
        {"A",
         "0123456789ABCDEF",
         "01",
         "1010"},
        {"10",
         "0123456789",
         "01",
         "1010"},
        {"1010",
         "01",
         "0123456789",
         "10"},
        {"12",
         "01234567",
         "0123456789ABCDEF",
         "A"},
        {"FF",
         "0123456789ABCDEF",
         "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
         "73"},
        {"ZZ",
         "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
         "0123456789",
         "1295"},
        {"0",
         "01",
         "0123456789ABCDEF",
         "0"},
        {"123456789",
         "0123456789",
         "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
         "21I3V9"},
        // 記号を含む基数のケース
        {"1234",
         "0123456789",
         "!@#$%^&*()_+",
         "(&_"},
        {"#&%)",
         "!@#$%^&*()_+",
         "0123456789",
         "4377"},
        // Unicodeを含む基数のケース
        {"1234",
         "0123456789",
         "🍎🍐🍊🍋🍌🍉🍇🍓🍒🍑",
         "🍐🍊🍋🍌"},
        {"🍐🍊🍒🍑",
         "🍎🍐🍊🍋🍌🍉🍇🍓🍒🍑",
         "0123456789",
         "1289"},
        // 混合ケース（数字、アルファベット、記号、Unicode）
        {"ABC123",
         "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
         "αβγδεζηθικλμνξοπρστυφχψω!@#$%^&*",
         "ττ@!#$"},
        {"λμ!@#",
         "αβγδεζηθικλμνξοπρστυφχψω!@#$%^&*",
         "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
         "6H0L6"}, // 32->36
    }

    for _, test := range tests {
        got := piscine.ConvertBase(test.nbr, test.baseFrom, test.baseTo)
        if got != test.want {
            t.Errorf("ConvertBase(%q, %q, %q) = %q; want %q\nconvert from %d to %d", test.nbr, test.baseFrom, test.baseTo, got, test.want, strLen(test.baseFrom), strLen(test.baseTo))
        }
    }
}

func strLen(s string) int {
  ret := 0
  for range s {
    ret++
  }
  return ret
}
