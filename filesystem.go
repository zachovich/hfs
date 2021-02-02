package hfs

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const IndexFile = "index.html"

type HttpFileSystem struct {
	RootDir    string
	TrimPrefix string
	Masks      []string
}

func (hfs *HttpFileSystem) Open(path string) (http.File, error) {
	if hfs.TrimPrefix != "" {
		path = strings.TrimPrefix(path, hfs.TrimPrefix)
	}

	path = filepath.Clean(filepath.Join(hfs.RootDir, path))

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		defer f.Close()

		index := filepath.Join(f.Name(), IndexFile)
		if _, err := os.Stat(index); err != nil {
			return nil, os.ErrPermission
		}

		return os.Open(index) // return index.html instead of directory
	}

	// it is a file
	bn := filepath.Base(f.Name())
	if len(hfs.Masks) > 0 { // TODO: need a better way to accept regex!
		for _, n := range hfs.Masks {
			if bn == n {
				return f, nil
			}
		}

		return nil, os.ErrPermission
	}

	return f, nil
}
