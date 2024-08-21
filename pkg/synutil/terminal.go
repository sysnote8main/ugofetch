package synutil

import (
	"fmt"
	"os"
)

func TitleStr(username string, host string) (string, string) {
	splitterLen := len(username) + len(host) + 1
	splitter := ""
	for i := 0; i < splitterLen; i++ {
		splitter += "-"
	}
	return fmt.Sprintf("\033[32m%s\033[39m@\033[32m%s\033[39m", username, host), splitter
}

func KvStr(name string, value string) string {
	if value == "" {
		value = "none"
	}
	return fmt.Sprintf("\033[33m%s\033[39m: %s", name, value)
}

func PrintFetch(arr []string, w int) {
	os.Stdout.Write([]byte("\033[25A"))
	for _, v := range arr {
		os.Stdout.WriteString("\033[51C" + v + LineBreak)
	}
	for i := 0; i < (25 - len(arr)); i++ {
		os.Stdout.Write([]byte("\n"))
	}
}
