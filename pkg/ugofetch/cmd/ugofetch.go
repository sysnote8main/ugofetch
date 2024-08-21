package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/sysnote8main/ugofetch/pkg/asciiart"
	"github.com/sysnote8main/ugofetch/pkg/misskey"
	"github.com/sysnote8main/ugofetch/pkg/misskey/param"
	"github.com/sysnote8main/ugofetch/pkg/synutil"
	"github.com/sysnote8main/ugofetch/pkg/ugofetch/util"
)

func Ugofetch(pipeInputStr string) {
	host := "misskey.systems"
	instance := misskey.Instance{
		Host: host,
	}

	// TODO add flag for username query
	queryStr := synutil.Question("Target username: ")

	searchParam := param.MisskeyUserSearchParam{
		Query:  queryStr,
		Limit:  100,
		Offset: 0,
		Origin: "local",
		Detail: true,
	}

	userArr, err := instance.User().SearchUserRecursive(searchParam)
	if err != nil {
		log.Fatal("Failed to search user: ", err)
	}

	choices := util.ExtractUsernameList(userArr)

	indexList, err := synutil.Select(choices)
	if err != nil {
		log.Fatal("Failed to select: ", err)
	}
	selectedUser := userArr[indexList[0]]

	// log.Printf("You selected: %s\n", selectedUser.Username)
	// log.Printf("AvatarUrl: %s", selectedUser.AvatarURL)

	w := 51
	// TODO check on windows
	// termW, _, err := term.GetSize(int(os.Stdout.Fd()))
	// if err == nil {
	// 	w = termW
	// }

	// fmt.Printf("Width: %d\n", w)
	fmt.Printf("%v\n", asciiart.Generate_AA(selectedUser.AvatarURL))

	if pipeInputStr != "" {
		synutil.PrintFetch(strings.Split(pipeInputStr, synutil.LineBreak), w)
	} else {
		// Ugomono fetch main logic
		name, splitter := synutil.TitleStr(selectedUser.Username, instance.Host)
		synutil.PrintFetch([]string{
			name,
			splitter,
			synutil.KvStr("Description", strings.Split(selectedUser.Description, "\n")[0]),
			synutil.KvStr("Birthday", selectedUser.Birthday),
			synutil.KvStr("Follower", fmt.Sprint(selectedUser.FollowersCount)),
			synutil.KvStr("Following", fmt.Sprint(selectedUser.FollowingCount)),
			synutil.KvStr("CreatedAt", selectedUser.CreatedAt.String()),
			synutil.KvStr("UpdatedAt", selectedUser.UpdatedAt.String()),
		}, w)
	}
}

func outputList(title string, content map[string]string) []string {
	// TODO implement this
	return []string{}
}
