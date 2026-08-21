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

	fmt.Println("out ports: \n" + midi.GetOutPorts().String())

	out, err := midi.FindOutPort("FLUID Synth")
	if err != nil {
		fmt.Printf("can't find fluidsynth")
		return
	}

	// create a SMF
	rd := bytes.NewReader(mkSMF())

	// read and play it
	smf.ReadTracksFrom(rd).Play(out)

}

// makes a SMF and returns the bytes
func mkSMF() []byte {
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

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.D(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.C(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.D(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))
	tr.Add(0, midi.E(5).NoteOff(0))

	tr.Add(clock.Ticks4th(), midi.D(5).NoteOn(0, 120))
	tr.Add(0, midi.D(5).NoteOff(0))

	tr.Add(clock.Ticks4th(), midi.C(5).NoteOn(0, 120))
	tr.Add(0, midi.C(5).NoteOff(0))

	tr.Add(clock.Ticks4th(), midi.D(5).NoteOn(0, 120))
	tr.Add(0, midi.D(5).NoteOff(0))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))
	tr.Add(0, midi.E(5).NoteOff(0))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))
	tr.Add(0, midi.E(5).NoteOff(0))

	tr.Add(clock.Ticks4th(), midi.E(5).NoteOn(0, 120))
	tr.Add(0, midi.E(5).NoteOff(0))

	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()
}
