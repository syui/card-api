package main

import (
	"context"
	"log"
	"net/http"

	"t/ent"
	"t/ent/ogent"
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"time"
	//"github.com/kyokomi/lottery"
	"math/rand"
)

func Random(i int) (l int){
	rand.Seed(time.Now().UnixNano())
	l = rand.Intn(20)
	for l == 0 {
		l = rand.Intn(20)
	}
	return
}

type User struct {
	username string `json:"username"`
	created_at time.Time `json:"created_at"`
}

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
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


func (h handler) DrawStart(ctx context.Context, params ogent.DrawStartParams) error {
	error := h.client.Card.UpdateOneID(params.ID).Exec(ctx)
	return (error)
}

//func (h handler) DrawStart(ctx context.Context, params ogent.DrawStartParams) (ogent.DrawStartNoContent, error) {
//	return ogent.DrawStartNoContent{}, h.client.Users.UpdateOneID(params.ID).Exec(ctx)
//}
//
//func (h handler) DrawDone(ctx context.Context, params ogent.DrawDoneParams) (ogent.DrawDoneNoContent, error) {
func (h handler) DrawDone(ctx context.Context, params ogent.DrawDoneParams) error {
	body := h.client.Card.GetX(ctx, params.ID)
	u_body := h.client.User.GetX(ctx, params.ID)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	t := time.Now().In(jst)
	tt := t.Format("20060102")
	f := body.CreatedAt.Add(time.Hour * 24 * 1).In(jst)
	ff := f.Format("20060102")
	fff := u_body.Next
	if tt < fff {
		error := h.client.User.UpdateOneID(params.ID).SetNext(ff).Exec(ctx)
		return	(error)
	} 
	error := h.client.User.UpdateOneID(params.ID).SetUpdatedAt(t).Exec(ctx)
	return	(error)
}

func main() {
	// Create ent client.
	client, err := ent.Open(dialect.SQLite, "file:/data/ent.sqlite?_fk=1")
	//client, err := ent.Open(dialect.SQLite, "file:data?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
	// Create the handler.
	h := handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}
	// Start listening.
	srv, err := ogent.NewServer(h)
	//srv,err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
