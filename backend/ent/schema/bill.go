package schema

import(
	"time"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Positive(),
		field.Time("added_time").Default(time.Now),
    }
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product",Product.Type).Ref("products").Unique(),
		edge.From("user",User.Type).Ref("sells").Unique(),
		edge.From("customer",Customer.Type).Ref("buys").Unique(),
	}
}