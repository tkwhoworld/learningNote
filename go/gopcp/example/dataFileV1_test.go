package example

import (
	"os"
	"testing"
)

func removeFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err){
			return err
		}
		return nil
	}
	file.Close()
	return os.Remove(path)
}

func TestIDataFile(t *testing.T) {
	t.Run("example/all")
}