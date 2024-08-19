package model

type Role struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Color           interface{} `json:"color"`
	IconURL         string      `json:"iconUrl"`
	Description     string      `json:"description"`
	IsModerator     bool        `json:"isModerator"`
	IsAdministrator bool        `json:"isAdministrator"`
	DisplayOrder    int         `json:"displayOrder"`
}

type BadgeRole struct {
	Name         string `json:"name"`
	IconURL      string `json:"iconUrl"`
	DisplayOrder int    `json:"displayOrder"`
}
