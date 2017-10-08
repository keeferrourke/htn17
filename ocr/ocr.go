package ocr

import (
	/* Standard library packages */
	"errors"
	"os"
	"strings"

	/* Third party */
	"github.com/otiai10/gosseract"
	/* Local packages */)

func Process(path string) ([]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("path: cannot stat file")
	}

	client, err := gosseract.NewClient()
	if err != nil {
		return nil, err
	}
	out, err := client.Src(path).Out()
	if err != nil {
		return nil, err
	}

	if out == "" {
		return nil, nil
	}
	return strings.Split(out, " "), nil
}
