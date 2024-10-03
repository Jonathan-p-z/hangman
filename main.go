package main

import (
	"gogame/game"
	"fmt"
	"bufio"
	"os"
)

func main() {
	game.DisplayMenu()
}

func displayWordWithFirstLetterRevealed(word string, lettersGuessed map[string]bool) {
	firstLetter := string(word[0])
	lettersGuessed[firstLetter] = true
	for _, letter := range word {
		letterStr := string(letter)
		if lettersGuessed[letterStr] {
			fmt.Printf("%s ", letterStr)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func reader() {
	file, err := os.Open("Hangman\\test\\mots_diffciles.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("Contenu du fichier:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
}
