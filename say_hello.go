package module_go

import "fmt"

func SayHello(name string) string {
	return name
}

func KalkulatorTambah(number1 int, number2 int) int {
	return number1 + number2
}

func KalkulatorKurang(number1 int, number2 int) int {
	return number1 - number2
}

func KalkulatorBagi(number1 int, number2 int) int {
	return number1 / number2
}

func loopS(loopstring [...]string) {
	for i := 0; i < len(loopstring); i++ {
		fmt.Println(loopstring[i])
	}
}

func loopI(loopint [...]int) {
	for i := 0; i < len(loopint); i++ {
		fmt.Println(loopint[i])
	}
}
