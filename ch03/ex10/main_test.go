package main

import "testing"

func TestMain(t *testing.T) {
	in := "1234567890"
	expect := "1,234,567,890"
	if comma(in) != expect {
		t.Error("fail", expect, comma(in))
	}

	in = "12"
	expect = "12"
	if comma(in) != expect {
		t.Error("fail", expect, comma(in))
	}

	in = "123"
	expect = "123"
	if comma(in) != expect {
		t.Error("fail", expect, comma(in))
	}

}
