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
		MaxLen(100).
		//Match(regexp.MustCompile("[a-z]+$")).
		Unique(),

		field.String("did").
		Optional(),

		field.Bool("member").
		Default(false).
		Optional(),

		field.Bool("book").
		Default(false).
		Optional(),

		field.Bool("manga").
		Default(false).
		Optional(),

		field.Bool("badge").
		Default(false).
		Optional(),

		field.Bool("bsky").
		Default(false).
		Optional(),

		field.Bool("mastodon").
		Default(true).
		Optional(),

		field.Bool("delete").
		Default(false).
		Optional(),

		field.Bool("handle").
		Default(false).
		Optional(),

		field.String("token").
		Optional().
		Sensitive(),

		field.String("password").
		NotEmpty().
		Immutable().
		Sensitive(),

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

		field.Time("raid_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.Time("server_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.Time("egg_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.Int("luck").
		Optional(),

		field.Time("luck_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.Int("like").
		Optional(),

		field.Int("like_rank").
		Optional(),

		field.Time("like_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.Int("fav").
		Optional(),

		field.Bool("ten").
		Optional(),

		field.Int("ten_su").
		Optional(),

		field.Int("ten_kai").
		Optional(),

		field.Int("aiten").
		Optional(),

		field.String("ten_card").
		Optional(),

		field.String("ten_delete").
		Optional(),

		field.String("ten_post").
		Optional(),

		field.String("ten_get").
		Optional(),

		field.Time("ten_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.String("next").
		Default(Nextime()).
		Optional(),

		field.Int("room").
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
