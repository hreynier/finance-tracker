package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Category item for list
type categoryItem struct {
	name string
}

func (i categoryItem) FilterValue() string { return i.name }

type categoryDelegate struct{}

func (d categoryDelegate) Height() int                             { return 1 }
func (d categoryDelegate) Spacing() int                            { return 0 }
func (d categoryDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

// TODO: Fix this being broken, what do we do here?
func (d categoryDelegate) Render(w *list.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(categoryItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.name)

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().
				Foreground(lipgloss.Color("212")).
				Background(lipgloss.Color("57")).
				PaddingLeft(2).
				PaddingRight(2).
				Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type transactionDelegate struct{}

func (d transactionDelegate) Height() int                             { return 2 }
func (d transactionDelegate) Spacing() int                            { return 1 }
func (d transactionDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

// TODO: Fix this being broken, what do we do here?
func (d transactionDelegate) Render(w *list.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Transaction)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s\n%s | %s %.2f", i.Name, i.DateString, i.Category, i.Amount)

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().
				Foreground(lipgloss.Color("212")).
				Background(lipgloss.Color("57")).
				PaddingLeft(2).
				PaddingRight(2).
				Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
