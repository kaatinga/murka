[![GitHub release](https://img.shields.io/github/release/kaatinga/murka.svg)](https://github.com/kaatinga/murka/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/murka/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/murka/branch/main/graph/badge.svg)](https://codecov.io/gh/kaatinga/murka)
[![lint workflow](https://github.com/kaatinga/murka/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/murka/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/murka/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# murka

A high performance string validator and string sanitizer. The package is intended for checking incorrect characters' existence in a string.
`Validate()` is 20 times faster than regexp package and 10 times faster than unicode package.

## 1. Installation

Use go get.

	go get github.com/kaatinga/murka

Then import the validator package into your own code.

	import "github.com/kaatinga/murka"

## 2. How to use

### Validation

By default, `Validate()` checks only a-z, A-z and 0-9 characters. Additional checks are done by functions that comply
with `func (value rune) bool`.

Prepare an additional function like this:

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

Benchmarks:

```
BenchmarkValidate
BenchmarkValidate-8    	        45240508	       26.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateByRegexp
BenchmarkValidateByRegexp-8    	 2072374	       578.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidateByUnicode
BenchmarkValidateByUnicode-8   	 4360131	       272.8 ns/op	       0 B/op	       0 allocs/op
```

### String sanitization

`Replace()` checks only a-z, A-z and 0-9 characters by default. Like with `Validate()`, additional checks are done by
functions that comply with `func (value rune) bool`. `Replace()` does not return any error, but substitute incorrect
characters with the input character of rune type. There is no sense to use this function without an additional checker.

`ReplaceNotaZ09()` replaces all characters except a-z, A-z and 0-9. The most efficient way to sanitize a string using
a-zA-Z0-9<input character> pattern.

Benchmarks:

```
BenchmarkReplace
BenchmarkReplace-8          	 5310733	       229.0 ns/op	      24 B/op	       2 allocs/op
BenchmarkReplaceNotAz09
BenchmarkReplaceNotAz09-8   	 9882952	       116.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringsReplace
BenchmarkStringsReplace-8   	 9212074	       133.3 ns/op	      16 B/op	       2 allocs/op
BenchmarkStringsMap
BenchmarkStringsMap-8       	 8836887	       130.8 ns/op	      32 B/op	       2 allocs/op
```[![Tests](https://github.com/kaatinga/luna/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kaatinga/luna/actions/workflows/test.yml)
