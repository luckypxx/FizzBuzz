package game

import (
	"fmt"
	"strconv"
	"strings"
)

const FIZZ = "Fizz"
const BUZZ = "Buzz"

var print string

func IsSpecialNumber(number int, specialNumber int, str string) bool {
	print = str
	return strings.Contains(strconv.Itoa(number), strconv.Itoa(specialNumber)) || (number%specialNumber) == 0
}

func IsFizzBuzz(number int, firstSprcialNumber int, secondSpecialNumber int) bool {
	if (number%firstSprcialNumber) == 0 && (number%secondSpecialNumber) == 0 {
		print = FIZZ + BUZZ
		return true
	}
	return false
}

func IsFizzOrBuzz(number int, firstSpecialNumber int, secondSpecialNumber int) bool {

	return IsSpecialNumber(number, firstSpecialNumber, FIZZ) || IsSpecialNumber(number, secondSpecialNumber, BUZZ)
}

func CheckNumber(number int, firstSpecialNumber int, secondSpecialNumber int) {
	if IsFizzBuzz(number, firstSpecialNumber, secondSpecialNumber) || IsFizzOrBuzz(number, firstSpecialNumber, secondSpecialNumber) {
		fmt.Println(print)
		return
	}
	fmt.Println(number)
}
