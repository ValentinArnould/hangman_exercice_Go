package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	          _______  _        _______  _______  _______  _       
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = ` 
     _______
    |       |
    |       O
    |      /|\
    |       |
    |      / \
  __|__
 |     |________
 |              |
 |______________|
`
	case 1:
		draw = ` 
     _______
    |       |
    |       
    |      /|\
    |       |
    |      / \
  __|__
 |     |________
 |              |
 |______________|
`
	case 2:
		draw = ` 
     _______
    |       |
    |       
    |      /|\
    |       |
    |      / 
  __|__
 |     |________
 |              |
 |______________|
`
	case 3:
		draw = ` 
     _______
    |       |
    |       
    |      /|\
    |       |
    |      
  __|__
 |     |________
 |              |
 |______________|
`
	case 4:
		draw = ` 
     _______
    |       |
    |       
    |      /|
    |       |
    |      
  __|__
 |     |________
 |              |
 |______________|
`
	case 5:
		draw = ` 
     _______
    |       |
    |       
    |       |
    |       |
    |      
  __|__
 |     |________
 |              |
 |______________|
`
	case 6:
		draw = ` 
     _______
    |       |
    |      
    |      
    |      
    |      
  __|__
 |     |________
 |              |
 |______________|
`
	case 7:
		draw = ` 
     
          
          
          
          
          
  __ __
 |     |________
 |              |
 |______________|
`
	case 8:
		draw = ` `
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess\n")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word\n", guess)

	case "lost":
		fmt.Print("You lost :( ! The word was: ")
		drawLetters((g.Letters))

	case "won":
		fmt.Print("You Won ! The word was: ")
		drawLetters((g.Letters))

	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
