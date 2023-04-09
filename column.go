package kanboard

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
  columnStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    Align(lipgloss.Center)
)

type Column struct {
  Status string
  Stories []Story
  Color string
  Focused bool
}

func (c Column) Display() string {
  sb := strings.Builder{}

  sb.WriteString(fmt.Sprintf("%s\n",c.Status))

  for i, s := range(c.Stories) {
    focused := i%3 == 0 && c.Focused
    sb.WriteString(s.Display(focused))
    sb.WriteString("\n")
  }

  return columnStyle.BorderForeground(lipgloss.Color(c.Color)).Render(sb.String())
}
