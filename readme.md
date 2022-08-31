heroku open-api ent example

- go-module-name : t

- onconflict

```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"user":"syui"}' api.syui.cf/users
...ok

$ !!
...err

$ heroku logs
```


```sh
# delete
$ curl -X DELETE https://api.syui.cf/users/1

# card draw
$ curl -X PUT api.syui.cf/users/1/d
$ curl api.syui.cf/users/1

# patch
$ curl -X PATCH -H "Content-Type: application/json" -d '{"battle":2}' api.syui.cf/users/1 
$ curl -X PATCH -H "Content-Type: application/json" -d '{"limit":false}' api.syui.cf/users/1 
$ d=`date "+%Y%m%d"`
$ curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$d\"}" api.syui.cf/users/1 
```

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

ref : 

- https://github.com/ent/ent/blob/master/dialect/sql/schema/postgres_test.go

- https://github.com/go-kratos/beer-shop/tree/main/app/catalog/service/internal/data/ent

- https://entgo.io/ja/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen/

- https://github.com/ariga/ogent/blob/main/example/todo/ent/entc.go

```sh
$ vim ent/entc.go
$ vim ent/schema/users.go
$ go generate ./...
$ go build
$ ./t
```
