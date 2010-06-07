package main

import (
	"fmt"
	"os"
	"readline"
	"exp/eval"
);

func main() {
	prompt := "go> ";
	world := eval.NewWorld()
	for {
		line := readline.ReadLine(&prompt);
		readline.AddHistory(*line);
		code, err := world.Compile(*line + ";")
		if err != nil {
			fmt.Fprintln(os.Stderr, "go-repl: ", err.String());
			continue
		}
		ret, err := code.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "go-repl: ", err.String());
			continue
		}
		if ret != nil {
			println(ret.String());
		}
	}
}
