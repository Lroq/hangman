package src

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)
var (
    blue = "\033[34m"
    reset = "\033[0m"
    green = "\033[32m"
    red = "\033[31m"
	yellow = "\033[33m"
    orange = "\033[33m"
	total_green = "\033[42m"
	total_yellow = "\033[43m"
	total_white = "\033[47m"
	total_blue = "\033[44m"
)

// Range structure to define the range of lines to read
type Range struct {
	Start int
	End   int
}

// Game structure to define the game
type Game struct {
	RandomWord         string
	RandomNum          int
	MyWord             string
	Underscores        string
	FalseLetters       string
	Letter             string
	Count              int
	Lives              int
	UsedLetters        string
	Score              int
	ConsecutiveCorrect int
}

// ChooseWord function to choose a word in the list of words
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ChooseWord function to choose a word in the list of words
func ReadFile() {
    rand.Seed(time.Now().UnixNano())

    language := ChooseLanguage()
    var content1, content2, content3 []byte
    var err error

    if language == "english" {
        content1, err = ioutil.ReadFile("../Ressources/english_words1.txt")
        content2, err = ioutil.ReadFile("../Ressources/english_words2.txt")
        content3, err = ioutil.ReadFile("../Ressources/english_words3.txt")
    } else if language == "french"{
        content1, err = ioutil.ReadFile("../Ressources/french_words1.txt")
        content2, err = ioutil.ReadFile("../Ressources/french_words2.txt")
        content3, err = ioutil.ReadFile("../Ressources/french_words3.txt")
    } else {
		content1, err = ioutil.ReadFile("../Ressources/spanish_words1.txt")
		content2, err = ioutil.ReadFile("../Ressources/spanish_words2.txt")
		content3, err = ioutil.ReadFile("../Ressources/spanish_words3.txt")
	}
    CheckError(err)

    words := strings.Split(string(content1), "\n")
    words = append(words, strings.Split(string(content2), "\n")...)
    words = append(words, strings.Split(string(content3), "\n")...)

    minLength, maxLength := Difficulty()
    Clear()
    randomWord := ChooseWord(words, minLength, maxLength)

    randomNum := len(randomWord) - 1
	fmt.Println(total_green+ "Your goal is to find the selected word from several basic word lists."+ reset)
	fmt.Println(total_green+ "You can only enter one letter at a time, but if you enter an entire word,"+ reset)
	fmt.Println(total_green+ "It is important to put a space between each letter for your word to be accepted."+ reset)
    fmt.Printf(blue +"The word to be guessed contains %d letters.\n" + reset, randomNum)
    fmt.Println(blue + "You have 10 lives to guess the word."+ reset)
	fmt.Println( green+ "Good luck!"+ reset)
	

   // fmt.Println(randomWord) // Consider removing this line for an actual game, as it reveals the word to be guessed.

    game := Game{RandomWord: randomWord, RandomNum: randomNum}
    game.PlayHangman()
    SaveGame(&game)
}

// ChooseWord function to choose a word in the list of words
func ReadLinesInRange(filename string, lineRange Range) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if lineCount >= lineRange.Start && lineCount <= lineRange.End {
			lines = append(lines, scanner.Text())
		}
		if lineCount >= lineRange.End {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// DisplayHangmanState function to display the hangman
func DisplayHangmanState(livesLeft int) {
	filename := "../Ressources/hangman.txt"
	ranges := []Range{
		{1, 2},
		{4, 9},
		{11, 17},
		{19, 25},
		{27, 34},
		{35, 41},
		{43, 49},
		{51, 57},
		{59, 65},
		{67, 73},
	}
	lineRange := ranges[9-livesLeft]
	lines, err := ReadLinesInRange(filename, lineRange)

	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	red := "\033[31m"
    reset := "\033[0m"

	for _, line := range lines {
		fmt.Println(red + line + reset)
	}
}
