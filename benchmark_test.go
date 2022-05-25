package murka

import (
	"testing"
	//"github.com/boyter/go-string"
)

//// nolint
//func BenchmarkValidate(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Validate("test")
//		Validate("12345")
//		Validate("TEST_PAGE", CheckUnderscore)
//	}
//}
//
//// nolint
//func BenchmarkValidateByRegexp(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		validateByRegexp("test")
//		validateByRegexp("12345")
//		validateByRegexp("TEST_PAGE")
//	}
//}
//
//// nolint
//func BenchmarkValidateByUnicode(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		validateByUnicode("test")
//		validateByUnicode("12345")
//		validateByUnicode("TEST_PAGE")
//	}
//}

// nolint
//func BenchmarkReplace(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Replace("tes:t", '_')
//		Replace("12:3:45", '_')
//	}
//}
//
//func BenchmarkReplaceNotAz09(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		ReplaceNotaZ09("tes:t", '_')
//		ReplaceNotaZ09("12:3:45", '_')
//	}
//}
//
//// nolint
//func BenchmarkStringsReplace(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		strings.ReplaceAll("tes:t", ":", "_")
//		strings.ReplaceAll("12:3:45", ":", "_")
//	}
//}
//
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
//// nolint
//func BenchmarkStringsMap(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		strings.Map(legalCharacters1, "tes:t")
//		strings.Map(legalCharacters1, "12:3:45")
//	}
//}

// nolint
func BenchmarkHighlight(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Highlight("test string", "<b>", "</b>", "str")
		Highlight("mamase", "<b>", "</b>", "mas")
		Highlight("mase", "<b>", "</b>", "mas")
		Highlight("mase", "<b>", "</b>", "ggg")
	}
}

// nolint
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
