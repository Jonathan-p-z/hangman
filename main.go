package main

import (
    "fmt"
    "main/jeu-du-pendu/game" 
    "github.com/eiannone/keyboard"
)

func main() {
    err := keyboard.Open()
    if err != nil {
        panic(err)
    }
    defer keyboard.Close()

    selectedOption := 0
    for {
        gogame.DisplayMainMenu(selectedOption)

        char, key, err := keyboard.GetKey()
        if err != nil {
            panic(err)
        }

        if key == keyboard.KeyArrowDown {
            selectedOption = (selectedOption + 1) % 3
        } else if key == keyboard.KeyArrowUp {
            selectedOption = (selectedOption - 1 + 3) % 3
        } else if key == keyboard.KeyEnter {
            switch selectedOption {
            case 0:
                gogame.StartGame()
            case 1:
                gogame.DisplayRules()
            case 2:
                fmt.Println("Merci d'avoir joué ! À bientôt !")
                return
            }
        }

        if char == 'q' || char == 'Q' {
            fmt.Println("Merci d'avoir joué ! À bientôt !")
            return
        }
    }
}
