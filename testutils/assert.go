package testutils

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T, context string) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v, want %v", context, got, want)
	}
}

func AssertSliceEqual[T comparable](t *testing.T, got, want []T, context string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s length = %d, want %d", context, len(got), len(want))
		return
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("%s[%d] = %v, want %v", context, i, got[i], v)
		}
	}
}

func AssertBool(t *testing.T, got, want bool, context string) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v, want %v", context, got, want)
	}
}
