package gogame

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

var maxErrors = 10

// StartGame lance une nouvelle partie
func StartGame() {
    word, err := getRandomWord("words/wordlist.txt")
    if err != nil {
        fmt.Println("Erreur lors du chargement du fichier de mots :", err)
        return
    }

    lettersGuessed := make(map[string]bool)
    errors := 0

    for errors < maxErrors {
        displayWordState(word, lettersGuessed)
        fmt.Println("Lettres déjà proposées :", getGuessedLetters(lettersGuessed))
        fmt.Println("Nombre d'erreurs :", errors, "/", maxErrors)

        fmt.Print("Proposez une lettre ou un mot : ")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))

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
        } else {
            fmt.Println("Entrée non valide.")
        }

        if isWordGuessed(word, lettersGuessed) {
            fmt.Println("Félicitations, vous avez deviné le mot :", word)
            return
        }

        if errors >= maxErrors {
            fmt.Println("Vous avez perdu. Le mot était :", word)
        }
    }
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

    rand.Seed(time.Now().Unix())
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
