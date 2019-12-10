package computer

func NewAmplifierChain(prgm []int, phases []int) (chain AmplifierChain) {
	for _, phase := range phases {
		chain = append(chain, Amplify(NewComputer(prgm), phase))
	}
	return
}

type AmplifierChain []Amplifier

func (chain AmplifierChain) Exec(input int) (int, error) {
	inputs, outputs := make(chan int), make(chan *Output)
	defer close(inputs)
	defer close(outputs)

	go func() {
		for _, amp := range chain {
			amp.Computer.Exec(inputs, outputs)
		}
	}()

	output := &Output{Value: input}
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
