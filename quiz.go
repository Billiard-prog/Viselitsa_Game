package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func viselitsa() {
	words := []string{
		"tiger", "london", "doctor", "laptop", "eagle",
		"paris", "teacher", "router", "dolphin", "tokyo",
		"engineer", "camera", "rabbit", "berlin", "lawyer",
		"tablet", "panda", "moscow", "chef", "printer",
	}

	rand.Seed(time.Now().UnixNano())

	hangman_stages := []string{
		`
  _____
  |     |
  |     
  |     
  |     
  |    
__|______
		`,
		`
  _____
  |     |
  |     O
  |     
  |     
  | 
__|______
		`,
		`
  _____
  |     |
  |     O
  |     |
  |     
  |
__|______
		`,
		`
  _____
  |     |
  |     O
  |    /|\
  |     
  |
__|______
		`,
		`
  _____
  |     |
  |     O
  |    /|\
  |    / 
  |
__|______
		`,
		`
  _____
  |     |
  |     O
  |    /|\
  |    / \
  |
__|______
		`,
	}

	fmt.Println("Welcome to my viselitsa!")
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Print("Now enter your real age: ")
	var age int
	fmt.Scan(&age)

	if age >= 12 {
		fmt.Println("You are allowed to play viselitsa! Have a nice game! XD")
		random_word := words[rand.Intn(len(words))]
		guessed_letters := make(map[rune]bool)
		used_letters := make(map[rune]bool)
		attempts := 6

		fmt.Println(hangman_stages[0])

		for attempts > 0 {
			word_progress := ""
			for _, letter := range random_word {
				if guessed_letters[letter] {
					word_progress += string(letter)
				} else {
					word_progress += "_"
				}
				word_progress += " "
			}
			fmt.Printf("Current word: %s\n", word_progress)

			fmt.Printf("%s, your letter! You only need to type 1 letter! (hint: your word is %v letters long): ", name, len(random_word))

			var answer string
			fmt.Scan(&answer)
			answer = strings.ToLower(answer)

			if len(answer) != 1 {
				fmt.Println("You only need to type 1 letter!")
				continue
			}

			first_answer := rune(answer[0])

			if first_answer < 'a' || first_answer > 'z' {
				fmt.Println("Please enter a valid letter!")
				continue
			}

			if used_letters[first_answer] {
				fmt.Println("You already used that letter!")
				continue
			}

			used_letters[first_answer] = true

			if strings.ContainsRune(random_word, first_answer) {
				guessed_letters[first_answer] = true
				fmt.Printf("Correct! Attempts left: %v\n", attempts)
			} else {
				attempts--
				fmt.Printf("Incorrect! Attempts left: %v\n", attempts)
				if attempts > 0 {
					fmt.Println(hangman_stages[6-attempts-1])
				}
			}

			if all_letters_guessed(random_word, guessed_letters) {
				fmt.Printf("Congrats, %s! You won my viselitsa game! You guessed the word: %s\n", name, random_word)
				return
			}

			time.Sleep(500 * time.Millisecond)
		}

		fmt.Println(hangman_stages[5]) // Показываем финальную виселицу
		fmt.Printf("Game over, %s! Thanks for playing my viselitsa game! The word was: %s\n", name, random_word)
	} else {
		fmt.Println("Sorry, you are not allowed to play this game(12+)")
		os.Exit(3)
	}
}

func all_letters_guessed(word string, guessed_letters map[rune]bool) bool {
	for _, letter := range word {
		if !guessed_letters[letter] {
			return false
		}
	}
	return true
}

func main() {
	viselitsa()
}
