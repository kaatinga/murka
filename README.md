[![GitHub release](https://img.shields.io/github/release/kaatinga/murka.svg)](https://github.com/kaatinga/murka/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/murka/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/murka/branch/main/graph/badge.svg)](https://codecov.io/gh/kaatinga/murka)
[![lint workflow](https://github.com/kaatinga/murka/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/murka/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/murka/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# murka
A high performance string validator. The package is intended for checking incorrect characters' existence in a string.
20 times faster than regexp package and 10 times faster than unicode package.

Benchmark run with Go 1.17:

```
BenchmarkValidate
BenchmarkValidate-8    	        45240508	       26.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateByRegexp
BenchmarkValidateByRegexp-8    	 2072374	       578.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateByUnicode
BenchmarkValidateByUnicode-8   	 4360131	       272.8 ns/op	       0 B/op	       0 allocs/op
```

## 1. Installation

Use go get.

	go get github.com/kaatinga/murka

Then import the validator package into your own code.

	import "github.com/kaatinga/murka"

## 2. How to use

By default, `Validate()` checks only a-z, A-z and 0-9 characters. Additional checks are done by functions that comply with
`func (value rune) bool`.

Example:
```go
// CheckUnderscore checks underscore character.
func CheckUnderscore(value rune) bool {
	return value == 0x5f
}
```

After that you can call `Validate()`.

```go
err := murka.Validate(text, CheckUnderscore)
if err != nil {
    return err
}
```