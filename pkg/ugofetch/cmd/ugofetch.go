package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sysnote8main/makemyaa/pkg/ugofetch/misskey"
	"github.com/sysnote8main/makemyaa/pkg/ugofetch/model/param"
	"github.com/sysnote8main/makemyaa/pkg/ugofetch/util"
)

const LineBreak = "\n"

func Ugofetch(pipeInputStr string) {
	host := "misskey.systems"
	instance := misskey.Instance{
		Host: host,
	}

	// TODO add flag for username query
	queryStr := util.Question("Target username: ")

	searchParam := param.MisskeyUserSearchParam{
		Query:  queryStr,
		Limit:  100,
		Offset: 0,
		Origin: "local",
		Detail: true,
	}

	userArr, err := instance.SearchUser(searchParam)
	if err != nil {
		log.Fatal("Failed to search user: ", err)
	}

	choices := util.ExtractUsernameList(userArr)

	indexList, err := util.Select(choices)
	if err != nil {
		log.Fatal("Failed to select: ", err)
	}
	selectedUser := userArr[indexList[0]]

	// log.Printf("You selected: %s\n", selectedUser.Username)
	// log.Printf("AvatarUrl: %s", selectedUser.AvatarURL)

	fmt.Printf("%v\n", util.Generate_AA(selectedUser.AvatarURL))

	if pipeInputStr != "" {
		PrintFetch(strings.Split(pipeInputStr, LineBreak))
	} else {
		// Ugomono fetch main logic
		name, splitter := titleStr(selectedUser.Username, instance.Host)
		PrintFetch([]string{
			name,
			splitter,
			kvstr("Description", strings.Split(selectedUser.Description, "\n")[0]),
			kvstr("Birthday", selectedUser.Birthday),
			kvstr("Follower", fmt.Sprint(selectedUser.FollowersCount)),
			kvstr("Following", fmt.Sprint(selectedUser.FollowingCount)),
			kvstr("CreatedAt", selectedUser.CreatedAt.String()),
			kvstr("UpdatedAt", selectedUser.UpdatedAt.String()),
		})
	}
}

func titleStr(username string, host string) (string, string) {
	splitterLen := len(username) + len(host) + 1
	splitter := ""
	for i := 0; i < splitterLen; i++ {
		splitter += "-"
	}
	return fmt.Sprintf("\033[32m%s\033[39m@\033[32m%s\033[39m", username, host), splitter
}

func kvstr(name string, value string) string {
	if value == "" {
		value = "none"
	}
	return fmt.Sprintf("\033[33m%s\033[39m: %s", name, value)
}

func PrintFetch(arr []string) {
	os.Stdout.Write([]byte("\033[25A"))
	for _, v := range arr {
		os.Stdout.Write([]byte("\033[51C"))
		os.Stdout.WriteString(v + LineBreak)
	}
	for i := 0; i < (25 - len(arr)); i++ {
		os.Stdout.Write([]byte("\n"))
	}
}
