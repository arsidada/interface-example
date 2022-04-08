package reader

import (
	"os"
)

type FsReader struct {
	FileName string
}

func (fs *FsReader) Read() ([]byte, error) {
	return os.ReadFile(fs.FileName)
}
