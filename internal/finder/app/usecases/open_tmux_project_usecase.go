package usecases_finder

import (
	"os/exec"
	"strings"
)

type OpenTmuxProjectUsecase struct{}

func NewOpenTmuxProjectUsecase() *OpenTmuxProjectUsecase {
	return &OpenTmuxProjectUsecase{}
}

func (o *OpenTmuxProjectUsecase) switchToProject(s string) {
	cmd := exec.Command("tmux", "switch-client", "-t", s)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (o *OpenTmuxProjectUsecase) Execute(project_folder string) {
	splited := strings.Split(project_folder, "/")
	project_name := splited[len(splited)-1]

	cmd := exec.Command("tmux", "has-session", "-t", project_name)
	if err := cmd.Run(); err == nil {
		o.switchToProject(project_name)
		return
	}

	cmd = exec.Command("tmux", "new-session", "-d", "-c", project_folder, "-s", project_name)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	o.switchToProject(project_name)
}
