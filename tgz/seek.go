package tgz

import (
	"archive/tar"
	"fmt"
	"io"
	"path/filepath"
)

// Seek attempts to find the given entry in the archive.
func (tgz *Reader) Seek(name string) (*tar.Header, error) {
	return tgz.seek(name, false)
}

// SeekRecursive attempts to find the given entry in the archive.
// Matches in subdirectories are also checked recursively.
func (tgz *Reader) SeekRecursive(name string) (*tar.Header, error) {
	return tgz.seek(name, true)
}

func (tgz *Reader) seek(name string, recursive bool) (*tar.Header, error) {
	for true {
		// Get the next header in the archive
		header, err := tgz.Next()
		if err == io.EOF {
			return nil, fmt.Errorf("file \"%s\" not found in archive", name)
		}
		if err != nil {
			return nil, err
		}
		// Return the first match in the archive
		if header.Name == name || (recursive && filepath.Base(header.Name) == name) {
			return header, nil
		}
	}
	return nil, nil
}
