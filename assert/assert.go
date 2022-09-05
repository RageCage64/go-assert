package assert

type TestingT interface {
	Helper()
	Fatal()
	Fatalf(string, ...any)
	Errorf(string, ...any)
}

// Assert that the passed condition is true. If not, fatally fail with
// `message` and format `args` into it.
func Assert(t TestingT, condition bool, message string, args ...any) {
	t.Helper()

	if !condition {
		t.Fatalf(message, args...)
	}
}

// Check that `got` equals `expected`.
func Equal[T comparable](t TestingT, got T, expected T) {
	EqualMsg(t, got, expected, "value did not equal expectation.\nexpected: %v\ngot: %v")
}

func DereferenceEqual[T comparable](t TestingT, got *T, expected *T) {
	DereferenceEqualMsg(t, got, expected, "go-assert: could not dereference nil pointer\ngot %v, expected %v")
}

func NilErr(t TestingT, err error) {
	NilErrMsg(t, err, "expected no error, got error:\n%v")
}

func SliceEqual[T comparable](t TestingT, got []T, expected []T) {
	SliceEqualMsg(
		t,
		got,
		expected,
		"slices were different sizes.\nexpected len:%d\ngot len:%d\n",
		"slices differed at index %d.\nexpected: %v\ngot: %v",
	)
}

func EqualMsg[T comparable](t TestingT, got T, expected T, message string) {
	t.Helper()

	if got != expected {
		t.Fatalf(message, expected, got)
	}
}

func DereferenceEqualMsg[T comparable](t TestingT, got *T, expected *T, message string) {
	t.Helper()

	if got == nil || expected == nil {
		t.Errorf(message, got, expected)
	} else {
		EqualMsg(t, *got, *expected, message)
	}
}

func NilErrMsg(t TestingT, err error, message string) {
	t.Helper()

	if err != nil {
		t.Fatalf(message, err)
	}
}

func SliceEqualMsg[T comparable](t TestingT, got []T, expected []T, sizeMessage, mismatchMessage string) {
	t.Helper()

	if len(got) != len(expected) {
		t.Fatalf(sizeMessage, len(expected), len(got))
	} else {
		for i := range got {
			if got[i] != expected[i] {
				t.Fatalf(sizeMessage, expected[i], got[i])
			}
		}
	}
}
