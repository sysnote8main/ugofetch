package model

import "time"

type UserData struct {
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	Username             string             `json:"username"`
	Host                 interface{}        `json:"host"`
	AvatarURL            string             `json:"avatarUrl"`
	AvatarBlurhash       string             `json:"avatarBlurhash"`
	AvatarDecorations    []AvatarDecoration `json:"avatarDecorations"`
	IsBot                bool               `json:"isBot"`
	IsCat                bool               `json:"isCat"`
	Emojis               interface{}        `json:"emojis"`
	OnlineStatus         string             `json:"onlineStatus"`
	BadgeRoles           []BadgeRole        `json:"badgeRoles"`
	URL                  interface{}        `json:"url"`
	URI                  interface{}        `json:"uri"`
	MovedTo              interface{}        `json:"movedTo"`
	AlsoKnownAs          interface{}        `json:"alsoKnownAs"`
	CreatedAt            time.Time          `json:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt"`
	LastFetchedAt        interface{}        `json:"lastFetchedAt"`
	BannerURL            interface{}        `json:"bannerUrl"`
	BannerBlurhash       interface{}        `json:"bannerBlurhash"`
	IsLocked             bool               `json:"isLocked"`
	IsSilenced           bool               `json:"isSilenced"`
	IsSuspended          bool               `json:"isSuspended"`
	Description          string             `json:"description"`
	Location             string             `json:"location"`
	Birthday             string             `json:"birthday"`
	Lang                 interface{}        `json:"lang"`
	Fields               []interface{}      `json:"fields"`
	VerifiedLinks        []interface{}      `json:"verifiedLinks"`
	FollowersCount       int                `json:"followersCount"`
	FollowingCount       int                `json:"followingCount"`
	NotesCount           int                `json:"notesCount"`
	PinnedNoteIds        []interface{}      `json:"pinnedNoteIds"`
	PinnedNotes          []interface{}      `json:"pinnedNotes"`
	PinnedPageID         interface{}        `json:"pinnedPageId"`
	PinnedPage           interface{}        `json:"pinnedPage"`
	PublicReactions      bool               `json:"publicReactions"`
	FollowersVisibility  string             `json:"followersVisibility"`
	FollowingVisibility  string             `json:"followingVisibility"`
	TwoFactorEnabled     bool               `json:"twoFactorEnabled"`
	UsePasswordLessLogin bool               `json:"usePasswordLessLogin"`
	SecurityKeys         bool               `json:"securityKeys"`
	Roles                []Role             `json:"roles"`
	Memo                 interface{}        `json:"memo"`
}
