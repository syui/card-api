package main

import (
"strconv"
	"time"
	"t/ent"
	"net/http"
	"math/rand"
	"context"
	"log"
	"os"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"t/ent/ogent"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/kyokomi/lottery"
)

type User struct {
	user string `json:"user"`
	created_at time.Time `json:"created_at"`
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func Random(i int) (l int){
	rand.Seed(time.Now().UnixNano())
	l = rand.Intn(i)
	for l == 0 {
		l = rand.Intn(i)
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

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}


func (h handler) DrawStart(ctx context.Context, params ogent.DrawStartParams) (ogent.DrawStartNoContent, error) {
	return ogent.DrawStartNoContent{}, h.client.Users.UpdateOneID(params.ID).Exec(ctx)
}

func (h handler) DrawDone(ctx context.Context, params ogent.DrawDoneParams) (ogent.DrawDoneNoContent, error) {
	body := h.client.Users.GetX(ctx, params.ID)
	total := body.Day
	total_n := total + 1
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	t := time.Now().In(jst)
	tt := t.Format("20060102")
	f := body.UpdatedAt.Add(time.Hour * 24 * 1).In(jst)
	ff := f.Format("20060102")
	fff := body.Next
	if tt < fff {
		return	ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetNext(ff).SetLimit(true).Exec(ctx)
	} 

	bb := h.client.Users.GetX(ctx, body.Battle)
	ba := bb.Attack
	aa := body.Attack
	attack_p := aa - ba
	win_n := body.Win + 1
	pat := float64(win_n) / float64(total_n) * 100
	var at int
	if attack_p > 0 {
		at = Random(55)
	} else {
		at = Random(45)
	}
	if at > 25 {
	b := Random(4)
	if b == 1 {
		b := Random(10)
		com := "attack+" + strconv.Itoa(b)
		a := body.Attack + b
		return ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetAttack(a).SetUpdatedAt(t).SetNext(ff).SetDay(total_n).SetWin(win_n).SetPercentage(pat).SetLimit(false).SetComment(com).Exec(ctx)
	} else if b == 2 {
		b := Random(10)
		com := "defense+" + strconv.Itoa(b)
		a := body.Defense + b
		return ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetDefense(a).SetUpdatedAt(t).SetNext(ff).SetDay(total_n).SetWin(win_n).SetPercentage(pat).SetLimit(false).SetComment(com).Exec(ctx)
	} else if b == 3 {
		b := Random(10)
		com := "hp+" + strconv.Itoa(b)
		a := body.Hp + b
		return ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetHp(a).SetUpdatedAt(t).SetNext(ff).SetDay(total_n).SetWin(win_n).SetPercentage(pat).SetComment(com).SetLimit(false).Exec(ctx)
	} else {
		b := Random(10)
		com := "critical+" + strconv.Itoa(b)
		a := body.Critical + b
		return ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetCritical(a).SetUpdatedAt(t).SetNext(ff).SetDay(total_n).SetWin(win_n).SetPercentage(pat).SetComment(com).SetLimit(false).SetWin(win_n).Exec(ctx)
	}
} else {
	com := "loss"
		return ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetNext(ff).SetDay(total_n).SetComment(com).SetLimit(false).Exec(ctx)
}
}

func main() {
	url := os.Getenv("DATABASE_URL") + "?sslmode=require"
	client, err := ent.Open("postgres", url)
	//client, err := Open(url)
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	h := handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}
	srv,err := ogent.NewServer(h)
	//srv,err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":" + port, srv); err != nil {
		log.Fatal(err)
	}
}
