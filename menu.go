package game

import "fmt"

func clearScreen() {

	fmt.Print("\033[H\033[2J")
}

func DisplayMenu() {
	for {
		clearScreen()
		fmt.Println("╔══════════════════════════════════════╗")
		fmt.Println("║            🎮 HANGMAN 2024 🎮        ║")
		fmt.Println("╠══════════════════════════════════════╣")
		fmt.Println("║ 1. Commencer une nouvelle partie     ║")
		fmt.Println("║ 2. Règles du jeu                     ║")
		fmt.Println("║ 3. Quitter                           ║")
		fmt.Println("╚══════════════════════════════════════╝")
		fmt.Print("Sélectionnez une option (1-3) : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			StartGame()
		case 2:
			DisplayRules()
		case 3:
			fmt.Println("Merci d'avoir joué ! À bientôt.")
			return
		default:
			fmt.Println("Choix invalide. Veuillez sélectionner une option valide.")
		}
	}
}
