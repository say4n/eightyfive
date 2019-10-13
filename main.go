package main

import "github.com/say4n/eightyfive/emulator"

func main() {
	emu := emulator.New()
	code := []string{"SHLD feff", "LHLD feff", "HLT"}

	emu.Execute(code)
	emu.DumpRegister()
	// emu.DumpMemory()
}
