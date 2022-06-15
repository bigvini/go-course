package main

import (
	"fmt"
)

func main() {

	const (
		oneApple = 5.99
		onePear  = 7.00
		myMoney  = 23.00
	)

	fmt.Println("Одне яблуко коштує ", oneApple, "грн.")
	fmt.Println("Одна груша коштує ", onePear, "грн.")
	fmt.Println("Ми маємо ", myMoney, "грн.")

	fmt.Println(" Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	nineApple := 9 * oneApple
	sevenPear := 9 * onePear
	fmt.Println("Щоб купити 9 яблук потрібно: ", nineApple, "грн")
	fmt.Println("Щоб купити 7 груш потрібно: ", sevenPear, "грн")

	fmt.Println(" Скільки груш ми можемо купити?")
	totalBuyPears := myMoney / onePear
	fmt.Println("Ми можемо купити", int(totalBuyPears), "груші.")

	fmt.Println("Скільки яблук ми можемо купити?")
	totalBuyAppels := myMoney / oneApple
	fmt.Println("Ми можемо купити", int(totalBuyAppels), "яблука.")

	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	twoAppelAndTwoPear := 2*oneApple + 2*onePear

	fmt.Println("Ми неможем купити 2 груші і 2 яблука. Тому що це буде ", twoAppelAndTwoPear, "грн.")

}
