package test

import (
  "testing"
  "strconv"
  "piscine"
)

func TestMax(t *testing.T) {
  got := piscine.Atoi("2147483647")
  want, _ := strconv.Atoi("2147483647")
  if got != want {
    t.Errorf("Atoi(\"2147483647\") == %d, want %d", got, want)
  }
}

func TestMin(t *testing.T) {
  got := piscine.Atoi("-2147483648")
  want, _ := strconv.Atoi("-2147483648")
  if got != want {
    t.Errorf("Atoi(\"-2147483648\") == %d, want %d", got, want)
  }
}

func TestOverFlow(t *testing.T) {
  got := piscine.Atoi("2147483649")
  want, _ := strconv.Atoi("2147483649")
  if got != want {
    t.Errorf("Atoi(\"2147483649\") == %d, want %d", got, want)
  }
}

func TestOverFlow2(t *testing.T) {
  got := piscine.Atoi("2147483648000000000000")
  want, _ := strconv.Atoi("2147483648000000000000")
  if got != want {
    t.Errorf("Atoi(\"2147483648000000000000\") == %d, want %d", got, want)
  }
}

func TestUnderFlow(t *testing.T) {
  got := piscine.Atoi("-2147483649")
  want, _ := strconv.Atoi("-2147483649")
  if got != want {
    t.Errorf("Atoi(\"-2147483649\") == %d, want %d", got, want)
  }
}

func TestUnderFlow2(t *testing.T) {
  got := piscine.Atoi("-2147483648000000000000")
  want, _ := strconv.Atoi("-2147483648000000000000")
  if got != want {
    t.Errorf("Atoi(\"-2147483648000000000000\") == %d, want %d", got, want)
  }
}

func TestNone(t *testing.T) {
  got := piscine.Atoi("")
  want, _ := strconv.Atoi("")
  if got != want {
    t.Errorf("Atoi(\"\") == %d, want %d", got, want)
  }
}
