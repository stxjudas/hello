package main

import (
	"fmt"
	"math"
)

type extremum struct {
	min int
	max int
}

type TExtremum struct {
	min      int
	indexMin int
	max      int
	indexMax int
	arr      []int
}

func NewTExtremum(sl []int) *TExtremum {
	return &TExtremum{arr: sl}
}

func (oExtr *TExtremum) calculate() (err error) {
	if len(oExtr.arr) == 0 {
		err = fmt.Errorf("empty slice not allowed")
		return
	}
	oExtr.min = oExtr.arr[0]
	oExtr.max = oExtr.arr[0]
	for i, value := range oExtr.arr {
		if value < oExtr.min {
			oExtr.min = value
			oExtr.indexMin = i
			fmt.Printf("%d asserted as minimum\n", oExtr.min)
		}
		if value > oExtr.max {
			oExtr.max = value
			oExtr.indexMax = i
		}
	}
	return
}

func (oExtr *TExtremum) GetMin() (int, int, error) {
	err := oExtr.calculate()
	return oExtr.indexMin, oExtr.min, err
}

func (oExtr *TExtremum) GetMax() (int, int, error) {
	return oExtr.indexMax, oExtr.max, oExtr.calculate()
}

func main() {
	fmt.Printf("This is hello world app \n")
	fmt.Println("Line without carriage return")

	var lambda int8
	fmt.Printf("emptyvalue is: %d\n", lambda)
	lambda = 125
	fmt.Printf("newvalue is: %d\n", lambda)

	pi := math.Pi
	fmt.Printf("pivalue is: %1.7f of type: %T\n", pi, pi)
	arrayPi()
	arraySort()

	var result []int
	var err error
	result, err = floatToArray(math.Ln2, 50)
	if err != nil {
		fmt.Printf("achtung: %s\n", err.Error())
		return
	}
	fmt.Printf("first %d digits of number are: %v\n", 10, result)

	arOrigin := []int{-2, 4, 9, 16, 25, 36, 49, 15}
	//	arOrigin := []int{}

	var oExtremum extremum
	oExtremum, err = newExtremum(arOrigin)
	if err != nil {
		fmt.Printf("alarm: %s\n", err.Error())
		return
	}
	fmt.Printf("our extremum: %+v\n", oExtremum)
	fmt.Printf("max is: %d\n", oExtremum.max)

	arModify := append(arOrigin, -10, 10, 20, 30, 40, 50, 60)
	oExtremum1 := NewTExtremum(arModify)
	var min, max, indexMin, indexMax int
	indexMin, min, err = oExtremum1.GetMin()
	if err != nil {
		fmt.Printf("danger: %s\n", err.Error())
		return
	}
	fmt.Printf("got minimum: %d at: %d\n", min, indexMin)

	fmt.Printf("object is: %+v\n", oExtremum1)
	indexMax, max, err = oExtremum1.GetMax()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("got maximum: %d at: %d\n", max, indexMax)
}

func newExtremum(sl []int) (oResult extremum, err error) {
	if len(sl) == 0 {
		err = fmt.Errorf("empty slice not allowed")
		return
	}
	var min, max int
	min = sl[0]
	max = sl[0]
	for _, value := range sl {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	oResult = extremum{min, max}
	return
}

func arraySort() {
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

func arrayPi() {
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

func floatToArray(floatNum float64, accuracy int) (slResult []int, err error) {
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
