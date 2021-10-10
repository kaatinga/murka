package murka

import (
	"github.com/kaatinga/assets"
	cerr "github.com/kaatinga/const-errs"
	"strconv"
	"strings"
)

const (
	maximumSampleLength       = 1 << 8
	maximumHighlightedSamples = 1 << 2
)

// incorrectSampleLength is used to compose error in case the sample is too long
type incorrectSampleLength int

func (length incorrectSampleLength) Error() string {
	return "the sample length is " + strconv.Itoa(int(length)) + " what exceeds the maximum " + assets.Byte2String(maximumSampleLength-1)
}

const ErrIncorrectLength = cerr.Error("the sample exceeds the maximum")

func (length incorrectSampleLength) Is(err error) bool {
	return err == ErrIncorrectLength
}

// Highlight takes in some content and locations and then inserts left/right, strings which can be used for
// highlighting around matching terms. For example, if you pass in "test" with sample "te", you will have
// "<strong>te</strong>st" as return.
func Highlight(text, left, right, sample string) (string, error) {

	// if the both inserts are empty, we return the input text string
	if len(left) == 0 && len(right) == 0 {
		return text, nil
	}

	// FIXME: check the maximum text length
	// FIXME: check the maximum left length
	// FIXME: check the maximum right length

	sampleLength := uint16(len(sample))
	if sampleLength >= maximumSampleLength {
		return "", incorrectSampleLength(sampleLength)
	}

	// stage 1: indexing
	sampleAsRunes := []rune(sample)

	var sampleFound bool
	var currentSampleIndex byte
	var lastSampleIndex = sampleLength - 1
	var howMany byte
	var items uint64
	var sampleIndex uint16
	var sections byte
	var lastIndexInSample uint16
	var key int
	var value rune

	// index the input string
	for key, value = range text {
		if !sampleFound {
			if value == sampleAsRunes[0] {
				// fmt.Println("beginning found", string([]rune{value}))
				sampleFound = true
				currentSampleIndex = 1
				sampleIndex = uint16(key)
				// fmt.Println("sampleIndex", sampleIndex)
			}
			continue
		}

		if sampleFound {

			if value != sampleAsRunes[currentSampleIndex] {

				if value == sampleAsRunes[0] {
					// fmt.Println("beginning found", string([]rune{value}))
					currentSampleIndex = 1
					sampleIndex = uint16(key)
					// fmt.Println("sampleIndex", sampleIndex)
					continue
				}

				currentSampleIndex = 0
				sampleFound = false
				// fmt.Println("it was not sample", string([]rune{value}), string([]rune{sampleAsRunes[currentSampleIndex]}))
				continue
			}

			if value == sampleAsRunes[lastSampleIndex] {
				currentSampleIndex = 0
				sampleFound = false
				items |= uint64(sampleIndex) << (howMany << 4)
				lastIndexInSample = uint16(key)

				// count number of sections
				if sampleIndex-lastIndexInSample > 1 {
					sections++
					// fmt.Println("gap detected")
				}
				sections += 3

				// fmt.Println("last character found", string([]rune{value}))
				// fmt.Printf("indexes, %b\n", items)

				// the maximum of samples is reached
				if howMany == maximumHighlightedSamples {
					// fmt.Println("maximum reached")
					break
				}

				// continue to seek samples
				howMany++
				// fmt.Println("found items", howMany)
				continue
			}

			currentSampleIndex++
			// fmt.Println("next character found", string([]rune{value}))
		}
	}

	if sections == 0 {
		//fmt.Println("early finish")
		return text, nil
	}

	// fmt.Println("key", key, "lastIndexInSample", lastIndexInSample)
	if uint16(key)-lastIndexInSample > 0 {
		sections++
		// fmt.Println("tail detected")
	}

	// fmt.Println("sections", sections)

	// stage 2: composing new string

	// wipe the values to reuse these variables
	lastIndexInSample = 0
	currentSampleIndex = 0

	sampledText := make([]string, sections)
	for i := byte(0); i < howMany; i++ {

		// fmt.Println("sample", i, "started to be processed")
		sampleIndex = uint16(items >> (i << 4))
		if sampleIndex-lastIndexInSample != 0 {
			sampledText[currentSampleIndex] = text[lastIndexInSample:sampleIndex]
			currentSampleIndex++
		}

		// we add sample
		sampledText[currentSampleIndex] = left
		currentSampleIndex++
		sampledText[currentSampleIndex] = text[sampleIndex : sampleIndex+sampleLength]
		// fmt.Println("sample added", text[sampleIndex:sampleIndex+uint16(len(left))])

		currentSampleIndex++
		sampledText[currentSampleIndex] = right
		currentSampleIndex++
		lastIndexInSample += sampleIndex + sampleLength
		// fmt.Println("lastIndexInSample was updated", lastIndexInSample)
		// fmt.Println("next unprocessed index is", currentSampleIndex)
	}

	if uint16(len(text)) > lastIndexInSample {
		sampledText[currentSampleIndex] = text[lastIndexInSample:]
	}

	return strings.Join(sampledText, ""), nil
}
