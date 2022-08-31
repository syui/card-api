// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateCard implements createCard operation.
//
// Creates a new Card and persists it to storage.
//
// POST /cards
func (UnimplementedHandler) CreateCard(ctx context.Context, req *CreateCardReq) (r CreateCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateGroup implements createGroup operation.
//
// Creates a new Group and persists it to storage.
//
// POST /groups
func (UnimplementedHandler) CreateGroup(ctx context.Context, req *CreateGroupReq) (r CreateGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUser implements createUser operation.
//
// Creates a new User and persists it to storage.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *CreateUserReq) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCard implements deleteCard operation.
//
// Deletes the Card with the requested ID.
//
// DELETE /cards/{id}
func (UnimplementedHandler) DeleteCard(ctx context.Context, params DeleteCardParams) (r DeleteCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteGroup implements deleteGroup operation.
//
// Deletes the Group with the requested ID.
//
// DELETE /groups/{id}
func (UnimplementedHandler) DeleteGroup(ctx context.Context, params DeleteGroupParams) (r DeleteGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteUser implements deleteUser operation.
//
// Deletes the User with the requested ID.
//
// DELETE /users/{id}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DrawDone implements drawDone operation.
//
// Draws a card item as done.
//
// PUT /cards/{id}/d
func (UnimplementedHandler) DrawDone(ctx context.Context, params DrawDoneParams) error {
	return ht.ErrNotImplemented
}

// DrawStart implements drawStart operation.
//
// Draws a card item as done.
//
// PATCH /users/{id}/card/start
func (UnimplementedHandler) DrawStart(ctx context.Context, params DrawStartParams) error {
	return ht.ErrNotImplemented
}

// ListCard implements listCard operation.
//
// List Cards.
//
// GET /cards
func (UnimplementedHandler) ListCard(ctx context.Context, params ListCardParams) (r ListCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGroup implements listGroup operation.
//
// List Groups.
//
// GET /groups
func (UnimplementedHandler) ListGroup(ctx context.Context, params ListGroupParams) (r ListGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGroupUsers implements listGroupUsers operation.
//
// List attached Users.
//
// GET /groups/{id}/users
func (UnimplementedHandler) ListGroupUsers(ctx context.Context, params ListGroupUsersParams) (r ListGroupUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUser implements listUser operation.
//
// List Users.
//
// GET /users
func (UnimplementedHandler) ListUser(ctx context.Context, params ListUserParams) (r ListUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserCard implements listUserCard operation.
//
// List attached Cards.
//
// GET /users/{id}/card
func (UnimplementedHandler) ListUserCard(ctx context.Context, params ListUserCardParams) (r ListUserCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCard implements readCard operation.
//
// Finds the Card with the requested ID and returns it.
//
// GET /cards/{id}
func (UnimplementedHandler) ReadCard(ctx context.Context, params ReadCardParams) (r ReadCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCardOwner implements readCardOwner operation.
//
// Find the attached User of the Card with the given ID.
//
// GET /cards/{id}/owner
func (UnimplementedHandler) ReadCardOwner(ctx context.Context, params ReadCardOwnerParams) (r ReadCardOwnerRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadGroup implements readGroup operation.
//
// Finds the Group with the requested ID and returns it.
//
// GET /groups/{id}
func (UnimplementedHandler) ReadGroup(ctx context.Context, params ReadGroupParams) (r ReadGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadUser implements readUser operation.
//
// Finds the User with the requested ID and returns it.
//
// GET /users/{id}
func (UnimplementedHandler) ReadUser(ctx context.Context, params ReadUserParams) (r ReadUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCard implements updateCard operation.
//
// Updates a Card and persists changes to storage.
//
// PATCH /cards/{id}
func (UnimplementedHandler) UpdateCard(ctx context.Context, req *UpdateCardReq, params UpdateCardParams) (r UpdateCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateGroup implements updateGroup operation.
//
// Updates a Group and persists changes to storage.
//
// PATCH /groups/{id}
func (UnimplementedHandler) UpdateGroup(ctx context.Context, req *UpdateGroupReq, params UpdateGroupParams) (r UpdateGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// Updates a User and persists changes to storage.
//
// PATCH /users/{id}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}
