package helper

import (
	"bufio"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// Function to check if the uploaded file is an image
func IsImageFile(file *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	return allowedTypes[ext]
}

// Function to save uploaded file to server and return its contents as a string
func SaveUploadedFile(file *multipart.FileHeader) (string, error) {
    // Open uploaded file
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    // Read file content as string
    scanner := bufio.NewScanner(src)
    var fileString string
    for scanner.Scan() {
        fileString += scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return fileString, nil
}
