package murka

import (
	"regexp"
	"testing"
	"unicode"
	//"github.com/boyter/go-string"
)

var re = regexp.MustCompile(`^[\w]+$`)

// validateByRegexp checks symbols in the input string.
func validateByRegexp(pagePath string) error {
	if re.MatchString(pagePath) {
		return nil
	}

	return ErrIncorrectCharacter
}

// validateByUnicode checks symbols in the input string.
func validateByUnicode(pagePath string) error {
	for _, value := range pagePath {
		if !unicode.IsOneOf(ranges, value) {
			return ErrIncorrectCharacter
		}
	}
	return nil
}

// range tables for a-zA-Z0-9
var ranges = []*unicode.RangeTable{
	{R16: []unicode.Range16{
		{0x61, 0x7a, 1},
	}},
	{R16: []unicode.Range16{
		{0x41, 0x5a, 1},
	}},
	{R16: []unicode.Range16{
		{0x30, 0x39, 1},
	}},
	{R16: []unicode.Range16{
		{0x5f, 0x5f, 1},
	}},
}

func BenchmarkValidate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Validate("test")                       //nolint: errcheck
		Validate("12345")                      //nolint: errcheck
		Validate("TEST_PAGE", CheckUnderscore) //nolint: errcheck
	}
}

func BenchmarkValidateByRegexp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validateByRegexp("test")      //nolint: errcheck
		validateByRegexp("12345")     //nolint: errcheck
		validateByRegexp("TEST_PAGE") //nolint: errcheck
	}
}

func BenchmarkValidateByUnicode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validateByUnicode("test")      //nolint: errcheck
		validateByUnicode("12345")     //nolint: errcheck
		validateByUnicode("TEST_PAGE") //nolint: errcheck
	}
}

//func BenchmarkReplace(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Replace("tes:t", '_')
//		Replace("12:3:45", '_')
//	}
//}

//
// func BenchmarkReplaceNonAlphanumeric(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		ReplaceNonAlphanumeric("tes:t", '_')
//		ReplaceNonAlphanumeric("12:3:45", '_')
//	}
//}
//
//func BenchmarkStringsReplace(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		strings.ReplaceAll("tes:t", ":", "_")
//		strings.ReplaceAll("12:3:45", ":", "_")
//	}
//}

//var legalCharacters1 = func(value rune) rune {
//	if !(value >= 0x61 && value <= 0x7A || // lowercase
//		value >= 0x41 && value <= 0x5A || // uppercase
//		value >= 0x30 && value <= 0x39) {
//		return '_'
//	}
//
//	return value
//}
//
// func BenchmarkStringsMap(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		strings.Map(legalCharacters1, "tes:t")
//		strings.Map(legalCharacters1, "12:3:45")
//	}
// }

//func BenchmarkHighlight(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Highlight("test string", "<b>", "</b>", "str")
//		Highlight("mamase", "<b>", "</b>", "mas")
//		Highlight("mase", "<b>", "</b>", "mas")
//		Highlight("mase", "<b>", "</b>", "ggg")
//	}
//}

//func BenchmarkGoString(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		locations := str.IndexAll("test string", "str", -1)
//		str.HighlightString("test string", locations, "<b>", "</b>")
//
//		locations = str.IndexAll("mamase", "mas", -1)
//		str.HighlightString("mamase", locations, "<b>", "</b>")
//
//		locations = str.IndexAll("mase", "mas", -1)
//		str.HighlightString("mase", locations, "<b>", "</b>")
//
//		locations = str.IndexAll("mase", "ggg", -1)
//		str.HighlightString("mase", locations, "<b>", "</b>")
//	}
//}
