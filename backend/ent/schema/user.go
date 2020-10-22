package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// User holds the schema definition for the User entity.
type User struct {
   ent.Schema
}
 
// Fields of the User.
func (User) Fields() []ent.Field {
   return []ent.Field{
		field.String("user_email").NotEmpty(),
		field.String("user_name").NotEmpty(),
   }
}
 
// Edges of the User.
func (User) Edges() []ent.Edge {
   return []ent.Edge{
      edge.To("sells",Bill.Type).StorageKey(edge.Column("user_id")),
      edge.From("role",Role.Type).Ref("roles").Unique(),
  }
}
