package main

import "github.com/say4n/eightyfive/emulator"

func main() {
	emu := emulator.New()
	code := []string{"MVI B ff", "MVI C ff", "LDAX B", "MVI C fe", "STAX B", "HLT"}

	emu.Execute(code)
	emu.DumpRegister()
	// emu.DumpMemory()
}
