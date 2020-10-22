// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NoT-Ton/app/ent/migrate"

	"github.com/NoT-Ton/app/ent/bill"
	"github.com/NoT-Ton/app/ent/customer"
	"github.com/NoT-Ton/app/ent/product"
	"github.com/NoT-Ton/app/ent/role"
	"github.com/NoT-Ton/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Bill is the client for interacting with the Bill builders.
	Bill *BillClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Bill = NewBillClient(c.config)
	c.Customer = NewCustomerClient(c.config)
	c.Product = NewProductClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Bill:     NewBillClient(cfg),
		Customer: NewCustomerClient(cfg),
		Product:  NewProductClient(cfg),
		Role:     NewRoleClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:   cfg,
		Bill:     NewBillClient(cfg),
		Customer: NewCustomerClient(cfg),
		Product:  NewProductClient(cfg),
		Role:     NewRoleClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Bill.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Bill.Use(hooks...)
	c.Customer.Use(hooks...)
	c.Product.Use(hooks...)
	c.Role.Use(hooks...)
	c.User.Use(hooks...)
}

// BillClient is a client for the Bill schema.
type BillClient struct {
	config
}

// NewBillClient returns a client for the Bill from the given config.
func NewBillClient(c config) *BillClient {
	return &BillClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bill.Hooks(f(g(h())))`.
func (c *BillClient) Use(hooks ...Hook) {
	c.hooks.Bill = append(c.hooks.Bill, hooks...)
}

// Create returns a create builder for Bill.
func (c *BillClient) Create() *BillCreate {
	mutation := newBillMutation(c.config, OpCreate)
	return &BillCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Bill.
func (c *BillClient) Update() *BillUpdate {
	mutation := newBillMutation(c.config, OpUpdate)
	return &BillUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BillClient) UpdateOne(b *Bill) *BillUpdateOne {
	mutation := newBillMutation(c.config, OpUpdateOne, withBill(b))
	return &BillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BillClient) UpdateOneID(id int) *BillUpdateOne {
	mutation := newBillMutation(c.config, OpUpdateOne, withBillID(id))
	return &BillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bill.
func (c *BillClient) Delete() *BillDelete {
	mutation := newBillMutation(c.config, OpDelete)
	return &BillDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BillClient) DeleteOne(b *Bill) *BillDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BillClient) DeleteOneID(id int) *BillDeleteOne {
	builder := c.Delete().Where(bill.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BillDeleteOne{builder}
}

// Create returns a query builder for Bill.
func (c *BillClient) Query() *BillQuery {
	return &BillQuery{config: c.config}
}

// Get returns a Bill entity by its id.
func (c *BillClient) Get(ctx context.Context, id int) (*Bill, error) {
	return c.Query().Where(bill.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BillClient) GetX(ctx context.Context, id int) *Bill {
	b, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// QueryProduct queries the product edge of a Bill.
func (c *BillClient) QueryProduct(b *Bill) *ProductQuery {
	query := &ProductQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bill.Table, bill.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bill.ProductTable, bill.ProductColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Bill.
func (c *BillClient) QueryUser(b *Bill) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bill.Table, bill.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bill.UserTable, bill.UserColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomer queries the customer edge of a Bill.
func (c *BillClient) QueryCustomer(b *Bill) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bill.Table, bill.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bill.CustomerTable, bill.CustomerColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BillClient) Hooks() []Hook {
	return c.hooks.Bill
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Create returns a create builder for Customer.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id int) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CustomerClient) DeleteOneID(id int) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Create returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{config: c.config}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id int) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id int) *Customer {
	cu, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return cu
}

// QueryBuys queries the buys edge of a Customer.
func (c *CustomerClient) QueryBuys(cu *Customer) *BillQuery {
	query := &BillQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.BuysTable, customer.BuysColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Create returns a create builder for Product.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id int) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProductClient) DeleteOneID(id int) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Create returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{config: c.config}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id int) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id int) *Product {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryProducts queries the products edge of a Product.
func (c *ProductClient) QueryProducts(pr *Product) *BillQuery {
	query := &BillQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.ProductsTable, product.ProductsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Create returns a create builder for Role.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Create returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{config: c.config}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	r, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return r
}

// QueryRoles queries the roles edge of a Role.
func (c *RoleClient) QueryRoles(r *Role) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.RolesTable, role.RolesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QuerySells queries the sells edge of a User.
func (c *UserClient) QuerySells(u *User) *BillQuery {
	query := &BillQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SellsTable, user.SellsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRole queries the role edge of a User.
func (c *UserClient) QueryRole(u *User) *RoleQuery {
	query := &RoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.RoleTable, user.RoleColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
