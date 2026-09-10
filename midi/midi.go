package midi

import (
	"bytes"
	"errors"
	"os"

	"github.com/adam-oprchal/riff/music"
	"github.com/adam-oprchal/riff/parser"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func SaveProgram(program parser.Program, outputFile string) error {

	err := os.WriteFile(outputFile, mkSMF(program), 0644)
	if err != nil {
		return errors.New("Error: failed to save MIDI file")
	}

	return nil
}

func PlayProgram(program parser.Program) error {

	defer midi.CloseDriver()

	//fmt.Println("out ports: \n" + midi.GetOutPorts().String())

	out, err := midi.FindOutPort("FLUID Synth")
	if err != nil {
		return errors.New("Error: can't find fluidsynth")
	}

	// create a SMF
	rd := bytes.NewReader(mkSMF(program))

	// read and play it
	go smf.ReadTracksFrom(rd).Play(out)

	return nil
}

func getDurationTicks(duration parser.Duration, clock smf.MetricTicks) uint32 {

	switch duration {
	case "whole":
		return clock.Ticks4th() * 4
	case "half":
		return clock.Ticks4th() * 2
	case "quarter":
		return clock.Ticks4th()
	case "eighth":
		return clock.Ticks8th()
	case "sixteenth":
		return clock.Ticks16th()
	default:
		return 0
	}
}

func playChord(chord parser.Chord, tr *smf.Track, clock smf.MetricTicks) {

	if len(chord.Pitches) == 0 {
		return
	}

	for _, p := range chord.Pitches {

		midiNote := getMIDINote(p)
		tr.Add(0, midiNote.NoteOn(0, 120))
	}

	midiNote := getMIDINote(chord.Pitches[0])
	tr.Add(getDurationTicks(chord.Duration, clock), midiNote.NoteOff(0))

	for _, p := range chord.Pitches[1:] {

		midiNote := getMIDINote(p)
		tr.Add(0, midiNote.NoteOff(0))
	}
}

func getMIDINote(note parser.Pitch) midi.Note {
	return midi.Note(music.GetPitchValue(note.Name) + note.Octave*12)
}

// makes a SMF and returns the bytes
func mkSMF(program parser.Program) []byte {
	var (
		bf    bytes.Buffer
		clock = smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also common
		tr    smf.Track
	)

	// first track must have tempo and meter informations
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(140))
	tr.Add(0, smf.MetaInstrument("Guitar"))

	tr.Add(0, midi.ProgramChange(0, gm.Instr_AcousticGuitarSteel.Value()))

	for _, e := range program.Events {

		switch event := e.(type) {
		case parser.Note:
			midiNote := getMIDINote(event.Pitch)

			tr.Add(0, midiNote.NoteOn(0, 120))
			tr.Add(getDurationTicks(event.Duration, clock), midiNote.NoteOff(0))
		case parser.TempoChange:
			tr.Add(0, smf.MetaTempo(float64(event.Value)))
		case parser.Chord:
			playChord(event, &tr, clock)
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
