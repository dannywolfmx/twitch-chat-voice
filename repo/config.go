package repo

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/dannywolfmx/twitch-chat-voice/model"
)

const DEFAULT_LANG string = "en"

type repoConfigFile struct {
	filename string
	fileMode fs.FileMode
	config   *model.Config
}

func NewRepoConfigFile(filename string) (*repoConfigFile, error) {
	config, err := getConfig(filename)

	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		config = &model.Config{
			Lang: "es",
		}
	}

	return &repoConfigFile{
		filename: filename,
		fileMode: os.FileMode(0666),
		config:   config,
	}, nil
}

func (r *repoConfigFile) AddChat(chat *model.Chat) error {
	if chat == nil {
		return errors.New("nil chat reference")
	}

	r.config.Chats = append(r.config.Chats, *chat)

	return r.save()
}

func (r *repoConfigFile) AddMuttedUser(user model.User) ([]model.User, error) {
	r.config.MuttedUsers = append(r.config.MuttedUsers, user)
	return r.GetMuttedUsers(), r.save()
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

func (r *repoConfigFile) GetConfig() *model.Config {
	return r.config
}

func (r *repoConfigFile) GetChats() []model.Chat {
	return r.config.Chats
}

func (r *repoConfigFile) GetLang() string {
	if r.config.Lang == "" {
		return DEFAULT_LANG
	}
	return r.config.Lang
}
func (r *repoConfigFile) GetMuttedUsers() []model.User {
	return r.config.MuttedUsers
}

func (r *repoConfigFile) GetTwitchUserInfo() model.TwitchUser {
	return r.config.TwitchInfo.TwitchUser
}

func (r *repoConfigFile) GetTwitchToken() string {
	return r.config.TwitchInfo.Token
}

func getConfig(filename string) (*model.Config, error) {
	buff, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &model.Config{}
	err = json.Unmarshal(buff, c)

	return c, err
}

func (r *repoConfigFile) RemoveChat(nameChannel string) error {
	if len(r.config.Chats) == 0 {
		return errors.New("no elements on the chats list")
	}
	for i, chat := range r.config.Chats {
		if chat.NameChannel == nameChannel {
			r.config.Chats = append(r.config.Chats[:i], r.config.Chats[i+1:]...)
			return r.save()
		}
	}

	return errors.New("missing channel name")
}

func (r *repoConfigFile) RemoveMuttedUser(user model.User) ([]model.User, error) {
	//No a error if the list is empty

	for i, userOnList := range r.config.MuttedUsers {
		if userOnList == user {
			r.config.MuttedUsers = append(r.config.MuttedUsers[:i], r.config.MuttedUsers[i+1:]...)
			return r.GetMuttedUsers(), r.save()
		}
	}

	return r.GetMuttedUsers(), nil
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

func (r *repoConfigFile) SaveTwitchInfo(info model.TwitchInfo) error {
	r.config.TwitchInfo = info
	return r.save()
}
