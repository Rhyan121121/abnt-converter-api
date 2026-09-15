package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"

	"abnt-converter-api/internal/abnt"
	"abnt-converter-api/internal/crossref"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var rawInput string

	clearScreen()
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Conversor ABNT").
				Description("Cole a citação APA completa aqui:").
				Value(&rawInput),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Operação cancelada pelo usuário.")
		os.Exit(0)
	}

	client := crossref.NewClient("seuemail@dominio.com")

	metadata, err := client.GetMetadata(rawInput)
	if err != nil {
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Bold(true)
		fmt.Println(errStyle.Render(fmt.Sprintf("Erro ao buscar metadados: %s", err)))
		os.Exit(1)
	}

	abntCitation := abnt.Formatter(*metadata)

	cardStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		MarginTop(1)

	if err := clipboard.WriteAll(abntCitation); err != nil {
		fmt.Println("Erro: na cópia para o clipboard")
		os.Exit(1)
	}
	fmt.Println(cardStyle.Render(abntCitation))
}
