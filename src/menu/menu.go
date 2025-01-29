package menu

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"expert-system/src/algo"
	"expert-system/src/factManager"
	"expert-system/src/models"
	"expert-system/src/v"
)

// ─────────────────────────────────────────────────────────────────────────────
// MainModel
// ─────────────────────────────────────────────────────────────────────────────
type MainModel struct {
	choices  []string
	cursor   int
	problem  *models.Problem
}

func InitialModel(problem *models.Problem) MainModel {
	return MainModel{
		choices: []string{
			"Run Resolution",
			"Modify Facts",
			"Show Rules AST",
			"Quit",
		},
		problem: problem,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			// Quit the entire program
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			switch m.choices[m.cursor] {
			case "Run Resolution":
				// Call your resolution algorithm here
				algo.Algo(m.problem.Rules)

				// After it finishes, just show the final facts or return
				// to the main menu, depending on your preference.
				// For simplicity, let's just print them and go back to menu:
				fmt.Println("\nResolution complete. Final facts are:")
				factManager.DisplayFacts(m.problem.Facts)

				return m, nil // stay in main menu

			case "Modify Facts":
				// Switch to the fact editor sub-model
				return NewFactModel(m.problem), nil

			case "Show Rules AST":
				// Switch to the AST viewer sub-model
				return NewASTModel(m.problem), nil

			case "Quit":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	s := "Expert System Menu\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	// Optional styling with lipgloss:
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

// ─────────────────────────────────────────────────────────────────────────────
// FactModel
// ─────────────────────────────────────────────────────────────────────────────

type FactModel struct {
	problem  *models.Problem
	cursor   int
	selected map[int]struct{}
}

func NewFactModel(problem *models.Problem) FactModel {
	return FactModel{
		problem:  problem,
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m FactModel) Init() tea.Cmd {
	return nil
}

func (m FactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			// Return to the main menu
			return InitialModel(m.problem), nil

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.problem.Facts)-1 {
				m.cursor++
			}

		case " ":
			// Toggle the selected state
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		case "enter":
			// Update the actual Fact objects in m.problem.Facts
			// to TRUE for any selected index, for demonstration
			for idx := range m.selected {
				f := &m.problem.Facts[idx]
				_ = factManager.SetFactValueByLetter(f.Letter, v.TRUE, true)
			}
			return InitialModel(m.problem), nil
		}
	}
	return m, nil
}

func (m FactModel) View() string {
	s := "Select Facts to Activate (SPACE to toggle, ENTER to confirm):\n\n"

	for i, fact := range m.problem.Facts {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %c = %s\n",
			cursor,
			checked,
			fact.Letter,
			fact.Value,
		)
	}

	s += "\nPress ENTER to confirm changes, ESC or Q to go back.\n"
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

// ─────────────────────────────────────────────────────────────────────────────
// ASTModel
// ─────────────────────────────────────────────────────────────────────────────

type ASTModel struct {
	problem *models.Problem
	index   int
}

func NewASTModel(problem *models.Problem) ASTModel {
	return ASTModel{
		problem: problem,
		index:   0,
	}
}

func (m ASTModel) Init() tea.Cmd {
	return nil
}

func (m ASTModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			// Return to the main menu
			return InitialModel(m.problem), nil

		case "left", "h":
			if m.index > 0 {
				m.index--
			}
		case "right", "l":
			if m.index < len(m.problem.Rules)-1 {
				m.index++
			}
		}
	}
	return m, nil
}

func (m ASTModel) View() string {
	if len(m.problem.Rules) == 0 {
		return "No rules to display. (Press Q to go back)\n"
	}

	s := fmt.Sprintf("Rule %d/%d AST:\n\n", m.index+1, len(m.problem.Rules))

	// Print the AST of the current rule
	m.problem.Rules[m.index].PrintAST()

	s += "\n\n←/→ Navigate  (Press Q/ESC to return to Menu)"
	return s
}

// ─────────────────────────────────────────────────────────────────────────────
// LaunchMenu
// ─────────────────────────────────────────────────────────────────────────────

func LaunchMenu(problem *models.Problem) {
	p := tea.NewProgram(InitialModel(problem))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running menu: %v\n", err)
		os.Exit(1)
	}
}
