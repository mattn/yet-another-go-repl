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
		func() {
			defer func() {
				recover()
			}()
			fmt.Print("# ")
			line, err := r.ReadString('\n')
			if err != nil {
				return
			}
			code, err := world.Compile(fset, line + "\n")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			ret, err := code.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			if ret != nil {
				fmt.Println(ret.String())
			}
		}()
	}
}
