package main

import (
	"os"

	usecases_add "github.com/Miguel-Panuto/tmux-go/internal/add/app/usecases"
	usecases_finder "github.com/Miguel-Panuto/tmux-go/internal/finder/app/usecases"
	finder_ui "github.com/Miguel-Panuto/tmux-go/internal/finder/ui"
	"github.com/Miguel-Panuto/tmux-go/internal/validations"
)

func main() {
	validations.ValidateFile()

	args := os.Args
	if len(args) > 1 && args[1] == "add" {
		add_project_source_usecase := usecases_add.NewAddProjectSourceUsecase()
		add_project_source_usecase.Execute()
		return
	}

	project_folder := finder_ui.Run()
	open_tmux_project_usecase := usecases_finder.NewOpenTmuxProjectUsecase()
	if project_folder == nil {
		return
	}
	open_tmux_project_usecase.Execute(*project_folder)
}
