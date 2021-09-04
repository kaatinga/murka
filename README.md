# murka
A high performance string validator. The package is intended to check incorrect characters' existence in a string.
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