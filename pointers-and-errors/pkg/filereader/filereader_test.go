package filereader_test

import (
	"path/filepath"
	"testing"

	"github.com/Marcelixoo/learn-go-with-tests/pointers-and-errors/pkg/filereader"
)

func TestReadFile(t *testing.T) {
	t.Run("reads content if file exists", func(t *testing.T) {
		config, err := filereader.ReadFile(filepath.Join("testdata", "settings.txt"))

		if err != nil {
			t.Errorf("unexpected error %q", err)
		}

		want := "test"
		if string(config) != want {
			t.Errorf("want %q got %q", want, config)
		}
	})

	t.Run("erroers if file does not exist", func(t *testing.T) {
		_, err := filereader.ReadFile("ops.txt")

		if err == nil {
			t.Error("expected error didn't occur")
		}
	})
}
