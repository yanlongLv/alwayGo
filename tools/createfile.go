package pkg

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

type project struct {
	dir         string
	fileName    string
	projectName string
}

func newProject(dir, fileName, projectName string) *project {
	return &project{dir, fileName, projectName}
}

// Create create new file and directory
func Create(projectfile, projectName string) error {
	_ = newProject(projectfile, "", projectName)

	boxs := packr.New("My new project", "./templates")
	os.MkdirAll(projectfile, 0755)
	for _, name := range boxs.List() {
		tmpl, _ := boxs.FindString(name)
		i := strings.LastIndex(name, string(os.PathSeparator))
		if i > 0 {
			dir := path.Join(projectfile, name[:i])
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		if strings.HasSuffix(name, ".tmpl") {
			name = strings.TrimSuffix(name, ".tmpl")
		}
		if err := write(projectfile, name, tmpl); err != nil {
			return err
		}

	}
	return nil
}

func write(projectFile, name, tmpl string) error {
	buf, err := parse(tmpl)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(projectFile, name), buf, 0644)
}

func parse(tmpl string) ([]byte, error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	var buff bytes.Buffer
	if err = t.Execute(&buff, t); err != nil {
		return nil, err
	}
	return buff.Bytes(), err
}
