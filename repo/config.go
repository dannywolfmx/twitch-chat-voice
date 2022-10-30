package repo

import (
	"encoding/json"
	"io/fs"
	"os"
)

type Config struct {
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
	return r.config.Username
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
	buff, err := json.Marshal(r.config)

	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, buff, r.fileMode)
}

func (r *repoConfigFile) SaveAnonymousUsername(username string) error {
	r.config.Username = username
	return r.save()
}
