package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func loadHangman(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var designs []string
	var currentDesign string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "`" {
			designs = append(designs, currentDesign)
			currentDesign = ""
		} else {
			currentDesign += line + "\n"
		}
	}

	// Append the last design if there's no trailing "`"
	if currentDesign != "" {
		designs = append(designs, currentDesign)
	}

	return designs, scanner.Err()
}

func displayPendu(errors int, designs []string) {
	if errors < len(designs) {
		fmt.Println(designs[errors])
	} else {
		fmt.Println(designs[len(designs)-1]) // Dernière étape en cas de dépassement
	}
}

func displayWinMessage(word string) {
	fmt.Println("Félicitation ! Vous avez trouvé le mot")
	fmt.Println("Vous avez gagné le jeu ! Bravo pour votre performance")
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

func clearScreenessais() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("\n" + strings.Repeat("\n", 100))
	}
}
