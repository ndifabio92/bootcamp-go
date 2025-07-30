package main

import "fmt"

func main() {
	fmt.Println("Enter a credit card number:")
	cardNumber := ""
	fmt.Scanln(&cardNumber)
	if luhnCheck(cardNumber) {
		fmt.Println("The credit card number is valid.")
	} else {
		fmt.Println("The credit card number is invalid.")
	}
}

func luhnCheck(cardNumber string) bool {
	sum := 0

	//Recorrer digitos de derecha a izquierda
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		//Duplicar cada segundo digito
		if (len(cardNumber)-i)%2 == 0 {
			digit *= 2

			//SI el resultado es mayor a 9, sumar sus digitos
			if digit > 9 {
				digit = digit/10 + digit%10
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
