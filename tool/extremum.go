package tool

import (
	"fmt"
)

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
