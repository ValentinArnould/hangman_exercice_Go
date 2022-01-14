package main

import (
	"bufio"
	"fmt"
	"os"

	"projects.go/hangman/dictionary"
	"projects.go/hangman/hangman"
)

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("Appuyez sur une touche pour continuer...")
		input.Scan()
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionary.PickWord())
	if err != nil {
		fmt.Printf("Could not launch the game: %v\n", err)
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("Appuyez sur une touche pour continuer...")
		input.Scan()
		os.Exit(1)
	}

	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			input := bufio.NewScanner(os.Stdin)
			fmt.Print("Appuyez sur une touche pour continuer...")
			input.Scan()
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			input := bufio.NewScanner(os.Stdin)
			fmt.Print("Appuyez sur une touche pour continuer...")
			input.Scan()
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
