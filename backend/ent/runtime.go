// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/NoT-Ton/app/ent/bill"
	"github.com/NoT-Ton/app/ent/customer"
	"github.com/NoT-Ton/app/ent/product"
	"github.com/NoT-Ton/app/ent/role"
	"github.com/NoT-Ton/app/ent/schema"
	"github.com/NoT-Ton/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	billFields := schema.Bill{}.Fields()
	_ = billFields
	// billDescQuantity is the schema descriptor for quantity field.
	billDescQuantity := billFields[0].Descriptor()
	// bill.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	bill.QuantityValidator = billDescQuantity.Validators[0].(func(int) error)
	// billDescAddedTime is the schema descriptor for added_time field.
	billDescAddedTime := billFields[1].Descriptor()
	// bill.DefaultAddedTime holds the default value on creation for the added_time field.
	bill.DefaultAddedTime = billDescAddedTime.Default.(func() time.Time)
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCustomerName is the schema descriptor for customer_name field.
	customerDescCustomerName := customerFields[0].Descriptor()
	// customer.CustomerNameValidator is a validator for the "customer_name" field. It is called by the builders before save.
	customer.CustomerNameValidator = customerDescCustomerName.Validators[0].(func(string) error)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescProductName is the schema descriptor for product_name field.
	productDescProductName := productFields[0].Descriptor()
	// product.ProductNameValidator is a validator for the "product_name" field. It is called by the builders before save.
	product.ProductNameValidator = productDescProductName.Validators[0].(func(string) error)
	// productDescProductCost is the schema descriptor for product_cost field.
	productDescProductCost := productFields[1].Descriptor()
	// product.ProductCostValidator is a validator for the "product_cost" field. It is called by the builders before save.
	product.ProductCostValidator = productDescProductCost.Validators[0].(func(int) error)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescRoleName is the schema descriptor for role_name field.
	roleDescRoleName := roleFields[0].Descriptor()
	// role.RoleNameValidator is a validator for the "role_name" field. It is called by the builders before save.
	role.RoleNameValidator = roleDescRoleName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserEmail is the schema descriptor for user_email field.
	userDescUserEmail := userFields[0].Descriptor()
	// user.UserEmailValidator is a validator for the "user_email" field. It is called by the builders before save.
	user.UserEmailValidator = userDescUserEmail.Validators[0].(func(string) error)
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
}
