package intcode

// NewAmplifier returns a pointer to a new instance of Amplifier with a new
// instance of Computer that has loaded its program.
func NewAmplifier(prgm string, phaseSetting int) Amplifier {
	comp := NewComputer()
	comp.Load(prgm)

	return Amplifier{
		PhaseSetting: phaseSetting,
		Controller:   comp,
	}
}

// Amplifier defines the phase setting and the controller
type Amplifier struct {
	PhaseSetting int
	Controller   *Computer
}

// Exec runs the controller with the given input
func (amp Amplifier) Exec(input int) Output {
	inputs := make(chan int)
	defer close(inputs)

	go func() {
		inputs <- amp.PhaseSetting
		inputs <- input
	}()

	return <-amp.Controller.Exec(inputs)
}

// NewAmplificationCircuit creates a new circuit of amplifiers as long as the
// number of given phaseSettings.
func NewAmplificationCircuit(prgm string, phaseSettings ...int) AmplificationCircuit {
	amps := make([]Amplifier, len(phaseSettings))

	for i, phaseSetting := range phaseSettings {
		amps[i] = NewAmplifier(prgm, phaseSetting)
	}

	return amps
}

// AmplificationCircuit is a slice of Amplifier
type AmplificationCircuit []Amplifier

// Exec chains the inputs and outputs of all amplifiers in the circuit,
// producing a final output.
func (amps AmplificationCircuit) Exec(input int) (output Output) {
	for _, amp := range amps {
		output = amp.Exec(input)
		if output.Error != nil {
			return output
		}
		input = output.Value
	}

	return output
}
