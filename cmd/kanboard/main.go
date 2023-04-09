package main

import (
	"fmt"

	"github.com/fsantand/kanboard"
)

func main()  {
  story_1 := kanboard.Story{
    Id: 1,
    Title: "Hello World",
    Status: "On going",
  }
  story_2 := kanboard.Story{
    Id: 2,
    Title: "Very long story title with information and stuff",
    Status: "On going",
  }

  column_1 := kanboard.Column{
    Focused: false,
    Status: "To Do",
    Color: "#bbbbbb",
    Stories: []kanboard.Story{
      story_1,
      story_2,
    },
  }

  column_2 := kanboard.Column{
    Focused: false,
    Status: "On going",
    Color: "#ffcc00",
    Stories: []kanboard.Story{
      story_1,
      story_2,
    },
  }

  column_3 := kanboard.Column{
    Focused: true,
    Status: "Completed",
    Color: "#65ff2d",
    Stories: []kanboard.Story{
      story_1,
      story_2,
    },
  }

  monitoring := kanboard.Column{
    Focused: false,
    Status: "Monitoring",
    Color: "#A020F0",
    Stories: []kanboard.Story{
    },
  }

  board := kanboard.Board{
    Columns: []kanboard.Column{
      column_1,
      column_2,
      monitoring,
      column_3,
    },
    Title: "Main Board",
  }
  
  fmt.Println(board.Display())
}
