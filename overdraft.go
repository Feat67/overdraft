package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var wallet uint32
	var balance, summa int32
	var operation string
	var overdraft = rand.Int31()
	fmt.Println("Введите номер счета:")
	fmt.Scan(&wallet)
	fmt.Println(overdraft, "ваш кредитный лимит!")
	fmt.Println(balance, "баланс вашего счёта")
	fmt.Println("Желайте пополнить счёт? Y/N")
	fmt.Scan(&operation)
	for overdraft != 0 {
		if operation == "Y" {
			fmt.Println("Введите желаймую сумму кредитования: доступно", overdraft)
			fmt.Scan(&summa)
			if summa < overdraft {
				balance += summa
				overdraft -= summa
				fmt.Println("Ваш баланс:", balance)
			} else if summa >= overdraft {
				balance += overdraft
				fmt.Println("Кредитный лимит исчерпан, баланс вашего счёта:", balance)
				overdraft = 0
			}
		} else {
			fmt.Println("Ваш баланс:", balance)
		}

	}
}
