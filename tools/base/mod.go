package base

import (
	"io/ioutil"

	"golang.org/x/mod/modfile"
)

func ModulePath(filename string) (string, error) {
	modBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil

}
