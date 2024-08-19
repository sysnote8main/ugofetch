package main

import (
	"io"
	"os"

	"github.com/sysnote8main/makemyaa/pkg/ugofetch/cmd"
	"golang.org/x/term"
)

var (
	pipeInputStr string = ""
)

func init() {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		stdInBytes, _ := io.ReadAll(os.Stdin)
		pipeInputStr = string(stdInBytes)
	}
}

func main() {
	cmd.Ugofetch(pipeInputStr)
}
