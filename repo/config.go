package repo

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

const DEFAULT_LANG string = "en"

type Config struct {
	ClientID      string        `json:"client_id"`
	Lang          string        `json:"lang"`
	TwitchInfo    TwitchInfo    `json:"twitch_info"`
	AnonymousUser AnonymousUser `json:"anonymous_user"`
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

type repoConfigFile struct {
	filename string
	fileMode fs.FileMode
	config   *Config
}

func NewRepoConfigFile(filename string) (*repoConfigFile, error) {
	config, err := getConfig(filename)

	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		config = &Config{}
	}

	return &repoConfigFile{
		filename: filename,
		fileMode: os.FileMode(0666),
		config:   config,
	}, nil
}

func (r *repoConfigFile) GetAnonymousUsername() string {
	return r.config.AnonymousUser.Username
}

func (r *repoConfigFile) GetClientID() (string, error) {
	if r.config.ClientID == "" {
		return "", errors.New("Missing ClientID in the config.json file")
	}
	return r.config.ClientID, nil
}

func (r *repoConfigFile) GetConfig() *Config {
	return r.config
}

func (r *repoConfigFile) GetLang() string {
	if r.config.Lang == "" {
		return DEFAULT_LANG
	}
	return r.config.Lang
}

func (r *repoConfigFile) GetTwitchToken() string {
	return r.config.TwitchInfo.Token
}

func getConfig(filename string) (*Config, error) {
	buff, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = json.Unmarshal(buff, c)

	return c, err
}

func (r *repoConfigFile) save() error {
	buff, err := json.MarshalIndent(r.config, "", "\t")

	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, buff, r.fileMode)
}

func (r *repoConfigFile) SaveAnonymousUsername(username string) error {
	r.config.AnonymousUser.Username = username
	return r.save()
}

func (r *repoConfigFile) SaveLang(lang string) error {
	r.config.Lang = lang
	return r.save()
}

func (r *repoConfigFile) SaveTwitchInfo(info TwitchInfo) error {
	r.config.TwitchInfo = info
	return r.save()
}
