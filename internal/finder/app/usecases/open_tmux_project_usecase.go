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

func (o *OpenTmuxProjectUsecase) isProjectAlreadyOpen(s string) bool {
	cmd := exec.Command("tmux", "has-session", "-t", s)

	if err := cmd.Run(); err == nil {
		return true
	}
	return false
}

func (o *OpenTmuxProjectUsecase) Execute(project_folder string) {
	splited := strings.Split(project_folder, "/")
	project_name := splited[len(splited)-1]

	if o.isProjectAlreadyOpen(project_name) {
		o.switchToProject(project_name)
		return
	}

	cmd := exec.Command("tmux", "new-session", "-d", "-c", project_folder, "-s", project_name)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	o.switchToProject(project_name)
}
