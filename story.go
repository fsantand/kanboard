package kanboard

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
  storyStyle = lipgloss.NewStyle().
    MarginBottom(0).
    BorderStyle(lipgloss.RoundedBorder()).
    Width(20)

  storyTitleStyle = lipgloss.NewStyle().
    Bold(true)

  statusStyle = lipgloss.NewStyle().
    Italic(true)

  focusedStyle = storyStyle.Copy().
    Foreground(lipgloss.Color("#FAFAFA")).
    BorderForeground(lipgloss.Color("#7D56F4"))
)

type Story struct {
  Id int
  Title, Status, Description string
}

func (s Story) Display(focused bool) string {
  sb := strings.Builder{}

  title := storyTitleStyle.Render(s.Title)

  sb.WriteString(fmt.Sprintf("#%d\n", s.Id))
  sb.WriteString(title)
  sb.WriteString("\n")
  
  if focused {
    return focusedStyle.Render(sb.String())
  }

  return storyStyle.Render(sb.String())
}
