package synutil

import (
	"bufio"
	"fmt"
	"os"

	"github.com/koki-develop/go-fzf"
	"golang.org/x/term"
)

func Question(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	result := scanner.Text()
	return result
}

func Select(choices []string) ([]int, error) {
	f, err := fzf.New()
	if err != nil {
		return nil, err
	}

	indexList, err := f.Find(choices, func(i int) string { return choices[i] })
	return indexList, err
}

func runedTerminalOutput(text string, terminalFd int, spacerLen int) error {
	w, _, err := term.GetSize(terminalFd)
	if err != nil {
		return err
	}
	r := runeStr(text, w-spacerLen)
	for _, v := range r {
		fmt.Println(v)
	}
	return nil
}

func runeStr(text string, splitLen int) []string {
	result := make([]string, 0)
	for i := 0; i < len(text); i += splitLen {
		if i+splitLen < len(text) {
			result = append(result, text[i:(i+splitLen)])
		} else {
			result = append(result, text[i:])
		}
	}
	return result
}
