package usecases_finder

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

type ListProjectsUsecase struct{}

func NewListProjectsUsecase() *ListProjectsUsecase {
	return &ListProjectsUsecase{}
}

func (l *ListProjectsUsecase) getPaths() []string {
	home_dir, _ := os.UserHomeDir()
	file, err := os.ReadFile(path.Join(home_dir, ".tmux", "files", "directories.json"))
	if err != nil {
		panic(err)
	}

	var paths []string
	if err := json.Unmarshal(file, &paths); err != nil {
		panic(err)
	}

	return paths
}

func (l *ListProjectsUsecase) parseOutput(output string) []string {
	var parsed string
	if runtime.GOOS == "windows" {
		panic("Not implemented for windows")
	}

	parsed = strings.ReplaceAll(output, "//", "/")
	s := strings.Split(parsed, "\n")

	for i := range s {
		if s[i] == "" {
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}

func (l *ListProjectsUsecase) Execute() []string {
	paths := l.getPaths()
	if len(paths) == 0 {
		return []string{}
	}

	var args []string

	for _, p := range paths {
		args = append(args, p)
	}
	args = append(args, "-type", "d")
	args = append(args, "-mindepth", "1")
	args = append(args, "-maxdepth", "1")

	cmd := exec.Command("find", args...)

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return l.parseOutput(string(out))
}
