package main

import (
	"github.com/fabulousduck/pisdui/pisdui"
)

func main() {

	pd := pisdui.NewInterpreter()
	pd.LoadFile("./test4.psd")
	pd.Parse()
}
