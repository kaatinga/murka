package murka

// Validate checks characters in the input string. The default charset is a-zA-Z0-9.
func Validate(pagePath string, additionalCheckers ...func(value rune) bool) error {

	for _, value := range pagePath {
		if !(value >= 0x61 && value <= 0x7A || // lowercase
			value >= 0x41 && value <= 0x5A || // uppercase
			value >= 0x30 && value <= 0x39 || // digits
			runAdditionalChecks(value, additionalCheckers...)) {
			return ErrIncorrectCharacter
		}
	}

	return nil
}
