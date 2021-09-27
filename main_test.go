package murka

import (
	"regexp"
	"testing"
	"unicode"
)

var re = regexp.MustCompile(`^[\w]+$`)

// validateByRegexp checks symbols in the input string.
func validateByRegexp(pagePath string) error {

	if re.MatchString(pagePath) {
		return nil
	}

	return ErrIncorrectCharacter
}

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

// validateByUnicode checks symbols in the input string.
func validateByUnicode(pagePath string) error {

	for _, value := range pagePath {
		if !unicode.IsOneOf(ranges, value) {
			return ErrIncorrectCharacter
		}
	}

	return nil
}

func Test_validatePagePath(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", nil},
		{"1", nil},
		{"p1", nil},
		{"az", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", nil},
		{"TEST_PAGE", nil},
		{"12345", nil},
		{"test", nil},
		{"ы", ErrIncorrectCharacter},
		{"-", ErrIncorrectCharacter},
		{".", ErrIncorrectCharacter},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := Validate(tt.pagePath, CheckUnderscore); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Clean_validatePagePath(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", ErrIncorrectCharacter},
		{"1", nil},
		{"p1", nil},
		{"az", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", ErrIncorrectCharacter},
		{"ы", ErrIncorrectCharacter},
		{"-", ErrIncorrectCharacter},
		{".", ErrIncorrectCharacter},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := Validate(tt.pagePath); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePagePath2(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", nil},
		{"1", nil},
		{"p1", nil},
		{"az", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", nil},
		{"ы", ErrIncorrectCharacter},
		{"-", ErrIncorrectCharacter},
		{".", ErrIncorrectCharacter},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := validateByRegexp(tt.pagePath); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePagePath3(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", nil},
		{"1", nil},
		{"p1", nil},
		{"aZ", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", nil},
		{"ы", ErrIncorrectCharacter},
		{"-", ErrIncorrectCharacter},
		{".", ErrIncorrectCharacter},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := validateByUnicode(tt.pagePath); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReplace(t *testing.T) {

	tests := []struct {
		name               string
		text               string
		character          rune
		additionalCheckers []func(value rune) bool
		want               string
	}{
		{
			"ok",
			"File-Test 09:29:2008 mac",
			'_',
			[]func(value rune) bool{CheckUnderscore},
			"File_Test_09_29_2008_mac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Replace(tt.text, tt.character, tt.additionalCheckers...)

			if got != tt.want {
				t.Errorf("Replace() got = %v, want %v", got, tt.want)
			}

			got = ReplaceNotaZ09(tt.text, tt.character)

			if got != tt.want {
				t.Errorf("ReplaceNotaZ09() got = %v, want %v", got, tt.want)
			}
		})
	}
}
