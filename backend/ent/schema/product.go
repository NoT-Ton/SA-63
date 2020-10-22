package schema

import(
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("product_name").NotEmpty(),
		field.Int("product_cost").Positive(),
    }
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("products",Bill.Type).StorageKey(edge.Column("product_id")),
    }
}