package main

import (
	"fmt"
	"math/rand"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	words = []string{"ABDEL", "DJOKOVIC", "PETITBABYFOOT"}
	
)
type Game struct {
	word           string
	guesses        []string
	maxAttempts    int
	currentAttempt int
}

func NewGame() *Game {
	word := words[rand.Intn(len(words))]
	guesses := make([]string, len(word))
	for i := range guesses {
		guesses[i] = "_"
	}
	return &Game{
		word:           word,
		guesses:        guesses,
		maxAttempts:    9,
		currentAttempt: 0,
	}
}

func (g *Game) makeGuess(guess string) {
	g.currentAttempt++
	found := false

	for i, letter := range g.word {
		if strings.ToUpper(string(letter)) == strings.ToUpper(guess) {
			g.guesses[i] = string(letter)
			found = true
		}
	}

	if !found {
	
	} else {
		fmt.Println("Correct guess!")
	}
}

func drawPendu(step int) {
	rl.DrawText("H A N G M A N", 10, 10, 30, rl.DarkBlue)
	
	// Draw the hangman
	switch step {

	case 1:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		 //racine
		 
	case 2:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // tronc
		
	case 3:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // branche
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)



	case 4:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // corde
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)
		corde := rl.LoadTexture("res/corde.png")
		rl.DrawTexture(corde,550, 215,rl.White)


		
	case 5:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // tete
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)
		corde := rl.LoadTexture("res/corde.png")
		rl.DrawTexture(corde,550, 215,rl.White)
		tete := rl.LoadTexture("res/tete.png")
		rl.DrawTexture(tete,500 , 370, rl.White)
	case 6:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // bras gauche
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)
		corde := rl.LoadTexture("res/corde.png")
		rl.DrawTexture(corde,550, 215,rl.White)
		tete := rl.LoadTexture("res/tete.png")
		rl.DrawTexture(tete,500 , 370, rl.White)
		brasgauche :=rl.LoadTexture("res/brasdroit.png")
		rl.DrawTexture(brasgauche,500, 370, rl.White)
	case 7:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // bras droit
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)
		corde := rl.LoadTexture("res/corde.png")
		rl.DrawTexture(corde,550, 215,rl.White)
		tete := rl.LoadTexture("res/tete.png")
		rl.DrawTexture(tete,500 , 370, rl.White)
		brasgauche :=rl.LoadTexture("res/brasdroit.png")
		rl.DrawTexture(brasgauche,500, 370, rl.White)
		bras :=rl.LoadTexture("res/brasgauche.png")
		rl.DrawTexture(bras,500 ,370 ,rl.White)
	case 8:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // jambe gauche
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)
		corde := rl.LoadTexture("res/corde.png")
		rl.DrawTexture(corde,550, 215,rl.White)
		tete := rl.LoadTexture("res/tete.png")
		rl.DrawTexture(tete,500 , 370, rl.White)
		brasgauche :=rl.LoadTexture("res/brasdroit.png")
		rl.DrawTexture(brasgauche,500, 370, rl.White)
		bras :=rl.LoadTexture("res/brasgauche.png")
		rl.DrawTexture(bras,500 ,370 ,rl.White)
		jambegauche:= rl.LoadTexture("res/jambegauche.png")
		rl.DrawTexture(jambegauche, 500, 370, rl.White)
	case 9:
		test := rl.LoadTexture("res/bas.png")
		rl.DrawTexture(test, 180, 700, rl.White)
		vertical := rl.LoadTexture("res/vertical.png")
		rl.DrawTexture(vertical, 180, 150,rl.Black) // jambes droite
		branche := rl.LoadTexture("res/branche.png")
		rl.DrawTexture(branche,400, 195, rl.Black)
		corde := rl.LoadTexture("res/corde.png")
		rl.DrawTexture(corde,550, 215,rl.White)
		tete := rl.LoadTexture("res/tete.png")
		rl.DrawTexture(tete,500 , 370, rl.White)
		brasgauche :=rl.LoadTexture("res/brasdroit.png")
		rl.DrawTexture(brasgauche,500, 370, rl.White)
		bras :=rl.LoadTexture("res/brasgauche.png")
		rl.DrawTexture(bras,500 ,370 ,rl.White)
		jambegauche:= rl.LoadTexture("res/jambegauche.png")
		rl.DrawTexture(jambegauche, 500, 370, rl.White)
		tout := rl.LoadTexture("res/tout.png")
		rl.DrawTexture(tout, 500, 370, rl.White)
	}
}

