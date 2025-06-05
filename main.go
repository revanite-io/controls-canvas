package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/revanite-io/sci/pkg/layer2"
)

var (
	catalogName string
	catalog     layer2.Catalog

	selectedCapabilities      map[string]item
	triedToReselectCapability map[string]bool // Just having fun with this one

	titleText = "Controls Canvas"

	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

func main() {
	selectedCapabilities = make(map[string]item)
	triedToReselectCapability = make(map[string]bool)

	if _, err := tea.NewProgram(newCatalogInputModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running model for catalog input:", err)
		os.Exit(1)
	}
	os.Exit(0)
}
