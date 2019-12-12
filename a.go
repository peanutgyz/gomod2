package gomod2

import (
	"fmt"

	"github.com/peanutgyz/gomod1"
)

func Gomod2Func() {
	fmt.Println("call Gomod2Func")
	gomod1.Gomod1Func()
}
