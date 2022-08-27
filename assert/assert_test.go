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

func (t *tMock) Fatalf(msg string, args ...any) {
	t.failed = true
	t.logs = append(t.logs, fmt.Sprintf(msg, args...))
}

func (t *tMock) Errorf(msg string, args ...any) {
	t.failed = true
	t.err = fmt.Errorf(msg, args...)
}

func TestAssertFail(t *testing.T) {
	testInstance := newTMock()
	failMsg := "expected %d to equal %d"
	assert.Assert(testInstance, 1 == 2, "expected %d to equal %d", 1, 2)
	if !testInstance.failed {
		t.Fatalf("Assert failed. %v", *testInstance)
	}
	if len(testInstance.logs) != 1 {
		t.Fatalf("Found %d logs. %v", len(testInstance.logs), testInstance.logs)
	}
	expectedFailLog := fmt.Sprintf(failMsg, 1, 2)
	if testInstance.logs[0] != expectedFailLog {
		t.Fatalf(
			"Failure log didn't match.\nexpected: %s\ngot: %s",
			expectedFailLog,
			testInstance.logs[0],
		)
	}
}

func TestEqual(t *testing.T) {

}
