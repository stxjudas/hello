package tool

import (
	"fmt"
	"math"
	)

func ArraySort() {
	var gobbot []int = []int{3, 5, 7, 9, 11, 13}
	fmt.Printf("static array: %v\n", gobbot)
	if len(gobbot) == 0 {
		return
	}
	max := gobbot[0]
	for _, value := range gobbot {
		if value > max {
			max = value
		}
	}
	fmt.Printf("maximum is: %d\n", max)
}

func ArrayPi() {
	var slPi []int8 = []int8{}
	var accuracy int = 10
	var tempPi float64 = math.Pi
	for i := 0; i < accuracy; i++ {
		slPi = append(slPi, int8(tempPi))
		//	tempPi = tempPi - float64(slPi[i])
		tempPi -= float64(slPi[i])
		//	tempPi = tempPi * 10
		tempPi *= 10
	}
	fmt.Printf("PiDigits: %v\n", slPi)
}

func FloatToArray(floatNum float64, accuracy int) (slResult []int, err error) {
	if accuracy < 0 {
		err = fmt.Errorf("negative errors accuracy not allowed")
		return
	}
	for i := 0; i < accuracy; i++ {
		slResult = append(slResult, int(floatNum))
		floatNum -= float64(slResult[i])
		floatNum *= 10
	}
	return
}
