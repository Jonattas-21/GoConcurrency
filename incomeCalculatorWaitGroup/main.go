package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var waitgroup sync.WaitGroup

func main() {

	//define variavéis
	var bankBalance int = 0
	var balanceMutex sync.Mutex

	//define fonte de receita
	incomes := []Income{
		{Source: "Salário", Amount: 500},
		{Source: "Mesada da mãe", Amount: 100},
		{Source: "Freelancer", Amount: 20},
	}

	//mostrar saldo inicial
	fmt.Printf("Valor inicial da conta bancária: %d", bankBalance)
	fmt.Println()

	waitgroup.Add(len(incomes))

	//Para cada fonte de receita calcula o recebimento nas 52 semanas do ano
	for i, income := range incomes {

		go func(i int, income Income) {
			defer waitgroup.Done()

			for week := 1; week <= 52; week++ {
				balanceMutex.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balanceMutex.Unlock()
			}
			fmt.Printf("Recebimento parcial do ano de %d.00: %s", bankBalance, income.Source)
			fmt.Println()

		}(i, income)
	}

	waitgroup.Wait()

	//mostrar recebimentos do ano
	fmt.Printf("Valor final da conta bancária: %d.00", bankBalance)

}
