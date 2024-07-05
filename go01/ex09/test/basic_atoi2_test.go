package test

import (
  "testing"
  "fmt"
  "strconv"
  "piscine"
)

func TestMax(t *testing.T) {
  got := piscine.BasicAtoi2("2147483647")
  want, _ := strconv.Atoi("2147483647")
  if got != want {
    t.Errorf("BasicAtoi2(\"2147483647\") == %d, want %d", got, want)
  }
}

func TestOver(t *testing.T) {
  got := piscine.BasicAtoi2("2147483648")
  want, _ := strconv.Atoi("2147483648")
  if got != want {
    t.Errorf("BasicAtoi2(\"2147483648\") == %d, want %d", got, want)
  }
}

func TestOver2(t *testing.T) {
  got := piscine.BasicAtoi2("2147483648000000000000")
  want, _ := strconv.Atoi("2147483648000000000000")
  fmt.Println(want)
  if got != want {
    t.Errorf("BasicAtoi2(\"2147483648000000000000\") == %d, want %d", got, want)
  }
}
