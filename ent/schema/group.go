package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
)

// Group holds the schema definition for the Group entity.
type Group struct {
    ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
    }
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("users", User.Type),
    }
}

// Indexes of the Group.
func (Group) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("name").Unique(),
    }
}

