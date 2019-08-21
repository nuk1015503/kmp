# README
===========================
This is a Knuth-Morris-Pratt algorithm implement.

Install
---------------
`go get github.com/nuk1015503/kmp`

Usage
---------------

```go
package main

import (
	"fmt"

	"github.com/nuk1015503/kmp"
)

func main() {
	test := "testeste23testte1testestest"
	pattern := "23testte1"

	isExist, position := kmp.KMP(test, pattern)
	// test string exists pattern string and start position is 8
	fmt.Println(isExist, position)
}

```