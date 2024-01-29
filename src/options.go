package src

import (
	"fmt"
	"math/rand"
	"os/exec"
	"os"
)

// ChooseWord function to choose a word in the list of words
func ChooseWord(words []string, minLength, maxLength int) string {
    suitableWords := []string{}
    for _, word := range words {
        if len(word) >= minLength && len(word) <= maxLength {
            suitableWords = append(suitableWords, word)
        }
    }
    if len(suitableWords) == 0 {
        return ""
    }
    return suitableWords[rand.Intn(len(suitableWords))] // choisie un mot au hasard dans la liste des mots dU TABLEAU suitableWords
}


// Difficulty function to choose the difficulty of the game
func Difficulty() (int, int) {
    var difficulty string
    for {
        fmt.Println(blue + "Please choose a difficulty" + reset, blue + "(" + reset, green + "easy," + reset, orange + "medium," + reset, red + "hard" + reset, blue + "):" + reset)
        fmt.Println(red + "hard: 9-12 letters" + reset)
        fmt.Println(orange + "medium: 6-8 letters" + reset)
        fmt.Println(green + "easy: 3-5 letters" + reset)
        fmt.Print(blue + "Your choice: "+ reset)
        fmt.Scanln(&difficulty)

        switch difficulty {
        case "easy":
            return 3, 5
        case "medium":
            return 6, 8
        case "hard":
            return 9, 12
        default:
            Clear()
            fmt.Println(total_yellow + "Please choose a valid difficulty." + reset)
        }
    }
}

// DisplayHangmanState function to display the hangman
func SaveGame(g *Game) error {
    file, err := os.Create("saved_hangman.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    fmt.Fprintf(file, "word: %s\nfalseletters: %s\nusedletters: %d\n%s lives: %d\nscore: %d\n",
        g.MyWord, g.FalseLetters, g.Count, g.UsedLetters, g.Lives, g.Score)

    return nil
}

// DisplayHangmanState function to display the hangman
func Clear() { // faut juste appeller la fonction
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// Replay function to ask the player if he wants to play again
func Replay() bool {
    var replay string
        fmt.Println(blue +"Do you want to play again ? (yes/no)"+ reset)
        for {
        fmt.Scanln(&replay)
        switch replay {
        case "yes":
            return true
        case "no":
            Clear()
            fmt.Println(total_green +"Thank you for playing !"+ reset)
            fmt.Println(total_green +"created by Hebrard Dylan and Louis Roques"+ reset)
            return false
        default:
            fmt.Println(total_yellow +"Please choose a valid answer." +reset)
        }
    }
}

// ChooseLanguage function to let the player choose the language of the words. It prompts the player to choose a language and returns the chosen language as a string.
func ChooseLanguage() string {
Clear()
    var choice int
    fmt.Println(blue + "Welcome to Hangman's Game !" + reset)
    fmt.Println(blue + "Choose the language for the words:" + reset)
    fmt.Println(total_green +"1: English " + reset)
    fmt.Println(total_white +"2: French  " + reset)
    fmt.Println(total_yellow +"3: Spanish " + reset)
    fmt.Print(blue +"Enter your choice ("+reset,green +"1" +reset,blue +" or " + reset, "2",blue +" or " + reset,yellow +"3"+ reset,blue +"): " + reset)
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        Clear()
        fmt.Println(total_yellow +"You have chosen English." + reset)
        return "english"
    case 2:
        Clear()
        fmt.Println(total_yellow +"You have chosen French." + reset)
        return "french"
    case 3:
        Clear()
        fmt.Println(total_yellow +"You have chosen Spanish." + reset)
        return "spanish"
    default:
        Clear()
        fmt.Println(total_yellow +"Invalid choice. Defaulting to English." + reset)
        return "english"
    }
}
