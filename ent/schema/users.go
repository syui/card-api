package schema

import (
	"time"
	"regexp"
	"entgo.io/ent"
	"math/rand"
	"entgo.io/ent/schema/field"
	"github.com/kyokomi/lottery"
)

type Users struct {
	ent.Schema
}

func Random(i int) (l int){
	rand.Seed(time.Now().UnixNano())
	l = rand.Intn(20)
	for l == 0 {
		l = rand.Intn(20)
	}
	return
}

func Kira(i int) (l bool){
	lot := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	if lot.Lot(i) {
		l = true
	} else {
		l = false
	}
	return
}

var jst,err = time.LoadLocation("Asia/Tokyo")

var a = Random(4)
var d = Random(20)
var h = Random(20)
var k = Random(20)
var s = Random(20)
var c = Random(20)
var battle = 2

func StatusR()(status string){
	if d == 1 {
		var status = "super"
		return status
	} else {
		var status = "normal"
		return status
	}
	return
}

func CharaR()(chara string){
if a == 1 {
	var chara = "drai"
return chara
} else if a == 2 {
	var chara = "octkat"
return chara
} else if a == 3 {
	var chara = "kyosuke"
return chara
} else {
	var chara = "ponta"
return chara
}
return
}

func Nextime() (tt string){
 t := time.Now().Add(time.Hour * 24 * 1).In(jst)
	tt = t.Format("20060102")
	return
}

func (Users) Fields() []ent.Field {
	return []ent.Field{

		field.String("user").
		NotEmpty().
		Immutable().
		MaxLen(7).
		Match(regexp.MustCompile("[a-z]+$")).
		Unique(),

		field.String("chara").
		Immutable().
		Default(CharaR()).
		Optional(),

		field.Int("skill").
		Immutable().
		Optional().
		Default(k),

		field.Int("hp").
		Optional().
		Default(h),

		field.Int("attack").
		Optional().
		Default(a),

		field.Int("defense").
		Optional().
		Default(d),

		field.Int("critical").
		Optional().
		Default(c),

		field.Int("battle").
		Optional().
		Default(1),

		field.Int("win").
		Optional().
		Default(0),

		field.Int("day").
		Optional().
		Default(0),

		field.Float("percentage").
		Optional().
		Default(0),

		field.Bool("limit").
		Optional().
		Default(false),

		field.String("status").
		Immutable().
		Optional().
		Default(StatusR()),

		field.String("comment").
		Optional().
		Default(""),

		field.Time("created_at").
		Immutable().
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.String("next").
		Default(Nextime()).
		Optional(),

		field.Time("updated_at").
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),

		field.String("url").
		Immutable().
		Optional().
		Default("https://syui.cf/api"),
	}

}

func (Users) Edges() []ent.Edge {
	return []ent.Edge{}
}

