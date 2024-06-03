package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	prompt  string
	options []string
	answer  int // Index of the correct answer in the options slice
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama Anda: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	questions := getQuestions()
	score := 0
	fmt.Println("\nSelamat datang di kuis,", name+"!")
	for _, question := range questions {
		fmt.Println("\n" + question.prompt)
		for i, option := range question.options {
			fmt.Printf("%d) %s\n", i+1, option)
		}
		fmt.Print("Jawaban Anda (masukkan nomor): ")
		var choice int
		fmt.Scanln(&choice)
		if choice-1 == question.answer {
			score++
			fmt.Println("Benar!")
		} else {
			fmt.Println("Salah. Jawaban yang benar adalah:", question.options[question.answer])
		}
	}
	if score > len(questions)/2 {
		fmt.Printf("\nSelamat, %s! Anda berhasil menjawab %d dari %d pertanyaan dengan benar.\n", name, score, len(questions))
	} else {
		fmt.Printf("\nMaaf, %s! Anda gagal menjawab %d dari %d pertanyaan dengan benar.\n", name, score, len(questions))
	}
}

func getQuestions() []Question {
	questions := []Question{
		{
			prompt: "Apa ibu kota Perancis?",
			options: []string{
				"Berlin",
				"Paris",
				"Madrid",
				"Roma",
			},
			answer: 1, // Paris
		},
		{
			prompt: "Apa gunung tertinggi di dunia?",
			options: []string{
				"K2",
				"Kangchenjunga",
				"Gunung Everest",
				"Lhotse",
			},
			answer: 2, // Mount Everest
		},
		{
			prompt: "Negara terkecil di dunia adalah?",
			options: []string{
				"Monako",
				"Nauru",
				"Tuvalu",
				"Vatikan",
			},
			answer: 3, // Vatican City
		},
	}
	return questions
}
