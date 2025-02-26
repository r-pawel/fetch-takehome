package models

import "testing"

func TestCheckMissingDataOneItem(t *testing.T) {
	retailer := "HMart"
	var items = []item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	testErr := checkMissingData(retailer, items)
	if testErr != nil {
		t.Fatalf(`checkMissingData(%v) should have no errors instead returned err = %v`, retailer, testErr)
	}
}

func TestCheckMissingDataMultipleItems(t *testing.T) {
	retailer := "HMart"
	var items = []item{
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
	testErr := checkMissingData(retailer, items)
	if testErr != nil {
		t.Fatalf(`checkMissingData(%v) should have no errors instead returned err = %v`, retailer, testErr)
	}
}

func TestCheckMissingDataNoItems(t *testing.T) {
	retailer := "HMart"
	var items = []item{}
	wantErr := "missing data"
	testErr := checkMissingData(retailer, items)
	if testErr.Error() != wantErr {
		t.Fatalf(`checkMissingData(%v) = %v, want match for, %v`, retailer, testErr, wantErr)
	}
}

func TestCheckMissingDataNoRetailer(t *testing.T) {
	retailer := " "
	var items = []item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantErr := "missing data"
	testErr := checkMissingData(retailer, items)
	if testErr.Error() != wantErr {
		t.Fatalf(`checkMissingData(%v) = %v, want match for, %v`, retailer, testErr, wantErr)
	}
}

func TestCheckMissingDataMissingBoth(t *testing.T) {
	retailer := " "
	var items = []item{}
	wantErr := "missing data"
	testErr := checkMissingData(retailer, items)
	if testErr.Error() != wantErr {
		t.Fatalf(`checkMissingData(%v) = %v, want match for, %v`, retailer, testErr, wantErr)
	}
}

func TestCheckDateEven(t *testing.T) {
	date := "2005-12-12"
	wantInt := 0
	testInt, testErr := checkDate(date)
	if testErr != nil {
		t.Fatalf(`checkDate(%v) returned the error: %v this test shouldnt return error`, date, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkDate(%v) = %v, want match for, %v`, date, testInt, wantInt)
	}
}

func TestCheckDateOdd(t *testing.T) {
	date := "2005-12-11"
	wantInt := 6
	testInt, testErr := checkDate(date)
	if testErr != nil {
		t.Fatalf(`checkDate(%v) returned the error: %v this test shouldnt return error`, date, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkDate(%v) = %v, want match for, %v`, date, testInt, wantInt)
	}
}

func TestCheckDateInvalid(t *testing.T) {
	date := "2001-500-120"
	wantInt := 0
	wantErr := "date parse error"
	testInt, testErr := checkDate(date)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkDate(%v) = %v, %v want match for %v, %v`, date, testInt, testErr, wantInt, wantErr)
	}
}

func TestCheckDateEmpty(t *testing.T) {
	date := ""
	wantInt := 0
	wantErr := "date parse error"
	testInt, testErr := checkDate(date)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkDate(%v) = %v, %v want match for %v, %v`, date, testInt, testErr, wantInt, wantErr)
	}
}

func TestCheckTimeMiddleBonus(t *testing.T) {
	time := "15:00"
	wantInt := 10
	testInt, testErr := checkTime(time)
	if testErr != nil {
		t.Fatalf(`checkTime(%v) returned an error: %v when it shouldn't have`, time, testErr)
	}
	if wantInt != testInt {
		t.Fatalf(`checkTime(%v) = %v, want match for, %v`, time, testInt, wantInt)
	}
}
func TestCheckTimeLowBonus(t *testing.T) {
	time := "14:00"
	wantInt := 0
	testInt, testErr := checkTime(time)
	if testErr != nil {
		t.Fatalf(`checkTime(%v) returned an error: %v when it shouldn't have`, time, testErr)
	}
	if wantInt != testInt {
		t.Fatalf(`checkTime(%v) = %v, want match for, %v`, time, testInt, wantInt)
	}
}
func TestCheckTimeHighBonus(t *testing.T) {
	time := "16:00"
	wantInt := 0
	testInt, testErr := checkTime(time)
	if testErr != nil {
		t.Fatalf(`checkTime(%v) returned an error: %v when it shouldn't have`, time, testErr)
	}
	if wantInt != testInt {
		t.Fatalf(`checkTime(%v) = %v, want match for, %v`, time, testInt, wantInt)
	}
}

func TestCheckTimeEmpty(t *testing.T) {
	time := ""
	wantInt := 0
	wantErr := "time parse error"
	testInt, testErr := checkTime(time)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkTime(%v) = %v, %v want match for %v, %v`, time, testInt, testErr, wantInt, wantErr)
	}
}

func TestCheckTimeBadHour(t *testing.T) {
	time := "50:00"
	wantInt := 0
	wantErr := "time parse error"
	testInt, testErr := checkTime(time)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkTime(%v) = %v, %v want match for %v, %v`, time, testInt, testErr, wantInt, wantErr)
	}
}

