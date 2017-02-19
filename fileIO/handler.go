package fileIO

import (
	"io/ioutil"
	"os"
)

// Write takes a byte array and writes it to disk
func Write(fileContent []byte, filePath string) error {
	err := ioutil.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func fileSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0
	}

	return fileInfo.Size()
}

// Load file reference and return byte array of content
func Load(file *os.File) ([]byte, error) {
	fileSize := fileSize(file)
	data := make([]byte, fileSize)
	_, err := file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Open files a filepath and returns a file reference
func Open(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
