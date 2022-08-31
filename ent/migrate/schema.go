// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeString, Unique: true, Size: 7},
		{Name: "chara", Type: field.TypeString, Nullable: true, Default: "ponta"},
		{Name: "skill", Type: field.TypeInt, Nullable: true, Default: 19},
		{Name: "hp", Type: field.TypeInt, Nullable: true, Default: 15},
		{Name: "attack", Type: field.TypeInt, Nullable: true, Default: 12},
		{Name: "defense", Type: field.TypeInt, Nullable: true, Default: 5},
		{Name: "critical", Type: field.TypeInt, Nullable: true, Default: 7},
		{Name: "battle", Type: field.TypeInt, Nullable: true, Default: 1},
		{Name: "win", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "day", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "percentage", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "limit", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "status", Type: field.TypeString, Nullable: true, Default: "normal"},
		{Name: "comment", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "next", Type: field.TypeString, Nullable: true, Default: "20220901"},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true, Default: "https://syui.cf/api"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
