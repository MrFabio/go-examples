package tui

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle      = lipgloss.NewStyle().Margin(2, 3)
	headerStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	codePaneStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, true, true, true).Padding(0, 1)
	outPaneStyle  = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, true, true, true).Padding(0, 1)
	gapStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("238")).SetString("│")
	viewportStyle = lipgloss.NewStyle()
)

type mode int

const (
	modeList mode = iota
	modeExample
)

type item struct {
	title    string
	run      func()
	file     string
	funcName string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

type model struct {
	list          list.Model
	mode          mode
	exampleFn     func()
	exampleTitle  string
	exampleOutput string
	outputVP      viewport.Model
	codeVP        viewport.Model
	width         int
	height        int
	codeRaw       string
	filePath      string
}

func captureOutput(fn func()) string {
	old := os.Stdout
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a pipe for stdin
	stdinR, stdinW, _ := os.Pipe()
	os.Stdin = stdinR

	// Start the function in a goroutine
	done := make(chan struct{})
	go func() {
		fn()
		close(done)
	}()

	// Wait for a short time to see if the function needs input
	select {
	case <-done:
		// Function completed without input
	case <-time.After(100 * time.Millisecond):
		// Function is waiting for input, provide some default input
		stdinW.WriteString("test input\n")
		<-done
	}

	w.Close()
	stdinW.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = old
	os.Stdin = oldStdin
	return buf.String()
}

func highlightGoCode(code string) string {
	var buf bytes.Buffer
	err := quick.Highlight(&buf, code, "go", "terminal16m", "monokai")
	if err != nil {
		return code
	}
	return buf.String()
}

func extractFunctionSource(filePath, funcName string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Could not read file: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	var src []string
	inFunc := false
	braceCount := 0
	funcPrefix := "func " + funcName + "("
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if !inFunc && strings.HasPrefix(trim, funcPrefix) {
			inFunc = true
		}
		if inFunc {
			src = append(src, line)
			braceCount += strings.Count(line, "{")
			braceCount -= strings.Count(line, "}")
			if braceCount == 0 && len(src) > 0 {
				break
			}
		}
	}
	if len(src) == 0 {
		return "Function not found."
	}
	return strings.Join(src, "\n")
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case modeList:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				i, ok := m.list.SelectedItem().(item)
				if ok {
					var output string
					if i.title == "IO Examples" || i.title == "Loops Examples" {
						output = "Example skipped - showing code only"
					} else {
						output = captureOutput(i.run)
					}
					codeRaw := extractFunctionSource(i.file, i.funcName)
					h, v := docStyle.GetFrameSize()
					codeWidth := (m.width - h) / 2
					outputWidth := m.width - h - codeWidth
					codeVP := viewport.New(codeWidth, m.height-v)
					outputVP := viewport.New(outputWidth, m.height-v)
					codeVP.SetContent(viewportStyle.Render(highlightGoCode(codeRaw)))
					outputVP.SetContent(viewportStyle.Render(output))
					return model{
						list: m.list, mode: modeExample, exampleFn: i.run, exampleTitle: i.title,
						exampleOutput: output, codeVP: codeVP, outputVP: outputVP, width: m.width, height: m.height, codeRaw: codeRaw,
						filePath: i.file,
					}, nil
				}
			}
		case tea.WindowSizeMsg:
			h, v := docStyle.GetFrameSize()
			m.list.SetSize(msg.Width-h, msg.Height-v)
			m.width = msg.Width
			m.height = msg.Height
		}
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	case modeExample:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" || msg.String() == "esc" {
				return model{list: m.list, mode: modeList, width: m.width, height: m.height}, nil
			}
		case tea.WindowSizeMsg:
			h, v := docStyle.GetFrameSize()
			codeWidth := (msg.Width - h) / 2
			outputWidth := msg.Width - h - codeWidth
			m.codeVP.Width = codeWidth
			m.codeVP.Height = msg.Height - v
			m.outputVP.Width = outputWidth
			m.outputVP.Height = msg.Height - v
			m.width = msg.Width
			m.height = msg.Height
		}
		var cmd tea.Cmd
		m.codeVP, cmd = m.codeVP.Update(msg)
		m.outputVP, _ = m.outputVP.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	switch m.mode {
	case modeList:
		return docStyle.Render(m.list.View())
	case modeExample:
		// Headers
		codeHeader := headerStyle.Render(fmt.Sprintf("Code (%s)", m.filePath))
		outHeader := headerStyle.Render(fmt.Sprintf("Output (%s)", m.exampleTitle))
		// Panes
		left := codePaneStyle.Width(m.codeVP.Width).Render(codeHeader + "\n" + m.codeVP.View())
		right := outPaneStyle.Width(m.outputVP.Width).Render(outHeader + "\n" + m.outputVP.View())
		gap := gapStyle.Render()
		row := lipgloss.JoinHorizontal(lipgloss.Top, left, gap, right)
		return "\n" + row
	}
	return ""
}
