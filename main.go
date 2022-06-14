package main

import "fmt"

func main() {

	var ApplePrice float64 = 5.99
	var PearPrice int = 7
	var OurBudget int = 23

	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? ", ApplePrice * float64(9),  PearPrice * 8)

	fmt.Println("2. Скільки груш ми можемо купити?", OurBudget / PearPrice)
	
	fmt.Println("3. Скільки яблук ми можемо купити?", float64(OurBudget) / 	ApplePrice)

	if float64(OurBudget) > ApplePrice * float64(2) + float64(PearPrice) * float64(2) {

		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", "Да, можем")
	} else {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", "Нет,не можем")
	}

}

