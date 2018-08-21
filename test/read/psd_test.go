package readtests

import (
	"fmt"
	"testing"

	"github.com/fabulousduck/pisdui/pisdui"
)

func TestPsdFileRead(t *testing.T) {
	_, err := pisdui.NewPSD("../files/test.psd")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("successfull file read")
}
