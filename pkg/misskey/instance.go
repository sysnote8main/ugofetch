package misskey

import "github.com/sysnote8main/ugofetch/pkg/misskey/endpoints"

type Instance struct {
	Host string
}

func (t Instance) User() endpoints.User {
	return endpoints.User{
		Host: t.Host,
	}
}
