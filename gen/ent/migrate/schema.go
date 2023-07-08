// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CollectionsColumns holds the columns for the "collections" table.
	CollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "index_type", Type: field.TypeString},
		{Name: "data_type", Type: field.TypeString},
		{Name: "embedder", Type: field.TypeString},
		{Name: "index_params", Type: field.TypeJSON},
	}
	// CollectionsTable holds the schema information for the "collections" table.
	CollectionsTable = &schema.Table{
		Name:       "collections",
		Columns:    CollectionsColumns,
		PrimaryKey: []*schema.Column{CollectionsColumns[0]},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "file_name", Type: field.TypeString, Unique: true},
		{Name: "collection_files", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_collections_files",
				Columns:    []*schema.Column{FilesColumns[2]},
				RefColumns: []*schema.Column{CollectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CollectionsTable,
		FilesTable,
	}
)

func init() {
	FilesTable.ForeignKeys[0].RefTable = CollectionsTable
}