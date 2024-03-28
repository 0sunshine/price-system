package db

import (
	"fmt"
	"testing"
)

func Test_getMarkByLevel(t *testing.T) {

	parentProjectKindID := "000001000002000003000004000005000006"

	for idx := 0; idx < 6; idx++ {
		n, err := getMarkByLevel(parentProjectKindID, idx)
		fmt.Println(n)

		if err != nil || n != idx+1 {
			t.Fatal("getMarkByLevel error")
		}
	}

}

func Test_getNextProjectKindIDRange(t *testing.T) {
	parentProjectKindID :=
		"000000" +
			"000000" +
			"000000" +
			"000000" +
			"000000" +
			"000000"
	min, max, err := getNextProjectKindIDRange(parentProjectKindID)
	println(min, " ", max)
	if err != nil || min != "000001000000000000000000000000000000" || max != "999999000000000000000000000000000000" {
		t.Fatal("getNextProjectKindIDRange1  error")
	}

	parentProjectKindID =
		"000800" +
			"000000" +
			"000000" +
			"000000" +
			"000000" +
			"000000"
	min, max, err = getNextProjectKindIDRange(parentProjectKindID)
	println(min, " ", max)
	if err != nil || min != "000800000001000000000000000000000000" || max != "000800999999000000000000000000000000" {
		t.Fatal("getNextProjectKindIDRange2  error")
	}

	parentProjectKindID =
		"000800" +
			"000700" +
			"000000" +
			"000000" +
			"000000" +
			"000000"
	min, max, err = getNextProjectKindIDRange(parentProjectKindID)
	println("3-- ", min, " ", max)
	if err != nil || min != "000800000700000001000000000000000000" || max != "000800000700999999000000000000000000" {
		t.Fatal("getNextProjectKindIDRange3  error")
	}

	parentProjectKindID =
		"000800" +
			"000700" +
			"000600" +
			"000000" +
			"000000" +
			"000000"
	min, max, err = getNextProjectKindIDRange(parentProjectKindID)
	println(min, " ", max)
	if err != nil || min != "000800000700000600000001000000000000" || max != "000800000700000600999999000000000000" {
		t.Fatal("getNextProjectKindIDRange4  error")
	}

	parentProjectKindID =
		"000800" +
			"000700" +
			"000600" +
			"000500" +
			"000000" +
			"000000"
	min, max, err = getNextProjectKindIDRange(parentProjectKindID)
	println(min, " ", max)
	if err != nil || min != "000800000700000600000500000001000000" || max != "000800000700000600000500999999000000" {
		t.Fatal("getNextProjectKindIDRange5  error")
	}

	parentProjectKindID =
		"000800" +
			"000700" +
			"000600" +
			"000500" +
			"000400" +
			"000000"
	min, max, err = getNextProjectKindIDRange(parentProjectKindID)
	println(min, " ", max, " ", err != nil)
	if err != nil || min != "000800000700000600000500000400000001" || max != "000800000700000600000500000400999999" {
		t.Fatal("getNextProjectKindIDRange6  error")
	}

	parentProjectKindID =
		"000800" +
			"000700" +
			"000600" +
			"000500" +
			"000400" +
			"000300"
	min, max, err = getNextProjectKindIDRange(parentProjectKindID)
	println(min, " ", max, " ", err != nil)
	if err == nil {
		t.Fatal("getNextProjectKindIDRange7  error")
	}

}
