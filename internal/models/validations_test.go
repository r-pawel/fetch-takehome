package models

import "testing"

func TestCheckMissingDataOneItem(t *testing.T) {
	retailer := "HMart"
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	testErr := CheckMissingData(retailer, items)
	if testErr != nil {
		t.Fatalf(`checkMissingData(%v) should have no errors instead returned err = %v`, retailer, testErr)
	}
}

func TestCheckMissingDataMultipleItems(t *testing.T) {
	retailer := "HMart"
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	testErr := CheckMissingData(retailer, items)
	if testErr != nil {
		t.Fatalf(`checkMissingData(%v) should have no errors instead returned err = %v`, retailer, testErr)
	}
}

func TestCheckMissingDataNoItems(t *testing.T) {
	retailer := "HMart"
	var items = []Item{}
	wantErr := "missing data"
	testErr := CheckMissingData(retailer, items)
	if testErr.Error() != wantErr {
		t.Fatalf(`checkMissingData(%v) = %v, want match for, %v`, retailer, testErr, wantErr)
	}
}

func TestCheckMissingDataNoRetailer(t *testing.T) {
	retailer := " "
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantErr := "missing data"
	testErr := CheckMissingData(retailer, items)
	if testErr.Error() != wantErr {
		t.Fatalf(`checkMissingData(%v) = %v, want match for, %v`, retailer, testErr, wantErr)
	}
}

func TestCheckMissingDataMissingBoth(t *testing.T) {
	retailer := " "
	var items = []Item{}
	wantErr := "missing data"
	testErr := CheckMissingData(retailer, items)
	if testErr.Error() != wantErr {
		t.Fatalf(`checkMissingData(%v) = %v, want match for, %v`, retailer, testErr, wantErr)
	}
}

func TestCalculateDatePointsEven(t *testing.T) {
	date := "2005-12-12"
	wantInt := 0
	testInt, testErr := CalculateDatePoints(date)
	if testErr != nil {
		t.Fatalf(`CalculateDatePoints(%v) returned the error: %v this test shouldnt return error`, date, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateDatePoints(%v) = %v, want match for, %v`, date, testInt, wantInt)
	}
}

func TestCalculateDatePointsOdd(t *testing.T) {
	date := "2005-12-11"
	wantInt := 6
	testInt, testErr := CalculateDatePoints(date)
	if testErr != nil {
		t.Fatalf(`CalculateDatePoints(%v) returned the error: %v this test shouldnt return error`, date, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateDatePoints(%v) = %v, want match for, %v`, date, testInt, wantInt)
	}
}

func TestCalculateDatePointsInvalid(t *testing.T) {
	date := "2001-500-120"
	wantInt := 0
	wantErr := "date parse error"
	testInt, testErr := CalculateDatePoints(date)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateDatePoints(%v) = %v, %v want match for %v, %v`, date, testInt, testErr, wantInt, wantErr)
	}
}

func TestCalculateDatePointsEmpty(t *testing.T) {
	date := ""
	wantInt := 0
	wantErr := "date parse error"
	testInt, testErr := CalculateDatePoints(date)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateDatePoints(%v) = %v, %v want match for %v, %v`, date, testInt, testErr, wantInt, wantErr)
	}
}

func TestCalculateTimePointsMiddleBonus(t *testing.T) {
	time := "15:00"
	wantInt := 10
	testInt, testErr := CalculateTimePoints(time)
	if testErr != nil {
		t.Fatalf(`CalculateTimePoints(%v) returned an error: %v when it shouldn't have`, time, testErr)
	}
	if wantInt != testInt {
		t.Fatalf(`CalculateTimePoints(%v) = %v, want match for, %v`, time, testInt, wantInt)
	}
}
func TestCalculateTimePointsLowBonus(t *testing.T) {
	time := "14:00"
	wantInt := 0
	testInt, testErr := CalculateTimePoints(time)
	if testErr != nil {
		t.Fatalf(`CalculateTimePoints(%v) returned an error: %v when it shouldn't have`, time, testErr)
	}
	if wantInt != testInt {
		t.Fatalf(`CalculateTimePoints(%v) = %v, want match for, %v`, time, testInt, wantInt)
	}
}
func TestCalculateTimePointsHighBonus(t *testing.T) {
	time := "16:00"
	wantInt := 0
	testInt, testErr := CalculateTimePoints(time)
	if testErr != nil {
		t.Fatalf(`CalculateTimePoints(%v) returned an error: %v when it shouldn't have`, time, testErr)
	}
	if wantInt != testInt {
		t.Fatalf(`CalculateTimePoints(%v) = %v, want match for, %v`, time, testInt, wantInt)
	}
}

