package assert

import "testing"

func Assert(t *testing.T, condition bool, message string, args ...interface{}) {
	t.Helper()

	if !condition {
		t.Fatalf(message, args...)
	}
}

func Equal[T comparable](t *testing.T, got T, expected T) {
	t.Helper()

	if got != expected {
		t.Fatalf("value did not equal expectation.\nexpected: %v\ngot: %v", expected, got)
	}
}

func DereferenceEqual[T comparable](t *testing.T, got *T, expected *T) {
	t.Helper()

	if got == nil || expected == nil {
		t.Errorf("go-assert: could not dereference nil pointer\ngot %v, expected %v", got, expected)
	}
	Equal(t, *got, *expected)
}

func NilErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got error:\n%v", err)
	}
}
