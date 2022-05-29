package mod1

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("1 + 2 != 3, out %d\n", r)
	}
}
