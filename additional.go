package murka

// CheckUnderscore checks underscore character.
func CheckUnderscore(value rune) bool {
	const underscore = '_'
	return value == underscore
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
