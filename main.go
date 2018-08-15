package main

import (
	"github.com/fabulousduck/pisdui/pisdui"
)

func main() {
	pisdui.NewPSD("./psd/test.psd").Parse()
}
