package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"os"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var wordsToGuess = []string{
	"golang",
	"programming",
	"computer",
	"developer",
	"engineer",
}

var maxAttempts = 10
var customImagePath = "custom_images/"
var backgroundImagePath = "res/CREEPYSCENE.png"

func main() {
	rand.Seed(time.Now().UnixNano())

	attempts := 0
	usedLetters := []string{}
	wordToGuess := selectRandomWord()
	guess := ""
	gameOver := false
	showImage := false

	rl.InitWindow(1350, 900, "Hangman Game")
	rl.SetTargetFPS(60)

	backgroundImage := rl.LoadTexture(backgroundImagePath)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(backgroundImage, 0, 0, rl.RayWhite)

		if !isWordGuessed(wordToGuess, usedLetters) && attempts < maxAttempts {
			displayWord(wordToGuess, usedLetters)
			rl.DrawText("Devinez une lettre:", 10, 10, 20, rl.DarkGray)
			rl.DrawText("Essais restants: "+fmt.Sprint(maxAttempts-attempts), 10, 40, 20, rl.DarkGray)

			guess = getGuess()

			if guess != "" && !strings.Contains(usedLettersJoined(usedLetters), guess) {
				usedLetters = append(usedLetters, guess)

				if !strings.Contains(wordToGuess, guess) {
					attempts++
					showImage = true
				}
			}
		} else {
			gameOver = true
		}

		if isWordGuessed(wordToGuess, usedLetters) {
			rl.DrawText("T a de la chance poto tu as survécu cette fois", 10, 80, 30, rl.DarkGray)
		} else if attempts >= maxAttempts {
			rl.DrawText("Tu es mort poto on se retrouve de l autre cotés.", 10, 80, 30, rl.DarkGray)
		}

		if showImage {
			imagePath := customImagePath + fmt.Sprintf("custom_image_%d.png", attempts)

			if fileExists(imagePath) {
				image := rl.LoadTexture(imagePath)
				rl.DrawTexture(image, 400, 300, rl.RayWhite)
			}
		}

		if gameOver {
			rl.DrawText("Le mot était: "+wordToGuess, 10, 120, 20, rl.Gray)
			rl.DrawText("Appuyez sur Entrée pour rejouer", 10, 160, 20, rl.Gray)

			if rl.IsKeyPressed(rl.KeyEnter) {
				wordToGuess = selectRandomWord()
				attempts = 0
				usedLetters = []string{}
				gameOver = false
				showImage = false
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func displayWord(wordToGuess string, usedLetters []string) {
	display := ""
	for _, letter := range wordToGuess {
		if strings.Contains(usedLettersJoined(usedLetters), string(letter)) {
			display += string(letter)
		} else {
			display += "_"
		}
	}
	rl.DrawText(display, 10, 80, 40, rl.DarkGray)
}

func usedLettersJoined(usedLetters []string) string {
	return strings.Join(usedLetters, "")
}

func isWordGuessed(wordToGuess string, usedLetters []string) bool {
	for _, letter := range wordToGuess {
		if !strings.Contains(usedLettersJoined(usedLetters), string(letter)) {
			return false
		}
	}
	return true
}

func selectRandomWord() string {
	return wordsToGuess[rand.Intn(len(wordsToGuess))]
}

func getGuess() string {
	guess := ""
	// Modify this function to handle letter input using the Enter key (KeyEnter).
	// You can use a similar logic as in the original code for other letter keys.
	if rl.IsKeyDown(rl.KeyA) {
		guess = "a"
	} else if rl.IsKeyDown(rl.KeyB) {
		guess = "b"
	} else if rl.IsKeyDown(rl.KeyC) {
		guess = "c"
	} else if rl.IsKeyDown(rl.KeyD) {
		guess = "d"
	} else if rl.IsKeyDown(rl.KeyE) {
		guess = "e"
	} else if rl.IsKeyDown(rl.KeyF) {
		guess = "f"
	} else if rl.IsKeyDown(rl.KeyG) {
		guess = "g"
	} else if rl.IsKeyDown(rl.KeyH) {
		guess = "h"
	} else if rl.IsKeyDown(rl.KeyI) {
		guess = "i"
	} else if rl.IsKeyDown(rl.KeyJ) {
		guess = "j"
	} else if rl.IsKeyDown(rl.KeyK) {
		guess = "k"
	} else if rl.IsKeyDown(rl.KeyL) {
		guess = "l"
	} else if rl.IsKeyDown(rl.KeyM) {
		guess = "m"
	} else if rl.IsKeyDown(rl.KeyN) {
		guess = "n"
	} else if rl.IsKeyDown(rl.KeyO) {
		guess = "o"
	} else if rl.IsKeyDown(rl.KeyP) {
		guess = "p"
	} else if rl.IsKeyDown(rl.KeyQ) {
		guess = "q"
	} else if rl.IsKeyDown(rl.KeyR) {
		guess = "r"
	} else if rl.IsKeyDown(rl.KeyS) {
		guess = "s"
	} else if rl.IsKeyDown(rl.KeyT) {
		guess = "t"
	} else if rl.IsKeyDown(rl.KeyU) {
		guess = "u"
	} else if rl.IsKeyDown(rl.KeyV) {
		guess = "v"
	} else if rl.IsKeyDown(rl.KeyW) {
		guess = "w"
	} else if rl.IsKeyDown(rl.KeyX) {
		guess = "x"
	} else if rl.IsKeyDown(rl.KeyY) {
		guess = "y"
	} else if rl.IsKeyDown(rl.KeyZ) {
		guess = "z"
	}
	return guess
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
