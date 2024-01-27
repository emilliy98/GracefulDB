package core

import (
	"crypto/rand"
	"fmt"
	"os"
)

// Name generation
func GenerateName() string {
	// This function is complete
	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

// Revision generation
func GenerateRev() string {
	// This function is complete
	b := make([]byte, 4)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

// Checking the folder name
func CheckFolderOrFile(patch, name string) bool {
	// This function is complete
	fullPath := fmt.Sprintf("%s%s", patch, name)
	f, err := os.Stat(fullPath)
	if err != nil {
		return false
	}

	if f.IsDir() {
		return true
	}

	return true
}
