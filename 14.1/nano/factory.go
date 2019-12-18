package nano

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
	157 ORE => 5 NZVS
	165 ORE => 6 DCFZ
	44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
	12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
	179 ORE => 7 PSHF
	177 ORE => 5 HKGWZ
	7 DCFZ, 7 PSHF => 2 XJWVT
	165 ORE => 2 GPVTF
	3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT
*/

type Input struct {
	Name     string
	Quantity int
}

type Chemical struct {
	Name   string
	Inputs []Input
	Output int
}

func NewFactory() Factory {
	return Factory{
		Chemicals: make(map[string]*Chemical),
		Stock:     make(map[string]int),
	}
}

type Factory struct {
	Chemicals map[string]*Chemical
	Stock     map[string]int
}

func (f Factory) OrePerFuel(reactions <-chan string) int {
	f.init(reactions)
	return f.need(1, f.Chemicals["FUEL"])
}

func (f Factory) need(qty int, chem *Chemical) (total int) {
	log.Printf("%s: HAVE %d; NEED %d\n", chem.Name, f.Stock[chem.Name], qty)

	for _, input := range chem.Inputs {
		if input.Name == "ORE" {
			f.Stock[chem.Name] += chem.Output
			return input.Quantity
		}

		amnt, rmdr := div(qty*input.Quantity, chem.Output)
		f.Stock[input.Name] -= amnt + rmdr
		for f.Stock[input.Name] < 0 {
			total += f.need(abs(f.Stock[input.Name]), f.Chemicals[input.Name])
		}
		log.Println("INPUT", f.Stock[input.Name], input.Name)
	}

	f.Stock[chem.Name] += qty

	log.Println("GET", qty, chem.Name, f.Stock)
	fmt.Println("")

	return total
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

func split(input string) (qty int, name string) {
	c := strings.Split(input, " ")
	qty, err := strconv.Atoi(c[0])
	if err != nil {
		log.Fatal(err)
	}
	return qty, c[1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func div(a, b int) (int, int) {
	return a / b, a % b
}
