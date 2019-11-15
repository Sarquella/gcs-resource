package gcsresource

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func DynamicPath(path, replaceKey, replaceValue string) string{
	return strings.Replace(path, replaceKey, replaceValue, -1)
}

func DynamicPathValue(file, sourceDir string) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(sourceDir, file))
	if err != nil {
		return "", err
	}

	value := string(data)

	value = strings.Replace(value, "\r\n", "\n", -1) //For Windows support
	value = strings.Split(value, "\n")[0] //Only takes first line

	return value, nil
}