func TestCalculateTimePointsEmpty(t *testing.T) {
	time := ""
	wantInt := 0
	wantErr := "time parse error"
	testInt, testErr := CalculateTimePoints(time)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateTimePoints(%v) = %v, %v want match for %v, %v`, time, testInt, testErr, wantInt, wantErr)
	}
}

func TestCalculateTimePointsBadHour(t *testing.T) {
	time := "50:00"
	wantInt := 0
	wantErr := "time parse error"
	testInt, testErr := CalculateTimePoints(time)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateTimePoints(%v) = %v, %v want match for %v, %v`, time, testInt, testErr, wantInt, wantErr)
	}
}

func TestCalculateTimePointsBadMinute(t *testing.T) {
	time := "16:90"
	wantInt := 0
	wantErr := "time parse error"
	testInt, testErr := CalculateTimePoints(time)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateTimePoints(%v) = %v, %v want match for %v, %v`, time, testInt, testErr, wantInt, wantErr)
	}
}

func TestCalculateTotalPointsNoPoints(t *testing.T) {
	total := "12.95"
	wantInt := 0
	testInt, testErr := CalculateTotalPoints(total)
	if testErr != nil {
		t.Fatalf(`CalculateTotalPoints(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCalculateTotalPointsEndIn00(t *testing.T) {
	total := "12.00"
	wantInt := 75
	testInt, testErr := CalculateTotalPoints(total)
	if testErr != nil {
		t.Fatalf(`CalculateTotalPoints(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCalculateTotalPointsEndIn25(t *testing.T) {
	total := "12.25"
	wantInt := 25
	testInt, testErr := CalculateTotalPoints(total)
	if testErr != nil {
		t.Fatalf(`CalculateTotalPoints(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCalculateTotalPointsEndIn50(t *testing.T) {
	total := "12.50"
	wantInt := 25
	testInt, testErr := CalculateTotalPoints(total)
	if testErr != nil {
		t.Fatalf(`CalculateTotalPoints(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCalculateTotalPointsEndIn75(t *testing.T) {
	total := "12.75"
	wantInt := 25
	testInt, testErr := CalculateTotalPoints(total)
	if testErr != nil {
		t.Fatalf(`CalculateTotalPoints(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCalculateTotalPointsExtraDecimal(t *testing.T) {
	total := "12.955"
	wantInt := 0
	wantErr := "total format error"
	testInt, testErr := CalculateTotalPoints(total)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, %v want match for %v, %v`, total, testInt, testErr, wantInt, wantErr)
	}
}
func TestCalculateTotalPointsEmpty(t *testing.T) {
	total := " "
	wantInt := 0
	wantErr := "total format error"
	testInt, testErr := CalculateTotalPoints(total)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, %v want match for %v, %v`, total, testInt, testErr, wantInt, wantErr)
	}
}
func TestCalculateTotalPointsBadFormat(t *testing.T) {
	total := "a12.95"
	wantInt := 0
	wantErr := "total format error"
	testInt, testErr := CalculateTotalPoints(total)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`CalculateTotalPoints(%v) = %v, %v want match for %v, %v`, total, testInt, testErr, wantInt, wantErr)
	}
}
func TestCalculateItemsPoints1Item(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 0
	testInt, testErr := CalculateItemsPoints(items)
	if testErr != nil {
		t.Fatalf(`CalculateItemsPoints(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCalculateItemsPoints2Items(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 5
	testInt, testErr := CalculateItemsPoints(items)
	if testErr != nil {
		t.Fatalf(`CalculateItemsPoints(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCalculateItemsPoints3Items(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 5
	testInt, testErr := CalculateItemsPoints(items)
	if testErr != nil {
		t.Fatalf(`CalculateItemsPoints(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCalculateItemsPoints10Items(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 25
	testInt, testErr := CalculateItemsPoints(items)
	if testErr != nil {
		t.Fatalf(`CalculateItemsPoints(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCalculateItemsPoints3Description(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	wantInt := 5
	testInt, testErr := CalculateItemsPoints(items)
	if testErr != nil {
		t.Fatalf(`CalculateItemsPoints(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCalculateItemsPointsMultipleDescription(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "16.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "200.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "2.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "50.15",
		},
	}
	wantInt := 72
	testInt, testErr := CalculateItemsPoints(items)
	if testErr != nil {
		t.Fatalf(`CalculateItemsPoints(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCalculateItemsPointsPriceError(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.155",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 0
	wantErr := "price format error"
	testInt, testErr := CalculateItemsPoints(items)
	if testErr.Error() != wantErr || testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, %v want match for %v, %v`, items, testInt, testErr, wantInt, wantErr)
	}
}
func TestCalculateItemsPointsDescriptionError(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
		{
			ShortDescription: "",
			Price:            "20.15",
		},
	}
	wantInt := 0
	wantErr := "description format error"
	testInt, testErr := CalculateItemsPoints(items)
	if testErr.Error() != wantErr || testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, %v want match for %v, %v`, items, testInt, testErr, wantInt, wantErr)
	}
}
func TestCalculateItemsPointsEmptyPriceError(t *testing.T) {
	var items = []Item{
		{
			ShortDescription: "Item",
			Price:            "",
		},
		{
			ShortDescription: "Item",
			Price:            "20.155",
		},
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 0
	wantErr := "price format error"
	testInt, testErr := CalculateItemsPoints(items)
	if testErr.Error() != wantErr || testInt != int64(wantInt) {
		t.Fatalf(`CalculateItemsPoints(%v) = %v, %v want match for %v, %v`, items, testInt, testErr, wantInt, wantErr)
	}
}
func TestCalculateAlphanumericPointsLetters(t *testing.T) {
	retailer := "Hmart"
	wantInt := 5
	testInt := CalculateAlphanumericPoints(retailer)
	if testInt != wantInt {
		t.Fatalf(`CalculateAlphanumericPoints(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCalculateAlphanumericPointsNumbers(t *testing.T) {
	retailer := "12345"
	wantInt := 5
	testInt := CalculateAlphanumericPoints(retailer)
	if testInt != wantInt {
		t.Fatalf(`CalculateAlphanumericPoints(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCalculateAlphanumericPointsSymbols(t *testing.T) {
	retailer := "!@#$%^&*"
	wantInt := 0
	testInt := CalculateAlphanumericPoints(retailer)
	if testInt != wantInt {
		t.Fatalf(`CalculateAlphanumericPoints(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCalculateAlphanumericPointsAll(t *testing.T) {
	retailer := "H&M45"
	wantInt := 4
	testInt := CalculateAlphanumericPoints(retailer)
	if testInt != wantInt {
		t.Fatalf(`CalculateAlphanumericPoints(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCheckTotalValid(t *testing.T) {
	total := "20.15"
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := CheckTotal(items, total)
	if err != nil {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestCheckTotalValidMultiple(t *testing.T) {
	total := "40.30"
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := CheckTotal(items, total)
	if err != nil {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestCheckTotalValidALot(t *testing.T) {
	total := "241.80"
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := CheckTotal(items, total)
	if err != nil {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestCheckTotalInvalid(t *testing.T) {
	total := "20.00"
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := CheckTotal(items, total)
	if err.Error() != "item total doesnt match actual total" {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestCheckTotalInvalidMultiple(t *testing.T) {
	total := "20.15"
	var items = []Item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := CheckTotal(items, total)
	if err.Error() != "item total doesnt match actual total" {
		t.Fatalf(`%v`, err.Error())
	}
}
