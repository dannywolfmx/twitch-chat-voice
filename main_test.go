package main

import (
	"testing"
)

func TestRepoDeleteConfig(t *testing.T) {

	repo := NewRepoConfig("prueba.json")

	if err := repo.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestRepoGetConfig(t *testing.T) {
	repo := NewRepoConfig("prueba.json")

	c := &Config{
		Username: "Prueba",
	}
	if err := repo.Save(c); err != nil {
		t.Fatal(err)
	}

	configFileData, err := repo.Get()

	if err != nil {
		t.Fatal(err)
	}

	if c.Username != configFileData.Username {
		t.Fatalf("El username %s es diferente al del archivo %s ", c.Username, configFileData.Username)
	}

	repo.Delete()
}

func TestRepoSaveConfig(t *testing.T) {
	c := &Config{
		Username: "user",
	}

	repo := NewRepoConfig("prueba.json")

	if err := repo.Save(c); err != nil {
		t.Fatal(err)
	}

	repo.Delete()
}
