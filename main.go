package main

import "github.com/say4n/eightyfive/emulator"

func main() {
	emu := emulator.New()
	code := []string{"MOV A M", "MOV B A", "HLT"}

	emu.Execute(code)
	emu.DumpRegister()
	// emu.DumpMemory()
}
