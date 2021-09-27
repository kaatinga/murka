package murka

// Validate checks characters in the input string. The default charset is a-zA-Z0-9.
func Validate(text string, additionalCheckers ...func(value rune) bool) error {

	for _, value := range text {
		if !(value >= 0x61 && value <= 0x7A || // lowercase
			value >= 0x41 && value <= 0x5A || // uppercase
			value >= 0x30 && value <= 0x39 || // digits
			runAdditionalChecks(value, additionalCheckers...)) {
			return ErrIncorrectCharacter
		}
	}

	return nil
}

// Replace replaces incorrect characters in the input string. The default charset is a-zA-Z0-9.
func Replace(text string, character rune, additionalCheckers ...func(value rune) bool) string {

	// we repeat the size of the new string
	var newString = []rune(text)

	for i := 0; i < len(text); i++ {
		if !(newString[i] >= 0x61 && newString[i] <= 0x7A || // lowercase
			newString[i] >= 0x41 && newString[i] <= 0x5A || // uppercase
			newString[i] >= 0x30 && newString[i] <= 0x39 || // digits
			runAdditionalChecks(newString[i], additionalCheckers...)) {
			newString[i] = character
		}
	}

	return string(newString)
}

// ReplaceNotaZ09 replaces not a-zA-Z0-9 characters in the input string ny in input character.
// The most efficient way to sanitize a string using a-zA-Z0-9<character> pattern.
func ReplaceNotaZ09(text string, character rune) string {

	// we repeat the size of the new string
	var newString = []rune(text)

	for i := 0; i < len(text); i++ {
		if !(newString[i] >= 0x61 && newString[i] <= 0x7A || // lowercase
			newString[i] >= 0x41 && newString[i] <= 0x5A || // uppercase
			newString[i] >= 0x30 && newString[i] <= 0x39) { // digits
			newString[i] = character
		}
	}

	return string(newString)
}
