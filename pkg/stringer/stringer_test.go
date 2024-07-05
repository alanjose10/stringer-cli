package stringer

import "testing"

func TestStringReverse(t *testing.T) {
	str := "abcdefghijkl"
	revStr := "lkjihgfedcba"

	if res := Reverse(str); revStr != res {
		t.Errorf("Got %s, wanted %s", res, revStr)
	} else {
		t.Log("String was reversed successfully")
	}
}
