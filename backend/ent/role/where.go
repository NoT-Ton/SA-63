// Code generated by entc, DO NOT EDIT.

package role

import (
	"github.com/NoT-Ton/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RoleName applies equality check predicate on the "role_name" field. It's identical to RoleNameEQ.
func RoleName(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleName), v))
	})
}

// RoleNameEQ applies the EQ predicate on the "role_name" field.
func RoleNameEQ(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleName), v))
	})
}

// RoleNameNEQ applies the NEQ predicate on the "role_name" field.
func RoleNameNEQ(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoleName), v))
	})
}

// RoleNameIn applies the In predicate on the "role_name" field.
func RoleNameIn(vs ...string) predicate.Role {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Role(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoleName), v...))
	})
}

// RoleNameNotIn applies the NotIn predicate on the "role_name" field.
func RoleNameNotIn(vs ...string) predicate.Role {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Role(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoleName), v...))
	})
}

// RoleNameGT applies the GT predicate on the "role_name" field.
func RoleNameGT(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoleName), v))
	})
}

// RoleNameGTE applies the GTE predicate on the "role_name" field.
func RoleNameGTE(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoleName), v))
	})
}

// RoleNameLT applies the LT predicate on the "role_name" field.
func RoleNameLT(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoleName), v))
	})
}

// RoleNameLTE applies the LTE predicate on the "role_name" field.
func RoleNameLTE(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoleName), v))
	})
}

// RoleNameContains applies the Contains predicate on the "role_name" field.
func RoleNameContains(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoleName), v))
	})
}

// RoleNameHasPrefix applies the HasPrefix predicate on the "role_name" field.
func RoleNameHasPrefix(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoleName), v))
	})
}

// RoleNameHasSuffix applies the HasSuffix predicate on the "role_name" field.
func RoleNameHasSuffix(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoleName), v))
	})
}

// RoleNameEqualFold applies the EqualFold predicate on the "role_name" field.
func RoleNameEqualFold(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoleName), v))
	})
}

// RoleNameContainsFold applies the ContainsFold predicate on the "role_name" field.
func RoleNameContainsFold(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoleName), v))
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RolesTable, RolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.User) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RolesTable, RolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		p(s.Not())
	})
}