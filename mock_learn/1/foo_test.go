package foo

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBaz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)
	m.EXPECT().Bar(gomock.Eq("a")).Return(123)

	expected := 124
	if got := Baz(m, "a"); got != expected {
		t.Errorf(`Baz(m, "abc") = %d, want %d`, got, expected)
	}
}
