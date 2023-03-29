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

func Random(i int) (l int){
	rand.Seed(time.Now().UnixNano())
	l = rand.Intn(20)
	for l == 0 {
		l = rand.Intn(20)
	}
	return
}

func CardR()(card int){
	var a = Random(10)
	if a == 1 {
		card = Random(12)
		return card
	} else {
		card = 0
		return card
	}
	return
}

var card = CardR()

func SuperR()(super string){
	var b = Random(100)
	if card == 0 || b != 1 {
		super = "normal"
		return super
	} else {
		if b == 1 {
			super = "super"
			return super
		}
		return
	}
}
var super = SuperR()

func CpR()(cp int){
	if super == "super" {
		var cp = Random(1000)
		return cp
	} else if card != 0 {
		var cp = Random(100)
		return cp
	} else {
		var cp = Random(20)
		return cp
	}
	return
}

var cp = CpR()

func (Card) Fields() []ent.Field {
	return []ent.Field{  
		field.Int("card").
		Immutable().
		//Default(card).
		DefaultFunc(func() int {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(10)
			if a == 1 {
				card = rand.Intn(12)
			} else {
				card = 0
			}
			return card
		}).
		Optional(),

		field.Int("cp").
		Immutable().
		//Default(cp).
		DefaultFunc(func() int {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(100)
			if a == 1 {
				a = rand.Intn(1000)
			}
			return a
		}).
		Optional(),

		field.String("status").
		Immutable().
		//Default(super).
		DefaultFunc(func() string {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(20)
			if a == 1 {
				super = "super"
			} else {
				super = "normal"
			}
			return super
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
