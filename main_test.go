package main

import (
	"strings"
	"testing"
)

var pancakesTest = []string{
	"+++++",
	"-----",
	"+",
	"-",
	"+-",
	"-+",
	"+-+-+-",
	"-+-+-+",
}
var pancakesRes = []string{
	"-----",
	"+++++",
	"-",
	"+",
	"-+",
	"+-",
	"-+-+-+",
	"+-+-+-",
}

var allPlusRes  = []bool{
	true,
	false,
	true,
	false,
	false,
	false,
	false,
	false,
}

func TestFlip1(t *testing.T) {
	for i, pancakes := range pancakesTest {
		pancakesSplit := strings.Split(pancakes, "")
		flip(pancakesSplit)
		if strings.Join(pancakesSplit, "") != pancakesRes[i] {
			t.Fail()
		}
		t.Logf("Pancake: %v", pancakesSplit)
	}
}

func TestAllPlus(t *testing.T) {
	for i, pancakes := range pancakesTest {
		pancakesSplit := strings.Split(pancakes, "")
		allPlus := allPlus(pancakesSplit)
		if allPlus != allPlusRes[i] {
			t.Logf("Pancake Fails: %v == %v", pancakesSplit, allPlus)
			t.Fail()
		}
	}
}