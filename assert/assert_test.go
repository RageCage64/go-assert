package assert_test

import (
	"fmt"
	"testing"

	"github.com/RageCage64/go-assert/assert"
)

type tMock struct {
	logs   []string
	failed bool
	err    error
}

func newTMock() *tMock {
	return &tMock{
		logs: []string{},
	}
}

func (t *tMock) Helper() {}

func (t *tMock) Fatal() {
	t.failed = true
}

func (t *tMock) Fatalf(msg string, args ...any) {
	t.logs = append(t.logs, fmt.Sprintf(msg, args...))
	t.Fatal()
}

func (t *tMock) Errorf(msg string, args ...any) {
	t.failed = true
	t.err = fmt.Errorf(msg, args...)
}

func TestAssertFail(t *testing.T) {
	testInstance := newTMock()
	failMsg := "expected %d to equal %d"
	a := 1
	b := 2
	assert.Assert(testInstance, a == b, failMsg, a, b)
	if !testInstance.failed {
		t.Fatalf("Assert failed. %v", *testInstance)
	}
	if len(testInstance.logs) != 1 {
		t.Fatalf("Found %d logs. %v", len(testInstance.logs), testInstance.logs)
	}
	expectedFailLog := fmt.Sprintf(failMsg, a, b)
	if testInstance.logs[0] != expectedFailLog {
		t.Fatalf(
			"Failure log didn't match.\nexpected: %s\ngot: %s",
			expectedFailLog,
			testInstance.logs[0],
		)
	}
}

func TestEqualFail(t *testing.T) {
	testInstance := newTMock()
	failMsg := "expected %v to equal %v"
	a := 1
	b := 2
	assert.EqualMsg(testInstance, a, b, failMsg)
	if len(testInstance.logs) != 1 {
		t.Fatalf("Found %d logs. %v", len(testInstance.logs), testInstance.logs)
	}
	expectedFailLog := fmt.Sprintf(failMsg, b, a)
	if testInstance.logs[0] != expectedFailLog {
		t.Fatalf(
			"Failure log didn't match.\nexpected: %s\ngot: %s",
			expectedFailLog,
			testInstance.logs[0],
		)
	}
}

func TestDerefenceEqualFail(t *testing.T) {
	testInstance := newTMock()
	type x struct {
		num int
	}
	failMsg := "%v not equal %v"
	a := &x{num: 1}
	b := &x{num: 2}
	assert.DereferenceEqualMsg(testInstance, a, b, failMsg)
	if len(testInstance.logs) != 1 {
		t.Fatalf("Found %d logs. %v", len(testInstance.logs), testInstance.logs)
	}
	expectedFailLog := fmt.Sprintf(failMsg, *b, *a)
	if testInstance.logs[0] != expectedFailLog {
		t.Fatalf(
			"Failure log didn't match.\nexpected: %s\ngot: %s",
			expectedFailLog,
			testInstance.logs[0],
		)
	}
}

func TestDereferenceEqualPass(t *testing.T) {
	testInstance := newTMock()
	type x struct {
		num int
	}
	a := &x{num: 1}
	b := &x{num: 1}
	assert.DereferenceEqualMsg(testInstance, a, b, "doesn't matter")
	if testInstance.failed {
		t.Fatalf("test failed when it should have passed")
	}
	if len(testInstance.logs) != 0 {
		t.Fatalf("test instance had logs when it shouldn't: %v", testInstance.logs)
	}
}

func TestSliceEqualFailDiffSize(t *testing.T) {
	testInstance := newTMock()
	failSizeMsg := "%v and %v"
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3}
	assert.SliceEqualMsg(testInstance, a, b, failSizeMsg, "something else")
	if len(testInstance.logs) != 1 {
		t.Fatalf("Found %d logs. %v", len(testInstance.logs), testInstance.logs)
	}
	expectedFailLog := fmt.Sprintf(failSizeMsg, len(b), len(a))
	if testInstance.logs[0] != expectedFailLog {
		t.Fatalf(
			"Failure log didn't match.\nexpected: %s\ngot: %s",
			expectedFailLog,
			testInstance.logs[0],
		)
	}
}

func TestSliceEqualMismatch(t *testing.T) {
	testInstance := newTMock()
	failSizeMsg := "%v and %v"
	a := []int{1, 2, 4}
	b := []int{1, 2, 3}
	assert.SliceEqualMsg(testInstance, a, b, failSizeMsg, "something else")
	if len(testInstance.logs) != 1 {
		t.Fatalf("Found %d logs. %v", len(testInstance.logs), testInstance.logs)
	}
	expectedFailLog := fmt.Sprintf(failSizeMsg, b[2], a[2])
	if testInstance.logs[0] != expectedFailLog {
		t.Fatalf(
			"Failure log didn't match.\nexpected: %s\ngot: %s",
			expectedFailLog,
			testInstance.logs[0],
		)
	}
}
