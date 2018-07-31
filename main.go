package main

import (
	"github.com/fabulousduck/pisdui/pisdui"
)

func main() {

	pd := pisdui.NewInterpreter()
	pd.LoadFile("./test.psd")
	pd.Parse()
}
