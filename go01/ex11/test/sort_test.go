package test

import (
  "testing"
  "piscine"
  "sort"
)

func TestSuccess1(t *testing.T) {
  got := []int{0, -1, -2, -3, -4, -5}
  want := []int{0, -1, -2, -3, -4, -5}
  piscine.SortIntegerTable(got)
  sort.Slice(want, func(i, j int) bool {
    return want[i] < want[j]
  })
  for i, g := range got {
    if g != want[i] {
      t.Errorf("Expected %d at index %d, got %d", want[i], i, g)
    }
  }
}

func TestSuccess2(t *testing.T) {
  got := []int{64, 34, 25, 12, 22, 11, 90}
  want := []int{64, 34, 25, 12, 22, 11, 90}
  piscine.SortIntegerTable(got)
  sort.Slice(want, func(i, j int) bool {
    return want[i] < want[j]
  })
  for i, g := range got {
    if g != want[i] {
      t.Errorf("Expected %d at index %d, got %d", want[i], i, g)
    }
  }
}

func TestError(t *testing.T) {
  got := []int{}
  want := []int{}
  piscine.SortIntegerTable(got)
  sort.Slice(want, func(i, j int) bool {
    return want[i] < want[j]
  })
  for i, g := range got {
    if g != want[i] {
      t.Errorf("Expected %d at index %d, got %d", want[i], i, g)
    }
  }
}
