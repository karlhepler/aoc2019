package computer

import (
	"github.com/karlhepler/aoc2019/5.2/computer"
)

func NewAmplifierChain(prgm []int, phaseSettings [5]int) (chain AmplifierChain) {
	for i := range chain {
		chain[i] = Amplifier{
			Program:      append(prgm[:0:0], prgm...),
			PhaseSetting: phaseSettings[i],
		}
	}
	return
}

type AmplifierChain [5]Amplifier

func (chain AmplifierChain) Exec(input int) (output int, err error) {
	for _, amp := range chain {
		if input, err = amp.Exec(input); err != nil {
			return
		}
	}

	output = input
	return
}

type Amplifier struct {
	Program      []int
	PhaseSetting int
}

func (amp Amplifier) Exec(input int) (output int, err error) {
	output, err = computer.Exec(amp.Program, amp.PhaseSetting, input)
	return
}
