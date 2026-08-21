package midi

import (
	"bytes"
	"fmt"

	"github.com/adam-oprchal/riff/parser"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func PlayProgram(program parser.Program) {

	defer midi.CloseDriver()

	//fmt.Println("out ports: \n" + midi.GetOutPorts().String())

	out, err := midi.FindOutPort("FLUID Synth")
	if err != nil {
		fmt.Printf("can't find fluidsynth")
		return
	}

	// create a SMF
	rd := bytes.NewReader(mkSMF(program))

	// read and play it
	smf.ReadTracksFrom(rd).Play(out)

}

var naturalNotes = map[byte]int{
	'C': 0,
	'D': 2,
	'E': 4,
	'F': 5,
	'G': 7,
	'A': 9,
	'B': 11,
}

func getMIDINote(note parser.Note) midi.Note {
	return midi.Note(naturalNotes[note.Pitch.Letter] + note.Pitch.Octave*12)
}

// makes a SMF and returns the bytes
func mkSMF(program parser.Program) []byte {
	var (
		bf    bytes.Buffer
		clock = smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also common
		tr    smf.Track
	)

	// first track must have tempo and meter informations
	tr.Add(0, smf.MetaMeter(3, 4))
	tr.Add(0, smf.MetaTempo(140))
	tr.Add(0, smf.MetaInstrument("Guitar"))

	tr.Add(0, midi.ProgramChange(0, gm.Instr_AcousticGuitarSteel.Value()))

	for _, e := range program.Events {

		switch event := e.(type) {
		case parser.Note:
			tr.Add(clock.Ticks4th(), getMIDINote(event).NoteOn(0, 120))
			tr.Add(0, getMIDINote(event).NoteOff(0))
		case parser.TempoChange:
		}
	}

	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()
}
