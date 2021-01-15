package liquibase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/flect"
)

type Generator struct {
	testPrefix string
}

var ErrName = errors.New("not valid path or name")

func (g Generator) Name() string {
	return "liquibase"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {

	ret, err := g.genPath(args, root)
	if err != nil {
		return err
	}
	path := ret[0]
	name := ret[1]
	timestamp := ret[2]
	_, err = os.Stat(path)
	if err == nil {
		fmt.Println("file/directory already exist ")

		return nil
	}

	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return (err)
		}

		tmpl, err := template.New("[timestamp]-[name-underscore].xml").Parse(mainTemplate)
		if err != nil {
			return err
		}

		data := struct {
			Name string
			Time string
		}{
			Name: name,
			Time: timestamp,
		}
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, data); err != nil {
			return err
		}

		err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
		if err != nil {
			return err
		}
	}
	return nil
}

//Genpath retunrs the path, the name of the file and the timestamp
func (g Generator) genPath(args []string, root string) ([]string, error) {
	var ret []string
	name := filepath.Base(args[3])
	if name == "." || name == "/" {
		return ret, ErrName
	}

	dir := filepath.Dir(args[3])
	if name == "." && dir == "." {
		return ret, ErrName
	}

	underscoreName := flect.Underscore(name)
	timestamp := time.Now().UTC().Format("20060102150405")
	if g.testPrefix != "" {
		timestamp = g.testPrefix
	}

	fullName := timestamp + "-" + underscoreName + ".xml"

	path := filepath.Join(root, "migrations", fullName)
	if dir != "." {
		path = filepath.Join(root, "migrations", dir, fullName)
	}
	ret = append(ret, path, underscoreName, timestamp)

	return ret, nil
}

func (g Generator) equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false

		}
	}
	return true
}