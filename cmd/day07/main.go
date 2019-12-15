package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
	. "github.com/asuahsahua/advent2019/cmd/intcode"
)

func main() {
	// There are five amplifiers connected in series; each one receives an input
	// signal and produces an output signal. They are connected such that the
	// first amplifier's output leads to the second amplifier's input, the
	// second amplifier's output leads to the third amplifier's input, and so
	// on. The first amplifier's input value is 0, and the last amplifier's
	// output leads to your ship's thrusters.

	//     O-------O  O-------O  O-------O  O-------O  O-------O
	// 0 ->| Amp A |->| Amp B |->| Amp C |->| Amp D |->| Amp E |-> (to thrusters)
	//     O-------O  O-------O  O-------O  O-------O  O-------O

	// The Elves have sent you some Amplifier Controller Software (your puzzle
	// input), a program that should run on your existing Intcode computer. Each
	// amplifier will need to run a copy of the program.

	// Try every combination of phase settings on the amplifiers. What is the
	// highest signal that can be sent to the thrusters?
	bestPhase := OptimizeAmpProgram(AmplifierControllerSoftware)
	Part1("%d", AmpProgramRun(AmplifierControllerSoftware, bestPhase))

	// It's no good - in this configuration, the amplifiers can't generate a
	// large enough output signal to produce the thrust you'll need. The Elves
	// quickly talk you through rewiring the amplifiers into a feedback loop:

	//       O-------O  O-------O  O-------O  O-------O  O-------O
	// 0 -+->| Amp A |->| Amp B |->| Amp C |->| Amp D |->| Amp E |-.
	//    |  O-------O  O-------O  O-------O  O-------O  O-------O |
	//    |                                                        |
	//    '--------------------------------------------------------+
	//                                                             |
	//                                                             v
	//                                                      (to thrusters)

	// Most of the amplifiers are connected as they were before; amplifier A's
	// output is connected to amplifier B's input, and so on. However, the
	// output from amplifier E is now connected into amplifier A's input. This
	// creates the feedback loop: the signal will be sent through the amplifiers
	// many times.


}

// The input
var AmplifierControllerSoftware string = `3,8,1001,8,10,8,105,1,0,0,21,30,47,60,81,102,183,264,345,426,99999,3,9,1002,9,5,9,4,9,99,3,9,1002,9,5,9,1001,9,4,9,1002,9,4,9,4,9,99,3,9,101,2,9,9,1002,9,4,9,4,9,99,3,9,1001,9,3,9,1002,9,2,9,101,5,9,9,1002,9,2,9,4,9,99,3,9,102,4,9,9,101,4,9,9,1002,9,3,9,101,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,99`