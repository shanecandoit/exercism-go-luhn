package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid is a card number valid?
func Valid(card string) bool {
	fmt.Println("Valid?", card)
	fmt.Println("")

	if len(card) <= 1 {
		return false
	}

	// strip spaces
	card = strings.ReplaceAll(card, " ", "")
	fmt.Println("no space", card)

	// 4539 1488 0343 6467
	// double ever other from right
	// 4_3_ 1_8_ 0_4_ 6_6_

	dubs := make([]int, len(card))
	for i := 1; i < len(card); i += 2 {
		ch := card[i]
		digi := string(ch)
		dubs[i], _ = strconv.Atoi(digi)
		// if err != nil {
		// 	return false, err
		// }

		dubs[i] += dubs[i]
		if dubs[i] > 9 {
			dubs[i] = dubs[i] - 9
		}

		if i+1 >= len(card) {
			break
		}
		ch = card[i+1]
		digi = string(ch)
		dubs[i+1], _ = strconv.Atoi(digi)
	}
	fmt.Println("dubs", dubs)

	sum := 0
	for _, n := range dubs {
		sum += n
	}
	fmt.Println("sum", sum)

	return sum%10 == 0
}

func main() {
	ex1 := "4539 1488 0343 6467"
	fmt.Println("Valid", Valid(ex1))
	fmt.Println("-")

	ex2 := "059"
	fmt.Println("Valid", Valid(ex2))
	fmt.Println("-")

	// this fails
	// luhn_test.go:8: Valid(59): a simple valid SIN that becomes invalid if reversed
	//     	 Expected: true
	//     	 Got: false
	ex3 := "59"
	fmt.Println("Valid", Valid(ex3))
	fmt.Println("-")
}
