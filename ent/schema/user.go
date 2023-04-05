package schema

import (
	"time"
	//"regexp"
	//"math/rand"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	//"entgo.io/ent/schema/mixin"
)

var jst,err = time.LoadLocation("Asia/Tokyo")

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

func Nextime() (tt string){
 t := time.Now().In(jst)
 //t := time.Now().Add(time.Hour * 24 * 1).In(jst)
	tt = t.Format("20060102")
	return
}

func (User) Fields() []ent.Field {
	return []ent.Field{

		field.String("username").
		NotEmpty().
		Immutable().
		MaxLen(30).
		//Match(regexp.MustCompile("[a-z]+$")).
		Unique(),

		field.String("password").
		NotEmpty().
		Immutable().
		Sensitive(),

		//field.Bool("limit").
		//Optional().
		//Default(false),

		field.Time("created_at").
		Immutable().
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.Time("updated_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.String("next").
		Default(Nextime()).
		Optional(),

	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("username").Unique(),
    }
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type),
			//Unique(),
	}
}
