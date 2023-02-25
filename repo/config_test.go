package repo

import (
	"os"
	"testing"
)

func TestNewRepoConfig(t *testing.T) {
	fileName := "prueba.json"
	_, err := NewRepoConfigFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	os.Remove(fileName)
}

func TestRepoConfigGetAnonymousUsername(t *testing.T) {
	fileName := "prueba.json"

	repo, err := NewRepoConfigFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	name := repo.GetAnonymousUsername()

	if name != "" {
		t.Fatalf("the username is %s", name)
	}

	os.Remove(fileName)
}

func TestRepoConfigSaveAnonymousUsername(t *testing.T) {
	fileName := "prueba.json"
	nameTest := "testName"

	repo, err := NewRepoConfigFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	name := repo.GetAnonymousUsername()

	if name != "" {
		t.Fatalf("the username is %s", name)
	}

	repo.SaveAnonymousUsername(nameTest)

	name = repo.GetAnonymousUsername()

	if name != nameTest {
		t.Fatalf("the username is %s", name)
	}

	repo, err = NewRepoConfigFile(fileName)

	if err != nil {
		t.Fatal(err)
	}

	name = repo.GetAnonymousUsername()

	if name != nameTest {
		t.Fatalf("the username is %s", name)
	}

	os.Remove(fileName)
}
