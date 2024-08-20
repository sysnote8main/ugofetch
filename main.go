package main

import (
	"fmt"
	"io"
	"os"

	// "github.com/sysnote8main/makemyaa/pkg/ugofetch/cmd"
	// "github.com/sysnote8main/makemyaa/pkg/bigaa"
	// "github.com/sysnote8main/makemyaa/pkg/ugofetch/util"
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
	//bigaa.BigAA(util.Question("Filepath: "))
	// testStr := "asdfghjasdfghjasdfghsaasjhgfdassdafghasdasljhgjkdlsy8ewiosdjklhfdsjklutwery89dsujdvsnmm,werudsyhukdlsfmerudfsyrtweyuifdlshjdm,ncvkhjsrnrlwtkem,cvbiuovbjiornm,ret m,cbujxcvbujiorstnrtn m,dfucvbuiorstnm,rsnm,fdjuiodfbuiosrjrtnm,srtkljfbuiocvjhkrsdnm,rtkljrstiouvbjm,cvnm,rtjilksrtuiorpsndfnm,cvbioufuisdojksrm,ndfgjicvbbuiocvbjilrtnklrstkljdfgjiofsdjklrtkL:srtmL:fguifuifsltj:srmL:rsjkofxjiojkrtkl:srgjiofgjiosfj:klsrtmlrS:fsipjfgijodjlR:stmlr:siopfsgjiopsrjtlR:stl:sfgjfiodjl:fkgl:"
	// f, err := os.Open("/dev/tty")
	// if err != nil {
	// 	panic(err)
	// }
	// w, h, err := term.GetSize(int(f.Fd()))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("WindowSize: ", h, "x", w)
	// r := runeStr(testStr, w-51)
	// for _, v := range r {
	// 	fmt.Println(v)
	// }
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
