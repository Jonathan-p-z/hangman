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

// StartGame démarre le jeu
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

		// Vérifie si l'utilisateur a appuyé sur 'ESC'
		if input == "\033" { // ASCII pour ESC
			fmt.Println("Retour au menu principal.")
			return
		}

		if len(input) == 1 { // Proposition d'une lettre
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
		} else if len(input) == len(word) { // Proposition d'un mot
			if input == word {
				fmt.Println("Félicitations, vous avez trouvé le mot :", word)
				return
			} else {
				fmt.Println("Mauvais mot.")
				errors += 2 // Pénalité pour mauvaise proposition
			}
		} else {
			fmt.Println("Entrée non valide.")
			continue // Ne pas afficher le pendu si l'entrée est invalide
		}

		if isWordGuessed(word, lettersGuessed) {
			fmt.Println("Félicitations, vous avez deviné le mot :", word)
			return
		}
	}

	fmt.Println("Vous avez perdu. Le mot était :", word)
}

// getRandomWord charge un mot aléatoire depuis le fichier
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

// displayWordState affiche l'état du mot avec les lettres devinées
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

// isWordGuessed vérifie si le mot a été deviné
func isWordGuessed(word string, lettersGuessed map[string]bool) bool {
	for _, letter := range word {
		if !lettersGuessed[string(letter)] {
			return false
		}
	}
	return true
}

// getGuessedLetters retourne les lettres déjà devinées
func getGuessedLetters(lettersGuessed map[string]bool) string {
	guessedLetters := []string{}
	for letter := range lettersGuessed {
		guessedLetters = append(guessedLetters, letter)
	}
	return strings.Join(guessedLetters, ", ")
}

// displayHangman dessine le bonhomme du pendu en fonction des erreurs
// displayHangman dessine le bonhomme du pendu en fonction des erreurs
func displayHangman(errors int) {
	hangmanStages := []string{
		`
          
          |   
          |  
          |  
          |  
          |
        - - - - - - - - - - - - - - - 
        `, // 0 erreurs
		`
          -----
          |   
          |   
          |  
          |  
          |
        - - - - - - - - - - - - - - - 
        `, // 1 erreur
		`
          -----
          |   |
          |   
          |   
          |  
          |
        - - - - - - - - - - - - - - - 
        `, // 2 erreurs
		`
          -----
          |   |
          |   O
          |  
          |  
          |
        - - - - - - - - - - - - - - - 
        `, // 3 erreurs
		`
          -----
          |   |
          |   O
          |   |
          |  
          |
        - - - - - - - - - - - - - - - 
        `, // 4 erreurs
		`
          -----
          |   |
          |   O
          |  /|
          |   
          |
        - - - - - - - - - - - - - - - 
        `, // 5 erreurs
		`
          -----
          |   |
          |   O
          |  /|\
          |  
          |
        - - - - - - - - - - - - - - - 
        `, // 6 erreurs
		`
          -----
          |   |
          |   O
          |  /|\
          |  / 
          |   
        - - - - - - - - - - - - - - - 
        `, // 7 erreurs
		`
          -----
          |   |
          |   O
          |  /|\
          |  / \
          |  
        - - - - - - - - - - - - - - - 
        `, // 8 erreurs
		`
          -----
          |   |
          |   O
          |  /|\/
          |  / \
          |  
        - - - - - - - - - - - - - - - 
        `, // 9 erreurs
		`
          -----
          |   |
          |   O
          | \/|\/
          |  / \
          |  
        - - - - - - - - - - - - - - - 
        `, // 10 erreurs
	}

	// Affiche le bonhomme en fonction du nombre d'erreurs
	if errors < len(hangmanStages) {
		fmt.Println(hangmanStages[errors])
	}
}
