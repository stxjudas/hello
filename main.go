package main

import (
	"fmt"
	"math"
	"github.com/stxjudas/hello/tool"
)

type extremum struct {
	min int
	max int
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
	tool.ArrayPi()
	tool.ArraySort()

	var result []int
	var err error
	result, err = tool.FloatToArray(math.Ln2, 50)
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
	oExtremum1 := tool.NewTExtremum(arModify)
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

