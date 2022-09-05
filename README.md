# go-assert

Assertion library for Go built with generics! 

This library is mostly meant to keep assertion patterns very minimal, and geared towards reducing the need to handroll `t.Fatalf` messages constantly by providing sensible default messages (that can be overriden if needed). 

## Usage

Install the module in your project:

```
go get github.com/RageCage64/go-assert
```

### Assert

A default `Assert` function is provided, which you can pass a condition and your own failure strings to.

```go
package main_test

import (
    "testing"

    "github.com/RageCage64/go-assert/assert"
)

func TestAssert(t *testing.T) {
    a := 1
    b := 2
    assert.Assert(t, a == b, "didn't work")
}
```

### NilErr

Checks if the error passed in is `nil`.

```go
package main_test

import (
    "testing"

    "github.com/RageCage64/go-assert/assert"
)

func TestAssert(t *testing.T) {
    err := errors.New("not nil")
    assert.NilErr(t, err)
}
```

### Equal

Using `Equal` you can check if two comparable types are equal.

```go
package main_test

import (
    "testing"

    "github.com/RageCage64/go-assert/assert"
)

type x struct {
    num int
}

func TestEqual(t *testing.T) {
    a := x{num: 1}
    b := x{num: 2}
    assert.Equal(t, a, b)
}
```

### DereferenceEqual

Using `DereferenceEqual` you can check if the value at two pointers of comparable type are the same.

```go
package main_test

import (
    "testing"

    "github.com/RageCage64/go-assert/assert"
)

type x struct {
    num int
}

func TestEqual(t *testing.T) {
    a := &x{num: 1}
    b := &x{num: 2}
    assert.DereferenceEqual(t, a, b)
}
```

### SliceEqual

Using `SliceEqual` you can check if two slices of comparable types are equal. Different messages will be used if the slices are different length or if they are mismatched.

```go
package main_test

import (
    "testing"

    "github.com/RageCage64/go-assert/assert"
)

func TestEqual(t *testing.T) {
    a := []int{1, 2}
    b := []int{1, 3}
    assert.SliceEqual(t, a, b)
}
```