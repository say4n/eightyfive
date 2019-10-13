package emulator

import (
	"encoding/hex"
	"fmt"
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
		addressL := uint16(e5.register["L"])
		addressH := uint16(e5.register["H"])
		address := addressH<<8 + addressL

		log.Printf("emulator.instructionset.handleMOV:address: %04x", address)
		log.Printf("emulator.instructionset.handleMOV:content: %02x", e5.memory[address])
		log.Printf("emulator.instructionset.handleMOV:target: %s", target)
		e5.register[target] = e5.memory[address]

	} else if target == "M" {
		// MOV m, r
		log.Println("emulator.instructionset.handleMOV:MOV m, r")
		addressL := uint16(e5.register["L"])
		addressH := uint16(e5.register["H"])
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

func (e5 *eightyfive) handleMVI(instruction string) {
	contents := strings.Split(instruction, " ")
	target := contents[1]
	data, err := hex.DecodeString(contents[2])
	if err != nil {
		panic(err)
	}
	log.Printf("emulator.instructionset.handleMVI:data: %d", data)
	data_8 := uint8(data[0])
	log.Printf("emulator.instructionset.handleMVI:data_8: %02x", data_8)

	if target == "M" {
		// MVI m, data_8
		log.Println("emulator.instructionset.handleMVI:MVI m, data_8")
		addressL := uint16(e5.register["L"])
		addressH := uint16(e5.register["H"])
		address := addressH<<8 + addressL

		log.Printf("emulator.instructionset.handleMVI:target: %s", target)
		log.Printf("emulator.instructionset.handleMVI:address: %04x", address)
		e5.memory[address] = data_8

	} else {
		// MVI r, data_8
		log.Println("emulator.instructionset.handleMVI:MVI r, data_8")
		log.Printf("emulator.instructionset.handleMVI:target: %s", target)
		e5.register[target] = data_8
	}

	e5.pc++
}

func (e5 *eightyfive) handleLXI(instruction string) {
	contents := strings.Split(instruction, " ")
	register_pair := contents[1]
	data, err := hex.DecodeString(contents[2])
	if err != nil {
		panic(err)
	}
	log.Printf("emulator.instructionset.handleLXI:data: %d", data)
	lobyte := uint8(data[0])
	hibyte := uint8(data[1])
	log.Printf("emulator.instructionset.handleLXI:data: %02x %02x", lobyte, hibyte)

	if register_pair == "B" {
		// LXI rp, data_16
		log.Printf("emulator.instructionset.handleLXI:BC")
		e5.register["B"] = lobyte
		e5.register["C"] = hibyte

	} else if register_pair == "D" {
		// LXI rp, data_16
		log.Printf("emulator.instructionset.handleLXI:DE")
		e5.register["D"] = lobyte
		e5.register["E"] = hibyte

	} else if register_pair == "H" {
		// LXI rp, data_16
		log.Printf("emulator.instructionset.handleLXI:HL")
		e5.register["H"] = lobyte
		e5.register["L"] = hibyte

	} else {
		panic(fmt.Sprintf("Invalid register pair '%s' for LXI", register_pair))
	}

	e5.pc++
}

func (e5 *eightyfive) handleLDA(instruction string) {
	// LDA addr
	contents := strings.Split(instruction, " ")
	addressbytes, err := hex.DecodeString(contents[1])
	if err != nil {
		panic(err)
	}
	addressL := uint16(addressbytes[0])
	addressH := uint16(addressbytes[1])
	address := addressH<<8 + addressL
	log.Printf("emulator.instructionset.handleLDA:address: %04x", address)
	log.Printf("emulator.instructionset.handleLDA:data: %02x", e5.memory[address])

	e5.register["A"] = e5.memory[address]

	e5.pc++
}

func (e5 *eightyfive) handleSTA(instruction string) {
	// STA addr
	contents := strings.Split(instruction, " ")
	addressbytes, err := hex.DecodeString(contents[1])
	if err != nil {
		panic(err)
	}
	addressL := uint16(addressbytes[0])
	addressH := uint16(addressbytes[1])
	address := addressH<<8 + addressL
	log.Printf("emulator.instructionset.handleSTA:address: %04x", address)
	log.Printf("emulator.instructionset.handleSTA:data: %02x", e5.register["A"])

	e5.memory[address] = e5.register["A"]

	e5.pc++
}
