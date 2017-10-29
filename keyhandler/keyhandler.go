package keyhandler

import (
	"math"
)

var alphabet = [62]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

var characterValues = map[byte]uint64{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
	'g': 16,
	'h': 17,
	'i': 18,
	'j': 19,
	'k': 20,
	'l': 21,
	'm': 22,
	'n': 23,
	'o': 24,
	'p': 25,
	'q': 26,
	'r': 27,
	's': 28,
	't': 29,
	'u': 30,
	'v': 31,
	'w': 32,
	'x': 34,
	'y': 35,
	'z': 36,
	'A': 37,
	'B': 38,
	'C': 39,
	'D': 40,
	'E': 41,
	'F': 42,
	'G': 43,
	'H': 44,
	'I': 45,
	'J': 46,
	'K': 47,
	'L': 48,
	'M': 49,
	'N': 50,
	'O': 51,
	'P': 52,
	'Q': 53,
	'R': 54,
	'S': 55,
	'T': 56,
	'U': 57,
	'V': 58,
	'W': 59,
	'X': 60,
	'Y': 61,
	'Z': 62,
}

func Encode(value uint64) string {
	var encodedValue string
	for value > 0 {
		var character = alphabet[calculateIndex(value)]
		value = uint64(math.Floor(float64(value) / float64(len(alphabet))))
		encodedValue = string(character) + encodedValue
	}
	return encodedValue
}

func Decode(value string) uint64 {
	var decodedValue uint64 = 0
	for position, character := range []byte (value) {
		var characterValue = characterValues[character]
		var power = len(value) - 1 - position
		decodedValue += uint64(float64(characterValue) * math.Pow(float64(len(alphabet)), float64(power)))
	}
	return decodedValue
}

func calculateIndex(value uint64) uint64 {
	var remainder = math.Mod(float64(value), float64(len(alphabet)))
	return uint64(math.Abs(remainder))
}