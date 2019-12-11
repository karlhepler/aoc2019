package intcode

import "fmt"

// NewAmplifier returns a pointer to a new instance of Amplifier with a new
// instance of Computer that has loaded its program.
func NewAmplifier(prgm string, phaseSetting int) *Amplifier {
	comp := NewComputer()
	comp.Load(prgm)

	return &Amplifier{
		PhaseSetting: phaseSetting,
		Controller:   comp,
	}
}

// Amplifier defines the phase setting and the controller
type Amplifier struct {
	PhaseSetting int
	Controller   *Computer
	Input        chan int
	Output       <-chan Output
}

// Exec runs the controller with the given input
func (amp *Amplifier) Exec() {
	amp.Input = make(chan int)
	amp.Output = amp.Controller.Exec(amp.Input)
}

// NewAmplificationCircuit creates a new circuit of amplifiers as long as the
// number of given phaseSettings.
func NewAmplificationCircuit(prgm string, phaseSettings ...int) AmplificationCircuit {
	amps := make([]*Amplifier, len(phaseSettings))

	for i, phaseSetting := range phaseSettings {
		amps[i] = NewAmplifier(prgm, phaseSetting)
	}

	return amps
}

// AmplificationCircuit is a slice of Amplifier
type AmplificationCircuit []*Amplifier

// Exec chains the inputs and outputs of all amplifiers in the circuit,
// producing a final output.
func (amps AmplificationCircuit) Exec(input int) (output Output) {
	for _, amp := range amps {
		amp.Exec()
	}

	output = Output{Value: input}
	looping := false

	for {
		fmt.Printf("[START] %d\n", output.Value)

		for _, amp := range amps {
			if !looping {
				amp.Input <- amp.PhaseSetting
			}
			amp.Input <- output.Value

			next, ok := <-amp.Output
			if !ok {
				break
			}
			if next.Error != nil {
				return
			}
			output = next
		}

		if amps[0].PhaseSetting < 5 {
			return
		}

		fmt.Printf("[END] %d\n", output.Value)
		looping = true
	}

	return
}
