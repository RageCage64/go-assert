package assert

var (
	// The failure format string for values not being equal. Formatted with `expected` then `got`.
	EqualMessage = "value did not equal expectation.\nexpected: %v\n     got: %v"

	// The error format string for one or both pointers being nil. Formatted with `got` then `expected`.
	DereferenceEqualErrMsg = "go-assert: could not dereference nil pointer\ngot %v, expected %v"

	// The failure format string if the err is not nil. Formatted with `err`.
	NilErrMessage = "expected no error, got error:\n%v"

	// The failure format string for slices being different sizes. Formatted with `expected` then `got`.
	SliceSizeMessage = "slices were different sizes.\nexpected len:%d\n     got len:%d\n"

	// The failure format string for slices not matching at some index. Formatted with the mismatched
	// index, then `expected`, then `got`.
	SliceMismatchMessage = "slices differed at index %d.\nexpected: %v\n     got: %v"
)

// The interface that represents the subset of `testing.T` that this package
// requires. Passing in a `testing.T` satisfies this interface.
type TestingT interface {
	Helper()
	Fatal(...any)
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

// Assert that `got` equals `expected`. The types between compared
// arguments must be the same. Uses `assert.EqualMessage`.
func Equal[T comparable](t TestingT, expected T, got T) {
	t.Helper()
	EqualMsg(t, expected, got, EqualMessage)
}

// Assert that the value at `got` equals the value at `expected`. Will
// error if either pointer is nil. Uses `assert.DereferenceEqualErrMsg`
// and `assert.EqualMessage`.
func DereferenceEqual[T comparable](t TestingT, expected *T, got *T) {
	t.Helper()
	DereferenceEqualMsg(t, expected, got, DereferenceEqualErrMsg, EqualMessage)
}

// Assert that that `err` is nil. Uses `assert.NilErrMessage`.
func NilErr(t TestingT, err error) {
	t.Helper()
	NilErrMsg(t, err, NilErrMessage)
}

// Assert that slices `got` and `expected` are equal. Will produce a
// different message if the lengths are different or if any element
// mismatches. Uses `assert.SliceSizeMessage` and
// `assert.SliceMismatchMessage`.
func SliceEqual[T comparable](t TestingT, expected []T, got []T) {
	t.Helper()
	SliceEqualMsg(
		t,
		expected,
		got,
		SliceSizeMessage,
		SliceMismatchMessage,
	)
}

// Assert that `got` equals `expected`. The types between compared
// arguments must be the same. Uses `message`.
func EqualMsg[T comparable](t TestingT, expected T, got T, message string) {
	t.Helper()

	if got != expected {
		t.Fatalf(message, expected, got)
	}
}

// Assert that the value at `got` equals the value at `expected`. Will
// error if either pointer is nil. Uses `errMessage` and `mismatchMessage`.
func DereferenceEqualMsg[T comparable](
	t TestingT,
	expected *T,
	got *T,
	errMessage,
	mismatchMessage string,
) {
	t.Helper()

	if got == nil || expected == nil {
		t.Errorf(errMessage, expected, got)
	} else {
		EqualMsg(t, *expected, *got, mismatchMessage)
	}
}

// Assert that that `err` is nil. Uses `message`.
func NilErrMsg(t TestingT, err error, message string) {
	t.Helper()

	if err != nil {
		t.Fatalf(message, err)
	}
}

// Assert that slices `got` and `expected` are equal. Will produce a
// different message if the lengths are different or if any element
// mismatches. Uses `sizeMessage` and `mismatchMessage`.
func SliceEqualMsg[T comparable](
	t TestingT,
	expected []T,
	got []T,
	sizeMessage, mismatchMessage string,
) {
	t.Helper()

	if len(got) != len(expected) {
		t.Fatalf(sizeMessage, len(expected), len(got))
	} else {
		for i := range got {
			if got[i] != expected[i] {
				t.Fatalf(mismatchMessage, i, expected[i], got[i])
			}
		}
	}
}
