package main

import (
	"regexp"
	"strconv"
	"strings"

	. "github.com/asuahsahua/advent2019/cmd/common"
)

// --- Day 14: Space Stoichiometry --- 

// As you approach the rings of Saturn, your ship's low fuel indicator turns on.
// There isn't any fuel here, but the rings have plenty of raw material. Perhaps
// your ship's Inter-Stellar Refinery Union brand nanofactory can turn these raw
// materials into fuel.
type NanoFactory struct {
	Reactions Reactions
	Inventory Inventory
	OreBill  int
	FuelProduced int
}
func NewNanoFactory() *NanoFactory {
	return &NanoFactory{
		Reactions: make(Reactions),
		Inventory: make(map[Chemical]int),
		OreBill:  0,
		FuelProduced: 0,
	}
}

type Inventory map[Chemical]int

// You ask the nanofactory to produce a list of the reactions it can perform
// that are relevant to this process (your puzzle input).
type Input string

// Every reaction turns some quantities of specific input chemicals into some
// quantity of an output chemical. Almost every chemical is produced by exactly
// one reaction.
type Chemical string
type Reactions map[Chemical]Reaction

// Ore: the raw material input to the entire process and is not produced by
// a reaction.
const Ore Chemical = "ORE"

// You just need to know how much ORE you'll need to collect before you can
// produce one unit of FUEL.
const Fuel Chemical = "FUEL"

// Each reaction gives specific quantities for its inputs and output; reactions
// cannot be partially run, so only whole integer multiples of these quantities
// can be used.
type ChemicalQuantity struct {
	Chemical Chemical
	Quantity int
}
func (c Chemical) Quantity(i int) ChemicalQuantity {
	return ChemicalQuantity{
		Chemical: c,
		Quantity: i,
	}
}

type Reaction struct {
	Inputs []ChemicalQuantity
	Output ChemicalQuantity
}

// Consume what is available and return an amount left to produce
func (nf *NanoFactory) Consume(cq ChemicalQuantity) ChemicalQuantity {
	consumed := MinI(nf.Inventory[cq.Chemical], cq.Quantity)
	cq.Quantity -= consumed
	nf.Inventory[cq.Chemical] -= consumed
	return cq
}

// Produce the requested quantity and return the current bill
func (nf *NanoFactory) Produce(req ChemicalQuantity) int {
	/* if we're consuming ore, mark production on the bill, consume what we have */
	if req.Chemical == Ore {
		nf.OreBill += req.Quantity
		return nf.OreBill
	}

	/* consume what we can */
	if req = nf.Consume(req); req.Quantity == 0 {
		return nf.OreBill
	}

	/* and produce the rest */
	reaction, has := nf.Reactions[req.Chemical]
	PanicIf(!has, "Should have reaction for %s", req.Chemical)

	reactTimes := CeilDiv(req.Quantity, reaction.Output.Quantity)
	for _, input := range reaction.Inputs {
		nf.Produce(ChemicalQuantity{
			Chemical: input.Chemical,
			Quantity: reactTimes * input.Quantity,
		})
	}

	/* stash the extras */
	produced := reactTimes * reaction.Output.Quantity
	leftover := produced - req.Quantity
	nf.Inventory[reaction.Output.Chemical] += leftover

	if req.Chemical == Fuel {
		nf.FuelProduced += req.Quantity
	}

	return nf.OreBill
}

// --- Part Two ---

// After collecting ORE for a while, you check your cargo hold: 1 trillion
// (1000000000000) units of ORE.
var OneTrillion = 1000000000000
// Given 1 trillion ORE, what is the maximum amount of FUEL you can produce?

func main() {
	// Given the list of reactions in your puzzle input, what is the minimum amount
	// of ORE required to produce exactly 1 FUEL?
	nf := PuzzleInput.NanoFactory()
	oneFuelOre := nf.Produce(Fuel.Quantity(1))
	Part1("Ore for fuel: %d", oneFuelOre)

	availableOre := OneTrillion

	for {
		/* well, its nice that this +1/-1 thing worked and I didn't have to dig deeper */
		p99 := (9999 * (availableOre - nf.OreBill)) / (oneFuelOre * 10000) + 1
		used := nf.Produce(Fuel.Quantity(p99))
		if used > availableOre {
			Part2("Fuel produced for a trillion ore: %d", nf.FuelProduced - 1)
			return
		}
	}
}

