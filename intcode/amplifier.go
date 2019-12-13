package intcode

// // NewAmplifier returns a pointer to a new instance of Amplifier with a new
// // instance of Computer that has loaded its program.
// func NewAmplifier(prgm string, phaseSetting int) *Amplifier {
// 	comp := NewComputer()
// 	comp.Load(prgm)

// 	return &Amplifier{
// 		PhaseSetting: phaseSetting,
// 		Controller:   comp,
// 	}
// }

// // Amplifier defines the phase setting and the controller
// type Amplifier struct {
// 	PhaseSetting int
// 	Controller   *Computer
// 	Input        chan int
// 	Output       <-chan int
// 	Halt         <-chan error
// }

// // Exec runs the controller with the given input
// func (amp *Amplifier) Exec() {
// 	amp.Input = make(chan int)
// 	amp.Output, amp.Halt = amp.Controller.Exec(amp.Input)
// }

// // NewAmplificationCircuit creates a new circuit of amplifiers as long as the
// // number of given phaseSettings.
// func NewAmplificationCircuit(prgm string, phaseSettings ...int) AmplificationCircuit {
// 	amps := make([]*Amplifier, len(phaseSettings))

// 	for i, phaseSetting := range phaseSettings {
// 		amps[i] = NewAmplifier(prgm, phaseSetting)
// 	}

// 	return amps
// }

// // AmplificationCircuit is a slice of Amplifier
// type AmplificationCircuit []*Amplifier

// // Exec chains the inputs and outputs of all amplifiers in the circuit,
// // producing a final output.
// func (amps AmplificationCircuit) Exec(input int) (output int, err error) {
// 	for _, amp := range amps {
// 		amp.Exec()
// 	}

// 	looping := false
// 	running := len(amps)
// 	output = input

// 	for running > 1 {
// 		for _, amp := range amps {
// 			if !looping {
// 				amp.Input <- amp.PhaseSetting
// 			}
// 			amp.Input <- output

// 			select {
// 			case output = <-amp.Output:
// 				if amp.Controller.Running == false {
// 					running--
// 				}
// 			case err = <-amp.Halt:
// 				return
// 			}
// 		}

// 		if amps[0].PhaseSetting < 5 {
// 			return
// 		}

// 		looping = true
// 	}

// 	return
// }
