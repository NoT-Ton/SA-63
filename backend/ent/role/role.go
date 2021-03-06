// Code generated by entc, DO NOT EDIT.

package role

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleName holds the string denoting the role_name field in the database.
	FieldRoleName = "role_name"

	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"

	// Table holds the table name of the role in the database.
	Table = "roles"
	// RolesTable is the table the holds the roles relation/edge.
	RolesTable = "users"
	// RolesInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RolesInverseTable = "users"
	// RolesColumn is the table column denoting the roles relation/edge.
	RolesColumn = "role_id"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldRoleName,
}

var (
	// RoleNameValidator is a validator for the "role_name" field. It is called by the builders before save.
	RoleNameValidator func(string) error
)
