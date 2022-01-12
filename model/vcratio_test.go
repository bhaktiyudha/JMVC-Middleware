package model

import (
	"testing"
)

func TestVCRatio(t *testing.T) {
	var counterData CounterData
	counterData.CarUp = 700
	counterData.BusSUp = 50
	counterData.BusLUp = 35
	counterData.TruckSUp = 0
	counterData.TruckMUp = 115
	counterData.TruckLUp = 100
	counterData.TruckXLUp = 0
	expectation := 0.5203289733721655
	actual, _ := countVCRatio(counterData)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
