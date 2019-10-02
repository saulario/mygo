package main

import (
	"fmt"

	"github.com/mygo/edi/eancom"
)

func main() {
	fmt.Println("hola")
	unb := eancom.UnbType{}

	unb.S0001 = eancom.S0001Type{
		De0001: "UNOC",
		De0002: "4",
	}

	fmt.Println(unb.S0001.Serialize(eancom.DefaultUna()))

}
