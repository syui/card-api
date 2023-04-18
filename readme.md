### build

```sh
$ vim ent/entc.go
$ vim ent/schema/users.go
$ go generate ./...
$ go build
$ ./card

$ go generate ./...
$ PASS=`cat ./token.json|jq -r .password` TOKEN=`cat ./token.json|jq -r .token` go run -mod=mod main.go
$ curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"syui\",\"password\":\"$pass\"}" localhost:8080/users
$ curl -X POST -H "Content-Type: application/json" -d "{\"owner\":1,\"password\":\"$pass\"}" localhost:8080/cards
$ curl -X POST -H "Content-Type: application/json" -d "{\"owner\":1,\"card\":1,\"cp\":11,\"status\":\"normal\",\"password\":\"$pass\"}" localhost:8080/cards
$ curl localhost:8080/users
$ curl localhost:8080/cards
$ curl localhost:8080/users/1
$ curl localhost:8080/users/1/card
```

### use

```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"username":"syui",\"password\":\"$pass\"}' https://api.syui.ai/users

# onconflict
$ !!

$ curl -sL https://api.syui.ai/users/1
```

```sh
# item select
$ curl -sL "https://api.syui.ai/users?itemsPerPage=255"
$ curl -sL "https://api.syui.ai/cards?itemsPerPage=255"
$ curl -sL "https://api.syui.ai/users/1/card?itemsPerPage=255"
```

### ref

```sh
$ vim ./ent/ogent/ogent.go
// 新規登録の停止
// CreateUsers handles POST /users-slice requests.
func (h *OgentHandler) CreateUsers(ctx context.Context, req CreateUsersReq) (CreateUsersRes, error) {
	b := h.client.Users.Create()
	//b.SetUser(req.User)
	b.SetUser("syui")
}

// 削除の無効
// DeleteUsers handles DELETE /users-slice/{id} requests.
func (h *OgentHandler) DeleteUsers(ctx context.Context, params DeleteUsersParams) (DeleteUsersRes, error) {
	if params.ID != 1 {
err := h.client.Users.DeleteOneID(params.ID).Exec(ctx)
	}
	return new(DeleteUsersNoContent), nil
}

// 要素の書き換えの禁止
// UpdateUsers handles PATCH /users-slice/{id} requests.
func (h *OgentHandler) UpdateUsers(ctx context.Context, req UpdateUsersReq, params UpdateUsersParams) (UpdateUsersRes, error) {
	b := h.client.Users.UpdateOneID(params.ID)
	// Add all fields.
	//if v, ok := req.Hp.Get(); ok {
	//	b.SetHp(v)
	//}
```

### link

- https://entgo.io/ja/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen/

- https://github.com/ariga/ogent/blob/main/example/todo/ent/entc.go

- https://github.com/ent/ent/blob/master/dialect/sql/schema/postgres_test.go

- https://github.com/go-kratos/beer-shop/tree/main/app/catalog/service/internal/data/ent

