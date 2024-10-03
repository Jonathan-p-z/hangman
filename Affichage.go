package game

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const maxAttempts = 10

func StartGame() {
	hangmanDesigns, err := loadHangman("pendu/Animation.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement du design du pendu :", err)
		return
	}

	word, err := getRandomWord("words/wordlist.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	lettersGuessed := make(map[string]bool)
	errors := 0

	fmt.Println("Appuyez sur 'ESC' à tout moment pour revenir au menu principal.")

	for errors < maxAttempts {
		clearScreen() // Nettoyer l'écran avant chaque nouvel affichage

		displayPendu(errors, hangmanDesigns)

		displayWordState(word, lettersGuessed)
		fmt.Println("Lettres déjà proposées :", getGuessedLetters(lettersGuessed))
		fmt.Printf("Nombre d'erreurs : %d / %d\n", errors, maxAttempts)

		fmt.Print("Proposez une lettre ou un mot : ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "\033" {
			fmt.Println("Retour au menu principal.")
			return
		}

		if len(input) == 1 {
			if _, exists := lettersGuessed[input]; exists {
				fmt.Println("Vous avez déjà proposé cette lettre.")
				continue
			}

			lettersGuessed[input] = true
			if strings.Contains(word, input) {
				fmt.Println("Bonne lettre !")
			} else {
				fmt.Println("Mauvaise lettre.")
				errors++
			}
		} else if len(input) == len(word) {
			if input == word {
				fmt.Println("Félicitations, vous avez trouvé le mot :", word)
				return
			} else {
				fmt.Println("Mauvais mot.")
				errors += 2
			}
		} else if len(input) > 1 {
			fmt.Printf("Vous avez entré %d lettres, cela vous coûte %d essais.\n", len(input), len(input))
			errors += len(input)
		} else {
			fmt.Println("Entrée non valide.")
		}

		if isWordGuessed(word, lettersGuessed) {
			fmt.Println("Félicitations, vous avez deviné le mot :", word)
			return
		}
	}

	fmt.Println("Vous avez perdu. Le mot était :", word)
}
