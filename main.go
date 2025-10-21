package murka

// Validate checks if all characters in the input string match the a-zA-Z0-9 pattern.
// Additional character validation can be provided through custom checker functions.
// Returns ErrIncorrectCharacter if any character doesn't match the pattern.
func Validate[I ~string](input I, additionalCheckers ...func(value rune) bool) error {
	for _, char := range input {
		if !isAlphanumeric(char) && !runAdditionalChecks(char, additionalCheckers...) {
			return ErrIncorrectCharacter
		}
	}
	return nil
}

// ValidateOnly checks if all characters in the input string match only the provided
// character validation functions. Returns ErrIncorrectCharacter if any character
// doesn't pass the provided validators.
func ValidateOnly[I ~string](input I, validators ...func(value rune) bool) error {
	for _, char := range input {
		if !runAdditionalChecks(char, validators...) {
			return ErrIncorrectCharacter
		}
	}
	return nil
}

// Replace substitutes characters that don't match the a-zA-Z0-9 pattern
// or additional validators with the specified replacement character.
// Returns the modified string.
func Replace[I ~string | ~[]rune](input I, replacementChar rune, additionalCheckers ...func(value rune) bool) string {
	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		if !isAlphanumeric(chars[i]) && !runAdditionalChecks(chars[i], additionalCheckers...) {
			chars[i] = replacementChar
		}
	}
	return string(chars)
}

// ReplaceNonAlphanumeric replaces all non-alphanumeric characters (not a-zA-Z0-9)
// in the input string with the specified replacement character.
// Provides an efficient way to sanitize a string to contain only alphanumeric characters.
func ReplaceNonAlphanumeric[I ~string | ~[]rune](input I, replacementChar rune) string {
	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		if !isAlphanumeric(chars[i]) {
			chars[i] = replacementChar
		}
	}
	return string(chars)
}

// isAlphanumeric checks if a character is an English letter (a-zA-Z) or a digit (0-9).
func isAlphanumeric(char rune) bool {
	return (char|0x20 >= 'a' && char|0x20 <= 'z') || // English letters (case-insensitive)
		(char >= '0' && char <= '9') // Digits
}
