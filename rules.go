package game

import "fmt"

func DisplayRules() {
	clearScreen()
	fmt.Println("=== Règles du Jeu ===")
	fmt.Println("1. Le but est de deviner un mot lettre par lettre.")
	fmt.Println("2. Vous avez 10 essais pour deviner le mot.")
	fmt.Println("3. Chaque mauvaise lettre réduit le nombre d'essais.")
	fmt.Println("4. Si vous devinez le mot entier et que c'est incorrect, vous perdez 2 essais.")
	fmt.Println("Bonne chance!")
	fmt.Println("Appuyez sur 'Entrée' pour revenir au menu principal...")

	var input string
	fmt.Scanln(&input)
}
