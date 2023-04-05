package main

import (
	"context"
	"log"
	"net/http"

	"t/ent"
	"t/ent/ogent"
	//"t/ent/migrate"
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"entgo.io/ent/dialect/sql/schema"
	"time"
)

type User struct {
	username string `json:"username"`
	created_at time.Time `json:"created_at"`
}

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (h handler) DrawStart(ctx context.Context, params ogent.DrawStartParams) error {
	error := h.client.Card.UpdateOneID(params.ID).Exec(ctx)
	return (error)
}

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
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatal(err)
	//}
	ctx := context.Background()
	err = client.Schema.Create(ctx, schema.WithAtlas(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
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
