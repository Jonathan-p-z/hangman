package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const maxAttempts = 10

func StartGame() {
	word, err := getRandomWord("words/wordlist.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	lettersGuessed := make(map[string]bool)
	errors := 0

	fmt.Println("Appuyez sur 'ESC' à tout moment pour revenir au menu principal.")

	for errors < maxAttempts {
		displayHangman(errors)
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
				errors += 4
			}
		} else if len(input) >= 2 && len(input) <= 15 {
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

func getRandomWord(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))], nil
}

func displayWordState(word string, lettersGuessed map[string]bool) {
	for _, letter := range word {
		if lettersGuessed[string(letter)] {
			fmt.Print(string(letter), " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func isWordGuessed(word string, lettersGuessed map[string]bool) bool {
	for _, letter := range word {
		if !lettersGuessed[string(letter)] {
			return false
		}
	}
	return true
}

func getGuessedLetters(lettersGuessed map[string]bool) string {
	guessedLetters := []string{}
	for letter := range lettersGuessed {
		guessedLetters = append(guessedLetters, letter)
	}
	return strings.Join(guessedLetters, ", ")
}

func displayHangman(errors int) {
	hangmanStages := []string{
		`
          
          |   
          |  
          |  
          |  
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   
          |   
          |  
          |  
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   
          |   
          |  
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |  
          |  
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |   |
          |  
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |  /|
          |   
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |  /|\
          |  
          |
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |  /|\
          |  / 
          |   
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |  /|\
          |  / \
          |  
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          |  /|\/
          |  / \
          |  
        - - - - - - - - - - - - - - - 
        `,
		`
          -----
          |   |
          |   O
          | \/|\/
          |  / \
          |  
        - - - - - - - - - - - - - - - 
        `,
	}

	// Affiche le bonhomme en fonction du nombre d'erreurs
	if errors < len(hangmanStages) {
		fmt.Println(hangmanStages[errors])
	}
}
