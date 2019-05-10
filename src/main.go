package main

import "fmt"
import "./game"

func main() {
	fmt.Println("please input two numbers")

	firstSpecialNumber,secondSpecialNumber := 0,0
	fmt.Scanf("%d %d",&firstSpecialNumber,&secondSpecialNumber)

	for i := 1; i < 101; i++ {
		game.CheckNumber(i, firstSpecialNumber, secondSpecialNumber)
	}
}
