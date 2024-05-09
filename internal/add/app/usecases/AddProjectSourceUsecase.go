package usecases_add

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"strings"
)

type AddProjectSourceUsecase struct{}

func NewAddProjectSourceUsecase() *AddProjectSourceUsecase {
	return &AddProjectSourceUsecase{}
}

func (a *AddProjectSourceUsecase) saveProjectFolder(s string) {
	s = strings.Split(s, "\n")[0]
	home_dir, _ := os.UserHomeDir()
	file, err := os.ReadFile(path.Join(home_dir, ".tmux", "files", "directories.json"))
	if err != nil {
		panic(err)
	}

	var paths []string
	if err := json.Unmarshal(file, &paths); err != nil {
		panic(err)
	}

	for _, p := range paths {
		if p == s {
			return
		}
	}

	paths = append(paths, s)
	file, err = json.Marshal(paths)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path.Join(home_dir, ".tmux", "files", "directories.json"), file, 0644)
	if err != nil {
		panic(err)
	}
}

func (a *AddProjectSourceUsecase) Execute() {
	cmd := exec.Command("pwd")

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	a.saveProjectFolder(string(out))
}
