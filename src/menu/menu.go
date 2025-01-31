package menu

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"

    "expert-system/src/algo"
    "expert-system/src/factManager"
    "expert-system/src/models"
    "expert-system/src/rules"
)

type MainModel struct {
	choices []string
	cursor  int
	problem *models.Problem

	showResolution bool
	reasoningLogs  []string
	resolutionDone bool
	ResolutionError string
}

func InitMainModel() MainModel {
	return MainModel{
		choices: []string{
			"Run Resolution",
			"Modify Facts",
			"Show Rules AST",
			"Quit",
		},
		cursor:  0,
		problem: nil,
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
				return m, tea.Quit
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case "b":
				if m.showResolution {
					m.showResolution = false
					return m, nil
				}
			case "enter":
				switch m.choices[m.cursor] {
					case "Run Resolution":
						factManager.FactList = m.problem.Facts
						formattedRules := rules.RulesConditionalOperatorFormatter(m.problem.Rules)

						success, logs := algo.Algo(formatedRules)

						m.showResolution = true
						m.reasoningLogs = logs
						m.resolutionDone = success

						if !success {
							m.ResolutionError = "No solution found"
						} else {
							m.ResolutionError = ""
						}

						return m, nil
					case "Modify Facts":
						// Something to modify facts in bubbletea between true, false and unknown
						return m, nil
					case "Show Rules AST":
						// Something to display the rules AST in bubbletea
						return m, nil
					case "Quit":
						return m, tea.Quit
					}
				}
			}
			return m, nil
	}


	func (m MainModel) View() string {
		style := lipgloss.NewStyle().Padding(1, 2)
		var s string

		if m.showResolution {
			s += "Resolution done: "
			if m.resolutionDone {
				s += "Success\n"
			} else {
				s += "Failure\n"
				if m.ResolutionError != "" {
					s += "Error: " + m.ResolutionError + "\n\n"
				}
			}

			s += "Reasoning logs:\n"
			if len(m.reasoningLogs) == 0 {
				s += "No logs\n"
			} else {
				for _, log := range m.reasoningLogs {
					s += log + "\n"
				}
			}

			s += "Facts:\n"
			factManager.SortFactListByAlphabet(factManager.FactList)
			s += factsToString(factManager.FactList) + "\n"

			return style.Render(s)
		}
	
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	
		s += "\nPress q to quit.\n"

		return lipgloss.NewStyle().Padding(1, 2).Render(s) // Make something wiht colors, rounded corners, etc.
	}
	
	func factsToString(facts []factManager.Fact) string {
		s := ""
		for _, f := range facts {
			s += fmt.Sprintf("%c = %s\n", f.Letter, f.Value)
		}
		return s
	}
	