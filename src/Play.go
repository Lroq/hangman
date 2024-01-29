package src

import (
	"fmt"
	"strings"
	"unicode"
)

// PlayHangman function to play the game
func (g *Game) InitializeGame() {
    trimmedRandomWord := strings.TrimSpace(g.RandomWord)
    g.Underscores = strings.Repeat("_", len(trimmedRandomWord))
    g.Lives = 0
}

// PrintPrompt function to print the prompt
func (g *Game) PrintPrompt() {
    fmt.Print(blue +"Enter a letter: " + reset)
}

// ValideLetter function to check if the letter is valid
func (g *Game) ValideLetter() bool {
    _, err := fmt.Scanf("%s", &g.Letter)
    if err != nil || len(g.Letter) != 1 || !unicode.IsLetter([]rune(g.Letter)[0]) {
        fmt.Println(red + "Please enter a valid letter." + reset)
        return false
    }
    return true
}

// UpdateGameState function to update the game state
func (g *Game) UpdateGameState(newLetter bool) {
    if strings.Contains(g.RandomWord, g.Letter) {
        if strings.Contains(g.UsedLetters, g.Letter) {
            fmt.Println(red + "This letter has already been used. Please choose another one."+ reset)
            return
        }
        Clear()
        fmt.Println(green + "Good Letter!" + reset)
        g.MyWord += g.Letter
        g.Count++ // Increment count when the letter is correct
        g.Score += 100 * (g.ConsecutiveCorrect + 1) // Increment score when the letter is correct
        g.ConsecutiveCorrect++
        g.UsedLetters += g.Letter

        occurrences := strings.Count(g.RandomWord, g.Letter)
        if occurrences >= 2 {
            g.Count += occurrences - 1
        }

        g.Underscores = ""
        g.Underscores = ""
        trimmedRandomWord := strings.TrimSpace(g.RandomWord)
        for j := 0; j < len(trimmedRandomWord); j++ {
            if strings.Contains(g.MyWord, string(trimmedRandomWord[j])) {
                g.Underscores += string(trimmedRandomWord[j])
            } else {
                g.Underscores += "_"
            }
        }
    } else {
        Clear()
        fmt.Println(red + "Bad Letter!" + reset)
        g.Lives++
        g.FalseLetters += g.Letter + ", " // add incorrect letter to falseLetters variable
        g.ConsecutiveCorrect = 0
        DisplayHangmanState(10 - g.Lives) // Display the hangman when the player makes a mistake
    }
}

// PrintStatus function to print the status of the game
func (g *Game) PrintStatus() {
    fmt.Println(red + "Incorrect letters:" + reset, g.FalseLetters)
    fmt.Println(green + "Correct letters:" + reset, g.MyWord)
    fmt.Println(blue + "Remaining lives:" + reset, 10-g.Lives)
    fmt.Println(blue + "Current SCORE:" + reset, g.Score)
    fmt.Println(blue + "Word to guess:" + reset, g.Underscores)
    fmt.Print("\n")
}

// CheckWinLoss function to check if the player won or lost
func (g *Game) CheckWinLoss() bool {
    if g.Lives >= 10 {
        fmt.Println(red + "Sorry, you lost. The word was:" + reset, g.RandomWord)
        return true
    }

    uniqueLetters := make(map[rune]bool)
    for _, letter := range g.RandomWord {
        if unicode.IsLetter(letter) {
            uniqueLetters[letter] = true
        }
    }

    for _, letter := range g.MyWord {
        if _, ok := uniqueLetters[letter]; ok {
            delete(uniqueLetters, letter)
        }
    }

    if len(uniqueLetters) == 0 {
        fmt.Println(green + "Victory! The word was:" + reset, g.RandomWord)
        fmt.Println(blue + "Your final score is:" + reset, g.Score)
        return true
    }

    return false
}

// PlayHangman function to play the game
func (g *Game) PlayHangman() {
    g.InitializeGame()
    newLetter := true

    for g.Lives < 10 && g.Count != g.RandomNum {
        if newLetter {
            g.PrintPrompt()
        }
        if !g.ValideLetter() {
            continue
        }
        newLetter = true
        g.UpdateGameState(newLetter)
        g.PrintStatus()

        if g.CheckWinLoss() {
            break
        }
    }
}
