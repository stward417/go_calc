package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var a, b *int

var oprs = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

func calc(text string) {
	var opr string
	for i := range oprs {
		for _, val := range text {
			if i == string(val) {
				opr += i
				if len(opr) > 1 || len(opr) < 1 {
					panic("Error. Input format")
				} else {
					data = strings.Split(text, opr)
				}

			}
		}
	}

	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)

	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic("Error. Number system ")
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := oprs[opr]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic("Error. Number range")
		}
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic("Error. Number range")
			}
		}
		if val, ok := oprs[opr]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}

func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic("Error. Zero")
	} else if romanResult < 0 {
		panic("Error. Unsigned")
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}

func main() {
	for {
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		text = strings.ReplaceAll(text, " ", "")
		text = strings.ToUpper(strings.TrimSpace(text))
		calc(text)
	}
}
