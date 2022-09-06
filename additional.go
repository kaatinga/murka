package murka

// CheckUnderscore checks underscore character.
func CheckUnderscore(value rune) bool {
	return value == 0x5f
}

// runAdditionalChecks uses additional functions to check input string.
func runAdditionalChecks(value rune, additionalCheckers ...func(value rune) bool) bool {
	if len(additionalCheckers) == 0 {
		return false
	}

	for _, check := range additionalCheckers {
		if check(value) {
			return true
		}
	}

	return false
}
