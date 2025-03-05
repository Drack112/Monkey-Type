package cmd

import (
	"log"
	"math/rand/v2"
	"os"

	"github.com/Drack112/MonkeyType-Golang/pkg/ui"
	"github.com/brianvoe/gofakeit/v7"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var words = gofakeit.Sentence(rand.IntN(40) + 1)

var rootCmd = &cobra.Command{
	Use:   "monkeyrc",
	Short: "A monkey type clone, write in your terminal!",
	Run: func(cmd *cobra.Command, args []string) {
		ui := ui.New(words)
		program := tea.NewProgram(ui)

		if _, err := program.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
