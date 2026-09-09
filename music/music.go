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

func GetPitchValue(noteName string) int {
	if len(noteName) == 1 {
		return NaturalPitches[noteName[0]]
	}

	if noteName[1] == Sharp {
		return NaturalPitches[noteName[0]] + 1
	}

	return NaturalPitches[noteName[0]] - 1
}
