package murka

import (
	"testing"
)

type testCase struct {
	sample  string
	wantErr error
}

var caseTemplates = []testCase{
	{"test_page", nil},
	{"p_1", nil},
	{"TEST_PAGE", nil},
	{"1", nil},
	{"p1", nil},
	{"az", nil},
	{"AZ", nil},
	{"19", nil},
	{"12345", nil},
	{"test", nil},
	{"ы", ErrIncorrectCharacter},
	{"-", ErrIncorrectCharacter},
	{".", ErrIncorrectCharacter},
}

func Test_validate(t *testing.T) {
	for i := range caseTemplates {
		t.Run(caseTemplates[i].sample, func(t *testing.T) {
			if err := Validate(caseTemplates[i].sample, CheckUnderscore); err != caseTemplates[i].wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, caseTemplates[i].wantErr)
			}
		})
	}
}

func Test_pure_validate(t *testing.T) {
	tests := caseTemplates
	tests[0].wantErr = ErrIncorrectCharacter
	tests[1].wantErr = ErrIncorrectCharacter
	tests[2].wantErr = ErrIncorrectCharacter
	for _, tt := range tests {
		t.Run(tt.sample, func(t *testing.T) {
			if err := Validate(tt.sample); err != tt.wantErr {
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

			got = ReplaceNonAlphanumeric(tt.text, tt.character)
			if got != tt.want {
				t.Errorf("ReplaceNonAlphanumeric() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateOnly(t *testing.T) {
	for i := range caseTemplates {
		if caseTemplates[i].sample != "_" {
			caseTemplates[i].wantErr = ErrIncorrectCharacter
		}
	}
	for _, tt := range caseTemplates {
		t.Run(tt.sample, func(t *testing.T) {
			if err := ValidateOnly(tt.sample, CheckUnderscore); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
