package nano

import (
	"log"
	"strconv"
	"strings"
)

/*
	9 ORE => 2 A
	8 ORE => 3 B
	7 ORE => 5 C
	3 A, 4 B => 1 AB
	5 B, 7 C => 1 BC
	4 C, 1 A => 1 CA
	2 AB, 3 BC, 4 CA => 1 FUEL

	A{9->2}
	B{8->3}
	C{7->5}

	A: 4
	B: 9
	C: 11
	AB: 2
	BC: 3
	CA: 4

	9 ORE
*/

type Input struct {
	Name     string
	Quantity float64
}

type Chemical struct {
	Name   string
	Inputs []Input
	Output float64
}

func NewFactory() Factory {
	return Factory{
		Chemicals: make(map[string]*Chemical),
		Stock:     make(map[string]float64),
	}
}

type Factory struct {
	Chemicals map[string]*Chemical
	Stock     map[string]float64
}

func (f Factory) OrePerFuel(reactions <-chan string) float64 {
	f.init(reactions)
	return f.need(1, f.Chemicals["FUEL"])
}

func (f Factory) need(qty float64, chem *Chemical) (total float64) {
	log.Println("START", qty, chem.Name, f.Stock)
	for _, input := range chem.Inputs {
		if input.Name == "ORE" {
			f.Stock[chem.Name] += chem.Output
			return input.Quantity
		}

		amount := (input.Quantity * qty) / chem.Output
		f.Stock[input.Name] -= amount
		for f.Stock[input.Name] < 0 {
			total += f.need(amount, f.Chemicals[input.Name])
		}
	}
	f.Stock[chem.Name] += qty
	log.Println("END", qty, chem.Name, f.Stock)
	return
}

func (f Factory) init(reactions <-chan string) {
	for reaction := range reactions {
		inputs, output := parts(reaction)

		oqty, oname := split(output)
		if _, ok := f.Chemicals[oname]; !ok {
			f.Chemicals[oname] = &Chemical{
				Name:   oname,
				Inputs: make([]Input, 0),
				Output: oqty,
			}
		}

		for _, input := range inputs {
			iqty, iname := split(input)
			(*f.Chemicals[oname]).Inputs = append(
				f.Chemicals[oname].Inputs,
				Input{iname, iqty},
			)
		}
	}
}

func parts(reaction string) (inputs []string, output string) {
	p := strings.Split(reaction, " => ")
	return strings.Split(p[0], ", "), p[1]
}

func split(input string) (qty float64, name string) {
	c := strings.Split(input, " ")
	qty, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		log.Fatal(err)
	}
	return qty, c[1]
}
