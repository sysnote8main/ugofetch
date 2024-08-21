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

	// AA Array
	aaStr := synutil.SplitByLineBreak(asciiart.Generate_AA(selectedUser.AvatarURL))

	// Information Array
	var infoList []string
	if pipeInputStr != "" {
		infoList = strings.Split(pipeInputStr, synutil.LineBreak)
	} else {
		// Ugomono fetch main logic
		name, splitter := synutil.TitleStr(selectedUser.Username, instance.Host)
		infoList = []string{
			name,
			splitter,
			synutil.KvStr("Description", strings.Split(selectedUser.Description, "\n")[0]),
			synutil.KvStr("Birthday", selectedUser.Birthday),
			synutil.KvStr("Follower", fmt.Sprint(selectedUser.FollowersCount)),
			synutil.KvStr("Following", fmt.Sprint(selectedUser.FollowingCount)),
			synutil.KvStr("CreatedAt", selectedUser.CreatedAt.String()),
			synutil.KvStr("UpdatedAt", selectedUser.UpdatedAt.String()),
		}
	}

	// Generate Output
	spacer := synutil.MakeSpacer(50)
	outputArr := synutil.ListJoin(
		aaStr,
		infoList,
		spacer,
		"",
		" ",
	)

	// Printout
	for _, v := range outputArr {
		fmt.Println(v)
	}
}