var PuzzleInput Input = `1 GZJM, 2 CQFGM, 20 SNPQ, 7 RVQG, 3 FBTV, 27 SQLH, 10 HFGCF, 3 ZQCH => 3 SZCN
4 FCDL, 6 NVPW, 21 GZJM, 1 FBTV, 1 NLSNB, 7 HFGCF, 3 SNPQ => 1 LRPK
15 FVHTD, 2 HBGFL => 4 BCVLZ
4 GFGS => 4 RVQG
5 BCVLZ, 4 LBQV => 7 TWSRV
6 DWKTF, 4 VCKL => 4 KDJV
16 WZJB => 4 RBGJQ
8 RBGJQ, 5 FCDL, 2 LWBQ => 1 MWSX
100 ORE => 7 WBRL
7 PGZGQ => 5 FVHTD
1 JCDML, 2 TWSRV => 9 JSQSB
3 WZJB, 1 NXNR => 6 XFPVS
7 JPCPK => 8 JCDML
11 LWBQ, 8 XFPVS => 9 PSPFR
2 TWSRV => 8 NVPW
2 LBQV => 1 PMJFD
2 LCZBD => 3 FBTV
1 WBQC, 1 ZPNKQ => 8 JPCPK
44 HFGCF, 41 PSPFR, 26 LMSCR, 14 MLMDC, 6 BWTHK, 3 PRKPC, 13 LRPK, 50 MWSX, 8 SZCN => 1 FUEL
1 XFPVS => 4 BJRSZ
1 GWBDR, 1 MBQC => 4 HZPRB
2 BJRSZ, 9 KDJV, 1 XFPVS => 8 SNVL
7 PMJFD, 30 SNVL, 1 BJRSZ => 2 JMTG
8 SNVL, 1 RBGJQ => 9 FCDL
2 HZPRB => 6 NLSNB
2 GRDG => 9 VCKL
1 FVHTD => 9 WZJB
130 ORE => 2 GRDG
3 WZJB, 1 GFGS, 1 NXNR => 9 SNPQ
9 VCKL => 5 WBQC
1 WBRL, 11 FPMPB => 7 PGZGQ
118 ORE => 3 LMSCR
3 SQLH, 1 PMJFD, 4 XJBL => 7 MLMDC
1 LMSCR, 10 GRDG => 2 TBDH
6 DWKTF => 2 SQLH
2 BJRSZ, 1 PGZGQ, 3 NXNR => 7 MBQC
5 PRKPC => 7 NXNR
9 SQLH => 5 LCZBD
1 FCDL => 9 CQFGM
5 PGZGQ, 1 TBDH => 8 HBGFL
15 JSQSB => 5 HFGCF
2 PGZGQ, 1 VCKL => 4 ZPNKQ
3 FBTV, 3 JMTG => 5 QLHKT
1 ZGZST, 2 LCZBD => 7 GFGS
2 RVQG => 4 ZQCH
1 ZPNKQ => 5 LBQV
3 LWBQ => 8 XJBL
1 LBQV, 9 JCDML => 3 GWBDR
8 VCKL, 6 FVHTD => 9 DWKTF
3 JCDML => 3 ZGZST
160 ORE => 5 FPMPB
3 SQLH, 22 LBQV, 5 BCVLZ => 6 PRKPC
1 WZJB => 2 GZJM
10 ZGZST => 2 LWBQ
5 TBDH, 19 NXNR, 9 QLHKT, 2 KDJV, 1 SQLH, 1 GWBDR, 6 HFGCF => 4 BWTHK`

// functions to parse the input into our relevant structures
func (pi Input) NanoFactory() *NanoFactory {
	nf := NewNanoFactory()
	nf.Reactions = pi.Reactions()
	return nf
}

func (pi Input) Reactions() map[Chemical]Reaction {
	cmap := make(map[Chemical]Reaction)

	for _, line := range SplitLines(string(pi)) {
		match := regexp.MustCompile(`^(.*) => (.*)$`).FindStringSubmatch(line)
		PanicIf(len(match) != 3, "Expected match on line %s", line)

		inputs := make([]ChemicalQuantity, 0)
		for _, s := range strings.Split(match[1], ",") {
			inputs = append(inputs, parseChemicalQuantity(s))
		}

		cr := Reaction{
			Inputs: inputs,
			Output: parseChemicalQuantity(match[2]),
		};

		cmap[cr.Output.Chemical] = cr
	}

	return cmap
}

func parseChemicalQuantity(s string) ChemicalQuantity {
	rex := regexp.MustCompile(`(\d+) (\w+),?`)
	match := rex.FindStringSubmatch(s)
	PanicIf(len(match) == 0, "Expected to match resourceQuantityRex: `%s`", s)

	quantity, err := strconv.Atoi(match[1])
	PanicIfErr(err)

	return ChemicalQuantity{
		Quantity: quantity,
		Chemical: Chemical(match[2]),
	}
}