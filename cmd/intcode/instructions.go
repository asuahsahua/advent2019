package intcode

const MAX_PARAMETER_COUNT = 3

type Instruction struct {
	paramCount int
	function func(ctx *InstructionContext)
}

var Instructions map[int]Instruction = map[int]Instruction{
	INST_ADD:           Instruction{3, I01_Add},
	INST_MULT:          Instruction{3, I02_Mult},
	INST_INPUT:         Instruction{1, I03_Input},
	INST_OUTPUT:        Instruction{1, I04_Output},
	INST_JUMP_IF_TRUE:  Instruction{2, I05_JumpIfTrue},
	INST_JUMP_IF_FALSE: Instruction{2, I06_JumpIfFalse},
	INST_LESS_THAN:     Instruction{3, I07_LessThan},
	INST_EQUALS:        Instruction{3, I08_Equals},

	INST_HCF:           Instruction{0, I99_HCF},
}

func (ctx *InstructionContext) Execute() {
	Instructions[ctx.FunctionID].function(ctx)
}