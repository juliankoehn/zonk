package http

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/juliankoehn/zonk/template"
)

type (
	httpParams struct {
		Encode  bool
		Package string
		Files   []*httpFile
	}
	httpFile struct {
		Base    string
		Name    string
		Path    string
		Ext     string
		Data    string
		Size    int64
		Time    int64
		Encoded bool
	}
)

func HttpHandle(pattern, packageName, output, prefix string) error {
	matches, err := doublestar.Glob(pattern)
	if err != nil {
		return err
	}

	params := httpParams{
		Encode:  false,
		Package: packageName,
	}

	for _, match := range matches {
		stat, oserr := os.Stat(match)
		if oserr != nil {
			return oserr
		}
		if stat.IsDir() {
			continue
		}
		raw, ioerr := ioutil.ReadFile(match)
		if ioerr != nil {
			return ioerr
		}
		encoded := true

		switch {
		case strings.HasSuffix(match, ".min.js"):
		case strings.HasSuffix(match, ".min.css"):
		case strings.HasSuffix(match, ".css"):
			encoded = false
		case strings.HasSuffix(match, ".js"):
			encoded = false
		case strings.HasSuffix(match, ".html"):
			encoded = false
		}

		data := string(raw)
		if !encoded {
			data = strings.Replace(data, "`", "`+\"`\"+`", -1)
		}

		params.Files = append(params.Files, &httpFile{
			Path:    strings.TrimPrefix(match, prefix),
			Name:    filepath.Base(match),
			Base:    strings.TrimSuffix(filepath.Base(match), filepath.Ext(match)),
			Ext:     filepath.Ext(match),
			Data:    data,
			Time:    stat.ModTime().Unix(),
			Size:    stat.Size(),
			Encoded: encoded,
		})
	}

	wr, err := os.Create(output)
	if err != nil {
		return err
	}
	defer wr.Close()

	for _, file := range params.Files {
		fmt.Println(file)
	}
	return template.Execute(wr, "http.tmpl", &params)
}
