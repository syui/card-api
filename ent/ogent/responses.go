// Code generated by ent, DO NOT EDIT.

package ogent

import "t/ent"

func NewCardCreate(e *ent.Card) *CardCreate {
	if e == nil {
		return nil
	}
	var ret CardCreate
	ret.ID = e.ID
	ret.Card = NewOptInt(e.Card)
	ret.Skill = NewOptString(e.Skill)
	ret.Status = NewOptString(e.Status)
	ret.Cp = NewOptInt(e.Cp)
	ret.URL = NewOptString(e.URL)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	return &ret
}

func NewCardCreates(es []*ent.Card) []CardCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CardCreate, len(es))
	for i, e := range es {
		r[i] = NewCardCreate(e).Elem()
	}
	return r
}

func (c *CardCreate) Elem() CardCreate {
	if c == nil {
		return CardCreate{}
	}
	return *c
}

func NewCardList(e *ent.Card) *CardList {
	if e == nil {
		return nil
	}
	var ret CardList
	ret.ID = e.ID
	ret.Card = NewOptInt(e.Card)
	ret.Skill = NewOptString(e.Skill)
	ret.Status = NewOptString(e.Status)
	ret.Cp = NewOptInt(e.Cp)
	ret.URL = NewOptString(e.URL)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	return &ret
}

func NewCardLists(es []*ent.Card) []CardList {
	if len(es) == 0 {
		return nil
	}
	r := make([]CardList, len(es))
	for i, e := range es {
		r[i] = NewCardList(e).Elem()
	}
	return r
}

func (c *CardList) Elem() CardList {
	if c == nil {
		return CardList{}
	}
	return *c
}

func NewCardRead(e *ent.Card) *CardRead {
	if e == nil {
		return nil
	}
	var ret CardRead
	ret.ID = e.ID
	ret.Card = NewOptInt(e.Card)
	ret.Skill = NewOptString(e.Skill)
	ret.Status = NewOptString(e.Status)
	ret.Cp = NewOptInt(e.Cp)
	ret.URL = NewOptString(e.URL)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	return &ret
}

func NewCardReads(es []*ent.Card) []CardRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]CardRead, len(es))
	for i, e := range es {
		r[i] = NewCardRead(e).Elem()
	}
	return r
}

func (c *CardRead) Elem() CardRead {
	if c == nil {
		return CardRead{}
	}
	return *c
}

func NewCardUpdate(e *ent.Card) *CardUpdate {
	if e == nil {
		return nil
	}
	var ret CardUpdate
	ret.ID = e.ID
	ret.Card = NewOptInt(e.Card)
	ret.Skill = NewOptString(e.Skill)
	ret.Status = NewOptString(e.Status)
	ret.Cp = NewOptInt(e.Cp)
	ret.URL = NewOptString(e.URL)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	return &ret
}

func NewCardUpdates(es []*ent.Card) []CardUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CardUpdate, len(es))
	for i, e := range es {
		r[i] = NewCardUpdate(e).Elem()
	}
	return r
}

func (c *CardUpdate) Elem() CardUpdate {
	if c == nil {
		return CardUpdate{}
	}
	return *c
}

func NewCardOwnerRead(e *ent.User) *CardOwnerRead {
	if e == nil {
		return nil
	}
	var ret CardOwnerRead
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Did = NewOptString(e.Did)
	ret.Member = NewOptBool(e.Member)
	ret.Book = NewOptBool(e.Book)
	ret.Manga = NewOptBool(e.Manga)
	ret.Badge = NewOptBool(e.Badge)
	ret.Bsky = NewOptBool(e.Bsky)
	ret.Mastodon = NewOptBool(e.Mastodon)
	ret.Delete = NewOptBool(e.Delete)
	ret.Handle = NewOptBool(e.Handle)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	ret.UpdatedAt = NewOptDateTime(e.UpdatedAt)
	ret.RaidAt = NewOptDateTime(e.RaidAt)
	ret.EggAt = NewOptDateTime(e.EggAt)
	ret.Luck = NewOptInt(e.Luck)
	ret.LuckAt = NewOptDateTime(e.LuckAt)
	ret.Like = NewOptInt(e.Like)
	ret.LikeRank = NewOptInt(e.LikeRank)
	ret.LikeAt = NewOptDateTime(e.LikeAt)
	ret.Fav = NewOptInt(e.Fav)
	ret.Ten = NewOptBool(e.Ten)
	ret.TenSu = NewOptInt(e.TenSu)
	ret.TenKai = NewOptInt(e.TenKai)
	ret.Aiten = NewOptInt(e.Aiten)
	ret.TenCard = NewOptString(e.TenCard)
	ret.TenDelete = NewOptString(e.TenDelete)
	ret.TenPost = NewOptString(e.TenPost)
	ret.TenGet = NewOptString(e.TenGet)
	ret.TenAt = NewOptDateTime(e.TenAt)
	ret.Next = NewOptString(e.Next)
	return &ret
}

