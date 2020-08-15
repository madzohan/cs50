// 		This mini library introduces crypto exchange oprations (Encode, Decode)
// using simple algorithm based on ASCII character shift on specified step

package crypto

import (
	"unicode"
)

var maxStep = int('z' - 'a') // lower or upper it eq 25

func Decode(encodedWord string, encodeStep int) string {
	decodeStep := maxStep - getStep(encodeStep) + 1
	return shifString(encodedWord, decodeStep)
}

func Encode(word string, step int) string {
	return shifString(word, getStep(step))
}

func shifString(word string, step int) (shifted string) {
	if step == 0 {
		return word
	}
	runes := []rune(word)
	shifted = ""
	for _, r := range runes {
		shifted += string(shift(r, step))
	}
	return shifted
}

func getStep(step int) (res int) {
	if step == 0 {
		res = 0
	} else if res = step % maxStep; res == 0 {
		res = maxStep
	}
	if step < 0 {
		res = maxStep + step + 1
	}
	return res
}

func makeShift(r rune, a rune, z rune, step int) rune {
	diff := int(z - r)
	if diff < step {
		r = a
		step -= diff + 1
	}
	return r + rune(step)
}

func shift(r rune, step int) (shifted rune) {
	stepLocal := step
	rLocal := r
	if unicode.IsLower(rLocal) {
		shifted = makeShift(rLocal, 'a', 'z', stepLocal)

	} else if unicode.IsUpper(rLocal) {
		shifted = makeShift(rLocal, 'A', 'Z', stepLocal)
	} else {
		shifted = r
	}
	return shifted
}
