package helloworld

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/version"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const EMPTY_DIR_PATH = "/volumes/empty-dir"
const FILE_CONFIG_MAP_PATH = "/volumes/config-map"
const SECRET_PATH = "/volumes/secrets"

func Say() string {
	fEmptyDir := filepath.Join(EMPTY_DIR_PATH, "data.txt")
	emptyDir, err := read(fEmptyDir)
	if err != nil {
		emptyDir = "1"
	}
	err = incr(fEmptyDir)
	if err != nil {
		return err.Error()
	}


	return fmt.Sprintf(
		"version=%s\n\n----------------------------------%s:= %s\n \n\n\n all env vars:= %s\n\n\n configmap1.properties=%s\n\n\n configmap2.properties=%s\n\n\n secret=%s\n\n\n",
		version.Get(),
		fEmptyDir,
		emptyDir,
		allEnvVars(),
		fileConfigMap("configmap1.properties"),
		fileConfigMap("configmap2.properties"),
		getSecret(SECRET_PATH),
	)
}

func read(f string) (string, error) {
	c, err := ioutil.ReadFile(f)
	if err != nil {
		err := ioutil.WriteFile(f, []byte("1"), 0777)
		if err != nil {
			return "", err
		}
		return "1", nil
	}
	return string(c), nil
}

func incr(f string) error {
	s, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	nb, err := strconv.Atoi(string(s))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(f, []byte(strconv.Itoa(nb + 1)), 0777)
	if err != nil {
		return err
	}
	return nil
}

func allEnvVars() string {
	out := ""
	for _, element := range os.Environ() {
		out = fmt.Sprintf("%s\n%s", out, element)
	}
	return out
}

func fileConfigMap(n string) string {
	f := filepath.Join(FILE_CONFIG_MAP_PATH, n)
	c, err := ioutil.ReadFile(f)
	if err != nil {
		return err.Error()
	}
	return string(c)
}

func getSecret(dir string) string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err.Error()
	}
	out := ""
	for _, file := range files {
		fullpath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			return fmt.Sprintf("%s\n%s\n\n", out, getSecret(fullpath))
		}
		content, err := ioutil.ReadFile(fullpath)
		if err != nil {
			out = out + err.Error()
		} else {
			out = fmt.Sprintf(
				"%s\n\nfile = %s\nbrut=%s",
				out,
				fullpath,
				content,
			)
		}
	}
	return out
}
