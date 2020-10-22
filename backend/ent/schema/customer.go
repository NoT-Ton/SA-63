package schema

import(
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)
// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_name").NotEmpty(),
    }
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("buys",Bill.Type).StorageKey(edge.Column("customer_id")),
    }
}
