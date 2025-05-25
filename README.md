# murka

[![Tests](https://github.com/kaatinga/murka/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kaatinga/murka/actions/workflows/test.yml)
[![GitHub release](https://img.shields.io/github/release/kaatinga/murka.svg)](https://github.com/kaatinga/murka/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/murka/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/murka/branch/main/graph/badge.svg)](https://codecov.io/gh/kaatinga/murka)
[![lint workflow](https://github.com/kaatinga/murka/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/murka/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/murka/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

A high-performance Go package for string validation and sanitization. Murka provides lightning-fast character validation and replacement functions, significantly outperforming standard Go libraries.

## Key Features

- **Blazing Fast**: Up to 20x faster than the standard regexp package and 10x faster than the unicode package
- **Zero Allocations**: Core validation functions allocate no memory
- **Customizable**: Easily extend with your own character validation functions
- **Dual Functionality**: Both validation and sanitization capabilities

## Installation

    bash go get github.com/kaatinga/murka

Then import the package into your code:

    import "github.com/kaatinga/murka"

## Usage Guide

### String Validation

Murka's validation functions check if a string contains only specified characters. By default, `Validate()` allows only alphanumeric characters (a-z, A-Z, 0-9).

#### Basic Validation

```go
    text := "Hello123"
	err := murka.Validate(text)
	if err != nil {
		// Handle invalid characters
	}
```

#### Custom Validation

Extend validation by providing custom character check functions:

```go
// CheckUnderscore allows underscore character func CheckUnderscore(value rune) bool { return value == '_' // or 0x5f in hex }
// CheckHyphen allows hyphen character func CheckHyphen(value rune) bool { return value == '-' }
// Now validate with multiple character sets
text := "user_name-123"
err := murka.Validate(text, CheckUnderscore, CheckHyphen)
if err != nil {
	// Handle invalid characters
}
```

### String Sanitization

Murka provides functions to clean strings by replacing unwanted characters.

#### Replace with Custom Rules

```go
// Allow alphanumeric characters and underscores, replace others with '-'
sanitized := murka.Replace(text, '-', CheckUnderscore)
```

#### Efficient Replacement

For maximum performance when you only want to keep alphanumeric characters:

```go
// Replace everything except a-z, A-Z, 0-9 with '-'
sanitized := murka.ReplaceNonAlphanumeric(text, '-')
```

## Performance

Murka is optimized for high-performance validation and sanitization with minimal memory usage.

### Benchmarks

```
BenchmarkReplace
BenchmarkReplace-8          	 5310733	       229.0 ns/op	      24 B/op	       2 allocs/op
BenchmarkReplaceNotAz09
BenchmarkReplaceNotAz09-8   	 9882952	       116.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringsReplace
BenchmarkStringsReplace-8   	 9212074	       133.3 ns/op	      16 B/op	       2 allocs/op
BenchmarkStringsMap
BenchmarkStringsMap-8       	 8836887	       130.8 ns/op	      32 B/op	       2 allocs/op
```

## Use Cases

- Input validation for usernames, file names, or other string fields
- Sanitizing user input before database storage
- Normalizing strings for consistent format
- High-performance text processing pipelines

## Contributing

Contributions are welcome! Check out the [help wanted](https://github.com/kaatinga/murka/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22) issues to get started.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/kaatinga/murka/blob/main/LICENSE) file for details.
