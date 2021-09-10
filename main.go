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
		if !(text[i] >= 0x61 && text[i] <= 0x7A || // lowercase
			text[i] >= 0x41 && text[i] <= 0x5A || // uppercase
			text[i] >= 0x30 && text[i] <= 0x39 || // digits
			runAdditionalChecks(newString[i], additionalCheckers...)) {
			newString[i] = character
		}
	}

	return string(newString)
}
