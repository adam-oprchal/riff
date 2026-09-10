package music

var NaturalPitches = map[byte]int{
	'C': 0,
	'D': 2,
	'E': 4,
	'F': 5,
	'G': 7,
	'A': 9,
	'B': 11,
}

const Sharp = '#'
const Flat = 'b'
const MinOctave = '0'
const MaxOctave = '9'

func IsPitch(word string) bool {
	if len(word) != 2 && len(word) != 3 {
		return false
	}

	_, exists := NaturalPitches[word[0]]
	if !exists {
		return false
	}

	if len(word) == 3 && word[1] != Flat && word[1] != Sharp {
		return false
	}

	octave := word[len(word)-1]

	if octave > MaxOctave || octave < MinOctave {
		return false
	}

	return true
}

func GetPitchValue(noteName string) int {
	if len(noteName) == 1 {
		return NaturalPitches[noteName[0]]
	}

	if noteName[1] == Sharp {
		return NaturalPitches[noteName[0]] + 1
	}

	return NaturalPitches[noteName[0]] - 1
}
