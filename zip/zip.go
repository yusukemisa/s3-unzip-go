package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src, dst string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dst, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			continue
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := io.Copy(f, rc); err != nil {
			return err
		}

	}
	return nil
}
