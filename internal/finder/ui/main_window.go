package finder_ui

import (
	"fmt"
	"os"

	usecases_finder "github.com/Miguel-Panuto/tmux-go/internal/finder/app/usecases"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(2).Background(lipgloss.Color("#4a7a96")).Padding(0, 1).Foreground(lipgloss.Color("#292831")).Bold(true)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

type model struct {
	list   list.Model
	choice string
	s      *string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" || msg.String() == "return" {
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i.Title())
			}
			return m, tea.Quit

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		*m.s = m.choice
		return quitTextStyle.Render(m.choice)
	}
	return docStyle.Render(m.list.View())
}

func newModel(s *string) model {
	folders := usecases_finder.NewListProjectsUsecase().Execute()

	items := make([]list.Item, len(folders))
	for i, folder := range folders {
		items[i] = item{title: folder}
	}

	l := list.New(items, itemDelegate{}, 20, 20)
	l.Title = "Folder to open project"
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{
		list: l,
		s:    s,
	}
}

func Run() string {
	var s string
	if _, err := tea.NewProgram(newModel(&s), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return s
}
