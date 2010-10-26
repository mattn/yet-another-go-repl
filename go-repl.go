package main

import (
	"fmt"
	"os"
	"github.com/kless/go-readin/readin"
	"exp/eval"
)

func main() {
	world := eval.NewWorld()
	for {
		line, err := readin.Input("go> ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "go-repl: ", err.String())
			continue
		}
		code, err := world.Compile(line + ";")
		if err != nil {
			fmt.Fprintln(os.Stderr, "go-repl: ", err.String())
			continue
		}
		ret, err := code.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "go-repl: ", err.String())
			continue
		}
		if ret != nil {
			println(ret.String())
		}
	}
}
