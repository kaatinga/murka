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
	var newString = make([]rune, len(text))

	for i, value := range text {
		if !(value >= 0x61 && value <= 0x7A || // lowercase
			value >= 0x41 && value <= 0x5A || // uppercase
			value >= 0x30 && value <= 0x39 || // digits
			runAdditionalChecks(value, additionalCheckers...)) {
			newString[i] = character
		} else {
			newString[i] = value
		}
	}

	return string(newString)
}
