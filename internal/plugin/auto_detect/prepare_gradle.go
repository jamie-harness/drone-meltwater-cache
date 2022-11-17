package auto_detect

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type gradlePreparer struct {}

func newGradlePreparer() *gradlePreparer {
	return &gradlePreparer{}
}
func (*gradlePreparer) PrepareRepo() error {
	fileName := "gradle.properties"
	pathToCache := filepath.Join("harness", "harnesscache")
	cmdToOverrideRepo := fmt.Sprintf("systemProp.gradle.user.home=/%s/\norg.gradle.caching=true\n", pathToCache)
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(cmdToOverrideRepo)
		return err
	} else {
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(cmdToOverrideRepo)

	}
	//fileName = "/root/.gradle/gradle.properties"
	//if _, err := os.Stat(fileName); !errors.Is(err, os.ErrNotExist) {
	//	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//	if err != nil {
	//		return err
	//	}
	//	defer f.Close()
	//	_, err = f.WriteString(cmdToOverrideRepo)
	//}
	return nil
}
