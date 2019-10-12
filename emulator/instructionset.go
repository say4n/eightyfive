package emulator

import (
	"log"
	"strings"
)

func (e5 *eightyfive) handleMOV(instruction string) {
	contents := strings.Split(instruction, " ")
	target := contents[1]
	source := contents[2]

	if source == "M" {
		// MOV r, m
		log.Println("emulator.instructionset.handleMOV:MOV r, m")
		addressL := int16(e5.register["L"])
		addressH := int16(e5.register["H"])
		address := addressH<<8 + addressL

		log.Printf("emulator.instructionset.handleMOV:address: %04x", address)
		log.Printf("emulator.instructionset.handleMOV:content: %02x", e5.memory[address])
		log.Printf("emulator.instructionset.handleMOV:target: %s", target)
		e5.register[target] = e5.memory[address]

	} else if target == "M" {
		// MOV m, r
		log.Println("emulator.instructionset.handleMOV:MOV m, r")
		addressL := int16(e5.register["L"])
		addressH := int16(e5.register["H"])
		address := addressH<<8 + addressL

		log.Printf("emulator.instructionset.handleMOV:address: %04x", address)
		log.Printf("emulator.instructionset.handleMOV:content: %02x", e5.memory[address])
		log.Printf("emulator.instructionset.handleMOV:source: %s", source)
		e5.memory[address] = e5.register[source]

	} else {
		// MOV r1, r2
		log.Println("emulator.instructionset.handleMOV:MOV r1, r2")
		log.Printf("emulator.instructionset.handleMOV:target: %s", target)
		log.Printf("emulator.instructionset.handleMOV:source: %s", source)
		e5.register[target] = e5.register[source]
	}

	e5.pc++
}
