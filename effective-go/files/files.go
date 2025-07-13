package files

import (
	"fmt"
	"os"
)

// Note: Always try to eliminate redundant else blocks.

func ReadFile() string {
	// ** Note: It is the file route from where the executable is! not where the package is! **
	filename := "./files/txt.txt"

	f, err := os.Open(filename)
	if err != nil {
		return "Error opening file: " + err.Error()
	}

	d, err := f.Stat()
	if err != nil {
		f.Close()
		return "Error getting file info: " + err.Error()
	}

	return fmt.Sprintf("File Name: %s\nFile Size: %d bytes\nFile Mode: %s\nFile Modified: %s",
		d.Name(), d.Size(), d.Mode(), d.ModTime().Format("2025-01-02 15:04:05"))
}
