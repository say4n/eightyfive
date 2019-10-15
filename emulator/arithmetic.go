package emulator

import (
	"log"
	"math"
	"math/bits"
	"strings"
)

func (e5 *eightyfive) updateFlags() {
	A := e5.register["A"]
	e5.flag["Z"] = A == 0
	e5.flag["P"] = bits.OnesCount(uint(A))%2 == 0
	e5.flag["AC"] = (A>>4 + A&0x0f) > 15
	e5.flag["S"] = ((A&0x80)>>7)&0x1 == 1
}

func (e5 *eightyfive) handleADD(instruction string) {
	contents := strings.Split(instruction, " ")
	operand := contents[1]

	if operand == "M" {
		// ADD m
		log.Println("emulator.instructionset.handleADD:ADD m")
		addressL := uint16(e5.register["L"])
		addressH := uint16(e5.register["H"])
		address := addressH<<8 + addressL

		log.Printf("emulator.instructionset.handleADD:address: %04x", address)
		log.Printf("emulator.instructionset.handleADD:content: %02x", e5.memory[address])
		log.Printf("emulator.instructionset.handleADD:A: %02x", e5.register["A"])

		sum := uint8(e5.register["A"] + e5.memory[address])
		e5.flag["CY"] = math.MaxInt8-e5.register["A"] < e5.memory[address]

		log.Printf("emulator.instructionset.handleADD:sum: %02x", sum)
		e5.register["A"] = sum

	} else {
		// ADD r
		log.Printf("emulator.instructionset.handleADD:ADD %s", operand)

		log.Printf("emulator.instructionset.handleADD:A: %02x", e5.register["A"])
		log.Printf("emulator.instructionset.handleADD:%s: %02x", operand, e5.register[operand])

		sum := uint8(e5.register["A"] + e5.register[operand])
		e5.flag["CY"] = math.MaxUint8-e5.register["A"] < e5.register[operand]

		log.Printf("emulator.instructionset.handleADD:sum: %02x", sum)
		e5.register["A"] = sum
	}

	e5.updateFlags()

	e5.pc++
}

func (e5 *eightyfive) handleADC(instruction string) {
	contents := strings.Split(instruction, " ")
	operand := contents[1]

	if operand == "M" {
		// ADC m
		log.Println("emulator.instructionset.handleADC:ADC m")
		addressL := uint16(e5.register["L"])
		addressH := uint16(e5.register["H"])
		address := addressH<<8 + addressL

		log.Printf("emulator.instructionset.handleADC:address: %04x", address)
		log.Printf("emulator.instructionset.handleADC:content: %02x", e5.memory[address])
		log.Printf("emulator.instructionset.handleADC:A: %02x", e5.register["A"])

		var carry uint8
		if e5.flag["CY"] {
			carry = 1
		} else {
			carry = 0
		}

		sum := uint8(e5.register["A"] + e5.memory[address] + carry)
		e5.flag["CY"] = math.MaxInt8-e5.register["A"]-carry < e5.memory[address]

		log.Printf("emulator.instructionset.handleADC:sum: %02x", sum)
		e5.register["A"] = sum

	} else {
		// ADC r
		log.Printf("emulator.instructionset.handleADD:ADC %s", operand)

		log.Printf("emulator.instructionset.handleADC:A: %02x", e5.register["A"])
		log.Printf("emulator.instructionset.handleADC:%s: %02x", operand, e5.register[operand])

		var carry uint8
		if e5.flag["CY"] {
			carry = 1
		} else {
			carry = 0
		}

		sum := uint8(e5.register["A"] + e5.register[operand] + carry)
		e5.flag["CY"] = math.MaxUint8-e5.register["A"]-carry < e5.register[operand]

		log.Printf("emulator.instructionset.handleADC:sum: %02x", sum)
		e5.register["A"] = sum
	}

	e5.updateFlags()

	e5.pc++
}
