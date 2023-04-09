package kanboard

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Board struct {
  Columns []Column
  Title string
}

var (
  boardStyle = lipgloss.NewStyle().Align(lipgloss.Center)

  boardTitleStyle = lipgloss.NewStyle().Bold(true).MarginBottom(3)
)

func (b Board) Display() string {
  var columns []string

  sb := strings.Builder{}

  for _, c := range(b.Columns) {
    if len(c.Stories) > 0 {
      columns = append(columns, c.Display())
    }
  }


  sb.WriteString(boardTitleStyle.Render(b.Title))
  sb.WriteString("\n")
  sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, columns...))
  sb.WriteString("\n")
  return boardStyle.Render(sb.String())
}
