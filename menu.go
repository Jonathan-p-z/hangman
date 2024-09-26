package gogame

import "fmt"

func DisplayMainMenu(selectedOption int) {
    clearScreen()
    fmt.Println("╔══════════════════════════════════════╗")
    fmt.Println("║            🎮 HANGMAN 2024 🎮          ║")
    fmt.Println("╠══════════════════════════════════════╣")

    options := []string{
        "Commencer une nouvelle partie",
        "Règles du jeu",
        "Quitter",
    }

    for i, option := range options {
        if i == selectedOption {
            fmt.Printf("║ ➤ %s %s\n", option, "⬅")
        } else {
            fmt.Printf("║   %s\n", option)
        }
    }

    fmt.Println("╚══════════════════════════════════════╝")
    fmt.Println("Utilisez les flèches pour naviguer, 'q' pour quitter.")
}

// clearScreen efface l'écran (pour un rendu plus propre lors du rafraîchissement)
func clearScreen() {
    fmt.Print("\033[H\033[2J")
}
