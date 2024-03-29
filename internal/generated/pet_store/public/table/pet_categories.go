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

var PetCategories = newPetCategoriesTable("public", "pet_categories", "")

type petCategoriesTable struct {
	postgres.Table

	// Columns
	PetID      postgres.ColumnInteger
	CategoryID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PetCategoriesTable struct {
	petCategoriesTable

	EXCLUDED petCategoriesTable
}

// AS creates new PetCategoriesTable with assigned alias
func (a PetCategoriesTable) AS(alias string) *PetCategoriesTable {
	return newPetCategoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PetCategoriesTable with assigned schema name
func (a PetCategoriesTable) FromSchema(schemaName string) *PetCategoriesTable {
	return newPetCategoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PetCategoriesTable with assigned table prefix
func (a PetCategoriesTable) WithPrefix(prefix string) *PetCategoriesTable {
	return newPetCategoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PetCategoriesTable with assigned table suffix
func (a PetCategoriesTable) WithSuffix(suffix string) *PetCategoriesTable {
	return newPetCategoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPetCategoriesTable(schemaName, tableName, alias string) *PetCategoriesTable {
	return &PetCategoriesTable{
		petCategoriesTable: newPetCategoriesTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newPetCategoriesTableImpl("", "excluded", ""),
	}
}

func newPetCategoriesTableImpl(schemaName, tableName, alias string) petCategoriesTable {
	var (
		PetIDColumn      = postgres.IntegerColumn("pet_id")
		CategoryIDColumn = postgres.IntegerColumn("category_id")
		allColumns       = postgres.ColumnList{PetIDColumn, CategoryIDColumn}
		mutableColumns   = postgres.ColumnList{PetIDColumn, CategoryIDColumn}
	)

	return petCategoriesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		PetID:      PetIDColumn,
		CategoryID: CategoryIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
