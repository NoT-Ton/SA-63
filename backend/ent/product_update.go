// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/NoT-Ton/app/ent/bill"
	"github.com/NoT-Ton/app/ent/predicate"
	"github.com/NoT-Ton/app/ent/product"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks      []Hook
	mutation   *ProductMutation
	predicates []predicate.Product
}

// Where adds a new predicate for the builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetProductName sets the product_name field.
func (pu *ProductUpdate) SetProductName(s string) *ProductUpdate {
	pu.mutation.SetProductName(s)
	return pu
}

// SetProductCost sets the product_cost field.
func (pu *ProductUpdate) SetProductCost(i int) *ProductUpdate {
	pu.mutation.ResetProductCost()
	pu.mutation.SetProductCost(i)
	return pu
}

// AddProductCost adds i to product_cost.
func (pu *ProductUpdate) AddProductCost(i int) *ProductUpdate {
	pu.mutation.AddProductCost(i)
	return pu
}

// AddProductIDs adds the products edge to Bill by ids.
func (pu *ProductUpdate) AddProductIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddProductIDs(ids...)
	return pu
}

// AddProducts adds the products edges to Bill.
func (pu *ProductUpdate) AddProducts(b ...*Bill) *ProductUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddProductIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// RemoveProductIDs removes the products edge to Bill by ids.
func (pu *ProductUpdate) RemoveProductIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveProductIDs(ids...)
	return pu
}

// RemoveProducts removes products edges to Bill.
func (pu *ProductUpdate) RemoveProducts(b ...*Bill) *ProductUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.ProductName(); ok {
		if err := product.ProductNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "product_name", err: fmt.Errorf("ent: validator failed for field \"product_name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.ProductCost(); ok {
		if err := product.ProductCostValidator(v); err != nil {
			return 0, &ValidationError{Name: "product_cost", err: fmt.Errorf("ent: validator failed for field \"product_cost\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ProductName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldProductName,
		})
	}
	if value, ok := pu.mutation.ProductCost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductCost,
		})
	}
	if value, ok := pu.mutation.AddedProductCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductCost,
		})
	}
	if nodes := pu.mutation.RemovedProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// SetProductName sets the product_name field.
func (puo *ProductUpdateOne) SetProductName(s string) *ProductUpdateOne {
	puo.mutation.SetProductName(s)
	return puo
}

// SetProductCost sets the product_cost field.
func (puo *ProductUpdateOne) SetProductCost(i int) *ProductUpdateOne {
	puo.mutation.ResetProductCost()
	puo.mutation.SetProductCost(i)
	return puo
}

// AddProductCost adds i to product_cost.
func (puo *ProductUpdateOne) AddProductCost(i int) *ProductUpdateOne {
	puo.mutation.AddProductCost(i)
	return puo
}

// AddProductIDs adds the products edge to Bill by ids.
func (puo *ProductUpdateOne) AddProductIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddProductIDs(ids...)
	return puo
}

// AddProducts adds the products edges to Bill.
func (puo *ProductUpdateOne) AddProducts(b ...*Bill) *ProductUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddProductIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// RemoveProductIDs removes the products edge to Bill by ids.
func (puo *ProductUpdateOne) RemoveProductIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveProductIDs(ids...)
	return puo
}

// RemoveProducts removes products edges to Bill.
func (puo *ProductUpdateOne) RemoveProducts(b ...*Bill) *ProductUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveProductIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	if v, ok := puo.mutation.ProductName(); ok {
		if err := product.ProductNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "product_name", err: fmt.Errorf("ent: validator failed for field \"product_name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.ProductCost(); ok {
		if err := product.ProductCostValidator(v); err != nil {
			return nil, &ValidationError{Name: "product_cost", err: fmt.Errorf("ent: validator failed for field \"product_cost\": %w", err)}
		}
	}

	var (
		err  error
		node *Product
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (pr *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.ProductName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldProductName,
		})
	}
	if value, ok := puo.mutation.ProductCost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductCost,
		})
	}
	if value, ok := puo.mutation.AddedProductCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductCost,
		})
	}
	if nodes := puo.mutation.RemovedProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Product{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}