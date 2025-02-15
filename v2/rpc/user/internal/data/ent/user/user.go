// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldUpdated holds the string denoting the updated field in the database.
	FieldUpdated = "updated"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldExtra holds the string denoting the extra field in the database.
	FieldExtra = "extra"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreated,
	FieldUpdated,
	FieldNickName,
	FieldSex,
	FieldPhone,
	FieldEmail,
	FieldExtra,
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

var (
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated time.Time
	// DefaultUpdated holds the default value on creation for the "updated" field.
	DefaultUpdated time.Time
	// UpdateDefaultUpdated holds the default value on update for the "updated" field.
	UpdateDefaultUpdated func() time.Time
	// NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	NickNameValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
