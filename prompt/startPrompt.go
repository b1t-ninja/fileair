package prompt

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"time"
)

var (
	FilePath string
	Duration time.Duration
)

func StartPrompt() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Blink(true).Background(lipgloss.Color("#45454")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(1)

	fmt.Println(style.Render("Welcome to File-Air"))
	time.Sleep(1 * time.Second)
	fmt.Println("          ______\n            _\\ _~-\\___\n    =  = ==(____AA____D\n                \\_____\\___________________,-~~~~~~~`-.._\n                /     o O o o o o O O o o o o o o O o  |\\_\n                `~-.__        ___..----..                  )\n                      `---~~\\___________/------------`````\n                      =  ===(_________D\n")
	time.Sleep(1 * time.Second)

	GetFilePath()
}

func GetFilePath() {
	form := huh.NewInput().Prompt("Path: ").Validate(ValidatePath).Value(&FilePath)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Create a map to associate the display strings with their corresponding durations
	var options = map[string]time.Duration{
		"1 Hour": time.Hour,
		"1 Day":  24 * time.Hour,
		"2 Days": 2 * 24 * time.Hour,
	}

	var selected string
	huh.NewSelect[string]().
		Title("Expiration: ").
		Options(
			huh.NewOption("1 Hour", "1 Hour"),
			huh.NewOption("1 Day", "1 Day"),
			huh.NewOption("2 Days", "2 Days"),
		).
		Value(&selected).Run()

	Duration, _ = options[selected]
}

func ValidatePath(path string) error {
	if len(path) == 0 {
		return errors.New("path can't be empty")
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("path doesn't exist")
	}
	return nil
}
