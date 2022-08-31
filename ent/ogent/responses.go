// Code generated by entc, DO NOT EDIT.

package ogent

import "t/ent"

func NewUsersCreate(e *ent.Users) *UsersCreate {
	if e == nil {
		return nil
	}
	return &UsersCreate{
		ID:         e.ID,
		User:       e.User,
		Chara:      NewOptString(e.Chara),
		Skill:      NewOptInt(e.Skill),
		Hp:         NewOptInt(e.Hp),
		Attack:     NewOptInt(e.Attack),
		Defense:    NewOptInt(e.Defense),
		Critical:   NewOptInt(e.Critical),
		Battle:     NewOptInt(e.Battle),
		Win:        NewOptInt(e.Win),
		Day:        NewOptInt(e.Day),
		Percentage: NewOptFloat64(e.Percentage),
		Limit:      NewOptBool(e.Limit),
		Status:     NewOptString(e.Status),
		Comment:    NewOptString(e.Comment),
		CreatedAt:  NewOptDateTime(e.CreatedAt),
		Next:       NewOptString(e.Next),
		UpdatedAt:  NewOptDateTime(e.UpdatedAt),
		URL:        NewOptString(e.URL),
	}
}

func NewUsersCreates(es []*ent.Users) []UsersCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersCreate, len(es))
	for i, e := range es {
		r[i] = NewUsersCreate(e).Elem()
	}
	return r
}

func (u *UsersCreate) Elem() UsersCreate {
	if u != nil {
		return UsersCreate{}
	}
	return *u
}

func NewUsersList(e *ent.Users) *UsersList {
	if e == nil {
		return nil
	}
	return &UsersList{
		ID:         e.ID,
		User:       e.User,
		Chara:      NewOptString(e.Chara),
		Skill:      NewOptInt(e.Skill),
		Hp:         NewOptInt(e.Hp),
		Attack:     NewOptInt(e.Attack),
		Defense:    NewOptInt(e.Defense),
		Critical:   NewOptInt(e.Critical),
		Battle:     NewOptInt(e.Battle),
		Win:        NewOptInt(e.Win),
		Day:        NewOptInt(e.Day),
		Percentage: NewOptFloat64(e.Percentage),
		Limit:      NewOptBool(e.Limit),
		Status:     NewOptString(e.Status),
		Comment:    NewOptString(e.Comment),
		CreatedAt:  NewOptDateTime(e.CreatedAt),
		Next:       NewOptString(e.Next),
		UpdatedAt:  NewOptDateTime(e.UpdatedAt),
		URL:        NewOptString(e.URL),
	}
}

func NewUsersLists(es []*ent.Users) []UsersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersList, len(es))
	for i, e := range es {
		r[i] = NewUsersList(e).Elem()
	}
	return r
}

func (u *UsersList) Elem() UsersList {
	if u != nil {
		return UsersList{}
	}
	return *u
}

func NewUsersRead(e *ent.Users) *UsersRead {
	if e == nil {
		return nil
	}
	return &UsersRead{
		ID:         e.ID,
		User:       e.User,
		Chara:      NewOptString(e.Chara),
		Skill:      NewOptInt(e.Skill),
		Hp:         NewOptInt(e.Hp),
		Attack:     NewOptInt(e.Attack),
		Defense:    NewOptInt(e.Defense),
		Critical:   NewOptInt(e.Critical),
		Battle:     NewOptInt(e.Battle),
		Win:        NewOptInt(e.Win),
		Day:        NewOptInt(e.Day),
		Percentage: NewOptFloat64(e.Percentage),
		Limit:      NewOptBool(e.Limit),
		Status:     NewOptString(e.Status),
		Comment:    NewOptString(e.Comment),
		CreatedAt:  NewOptDateTime(e.CreatedAt),
		Next:       NewOptString(e.Next),
		UpdatedAt:  NewOptDateTime(e.UpdatedAt),
		URL:        NewOptString(e.URL),
	}
}

func NewUsersReads(es []*ent.Users) []UsersRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersRead, len(es))
	for i, e := range es {
		r[i] = NewUsersRead(e).Elem()
	}
	return r
}

func (u *UsersRead) Elem() UsersRead {
	if u != nil {
		return UsersRead{}
	}
	return *u
}

func NewUsersUpdate(e *ent.Users) *UsersUpdate {
	if e == nil {
		return nil
	}
	return &UsersUpdate{
		ID:         e.ID,
		User:       e.User,
		Chara:      NewOptString(e.Chara),
		Skill:      NewOptInt(e.Skill),
		Hp:         NewOptInt(e.Hp),
		Attack:     NewOptInt(e.Attack),
		Defense:    NewOptInt(e.Defense),
		Critical:   NewOptInt(e.Critical),
		Battle:     NewOptInt(e.Battle),
		Win:        NewOptInt(e.Win),
		Day:        NewOptInt(e.Day),
		Percentage: NewOptFloat64(e.Percentage),
		Limit:      NewOptBool(e.Limit),
		Status:     NewOptString(e.Status),
		Comment:    NewOptString(e.Comment),
		CreatedAt:  NewOptDateTime(e.CreatedAt),
		Next:       NewOptString(e.Next),
		UpdatedAt:  NewOptDateTime(e.UpdatedAt),
		URL:        NewOptString(e.URL),
	}
}

func NewUsersUpdates(es []*ent.Users) []UsersUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersUpdate, len(es))
	for i, e := range es {
		r[i] = NewUsersUpdate(e).Elem()
	}
	return r
}

func (u *UsersUpdate) Elem() UsersUpdate {
	if u != nil {
		return UsersUpdate{}
	}
	return *u
}
