package slug

import "testing"

func Test_Slug(t *testing.T) {
	in := "Zażółć gęślą jaźń & coś"
	expected := "zazolc-gesla-jazn-i-cos"
	out := Slug(in)
	if out != expected {
		t.Errorf("Wrong slug: expected %s got %s", expected, out)
	}
}
