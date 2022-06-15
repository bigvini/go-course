package main

import (
	"fmt"
)

func main() {

	const (
		oneApplePrice = 5.99
		onePearPrice  = 7.00
		myMoney       = 23.00
	)

	fmt.Println("Одне яблуко коштує ", oneApplePrice, "грн.")
	fmt.Println("Одна груша коштує ", onePearPrice, "грн.")
	fmt.Println("Ми маємо ", myMoney, "грн.")

	fmt.Println(" Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	nineApple := 9 * oneApplePrice
	sevenPear := 9 * onePearPrice
	fmt.Println("Щоб купити 9 яблук потрібно: ", nineApple, "грн")
	fmt.Println("Щоб купити 7 груш потрібно: ", sevenPear, "грн")

	fmt.Println(" Скільки груш ми можемо купити?")
	totalBuyPears := myMoney / onePearPrice
	fmt.Println("Ми можемо купити", int(totalBuyPears), "груші.")

	fmt.Println("Скільки яблук ми можемо купити?")
	totalBuyAppels := myMoney / oneApplePrice
	fmt.Println("Ми можемо купити", int(totalBuyAppels), "яблука.")

	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	twoAppelAndTwoPear := 2*oneApplePrice + 2*onePearPrice

	fmt.Println("Ми неможем купити 2 груші і 2 яблука. Тому що це буде ", twoAppelAndTwoPear, "грн.")

}
