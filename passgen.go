// passgen v1.0.0
// Copyright 2021 arataca89@gmail.com
// Licensed under the MIT License
// 20210429

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
)

func help() {
	fmt.Println("USO:")
	fmt.Printf("passgen\n\tExibe o help.\n\n")

	fmt.Printf("passgen char\n\tExibe o conjunto de caractres ASCII")
	fmt.Printf(" utilizados para gerar a senha.\n\n")

	fmt.Printf("passgen c nr\n\tCria uma senha de nr caracteres.\n")
	fmt.Printf("\tCaso nr não seja definido cria uma senha de 16 caracteres.\n")

	fmt.Printf("\nA senha é criada e copiada para o clipboard.\n")
	fmt.Printf("Use CTRL + V para colar a senha onde você quiser.\n")
}

func randInt(min, max int) int {
	if min < 0 || max <= 0 || min > max {
		return 0
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func printChars(start rune, end rune) {
	for i := start; i <= end; i++ {
		fmt.Print(string(i), " ")
	}
	fmt.Println()
}

func usedChars() {
	fmt.Printf("Caracteres usados:\n")
	printChars(33, 47)
	printChars(48, 57)
	printChars(58, 64)
	printChars(65, 90)
	printChars(91, 96)
	printChars(97, 122)
	printChars(123, 126)
}

func createPassword() {
	var pass string
	if len(os.Args) < 3 {
		for i := 1; i <= 16; i++ {
			pass = pass + string(randInt(33, 126))
		}
		fmt.Println(pass)
		clipboard.WriteAll(pass)
	} else {
		n, _ := strconv.Atoi(os.Args[2])
		for i := 1; i <= n; i++ {
			pass = pass + string(randInt(33, 126))
		}
		fmt.Println(pass)
		clipboard.WriteAll(pass)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 || len(os.Args) > 3 {
		help()
	} else if os.Args[1] == "char" {
		usedChars()
	} else if os.Args[1] == "c" {
		createPassword()
	} else {
		help()
	}

}
