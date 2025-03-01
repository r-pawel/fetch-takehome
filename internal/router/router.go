package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/r-pawel/fetch-takehome/internal/models"
	"net/http"
)

var receiptsByID = make(map[string]int64)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", sendPoints)
	return router
}

func processReceipt(c *gin.Context) {
	var (
		newReceipt models.Receipt
		pointTotal int64 = 0
	)

	if err := c.BindJSON(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}

	err := models.CheckTotal(newReceipt.Items, newReceipt.Total)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}

	err = models.CheckMissingData(newReceipt.Retailer, newReceipt.Items)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}

	datePointAdd, err := models.CalculateDatePoints(newReceipt.PurchaseDate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}
	pointTotal += int64(datePointAdd)

	timePointAdd, err := models.CalculateTimePoints(newReceipt.PurchaseTime)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}
	pointTotal += int64(timePointAdd)

	totalPointAdd, err := models.CalculateTotalPoints(newReceipt.Total)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}
	pointTotal += int64(totalPointAdd)

	itemPointAdd, err := models.CalculateItemsPoints(newReceipt.Items)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}
	pointTotal += itemPointAdd

	pointTotal += int64(models.CalculateAlphanumericPoints(newReceipt.Retailer))

	newID := uuid.New()
	returnID := models.ID{ID: newID.String()}
	receiptsByID[newID.String()] = pointTotal

	c.IndentedJSON(http.StatusOK, returnID)
}

func sendPoints(c *gin.Context) {
	checkID := c.Param("id")
	if receiptsByID[checkID] != 0 {
		returnPoints := models.Points{Points: receiptsByID[checkID]}
		c.IndentedJSON(http.StatusOK, returnPoints)
	} else {
		c.IndentedJSON(http.StatusNotFound, "No receipt found for that id")
	}
}
