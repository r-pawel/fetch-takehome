package models

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// TODO separate validation and calculating points into separate functions

func CheckMissingData(retailer string, items []Item) error {
	if strings.Trim(retailer, " ") == "" || len(items) == 0 {
		return fmt.Errorf("missing data")
	}
	return nil
}

func CalculateDatePoints(purchaseDateString string) (int, error) {
	purchaseDate, err := time.Parse("2006-01-02", purchaseDateString)
	if err != nil {
		return 0, fmt.Errorf("date parse error")
	}
	purchaseDay := purchaseDate.Day()
	if purchaseDay%2 == 1 {
		return 6, nil
	}
	return 0, nil
}

func CalculateTimePoints(purchaseTimeString string) (int, error) {
	purchaseTime, err := time.Parse("15:04", purchaseTimeString)
	if err != nil {
		return 0, fmt.Errorf("time parse error")
	}
	early, _ := time.Parse("15:04", "14:00")
	late, _ := time.Parse("15:04", "16:00")
	if purchaseTime.Before(late) && purchaseTime.After(early) {
		return 10, nil
	}
	return 0, nil
}

func CalculateTotalPoints(purchaseTotalString string) (int, error) {
	points := 0

	if !(regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(purchaseTotalString)) {
		return 0, fmt.Errorf("total format error")
	}

	totalFloat, err := strconv.ParseFloat(purchaseTotalString, 64)
	if err != nil {
		return 0, fmt.Errorf("total parse error")
	}

	if math.Mod(totalFloat, 1) == 0 {
		points += 50
	}

	if math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	return points, nil
}

func CalculateItemsPoints(items []Item) (int64, error) {
	var pointTotal int64 = 0
	for _, item := range items {
		description := strings.Trim(item.ShortDescription, " ")
		if !(regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(item.Price)) {
			return 0, fmt.Errorf("price format error")
		}
		if len(description) == 0 {
			return 0, fmt.Errorf("description format error")
		}
		if len(description)%3 == 0 {
			numberPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("parse price error")
			}
			pointTotal += (int64)(math.Ceil(numberPrice * .2))
		}
	}
	pointTotal += (int64)(len(items)/2) * 5
	return pointTotal, nil
}

func CalculateAlphanumericPoints(retailer string) int {
	alphanumeric := 0

	for _, char := range retailer {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			alphanumeric += 1
		}
	}

	return alphanumeric
}

func CheckTotal(items []Item, purchaseTotalString string) error {
	doubleTotal, err := strconv.ParseFloat(purchaseTotalString, 64)
	if err != nil {
		return fmt.Errorf("total parse error")
	}

	itemsTotalCost := 0.00

	for _, item := range items {
		itemCost, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return fmt.Errorf("individual cost parse error")
		}
		itemsTotalCost += itemCost
	}

	if (itemsTotalCost - doubleTotal) > 0.001 {
		return fmt.Errorf("item total doesnt match actual total")
	}

	return nil
}