func TestCheckTimeBadMinute(t *testing.T) {
	time := "16:90"
	wantInt := 0
	wantErr := "time parse error"
	testInt, testErr := checkTime(time)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkTime(%v) = %v, %v want match for %v, %v`, time, testInt, testErr, wantInt, wantErr)
	}
}

func TestCheckTotalNoPoints(t *testing.T) {
	total := "12.95"
	wantInt := 0
	testInt, testErr := checkTotal(total)
	if testErr != nil {
		t.Fatalf(`checkTotal(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCheckTotalEndIn00(t *testing.T) {
	total := "12.00"
	wantInt := 75
	testInt, testErr := checkTotal(total)
	if testErr != nil {
		t.Fatalf(`checkTotal(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCheckTotalEndIn25(t *testing.T) {
	total := "12.25"
	wantInt := 25
	testInt, testErr := checkTotal(total)
	if testErr != nil {
		t.Fatalf(`checkTotal(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCheckTotalEndIn50(t *testing.T) {
	total := "12.50"
	wantInt := 25
	testInt, testErr := checkTotal(total)
	if testErr != nil {
		t.Fatalf(`checkTotal(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCheckTotalEndIn75(t *testing.T) {
	total := "12.75"
	wantInt := 25
	testInt, testErr := checkTotal(total)
	if testErr != nil {
		t.Fatalf(`checkTotal(%v) returned an error: %v when it shouldn't have`, total, testErr)
	}
	if testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, want match for, %v`, total, testInt, wantInt)
	}
}
func TestCheckTotalExtraDecimal(t *testing.T) {
	total := "12.955"
	wantInt := 0
	wantErr := "total format error"
	testInt, testErr := checkTotal(total)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, %v want match for %v, %v`, total, testInt, testErr, wantInt, wantErr)
	}
}
func TestCheckTotalEmpty(t *testing.T) {
	total := " "
	wantInt := 0
	wantErr := "total format error"
	testInt, testErr := checkTotal(total)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, %v want match for %v, %v`, total, testInt, testErr, wantInt, wantErr)
	}
}
func TestCheckTotalBadFormat(t *testing.T) {
	total := "a12.95"
	wantInt := 0
	wantErr := "total format error"
	testInt, testErr := checkTotal(total)
	if testErr.Error() != wantErr || testInt != wantInt {
		t.Fatalf(`checkTotal(%v) = %v, %v want match for %v, %v`, total, testInt, testErr, wantInt, wantErr)
	}
}
func TestCheckItems1Item(t *testing.T) {
	var items = []item{
		{
			ShortDescription: "Item",
			Price:            "20.15",
		},
	}
	wantInt := 0
	testInt, testErr := checkItems(items)
	if testErr != nil {
		t.Fatalf(`checkItems(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCheckItems2Items(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr != nil {
		t.Fatalf(`checkItems(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCheckItems3Items(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr != nil {
		t.Fatalf(`checkItems(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCheckItems10Items(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr != nil {
		t.Fatalf(`checkItems(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCheckItems3Description(t *testing.T) {
	var items = []item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	wantInt := 5
	testInt, testErr := checkItems(items)
	if testErr != nil {
		t.Fatalf(`checkItems(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCheckItemsMultipleDescription(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr != nil {
		t.Fatalf(`checkItems(%v) returned an error: %v when it shouldn't have`, items, testErr)
	}
	if testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, want match for, %v`, items, testInt, wantInt)
	}
}
func TestCheckItemsPriceError(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr.Error() != wantErr || testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, %v want match for %v, %v`, items, testInt, testErr, wantInt, wantErr)
	}
}
func TestCheckItemsDescriptionError(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr.Error() != wantErr || testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, %v want match for %v, %v`, items, testInt, testErr, wantInt, wantErr)
	}
}
func TestCheckItemsEmptyPriceError(t *testing.T) {
	var items = []item{
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
	testInt, testErr := checkItems(items)
	if testErr.Error() != wantErr || testInt != int64(wantInt) {
		t.Fatalf(`checkItems(%v) = %v, %v want match for %v, %v`, items, testInt, testErr, wantInt, wantErr)
	}
}
func TestCheckAlphanumericLetters(t *testing.T) {
	retailer := "Hmart"
	wantInt := 5
	testInt := checkAlphanumeric(retailer)
	if testInt != wantInt {
		t.Fatalf(`checkAlphanumeric(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCheckAlphanumericNumbers(t *testing.T) {
	retailer := "12345"
	wantInt := 5
	testInt := checkAlphanumeric(retailer)
	if testInt != wantInt {
		t.Fatalf(`checkAlphanumeric(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCheckAlphanumericSymbols(t *testing.T) {
	retailer := "!@#$%^&*"
	wantInt := 0
	testInt := checkAlphanumeric(retailer)
	if testInt != wantInt {
		t.Fatalf(`checkAlphanumeric(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestCheckAlphanumericAll(t *testing.T) {
	retailer := "H&M45"
	wantInt := 4
	testInt := checkAlphanumeric(retailer)
	if testInt != wantInt {
		t.Fatalf(`checkAlphanumeric(%v) = %v, want match for, %v`, retailer, testInt, wantInt)
	}
}
func TestValidateTotalValid(t *testing.T) {
	total := "20.15"
	var items = []item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := validateTotal(items, total)
	if err != nil {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestValidateTotalValidMultiple(t *testing.T) {
	total := "40.30"
	var items = []item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := validateTotal(items, total)
	if err != nil {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestValidateTotalValidALot(t *testing.T) {
	total := "241.80"
	var items = []item{
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
	err := validateTotal(items, total)
	if err != nil {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestValidateTotalInvalid(t *testing.T) {
	total := "20.00"
	var items = []item{
		{
			ShortDescription: "Ite",
			Price:            "20.15",
		},
	}
	err := validateTotal(items, total)
	if err.Error() != "item total doesnt match actual total" {
		t.Fatalf(`%v`, err.Error())
	}
}
func TestValidateTotalInvalidMultiple(t *testing.T) {
	total := "20.15"
	var items = []item{
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
	err := validateTotal(items, total)
	if err.Error() != "item total doesnt match actual total" {
		t.Fatalf(`%v`, err.Error())
	}
}
