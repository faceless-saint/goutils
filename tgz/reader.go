package tgz

import (
	"archive/tar"
	"compress/gzip"
	"io"
)

// Reader reads from a tar archive compressed using gzip. All
// methods from the internal tar.Reader are exposed for direct
// use. Decompression of the input is performed transparently.
type Reader struct {
	r  *tar.Reader
	gz *gzip.Reader
}

// NewReader returns a new Reader reading from the given input.
func NewReader(r io.Reader) (*Reader, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{tar.NewReader(gz), gz}, nil
}

// Next exposes the corresponding method from tar's Reader.
func (tgz *Reader) Next() (*tar.Header, error) {
	return tgz.r.Next()
}

// Read exposes the corresponding method from tar's Reader.
func (tgz *Reader) Read(b []byte) (int, error) {
	return tgz.r.Read(b)
}

// Close exposes the corresponding method from gzip's Reader.
func (tgz *Reader) Close() error {
	return tgz.gz.Close()
}

// Reset exposes the corresponding method from gzip's Reader.
func (tgz *Reader) Reset(r *io.Reader) error {
	return tgz.gz.Reset(*r)
}
