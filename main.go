package main

import "github.com/say4n/eightyfive/emulator"

func main() {
	emu := emulator.New()
	code := []string{"NOP", "NOP", "HLT"}

	emu.Execute(code)
	// emu.DumpMemory()
}
