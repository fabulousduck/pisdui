package main

import (
	"github.com/fabulousduck/pisdui/pisdui"
)

func main() {

	pd := pisdui.NewPSD()
	pd.LoadFile("./psd/test.psd")
	pd.Parse()
}
