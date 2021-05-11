package udt

import (
	"testing"
	"github.com/getlantern/go-udtwrapper"
)

func TestStuff(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:11000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
	} else {
		t.Logf("Socket is: %s", s)
	}
}

func main() {
	s, err := Dial("ip4", "localhost:1900")
	if err != nil {
		panic(err)
	}

}
