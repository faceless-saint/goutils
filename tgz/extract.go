package tgz

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// Extract the target referenced by the given header. Data for regular
// files is read from the Reader. Supported target types are files,
// directories, symbolic links, and hard links.
func (tgz *Reader) Extract(header *tar.Header) error {
	switch header.Typeflag {

	case tar.TypeReg:
		if err := os.MkdirAll(filepath.Dir(header.Name), 0755); err != nil {
			return err
		}
		file, err := os.Create(header.Name)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tgz)
		return err

	case tar.TypeDir:
		return os.MkdirAll(header.Name, 0755)

	case tar.TypeSymlink:
		return os.Symlink(header.Linkname, header.Name)

	case tar.TypeLink:
		return os.Link(header.Linkname, header.Name)

	default:
		break
	}
	return nil
}

// ExtractAll iterates through the archive headers and extracts
// all referenced targets until encountering EOF.
func (tgz *Reader) ExtractAll() ([]*tar.Header, error) {
	var headers []*tar.Header
	for true {
		header, err := tgz.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return headers, err
		}
		if err := tgz.Extract(header); err != nil {
			return headers, err
		}
		headers = append(headers, header)
	}
	return headers, nil
}
