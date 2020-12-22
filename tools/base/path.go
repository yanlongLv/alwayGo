package base

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
)

func copyFile(src, dst string, replaces []string) error {
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	var old string
	for i, next := range replaces {
		if i%2 == 0 {
			old = next
			continue
		}

		buf = bytes.ReplaceAll(buf, []bytes(old), []byte(next))
	}
	return ioutil.WriteFile(dst, buf, fileInfo.Mode())
}

func copyDir(src, dst string, replaces, ignores []string) error {
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dst, fileInfo.Mode()); err != nil {
		return err
	}
	readResult, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, result := range readResult {
		if hasSet(result.Name(), ignores) {
			continue
		}
		srcfile := path.Join(src, result.Name())
		dstfile := path.Join(src, result.Name())
		if result.IsDir() {
			if err = copyDir(srcfile, dstfile, replaces, ignores); err != nil {
				return err
			}
		} else {
			if err = copyFile(srcfile, dstfile, replaces); err != nil {
				return err
			}
		}
	}
	return nil
}

func hasSet(name string, sets []string) bool {
	for _, ig := range sets {
		if ig == name {
			return true
		}
	}
	return false
}
