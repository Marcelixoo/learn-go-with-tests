package filereader

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}

	return buf, nil
}
