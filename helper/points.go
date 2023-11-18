package helper

import (
	"fmt"
	"math"
	"receipt-processor-backend/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(receipt models.Receipt) (int, error) {
	points := 0

	points += calculatePointsFromRetailerName(receipt)

	pointsFromReceiptTotal, err := calculatePointsFromReceiptTotal(receipt)
	if err != nil {
		return 0, err
	}
	points += pointsFromReceiptTotal

	fmt.Printf("\n Points after receipt total : %d \n", points)

	pointsFromReceiptItems, err := calculatePointsFromItems(receipt)
	if err != nil {
		return 0, err
	}
	points += pointsFromReceiptItems
	fmt.Printf("\n Points after receipt items : %d \n", points)

	pointsFromReceiptPurchaseDateTime, err := calculatePointsFromPurchaseDateTime(receipt)
	if err != nil {
		return 0, err
	}
	points += pointsFromReceiptPurchaseDateTime

	return points, nil
}

func calculatePointsFromRetailerName(receipt models.Receipt) int {
	count := 0

	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}

	return count
}

func calculatePointsFromReceiptTotal(receipt models.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, err
	}

	returnPoints := 0

	if total == math.Trunc(total) {
		returnPoints += 50
	}

	if math.Mod(total, 0.25) == 0 {
		returnPoints += 25
	}

	return returnPoints, nil

}

func calculatePointsFromItems(receipt models.Receipt) (int, error) {

	returnPoints := 0

	numItems := len(receipt.Items)

	returnPoints += 5 * int(numItems/2)

	for i := 0; i < numItems; i++ {

		currItem := receipt.Items[i]

		if len(strings.Trim(currItem.ShortDescription, " "))%3 == 0 {
			price, err := strconv.ParseFloat(currItem.Price, 64)
			if err != nil {
				return 0, err
			}
			returnPoints += int(math.Trunc(price*0.2)) + 1
		}
	}

	return returnPoints, nil

}

func calculatePointsFromPurchaseDateTime(receipt models.Receipt) (int, error) {

	returnPoints := 0
	purchaseDateTime, err := time.Parse("2006-01-02 15:04", receipt.PurchaseDate+" "+receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}

	if purchaseDateTime.Day()%2 == 1 {
		returnPoints += 6
	}

	if (purchaseDateTime.Hour() >= 14 && purchaseDateTime.Minute() != 0) && purchaseDateTime.Hour() < 16 {
		returnPoints += 10
	}

	return returnPoints, nil

}
