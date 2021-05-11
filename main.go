package main

import (
	"github.com/getlantern/go-udtwrapper/udt"
)


func main() {
	s, err := Dial("ip4", "localhost:1900")
	if err != nil {
		panic(err)
	}

}
