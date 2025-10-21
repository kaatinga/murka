package murka

import (
	"errors"
	"strconv"

	faststrconv "github.com/kaatinga/strconv"
)

var (
	ErrIncorrectCharacter = errors.New("input string contains an incorrect character")
	ErrTooLongString      = errors.New("input string is too long, the maximum is 65535 characters")
	ErrIncorrectLength    = errors.New("the sample exceeds the maximum")
)

// incorrectSampleLength is used to compose error in case the sample is too long
type incorrectSampleLength int

func (length incorrectSampleLength) Error() string {
	return "the sample length is " + strconv.Itoa(int(length)) + " what exceeds the maximum " + faststrconv.Byte2String(maximumSampleLength)
}

func (length incorrectSampleLength) Is(err error) bool {
	return err == ErrIncorrectLength
}
