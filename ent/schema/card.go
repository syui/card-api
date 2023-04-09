package schema

import (
	"time"
	"math/rand"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}
var url = "https://card.syui.ai"

var card int
var super string
var cp int

func (Card) Fields() []ent.Field {
	return []ent.Field{  

		field.String("password").
		NotEmpty().
		Immutable().
		Sensitive(),

		field.Int("card").
		Immutable().
		DefaultFunc(func() int {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(10)
			if a == 1 {
				card = rand.Intn(16)
			} else {
				card = 0
			}
			if card == 13 {
				card = 14
			}
			return card
		}).
		Optional(),

		field.String("status").
		Immutable().
		DefaultFunc(func() string {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(10)
			if a == 1 {
				super = "super"
			} else {
				super = "normal"
			}
			if card == 0 {
				super = "normal"
			}
			return super
		}).
		Optional(),

		field.Int("cp").
		Immutable().
		DefaultFunc(func() int {
			rand.Seed(time.Now().UnixNano())
			var cp = 1 + rand.Intn(100)
			if cp == 2 {
				cp = 50 + rand.Intn(100)
			}
			if card >= 1 {
				cp = 50 + cp + rand.Intn(200)
			}
			if super == "super" {
				cp = 200 + cp + rand.Intn(300)
			}

			return cp
		}).
		Optional(),

		field.String("url").
		Immutable().
		Default(url).
		Optional(),

		field.Time("created_at").
		Immutable().
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),
	}  
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Unique().
			Required(),

	}
}
