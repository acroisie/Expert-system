package menu

import (
	"expert-system/src/algo"
	"expert-system/src/factManager"
	"expert-system/src/models"
	"expert-system/src/rules"
	"expert-system/src/v"
	"fmt"
	"strings"

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
	scrollOffset     int
	maxVisibleLines  int
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
		scrollOffset:     0,
		maxVisibleLines:  10,
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
			switch msg.String() {
			case "up":
				if m.scrollOffset > 0 {
					m.scrollOffset--
				}
			case "down":
				totalLines := len(m.reasoningLogs) + len(m.problem.Facts) + len(m.problem.Queries) + 2
				if m.scrollOffset < totalLines-m.maxVisibleLines {
					m.scrollOffset++
				}
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
			factManager.SortFactListByAlphabet(factManager.FactList)
			var lines []string
			if len(m.reasoningLogs) == 0 {
				lines = append(lines, "Reasoning logs:", "No logs")
			} else {
				lines = append(lines, "Reasoning logs:")
				lines = append(lines, m.reasoningLogs...)
			}
			lines = append(lines, "Facts:")
			lines = append(lines, factsToList(factManager.FactList)...)
			lines = append(lines, "Queries:")
			lines = append(lines, queryResultsToStringList(m.problem.Queries, factManager.FactList)...)
			cleanLines := make([]string, 0, len(lines))
			for _, l := range lines {
				l = strings.TrimSpace(l)
				cleanLines = append(cleanLines, l)
			}
			start := m.scrollOffset
			end := start + m.maxVisibleLines
			if end > len(cleanLines) {
				end = len(cleanLines)
			}
			for _, line := range cleanLines[start:end] {
				str += line + "\n"
			}
			if len(cleanLines) > m.maxVisibleLines {
				str += instructionsStyle.Render("\nUse ↑/↓ to scroll.\n")
			}
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

func factsToList(facts []factManager.Fact) []string {
	var lines []string
	for _, f := range facts {
		lines = append(lines, fmt.Sprintf("%c = %s", f.Letter, f.Value))
	}
	return lines
}

func queryResultsToStringList(queries []models.Query, facts []factManager.Fact) []string {
	var result []string
	queryMap := make(map[rune]bool)
	for _, query := range queries {
		queryMap[query.Letter] = true
	}
	for _, f := range facts {
		if queryMap[f.Letter] {
			result = append(result, fmt.Sprintf("%c = %s", f.Letter, f.Value))
		}
	}
	if len(result) == 0 {
		result = append(result, "No queries found in facts.")
	}
	return result
}
