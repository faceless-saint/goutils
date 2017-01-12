package tgz

import (
	"archive/tar"
	"compress/gzip"
	"io"
)

// Writer writes to a tar archive compressed using gzip. All
// methods from the internal tar.Writer are exposed for direct
// use. Compression of the output is performed transparently.
type Writer struct {
	w  *tar.Writer
	gz *gzip.Writer
}

// NewWriter returns a new Writer writing to the given output.
func NewWriter(w io.Writer) *Writer {
	gz := gzip.NewWriter(w)
	return &Writer{tar.NewWriter(gz), gz}
}

// NewWriter returns a new Writer writing to the given output,
// using the given gzip compression level instead of the default.
func NewWriterLevel(w io.Writer, level int) (*Writer, error) {
	gz, err := gzip.NewWriterLevel(w, level)
	if err != nil {
		return nil, err
	}
	return &Writer{tar.NewWriter(gz), gz}, nil
}

// Write exposes the corresponding method from tar's Writer.
func (tgz *Writer) Write(b []byte) (int, error) {
	return tgz.w.Write(b)
}

// WriteHeader exposes the corresponding method from tar's Writer.
func (tgz *Writer) WriteHeader(hdr *tar.Header) error {
	return tgz.w.WriteHeader(hdr)
}

// Close exposes the corresponding method from tar's Writer
// and gzip's Writer, calling each in the appropriate order.
func (tgz *Writer) Close() error {
	if err := tgz.w.Close(); err != nil {
		defer tgz.gz.Close()
		return err
	}
	return tgz.gz.Close()
}

// Flush exposes the corresponding method from tar's Writer
// and gzip's Writer, calling each in the appropriate order.
func (tgz *Writer) Flush() error {
	if err := tgz.w.Flush(); err != nil {
		defer tgz.gz.Flush()
		return err
	}
	return tgz.gz.Flush()
}

// Reset exposes the corresponding method from gzip's Writer.
func (tgz *Writer) Reset(w *io.Writer) {
	tgz.gz.Reset(*w)
}
