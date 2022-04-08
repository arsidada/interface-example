package reader

import (
	"bufio"
	"fmt"
	"os"
)

type StdinReader struct{}

func (sir *StdinReader) Read() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return []byte(text), nil
}
