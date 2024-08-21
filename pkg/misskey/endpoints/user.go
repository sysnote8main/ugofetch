package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sysnote8main/ugofetch/pkg/misskey/model"
	"github.com/sysnote8main/ugofetch/pkg/misskey/param"
)

type User struct {
	Host string
}

// TODO fix duplicate user
func (u User) SearchUserRecursive(param param.MisskeyUserSearchParam) ([]model.UserData, error) {
	result := make([]model.UserData, 0)
	for {
		d, err := u.SearchUser(param)
		if err != nil {
			return nil, err
		}
		result = append(result, d...)
		if len(d) < param.Limit {
			break
		}
		param.Offset += param.Limit
	}
	return result, nil
}

func (u User) SearchUser(param param.MisskeyUserSearchParam) ([]model.UserData, error) {
	paramJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s/api/users/search", u.Host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request. status_code: %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var userArr []model.UserData
	err = json.Unmarshal(b, &userArr)
	if err != nil {
		return nil, err
	}

	if len(userArr) == 0 {
		return nil, fmt.Errorf("no user has matched to this query")
	}

	return userArr, nil
}
