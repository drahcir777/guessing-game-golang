package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação!")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Printf("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInteiro, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O valor informado não é um número inteiro válido.")
			return
		}

		switch {
			case chuteInteiro < x:
				fmt.Println("Seu chute foi menor que o número sorteado.", chuteInteiro)
			case chuteInteiro > x:
				fmt.Println("Seu chute foi maior que o número sorteado.", chuteInteiro)
			case chuteInteiro == x:
				fmt.Println("Parabéns! Você acertou o número sorteado.")
				return
		}

		chutes[i] = chuteInteiro
	}

	fmt.Printf("Você não acertou o número sorteado. O número sorteado foi %d. Seus chutes foram: %v\n", x, chutes)
}
