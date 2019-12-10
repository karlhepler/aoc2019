package computer

func NewAmplifierChain(prgm []int, phases []int) (chain AmplifierChain) {
	for _, phase := range phases {
		chain = append(chain, Amplify(NewComputer(prgm), phase))
	}
	return
}

type AmplifierChain []Amplifier

func (chain AmplifierChain) Exec(input int) (int, error) {
	inputs, outputs := make(chan int), make(chan Output)
	defer close(inputs)
	defer close(outputs)

	go func() {
		for _, amp := range chain {
			amp.Computer.Exec(inputs, outputs)
		}
	}()

	output := Output{Value: input}
	for _, amp := range chain {
		inputs <- amp.PhaseSetting
		inputs <- output.Value

		output = <-outputs
		if output.Error != nil {
			return -1, output.Error
		}
	}

	return output.Value, output.Error
}

func Amplify(comp Computer, phase int) Amplifier {
	return Amplifier{
		Computer:     comp,
		PhaseSetting: phase,
	}
}

type Amplifier struct {
	Computer
	PhaseSetting int
}

// import (
// 	"github.com/karlhepler/aoc2019/5.2/computer"
// )

// func NewAmplifierChain(prgm []int, phaseSettings [5]int) (chain AmplifierChain) {
// 	for i := range chain {
// 		chain[i] = Amplifier{
// 			Program:      clone(prgm),
// 			PhaseSetting: phaseSettings[i],
// 		}
// 	}
// 	return
// }

// type AmplifierChain [5]Amplifier

// func (chain AmplifierChain) Exec(input int) (output int, err error) {
// 	for range chain {
// 		output, err = chain.exec(input)
// 		if err != nil || output == 0 {
// 			return
// 		}
// 		input = output
// 	}
// 	return
// }

// func (chain AmplifierChain) exec(input int) (output int, err error) {
// 	for _, amp := range chain {
// 		if input, err = amp.Exec(input); err != nil {
// 			return
// 		}
// 	}

// 	output = input
// 	return
// }

// type Amplifier struct {
// 	Program      []int
// 	PhaseSetting int
// }

// func (amp Amplifier) Exec(input int) (output int, err error) {
// 	output, err = computer.Exec(amp.Program, amp.PhaseSetting, input)
// 	return
// }

// func clone(prgm []int) []int {
// 	return
// }
