package emulator

import (
	"log"
	"strings"
)

type eightyfive struct {
	register map[string]uint8 // 7x 8 bit registers
	flag     map[string]bool  // 5x 1 bit flag
	memory   [64 * 1024]uint8 // 64 KB memory
	pc       uint16           // 16 bit program counter
	sp       uint16           // 16 bit stack pointer
}

func New() *eightyfive {
	log.Println("emulator.emulator.New:init")
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
		line := code[e5.pc]
		if line == "HLT" {
			log.Printf("emulator.emulator.Execute:HLT, PC=%d\n", e5.pc)
			break
		} else if line == "NOP" {
			log.Printf("emulator.emulator.Execute:NOP, PC=%d\n", e5.pc)
			e5.pc++
		} else if strings.HasPrefix(line, "MOV ") {
			e5.handleMOV(line)
		} else if strings.HasPrefix(line, "MVI ") {
			e5.handleMVI(line)
		} else if strings.HasPrefix(line, "LXI ") {
			e5.handleLXI(line)
		} else if strings.HasPrefix(line, "LDA ") {
			e5.handleLDA(line)
		} else if strings.HasPrefix(line, "STA ") {
			e5.handleSTA(line)
		} else if strings.HasPrefix(line, "LHLD ") {
			e5.handleLHLD(line)
		} else if strings.HasPrefix(line, "SHLD ") {
			e5.handleSHLD(line)
		} else if strings.HasPrefix(line, "LDAX ") {
			e5.handleLDAX(line)
		} else if strings.HasPrefix(line, "STAX ") {
			e5.handleSTAX(line)
		} else {
			log.Printf("emulator.emulator.Execute:PC=%d\n", e5.pc)
		}
	}
}

func (e5 *eightyfive) DumpMemory() {
	for address, membyte := range e5.memory {
		log.Printf("emulator.emulator.DumpMemory:%04x: %02x", address, membyte)
	}
}

func (e5 *eightyfive) DumpRegister() {
	for register, content := range e5.register {
		log.Printf("emulator.emulator.DumpRegister:%s: %02x", register, content)
	}
}
