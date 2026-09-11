package music

var naturalPitches = map[byte]int{
	'C': 0,
	'D': 2,
	'E': 4,
	'F': 5,
	'G': 7,
	'A': 9,
	'B': 11,
}

const sharp = '#'
const flat = 'b'
const minOctave = '0'
const maxOctave = '9'

func IsPitch(word string) bool {
	if len(word) != 2 && len(word) != 3 {
		return false
	}

	_, exists := naturalPitches[word[0]]
	if !exists {
		return false
	}

	if len(word) == 3 && word[1] != flat && word[1] != sharp {
		return false
	}

	octave := word[len(word)-1]

	if octave > maxOctave || octave < minOctave {
		return false
	}

	return true
}

// GetPitchValue(name) assumes IsPitch(name) is true
func GetPitchValue(name string) int {

	nameWithoutOctave := name[0 : len(name)-1]
	octave := int(name[len(name)-1] - '0')

	if len(nameWithoutOctave) == 1 {
		return naturalPitches[nameWithoutOctave[0]] + octave*12
	}

	if nameWithoutOctave[1] == sharp {
		return naturalPitches[nameWithoutOctave[0]] + 1 + octave*12
	}

	return naturalPitches[nameWithoutOctave[0]] - 1 + octave*12
}
