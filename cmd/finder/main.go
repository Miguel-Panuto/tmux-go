package main

import (
	usecases_finder "github.com/Miguel-Panuto/tmux-go/internal/finder/app/usecases"
	finder_ui "github.com/Miguel-Panuto/tmux-go/internal/finder/ui"
)

func main() {
	project_folder := finder_ui.Run()
	open_tmux_project_usecase := usecases_finder.NewOpenTmuxProjectUsecase()

	open_tmux_project_usecase.Execute(project_folder)
}
