package util

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/koki-develop/go-fzf"
	"github.com/sysnote8main/makemyaa/pkg/ugofetch/model"
)

func Question(msg string) string {
	tty, err := os.Open("/dev/tty")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(tty)
	defer tty.Close()
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

func ExtractUsernameList(userList []model.UserData) []string {
	choices := make([]string, 0)
	for _, v := range userList {
		choices = append(choices, v.Username)
	}
	return choices
}
