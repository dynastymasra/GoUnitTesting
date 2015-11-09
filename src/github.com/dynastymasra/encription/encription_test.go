package encription

import "testing"

type testCase struct {
  in string
  out string
  shift int
}

var testEncrip = []testCase {
  {"abc", "def", 3}, {"xyz", "abc", 3}, {"xyzabc", "defghi", 6}, {"ABC", "DEF", 3},
  {"XYZ", "ABC", 3}, {"XYZABC", "DEFGHI", 6},
}

var testDecrip = []testCase {
  {"def", "abc", 3}, {"abc", "xyz", 3}, {"defghi", "xyzabc", 6}, {"DEF", "ABC", 3},
  {"ABC", "XYZ", 3}, {"DEFGHI", "XYZABC", 6},
}

func TestEncription(text *testing.T) {
  for _, item := range testEncrip {
    res := Encription(item.in, item.shift)
    if res != item.out {
      text.Fatalf("Testing Encription Fail Text %s, Out %s, Must %s\n", item.in, res, item.out)
    }
  }
}

func TestDecription(text *testing.T) {
  for _, item := range testDecrip {
    res := Decription(item.in, item.shift)
    if res != item.out {
      text.Fatalf("Testing Description Fail Text %s, Out %s, Must %s\n", item.in, res, item.out)
    }
  }
}
