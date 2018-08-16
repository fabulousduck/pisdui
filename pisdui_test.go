package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/fabulousduck/pisdui/pisdui"
)

func TestPSD(t *testing.T) {
	psd, err := pisdui.NewPSD("./test.psd")
	fmt.Println("fugg")
	if err != nil {
		log.Fatalln(err)
		t.FailNow()
	}

	psd.Parse()
}
