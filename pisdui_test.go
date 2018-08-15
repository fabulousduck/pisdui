package main

import (
	"github.com/fabulousduck/pisdui/pisdui"
	"log"
	"testing"
)

func TestPSD(t *testing.T) {
	psd, err := pisdui.NewPSD("./psd/test.psd")

	if err != nil {
		log.Fatalln(err)
		t.FailNow()
	}

	psd.Parse()
}
