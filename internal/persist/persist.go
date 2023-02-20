package persist

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var ErrDirDoesntExist = errors.New("persist directory does not exist")

func PersistForecastResponse(persistDir string, body []byte) error {
	//check if persistDir exists
	exists, err := exists(persistDir)
	if err != nil || !exists {
		return ErrDirDoesntExist
	}

	filename := constructFilename(persistDir)
	//write body to file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func constructFilename(persistDir string) string {
	filename := fmt.Sprintf("meteoblue_%s.json", time.Now().Format(time.RFC3339))
	return filepath.Join(persistDir, filename)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
