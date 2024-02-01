package main

import "github.com/Lroq/hangman/src"

func main() {
    for {
        src.Clear()
        src.ReadFile()
        if !src.Replay() {
            break
        }
    }
}