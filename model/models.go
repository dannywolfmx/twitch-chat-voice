package model

type User = string

type Config struct {
	ClientID      string        `json:"client_id"`
	Lang          string        `json:"lang"`
	TwitchInfo    TwitchInfo    `json:"twitch_info"`
	AnonymousUser AnonymousUser `json:"anonymous_user"`
	Chats         []Chat        `json:"chats"`
	MuttedUsers   []User        `json:"mutted_users"`
	SampleRateTTS int           `json:"samplerate_tts"`
}

type Chat struct {
	NameChannel string `json:"name_channel"`
}

type TwitchInfo struct {
	Token      string     `json:"token"`
	TwitchUser TwitchUser `json:"twitch_user"`
}

type TwitchUser struct {
	ID              string `json:"id"`
	BroadcasterType string `json:"broadcaster_type"`
	CreatedAt       string `json:"created_at"`
	Description     string `json:"description"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	Login           string `json:"login"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Type            string `json:"type"`
}

type AnonymousUser struct {
	Username string `json:"username"`
}
