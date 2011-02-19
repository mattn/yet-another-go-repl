package main

import (
	"fmt"
	"os"
	"github.com/kless/go-readin/readin"
	"exp/eval"
    "go/token"
)

func main() {
	world := eval.NewWorld()
    var fset = token.NewFileSet()
	for {
		line, err := readin.RepeatPrompt("go> ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "go-repl: ", err.String())
			continue
		}
        code, err := world.Compile(fset, line + ";")
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
