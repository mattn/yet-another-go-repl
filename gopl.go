package main

import (
	"bufio"
	"fmt"
	"os"
	"bitbucket.org/binet/go-eval/pkg/eval"
	"go/token"
)

func main() {
	world := eval.NewWorld()
	var fset = token.NewFileSet()
	r := bufio.NewReader(os.Stdin)
	for {
		print("# ")
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		code, err := world.Compile(fset, line+";")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.String())
			continue
		}
		ret, err := code.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.String())
			continue
		}
		if ret != nil {
			println(ret.String())
		}
	}
}