func NewCardOwnerReads(es []*ent.User) []CardOwnerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]CardOwnerRead, len(es))
	for i, e := range es {
		r[i] = NewCardOwnerRead(e).Elem()
	}
	return r
}

func (u *CardOwnerRead) Elem() CardOwnerRead {
	if u == nil {
		return CardOwnerRead{}
	}
	return *u
}

func NewGroupCreate(e *ent.Group) *GroupCreate {
	if e == nil {
		return nil
	}
	var ret GroupCreate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupCreates(es []*ent.Group) []GroupCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupCreate, len(es))
	for i, e := range es {
		r[i] = NewGroupCreate(e).Elem()
	}
	return r
}

func (gr *GroupCreate) Elem() GroupCreate {
	if gr == nil {
		return GroupCreate{}
	}
	return *gr
}

func NewGroupList(e *ent.Group) *GroupList {
	if e == nil {
		return nil
	}
	var ret GroupList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupLists(es []*ent.Group) []GroupList {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupList, len(es))
	for i, e := range es {
		r[i] = NewGroupList(e).Elem()
	}
	return r
}

func (gr *GroupList) Elem() GroupList {
	if gr == nil {
		return GroupList{}
	}
	return *gr
}

func NewGroupRead(e *ent.Group) *GroupRead {
	if e == nil {
		return nil
	}
	var ret GroupRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupReads(es []*ent.Group) []GroupRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupRead, len(es))
	for i, e := range es {
		r[i] = NewGroupRead(e).Elem()
	}
	return r
}

func (gr *GroupRead) Elem() GroupRead {
	if gr == nil {
		return GroupRead{}
	}
	return *gr
}

func NewGroupUpdate(e *ent.Group) *GroupUpdate {
	if e == nil {
		return nil
	}
	var ret GroupUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupUpdates(es []*ent.Group) []GroupUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupUpdate, len(es))
	for i, e := range es {
		r[i] = NewGroupUpdate(e).Elem()
	}
	return r
}

func (gr *GroupUpdate) Elem() GroupUpdate {
	if gr == nil {
		return GroupUpdate{}
	}
	return *gr
}

func NewGroupUsersList(e *ent.User) *GroupUsersList {
	if e == nil {
		return nil
	}
	var ret GroupUsersList
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Did = NewOptString(e.Did)
	ret.Member = NewOptBool(e.Member)
	ret.Book = NewOptBool(e.Book)
	ret.Manga = NewOptBool(e.Manga)
	ret.Badge = NewOptBool(e.Badge)
	ret.Bsky = NewOptBool(e.Bsky)
	ret.Mastodon = NewOptBool(e.Mastodon)
	ret.Delete = NewOptBool(e.Delete)
	ret.Handle = NewOptBool(e.Handle)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	ret.UpdatedAt = NewOptDateTime(e.UpdatedAt)
	ret.RaidAt = NewOptDateTime(e.RaidAt)
	ret.EggAt = NewOptDateTime(e.EggAt)
	ret.Luck = NewOptInt(e.Luck)
	ret.LuckAt = NewOptDateTime(e.LuckAt)
	ret.Like = NewOptInt(e.Like)
	ret.LikeRank = NewOptInt(e.LikeRank)
	ret.LikeAt = NewOptDateTime(e.LikeAt)
	ret.Fav = NewOptInt(e.Fav)
	ret.Ten = NewOptBool(e.Ten)
	ret.TenSu = NewOptInt(e.TenSu)
	ret.TenKai = NewOptInt(e.TenKai)
	ret.Aiten = NewOptInt(e.Aiten)
	ret.TenCard = NewOptString(e.TenCard)
	ret.TenDelete = NewOptString(e.TenDelete)
	ret.TenPost = NewOptString(e.TenPost)
	ret.TenGet = NewOptString(e.TenGet)
	ret.TenAt = NewOptDateTime(e.TenAt)
	ret.Next = NewOptString(e.Next)
	return &ret
}

