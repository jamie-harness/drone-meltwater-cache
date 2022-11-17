package auto_detect

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type mavenPreparer struct {}

func newMavenPreparer() *mavenPreparer {
	return &mavenPreparer{}
}
func (*mavenPreparer) PrepareRepo() error {
	path := filepath.Join(".mvn")
	fileName := "maven.config"
	pathToCache := filepath.Join("harness", ".m2", "repository")
	cmdToOverrideRepo := fmt.Sprintf(" -Dmaven.repo.local=/%s ", pathToCache)
	if _, err := os.Stat(filepath.Join(path, fileName)); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(path, fileName))
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(cmdToOverrideRepo)
		return err
	} else {
		f, err := os.OpenFile(path + fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(cmdToOverrideRepo)

	}
	return nil
}
