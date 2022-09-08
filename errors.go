package murka

import cerr "github.com/kaatinga/const-errs"

const (
	ErrIncorrectCharacter = cerr.Error("input string contains an incorrect character")
	ErrTooLongString      = cerr.Error("input string is too long, the maximum is 65535 characters")
)
