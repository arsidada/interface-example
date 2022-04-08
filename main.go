package main

import (
	"fmt"
	"os"

	"github.com/arsidada/interface-example/pkg/reader"
)

type Reader struct {
	reader.Reader
}

func (r *Reader) Read() ([]byte, error) {
	return r.Reader.Read()
}

func main() {
	FsReader := Reader{
		Reader: &reader.FsReader{
			FileName: "test.json",
		},
	}

	data, err := FsReader.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(data))

	StdinReader := Reader{
		Reader: &reader.StdinReader{},
	}
	data, err = StdinReader.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(data))

}
