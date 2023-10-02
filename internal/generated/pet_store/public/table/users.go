//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Users = newUsersTable("public", "users", "")

type usersTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	Username  postgres.ColumnString
	FirstName postgres.ColumnString
	LastName  postgres.ColumnString
	Email     postgres.ColumnString
	Password  postgres.ColumnString
	Phone     postgres.ColumnString
	Status    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UsersTable struct {
	usersTable

	EXCLUDED usersTable
}

// AS creates new UsersTable with assigned alias
func (a UsersTable) AS(alias string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersTable with assigned schema name
func (a UsersTable) FromSchema(schemaName string) *UsersTable {
	return newUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersTable with assigned table prefix
func (a UsersTable) WithPrefix(prefix string) *UsersTable {
	return newUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersTable with assigned table suffix
func (a UsersTable) WithSuffix(suffix string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersTable(schemaName, tableName, alias string) *UsersTable {
	return &UsersTable{
		usersTable: newUsersTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newUsersTableImpl("", "excluded", ""),
	}
}

func newUsersTableImpl(schemaName, tableName, alias string) usersTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UsernameColumn  = postgres.StringColumn("username")
		FirstNameColumn = postgres.StringColumn("first_name")
		LastNameColumn  = postgres.StringColumn("last_name")
		EmailColumn     = postgres.StringColumn("email")
		PasswordColumn  = postgres.StringColumn("password")
		PhoneColumn     = postgres.StringColumn("phone")
		StatusColumn    = postgres.StringColumn("status")
		allColumns      = postgres.ColumnList{IDColumn, UsernameColumn, FirstNameColumn, LastNameColumn, EmailColumn, PasswordColumn, PhoneColumn, StatusColumn}
		mutableColumns  = postgres.ColumnList{UsernameColumn, FirstNameColumn, LastNameColumn, EmailColumn, PasswordColumn, PhoneColumn, StatusColumn}
	)

	return usersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Username:  UsernameColumn,
		FirstName: FirstNameColumn,
		LastName:  LastNameColumn,
		Email:     EmailColumn,
		Password:  PasswordColumn,
		Phone:     PhoneColumn,
		Status:    StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}