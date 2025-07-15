package main

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model struct {
	status int
	err    error
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}

	res, err := c.Get(url)
	if err != nil {
		// There was an error making our request, wrap in a message and return
		return errMsg{err}
	}

	// Received response. Wrap in message and return.
	return statusMsg(res.StatusCode)
}

type statusMsg int
type errMsg struct{ err error }

// For messages with error it's useful to implement the error interface on the message
func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		// The server returned a status message. Save to our model.
		// Also tell Bubbletea to quit. We can still render a final view
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		// The server returned an error. Save to our model.
		// Tell bubbletea to quit.
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		// Ctrl+c quits. Even with short programs we should provide a way to quit easily.
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	// any other messages, don't do anything
	return m, nil
}

// We look at the model, and build accordingly.
func (m model) View() string {
	// If there's an error, print it out and don't do anything else
	if m.err != nil {
		return fmt.Sprintf("\n We've had some trouble: %v\n\n", m.err)
	}

	// Tell user we are doing something
	s := fmt.Sprintf("Checking %s ...\n\n", url)

	// When the server responds, we can add to current line
	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	// Send off string
	return "\n" + s + "\n\n"
}
