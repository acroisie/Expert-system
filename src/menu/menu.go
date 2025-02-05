package menu

import (
	"expert-system/src/algo"
	"expert-system/src/factManager"
	"expert-system/src/models"
	"expert-system/src/rules"
	"expert-system/src/v"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	screenMenu screen = iota
	screenResolution
	screenAST
	screenEditFacts
)

var (
	titleStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#00bcd4")).Bold(true).Margin(0, 2)
	menuItemStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#bbbbbb"))
	menuCursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4caf50")).Bold(true)
	borderStyle       = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(lipgloss.Color("#444444"))
	instructionsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true)
)

type MainModel struct {
	choices          []string
	cursor           int
	problem          *models.Problem
	showResolution   bool
	reasoningLogs    []string
	resolutionDone   bool
	ResolutionError  string
	screen           screen
	astString        string
	currentASTIndex  int
	currentFactIndex int
}

func InitMainModel(problem *models.Problem) MainModel {
	return MainModel{
		choices: []string{
			"Run Resolution",
			"Modify Facts",
			"Show Rules AST",
			"Quit",
		},
		cursor:           0,
		problem:          problem,
		screen:           screenMenu,
		currentASTIndex:  0,
		currentFactIndex: 0,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "b" {
			m.screen = screenMenu
			return m, nil
		}
		switch m.screen {
		case screenMenu:
			switch msg.String() {
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
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
				case "Modify Facts":
					m.screen = screenEditFacts
					m.currentFactIndex = 0
				case "Show Rules AST":
					m.astString = buildASTForRule(m.problem.Rules, m.currentASTIndex)
					m.screen = screenAST
				case "Quit":
					return m, tea.Quit
				}
			}
		case screenEditFacts:
			switch msg.String() {
			case "up":
				if m.currentFactIndex > 0 {
					m.currentFactIndex--
				}
			case "down":
				if m.currentFactIndex < len(m.problem.Facts)-1 {
					m.currentFactIndex++
				}
			case "t":
				m.problem.Facts[m.currentFactIndex].Value = v.TRUE
			case "f":
				m.problem.Facts[m.currentFactIndex].Value = v.FALSE
			case "u":
				m.problem.Facts[m.currentFactIndex].Value = v.UNKNOWN
			}
		case screenAST:
			switch msg.String() {
			case "left":
				if m.currentASTIndex > 0 {
					m.currentASTIndex--
					m.astString = buildASTForRule(m.problem.Rules, m.currentASTIndex)
				}
			case "right":
				if m.currentASTIndex < len(m.problem.Rules)-1 {
					m.currentASTIndex++
					m.astString = buildASTForRule(m.problem.Rules, m.currentASTIndex)
				}
			}
		case screenResolution:
			// Aucune mise à jour spécifique pour cet écran pour l'instant.
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
	var str string
	switch m.screen {
	case screenAST:
		str += "---------- Rules AST ----------\n\n"
		str += m.astString
		str += "\n" + instructionsStyle.Render("Use left/right arrow keys to navigate rules.")
		str += "\n" + instructionsStyle.Render("Press b to go back. Press q to quit.")
		return borderStyle.Render(str)
	case screenResolution:
		if m.showResolution {
			str += "Resolution: "
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
			str += instructionsStyle.Render("Press b to go back. Press q to quit.")
			return borderStyle.Render(str)
		}
	case screenEditFacts:
		str += "---------- Modify Facts ----------\n\n"
		for i, fact := range m.problem.Facts {
			cursor := " "
			if i == m.currentFactIndex {
				cursor = ">"
			}
			str += fmt.Sprintf("%s %c = %s\n", cursor, fact.Letter, fact.Value)
		}
		str += "\n" + instructionsStyle.Render("Use up/down arrows to select a fact.")
		str += "\n" + instructionsStyle.Render("Press 't' for TRUE, 'f' for FALSE, 'u' for UNKNOWN.")
		str += "\n" + instructionsStyle.Render("Press b to go back to the menu.")
		return borderStyle.Render(str)
	}
	str += titleStyle.Render("Expert System")
	str += "\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = menuCursorStyle.Render(">")
			choice = menuItemStyle.Render(choice)
		} else {
			choice = menuItemStyle.Render(choice)
		}
		str += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	str += "\n" + instructionsStyle.Render("Press q to quit.")
	return borderStyle.Render(str)
}

func factsToString(facts []factManager.Fact) string {
	s := ""
	for _, f := range facts {
		s += fmt.Sprintf("%c = %s\n", f.Letter, f.Value)
	}
	return s
}
