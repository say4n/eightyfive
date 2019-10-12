package emulator

import "log"

type eightyfive struct {
	register map[string]uint8 // 7x 8 bit registers
	flag     map[string]bool  // 5x 1 bit flag
	memory   [64 * 1024]bool  // 64 Kb memory
	pc       uint16           // 16 bit program counter
	sp       uint16           // 16 bit stack pointer
}

func New() *eightyfive {
	log.Println("eightyfive.emulator.New:init")
	e5 := new(eightyfive)

	e5.register = make(map[string]uint8)
	e5.register["A"] = 0 // Accumulator
	e5.register["B"] = 0 // B register
	e5.register["C"] = 0 // C register
	e5.register["D"] = 0 // D register
	e5.register["E"] = 0 // E register
	e5.register["H"] = 0 // H register
	e5.register["L"] = 0 // L register

	e5.flag = make(map[string]bool)
	e5.flag["CY"] = false // Carry flag
	e5.flag["AC"] = false // Auxiliary carry flag
	e5.flag["S"] = false  // Sign flag
	e5.flag["Z"] = false  // Zero flag
	e5.flag["P"] = false  // Parity flag

	e5.pc = 0      // Program counter
	e5.sp = 0xffff // Stack pointer

	return e5
}

func (e5 *eightyfive) Execute(code []string) {
	for {
		if code[e5.pc] == "HLT" {
			log.Printf("eightyfive.emulator.Execute:HLT, PC=%d\n", e5.pc)
			break
		} else if code[e5.pc] == "NOP" {
			log.Printf("eightyfive.emulator.Execute:NOP, PC=%d\n", e5.pc)
			e5.pc++
		} else {
			log.Printf("eightyfive.emulator.Execute:PC=%d\n", e5.pc)
		}
	}
}

func (e5 *eightyfive) DumpMemory() {
	var bitset int8
	for address, bit := range e5.memory {
		if (address+1)%8 == 0 {
			log.Printf("eightyfive.emulator.DumpMemory:%04x: %02x", address-7, bitset)
			bitset = 0
		}
		if bit {
			bitset |= 1
		}

		bitset <<= 1
	}
}
