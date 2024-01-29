package main

import "cmd/src"

func main() {
    for {
        src.Clear()
        src.ReadFile()
        if !src.Replay() {
            break
        }
    }
}