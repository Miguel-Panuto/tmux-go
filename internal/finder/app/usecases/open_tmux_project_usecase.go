package usecases_finder

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type OpenTmuxProjectUsecase struct{}

func NewOpenTmuxProjectUsecase() *OpenTmuxProjectUsecase {
	return &OpenTmuxProjectUsecase{}
}

func (o *OpenTmuxProjectUsecase) isTmuxRunning() bool {
	cmd := exec.Command("pgrep", "tmux")

	out, err := cmd.Output()
	if err != nil {
		return false
	}

	value := strings.TrimSpace(string(out))
	value = strings.TrimSuffix(value, "\n")
	num, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return num > 0
}

func (o *OpenTmuxProjectUsecase) attachProject(s string) {
	cmd := exec.Command("tmux", "attach", "-t", s)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if err := cmd.Wait(); err != nil {
			panic(err)
		}
	}

}

func (o *OpenTmuxProjectUsecase) switchToProject(s string) {
	cmd := exec.Command("tmux", "switch-client", "-t", s)
	if err := cmd.Run(); err != nil {
		o.attachProject(s)
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

	is_tmux_running := o.isTmuxRunning()

	cmd := exec.Command("tmux", "new-session", "-d", "-c", project_folder, "-s", project_name)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	if !is_tmux_running {
		o.attachProject(project_name)
		return
	}

	o.switchToProject(project_name)
}
