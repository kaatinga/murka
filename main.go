package murka

// Validate checks characters in the input string. The most efficient way to sanitize a string using a-zA-Z0-9 pattern.
// a-zA-Z0-9 is the default pattern, but you can add your own patterns.
func Validate[I ~string](it I, additionalCheckers ...func(value rune) bool) error {
	for _, value := range it {
		if !(value|0x20 >= 0x61 && value|0x20 <= 0x7A || // english letters
			value >= 0x30 && value <= 0x39 || // digits
			runAdditionalChecks(value, additionalCheckers...)) {
			return ErrIncorrectCharacter
		}
	}
	return nil
}

// ValidateOnly checks characters in the input string.
func ValidateOnly[I ~string](it I, additionalCheckers ...func(value rune) bool) error {
	for _, value := range it {
		if !runAdditionalChecks(value, additionalCheckers...) {
			return ErrIncorrectCharacter
		}
	}
	return nil
}

// Replace checks characters in the input string.
func Replace[I ~string | []rune](it I, character rune, additionalCheckers ...func(value rune) bool) string {
	// we repeat the size of the new string
	var newString = []rune(it)
	for i := 0; i < len(newString); i++ {
		if !(newString[i]|0x20 >= 0x61 && newString[i]|0x20 <= 0x7A || // english letters
			newString[i] >= 0x30 && newString[i] <= 0x39 || // digits
			runAdditionalChecks(newString[i], additionalCheckers...)) {
			newString[i] = character
		}
	}
	return string(newString)
}

// ReplaceNotaZ09 replaces not a-zA-Z0-9 characters in the input string ny in input character.
// The most efficient way to sanitize a string using a-zA-Z0-9<character> pattern.
func ReplaceNotaZ09[I ~string | []rune](it I, character rune) string {
	// we repeat the size of the new string
	var newString = []rune(it)
	for i := 0; i < len(newString); i++ {
		if !(newString[i]|0x20 >= 0x41|0x20 && newString[i] <= 0x5A|0x20 || // english letters
			newString[i] >= 0x30 && newString[i] <= 0x39) { // digits
			newString[i] = character
		}
	}
	return string(newString)
}
