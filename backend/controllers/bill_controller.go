package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/NoT-Ton/app/ent"
	"github.com/NoT-Ton/app/ent/customer"
	"github.com/NoT-Ton/app/ent/product"
	"github.com/NoT-Ton/app/ent/user"
	"github.com/gin-gonic/gin"
)

// BillController defines the struct for the bill controller
type BillController struct {
	client *ent.Client
	router gin.IRouter
}

// Bill defines the struct for the bill
type Bill struct {
	ProductID  int
	QuantityID int
	UserID     int
	CustomerID int
	Addtime    string
}

// CreateBill handles POST requests for adding bill entities
// @Summary Create bill
// @Description Create bill
// @ID create-bill
// @Accept   json
// @Produce  json
// @Param bill body Bill true "Bill entity"
// @Success 200 {object} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [post]
func (ctl *BillController) CreateBill(c *gin.Context) {
	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
		})
		return
	}

	p, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.ProductID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "product not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	cu, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.CustomerID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Addtime)
	b, err := ctl.client.Bill.
		Create().
		SetProduct(p).
		SetUser(u).
		SetCustomer(cu).
		SetAddedTime(times).
		SetQuantity(obj.QuantityID).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// ListBill handles request to get a list of bill entities
// @Summary List bill entities
// @Description list bill entities
// @ID list-bill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [get]
func (ctl *BillController) ListBill(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	bills, err := ctl.client.Bill.
		Query().
		WithCustomer().
		WithProduct().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bills)
}

// DeleteBill handles DELETE requests to delete a bill entity
// @Summary Delete a bill entity by ID
// @Description get bill by ID
// @ID delete-bill
// @Produce  json
// @Param id path int true "Bill ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills/{id} [delete]
func (ctl *BillController) DeleteBill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bill.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewBillController creates and registers handles for the bill controller
func NewBillController(router gin.IRouter, client *ent.Client) *BillController {
	bc := &BillController{
		client: client,
		router: router,
	}
	bc.register()
	return bc
}

// InitBillController registers routes to the main engine
func (ctl *BillController) register() {
	bills := ctl.router.Group("/bills")

	bills.GET("", ctl.ListBill)
	bills.POST("", ctl.CreateBill)
	bills.DELETE(":id", ctl.DeleteBill)
}
