// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhone,
	FieldPassword,
	FieldSex,
	FieldAge,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Sex defines the type for the "sex" enum field.
type Sex string

// Sex values.
const (
	SexMALE   Sex = "MALE"
	SexFEMALE Sex = "FEMALE"
)

func (s Sex) String() string {
	return string(s)
}

// SexValidator is a validator for the "sex" field enum values. It is called by the builders before save.
func SexValidator(s Sex) error {
	switch s {
	case SexMALE, SexFEMALE:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for sex field: %q", s)
	}
}