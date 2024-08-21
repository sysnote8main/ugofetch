package util

import "github.com/sysnote8main/ugofetch/pkg/misskey/model"

func ExtractUsernameList(userList []model.UserData) []string {
	choices := make([]string, 0)
	for _, v := range userList {
		choices = append(choices, v.Username)
	}
	return choices
}
