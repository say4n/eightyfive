package main

import "github.com/say4n/eightyfive/emulator"

func main() {
	emu := emulator.New()
	code := []string{"MOV A M", "MOV B A", "MVI H ff", "MVI M e2", "LXI D abcd", "STA f0ff", "LDA f0ff", "HLT"}

	emu.Execute(code)
	emu.DumpRegister()
	// emu.DumpMemory()
}