func (g *Game) checkWin() bool {
	return strings.Join(g.guesses, "") == g.word
}

func main() {
	rand.Seed(42)

	game := NewGame()

	rl.InitWindow(1350, 900, "Hangman Game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	image := rl.LoadImage("res/CREEPYSCENE.png")
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)

	for !rl.WindowShouldClose() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyPressed(rl.KeyEnter) {
			game = NewGame()
		}

		
		if rl.IsKeyPressed(rl.KeyA) {
			game.makeGuess("A")
		} else if rl.IsKeyPressed(rl.KeyB) {
			game.makeGuess("B")
		} else if rl.IsKeyPressed(rl.KeyC) {
			game.makeGuess("C")
		} else if rl.IsKeyPressed(rl.KeyD) {
			game.makeGuess("D")
		} else if rl.IsKeyPressed(rl.KeyE) {
			game.makeGuess("E")
		} else if rl.IsKeyPressed(rl.KeyF) {
			game.makeGuess("F")
		} else if rl.IsKeyPressed(rl.KeyG) {
			game.makeGuess("G")
		} else if rl.IsKeyPressed(rl.KeyH) {
			game.makeGuess("H")
		} else if rl.IsKeyPressed(rl.KeyI) {
			game.makeGuess("I")
		} else if rl.IsKeyPressed(rl.KeyJ) {
			game.makeGuess("J")
		} else if rl.IsKeyPressed(rl.KeyK) {
			game.makeGuess("K")
		} else if rl.IsKeyPressed(rl.KeyL) {
			game.makeGuess("L")
		} else if rl.IsKeyPressed(rl.KeyM) {
			game.makeGuess("M")
		} else if rl.IsKeyPressed(rl.KeyN) {
			game.makeGuess("N")
		} else if rl.IsKeyPressed(rl.KeyO) {
			game.makeGuess("O")
		} else if rl.IsKeyPressed(rl.KeyP) {
			game.makeGuess("P")
		} else if rl.IsKeyPressed(rl.KeyQ) {
			game.makeGuess("Q")
		} else if rl.IsKeyPressed(rl.KeyR) {
			game.makeGuess("R")
		} else if rl.IsKeyPressed(rl.KeyS) {
			game.makeGuess("S")
		} else if rl.IsKeyPressed(rl.KeyH) {
			game.makeGuess("H")
		} else if rl.IsKeyPressed(rl.KeyT) {
			game.makeGuess("T")
		} else if rl.IsKeyPressed(rl.KeyU) {
			game.makeGuess("U")
		} else if rl.IsKeyPressed(rl.KeyV) {
			game.makeGuess("V")
		} else if rl.IsKeyPressed(rl.KeyX) {
			game.makeGuess("X")
		} else if rl.IsKeyPressed(rl.KeyY) {
			game.makeGuess("Y")
		} else if rl.IsKeyPressed(rl.KeyZ) {
			game.makeGuess("Z")
		}
		rl.DrawTexture(texture, 0, 0, rl.White)

		rl.BeginDrawing()
		

		// Draw the current state of the game
		drawPendu(game.currentAttempt)

		// Display the word
		rl.DrawText(strings.Join(game.guesses, " "), 10, 800, 30, rl.DarkGray)

		// Display attempts left
		rl.DrawText(fmt.Sprintf("Attempts left: %d", game.maxAttempts-game.currentAttempt), 10, 850, 20, rl.DarkGray)

		if game.checkWin() {
			rl.DrawText("T'as frôlé la mort poto, bien vu", 10, 800, 20, rl.DarkGreen)
		} else if game.currentAttempt >= game.maxAttempts {
			rl.DrawText("T'es mort pendu ^^, appuie sur 'R' si tu veux revivre", 10, 800, 20, rl.Red)
		}

		rl.EndDrawing()
	}
	rl.UnloadTexture(texture)
	rl.CloseWindow()

	fmt.Println("Game Over!")
}
