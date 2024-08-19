package param

type MisskeyUserSearchParam struct {
	Query  string `json:"query"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Origin string `json:"origin"`
	Detail bool   `json:"detail"`
}
