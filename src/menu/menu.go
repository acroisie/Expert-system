package menu

import (
	"expert-system/src/algo"
	"expert-system/src/factManager"
	"expert-system/src/models"
	"expert-system/src/rules"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	screenMenu screen = iota
	screenResolution
	screenAST
)

type MainModel struct {
	choices         []string
	cursor          int
	problem         *models.Problem
	showResolution  bool
	reasoningLogs   []string
	resolutionDone  bool
	ResolutionError string
	screen          screen
	astString       string
	currentASTIndex int
}

func InitMainModel(problem *models.Problem) MainModel {
	return MainModel{
		choices: []string{
			"Run Resolution",
			"Modify Facts",
			"Show Rules AST",
			"Quit",
		},
		cursor:          0,
		problem:         problem,
		screen:          screenMenu,
		currentASTIndex: 0,
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
		case "left":
			if m.screen == screenAST && m.currentASTIndex > 0 {
				m.currentASTIndex--
				m.astString = buildASTForRule(m.problem.Rules, m.currentASTIndex)
				return m, nil
			}
		case "right":
			if m.screen == screenAST && m.currentASTIndex < len(m.problem.Rules)-1 {
				m.currentASTIndex++
				m.astString = buildASTForRule(m.problem.Rules, m.currentASTIndex)
				return m, nil
			}
		case "b":
			m.screen = screenMenu
			return m, nil
		case "enter":
			switch m.choices[m.cursor] {
			case "Run Resolution":
				factManager.FactList = m.problem.Facts
				formattedRules := rules.RulesConditionalOperatorFormatter(m.problem.Rules)
				success, logs := algo.Algo(formattedRules)

				m.showResolution = true
				m.reasoningLogs = logs
				m.resolutionDone = success
				if !success {
					m.ResolutionError = "No solution found"
				} else {
					m.ResolutionError = ""
				}
				m.screen = screenResolution
				return m, nil

			case "Modify Facts":
				// Something to modify facts in bubbletea between true, false and unknown
				return m, nil
			case "Show Rules AST":
				m.astString = buildASTForRule(m.problem.Rules, m.currentASTIndex)
				m.screen = screenAST
				return m, nil
			case "Quit":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func buildASTForRule(rules []rules.Rule, index int) string {
	if index < 0 || index >= len(rules) {
		return "Index out of bounds"
	}
	var s string
	s += fmt.Sprintf("Rule %d:\n", index+1)
	s += rules[index].PrintAST()
	return s
}

func (m MainModel) View() string {
	style := lipgloss.NewStyle().Padding(1, 2)
	var str string

	switch m.screen {
	case screenAST:
		str += "---------- Rules AST ----------\n\n"
		str += m.astString
		str += "\nUse left/right arrow keys to navigate rules.\n"
		str += "\nPress b to go back.\nPress q to quit.\n"
		return style.Render(str)

	case screenResolution:
		if m.showResolution {
			str += "Resolution done: "
			if m.resolutionDone {
				str += "Success\n"
			} else {
				str += "Failure\n"
				if m.ResolutionError != "" {
					str += "Error: " + m.ResolutionError + "\n\n"
				}
			}

			str += "Reasoning logs:\n"
			if len(m.reasoningLogs) == 0 {
				str += "No logs\n"
			} else {
				for _, log := range m.reasoningLogs {
					str += log + "\n"
				}
			}

			str += "Facts:\n"
			factManager.SortFactListByAlphabet(factManager.FactList)
			str += factsToString(factManager.FactList) + "\n"

			str += "Press b to go back.\n"
			str += "Press q to quit.\n"

			return style.Render(str)
		}
	}

	str += "Expert System\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		str += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	str += "\nPress q to quit.\n"

	return style.Render(str)
}

func factsToString(facts []factManager.Fact) string {
	s := ""
	for _, f := range facts {
		s += fmt.Sprintf("%c = %s\n", f.Letter, f.Value)
	}
	return s
}