func NewGroupUsersLists(es []*ent.User) []GroupUsersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupUsersList, len(es))
	for i, e := range es {
		r[i] = NewGroupUsersList(e).Elem()
	}
	return r
}

func (u *GroupUsersList) Elem() GroupUsersList {
	if u == nil {
		return GroupUsersList{}
	}
	return *u
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	var ret UserCreate
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Did = NewOptString(e.Did)
	ret.Member = NewOptBool(e.Member)
	ret.Book = NewOptBool(e.Book)
	ret.Manga = NewOptBool(e.Manga)
	ret.Badge = NewOptBool(e.Badge)
	ret.Bsky = NewOptBool(e.Bsky)
	ret.Mastodon = NewOptBool(e.Mastodon)
	ret.Delete = NewOptBool(e.Delete)
	ret.Handle = NewOptBool(e.Handle)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	ret.UpdatedAt = NewOptDateTime(e.UpdatedAt)
	ret.RaidAt = NewOptDateTime(e.RaidAt)
	ret.EggAt = NewOptDateTime(e.EggAt)
	ret.Luck = NewOptInt(e.Luck)
	ret.LuckAt = NewOptDateTime(e.LuckAt)
	ret.Like = NewOptInt(e.Like)
	ret.LikeRank = NewOptInt(e.LikeRank)
	ret.LikeAt = NewOptDateTime(e.LikeAt)
	ret.Fav = NewOptInt(e.Fav)
	ret.Ten = NewOptBool(e.Ten)
	ret.TenSu = NewOptInt(e.TenSu)
	ret.TenKai = NewOptInt(e.TenKai)
	ret.Aiten = NewOptInt(e.Aiten)
	ret.TenCard = NewOptString(e.TenCard)
	ret.TenDelete = NewOptString(e.TenDelete)
	ret.TenPost = NewOptString(e.TenPost)
	ret.TenGet = NewOptString(e.TenGet)
	ret.TenAt = NewOptDateTime(e.TenAt)
	ret.Next = NewOptString(e.Next)
	return &ret
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u == nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	var ret UserList
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Did = NewOptString(e.Did)
	ret.Member = NewOptBool(e.Member)
	ret.Book = NewOptBool(e.Book)
	ret.Manga = NewOptBool(e.Manga)
	ret.Badge = NewOptBool(e.Badge)
	ret.Bsky = NewOptBool(e.Bsky)
	ret.Mastodon = NewOptBool(e.Mastodon)
	ret.Delete = NewOptBool(e.Delete)
	ret.Handle = NewOptBool(e.Handle)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	ret.UpdatedAt = NewOptDateTime(e.UpdatedAt)
	ret.RaidAt = NewOptDateTime(e.RaidAt)
	ret.EggAt = NewOptDateTime(e.EggAt)
	ret.Luck = NewOptInt(e.Luck)
	ret.LuckAt = NewOptDateTime(e.LuckAt)
	ret.Like = NewOptInt(e.Like)
	ret.LikeRank = NewOptInt(e.LikeRank)
	ret.LikeAt = NewOptDateTime(e.LikeAt)
	ret.Fav = NewOptInt(e.Fav)
	ret.Ten = NewOptBool(e.Ten)
	ret.TenSu = NewOptInt(e.TenSu)
	ret.TenKai = NewOptInt(e.TenKai)
	ret.Aiten = NewOptInt(e.Aiten)
	ret.TenCard = NewOptString(e.TenCard)
	ret.TenDelete = NewOptString(e.TenDelete)
	ret.TenPost = NewOptString(e.TenPost)
	ret.TenGet = NewOptString(e.TenGet)
	ret.TenAt = NewOptDateTime(e.TenAt)
	ret.Next = NewOptString(e.Next)
	return &ret
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u == nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	var ret UserRead
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Did = NewOptString(e.Did)
	ret.Member = NewOptBool(e.Member)
	ret.Book = NewOptBool(e.Book)
	ret.Manga = NewOptBool(e.Manga)
	ret.Badge = NewOptBool(e.Badge)
	ret.Bsky = NewOptBool(e.Bsky)
	ret.Mastodon = NewOptBool(e.Mastodon)
	ret.Delete = NewOptBool(e.Delete)
	ret.Handle = NewOptBool(e.Handle)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	ret.UpdatedAt = NewOptDateTime(e.UpdatedAt)
	ret.RaidAt = NewOptDateTime(e.RaidAt)
	ret.EggAt = NewOptDateTime(e.EggAt)
	ret.Luck = NewOptInt(e.Luck)
	ret.LuckAt = NewOptDateTime(e.LuckAt)
	ret.Like = NewOptInt(e.Like)
	ret.LikeRank = NewOptInt(e.LikeRank)
	ret.LikeAt = NewOptDateTime(e.LikeAt)
	ret.Fav = NewOptInt(e.Fav)
	ret.Ten = NewOptBool(e.Ten)
	ret.TenSu = NewOptInt(e.TenSu)
	ret.TenKai = NewOptInt(e.TenKai)
	ret.Aiten = NewOptInt(e.Aiten)
	ret.TenCard = NewOptString(e.TenCard)
	ret.TenDelete = NewOptString(e.TenDelete)
	ret.TenPost = NewOptString(e.TenPost)
	ret.TenGet = NewOptString(e.TenGet)
	ret.TenAt = NewOptDateTime(e.TenAt)
	ret.Next = NewOptString(e.Next)
	return &ret
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u == nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	var ret UserUpdate
	ret.ID = e.ID
	ret.Username = e.Username
	ret.Did = NewOptString(e.Did)
	ret.Member = NewOptBool(e.Member)
	ret.Book = NewOptBool(e.Book)
	ret.Manga = NewOptBool(e.Manga)
	ret.Badge = NewOptBool(e.Badge)
	ret.Bsky = NewOptBool(e.Bsky)
	ret.Mastodon = NewOptBool(e.Mastodon)
	ret.Delete = NewOptBool(e.Delete)
	ret.Handle = NewOptBool(e.Handle)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	ret.UpdatedAt = NewOptDateTime(e.UpdatedAt)
	ret.RaidAt = NewOptDateTime(e.RaidAt)
	ret.EggAt = NewOptDateTime(e.EggAt)
	ret.Luck = NewOptInt(e.Luck)
	ret.LuckAt = NewOptDateTime(e.LuckAt)
	ret.Like = NewOptInt(e.Like)
	ret.LikeRank = NewOptInt(e.LikeRank)
	ret.LikeAt = NewOptDateTime(e.LikeAt)
	ret.Fav = NewOptInt(e.Fav)
	ret.Ten = NewOptBool(e.Ten)
	ret.TenSu = NewOptInt(e.TenSu)
	ret.TenKai = NewOptInt(e.TenKai)
	ret.Aiten = NewOptInt(e.Aiten)
	ret.TenCard = NewOptString(e.TenCard)
	ret.TenDelete = NewOptString(e.TenDelete)
	ret.TenPost = NewOptString(e.TenPost)
	ret.TenGet = NewOptString(e.TenGet)
	ret.TenAt = NewOptDateTime(e.TenAt)
	ret.Next = NewOptString(e.Next)
	return &ret
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u == nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserCardList(e *ent.Card) *UserCardList {
	if e == nil {
		return nil
	}
	var ret UserCardList
	ret.ID = e.ID
	ret.Card = NewOptInt(e.Card)
	ret.Skill = NewOptString(e.Skill)
	ret.Status = NewOptString(e.Status)
	ret.Cp = NewOptInt(e.Cp)
	ret.URL = NewOptString(e.URL)
	ret.CreatedAt = NewOptDateTime(e.CreatedAt)
	return &ret
}

func NewUserCardLists(es []*ent.Card) []UserCardList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCardList, len(es))
	for i, e := range es {
		r[i] = NewUserCardList(e).Elem()
	}
	return r
}

func (c *UserCardList) Elem() UserCardList {
	if c == nil {
		return UserCardList{}
	}
	return *c
}
