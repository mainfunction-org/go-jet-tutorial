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

var Pets = newPetsTable("public", "pets", "")

type petsTable struct {
	postgres.Table

	// Columns
	ID     postgres.ColumnInteger
	Name   postgres.ColumnString
	Status postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PetsTable struct {
	petsTable

	EXCLUDED petsTable
}

// AS creates new PetsTable with assigned alias
func (a PetsTable) AS(alias string) *PetsTable {
	return newPetsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PetsTable with assigned schema name
func (a PetsTable) FromSchema(schemaName string) *PetsTable {
	return newPetsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PetsTable with assigned table prefix
func (a PetsTable) WithPrefix(prefix string) *PetsTable {
	return newPetsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PetsTable with assigned table suffix
func (a PetsTable) WithSuffix(suffix string) *PetsTable {
	return newPetsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPetsTable(schemaName, tableName, alias string) *PetsTable {
	return &PetsTable{
		petsTable: newPetsTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newPetsTableImpl("", "excluded", ""),
	}
}

func newPetsTableImpl(schemaName, tableName, alias string) petsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		NameColumn     = postgres.StringColumn("name")
		StatusColumn   = postgres.StringColumn("status")
		allColumns     = postgres.ColumnList{IDColumn, NameColumn, StatusColumn}
		mutableColumns = postgres.ColumnList{NameColumn, StatusColumn}
	)

	return petsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:     IDColumn,
		Name:   NameColumn,
		Status: StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}