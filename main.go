package main

import "github.com/say4n/eightyfive/emulator"

func main() {
	emu := emulator.New()
	code := []string{"ADD B", "ADD C", "HLT"}

	emu.Execute(code)
	emu.DumpRegister()
	emu.DumpFlags()
	// emu.DumpMemory()
}
