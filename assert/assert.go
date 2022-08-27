package assert

type TestingT interface {
	Helper()
	Fatalf(string, ...any)
	Errorf(string, ...any)
}

func Assert(t TestingT, condition bool, message string, args ...any) {
	t.Helper()

	if !condition {
		t.Fatalf(message, args...)
	}
}

func Equal[T comparable](t TestingT, got T, expected T) {
	t.Helper()

	if got != expected {
		t.Fatalf("value did not equal expectation.\nexpected: %v\ngot: %v", expected, got)
	}
}

func DereferenceEqual[T comparable](t TestingT, got *T, expected *T) {
	t.Helper()

	if got == nil || expected == nil {
		t.Errorf("go-assert: could not dereference nil pointer\ngot %v, expected %v", got, expected)
	}
	Equal(t, *got, *expected)
}

func NilErr(t TestingT, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got error:\n%v", err)
	}
}
