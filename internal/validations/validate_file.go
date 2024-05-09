package validations

import (
	"os"
	"path"
)

func validateFolderWithPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

func validateFolder() {
	home_dir, _ := os.UserHomeDir()
	validateFolderWithPath(path.Join(home_dir, ".tmux"))
	validateFolderWithPath(path.Join(home_dir, ".tmux", "files"))
}

func validateFile() {
	home_dir, _ := os.UserHomeDir()
	if _, err := os.Stat(path.Join(home_dir, ".tmux", "files", "directories.json")); os.IsNotExist(err) {
		os.WriteFile(path.Join("files", "directories.json"), []byte("[]"), 0644)
	}
}

func ValidateFile() {
	validateFolder()
	validateFile()
}
