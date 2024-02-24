package protoss

import (
	"os"
)

// Read
type Reader struct {
	binary []byte
	offset int
}

func OpenBinary(path string) (reader *Reader, err error) {
	reader = new(Reader)
	reader.binary, err = os.ReadFile(path)
	return
}

func (reader *Reader) Close() (err error) {
	err = nil
	return
}
