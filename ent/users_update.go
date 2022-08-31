// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"t/ent/predicate"
	"t/ent/users"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersUpdate is the builder for updating Users entities.
type UsersUpdate struct {
	config
	hooks    []Hook
	mutation *UsersMutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uu *UsersUpdate) Where(ps ...predicate.Users) *UsersUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetChara sets the "chara" field.
func (uu *UsersUpdate) SetChara(s string) *UsersUpdate {
	uu.mutation.SetChara(s)
	return uu
}

// SetNillableChara sets the "chara" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableChara(s *string) *UsersUpdate {
	if s != nil {
		uu.SetChara(*s)
	}
	return uu
}

// ClearChara clears the value of the "chara" field.
func (uu *UsersUpdate) ClearChara() *UsersUpdate {
	uu.mutation.ClearChara()
	return uu
}

// SetHp sets the "hp" field.
func (uu *UsersUpdate) SetHp(i int) *UsersUpdate {
	uu.mutation.ResetHp()
	uu.mutation.SetHp(i)
	return uu
}

// SetNillableHp sets the "hp" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableHp(i *int) *UsersUpdate {
	if i != nil {
		uu.SetHp(*i)
	}
	return uu
}

// AddHp adds i to the "hp" field.
func (uu *UsersUpdate) AddHp(i int) *UsersUpdate {
	uu.mutation.AddHp(i)
	return uu
}

// ClearHp clears the value of the "hp" field.
func (uu *UsersUpdate) ClearHp() *UsersUpdate {
	uu.mutation.ClearHp()
	return uu
}

// SetAttack sets the "attack" field.
func (uu *UsersUpdate) SetAttack(i int) *UsersUpdate {
	uu.mutation.ResetAttack()
	uu.mutation.SetAttack(i)
	return uu
}

// SetNillableAttack sets the "attack" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableAttack(i *int) *UsersUpdate {
	if i != nil {
		uu.SetAttack(*i)
	}
	return uu
}

// AddAttack adds i to the "attack" field.
func (uu *UsersUpdate) AddAttack(i int) *UsersUpdate {
	uu.mutation.AddAttack(i)
	return uu
}

// ClearAttack clears the value of the "attack" field.
func (uu *UsersUpdate) ClearAttack() *UsersUpdate {
	uu.mutation.ClearAttack()
	return uu
}

// SetDefense sets the "defense" field.
func (uu *UsersUpdate) SetDefense(i int) *UsersUpdate {
	uu.mutation.ResetDefense()
	uu.mutation.SetDefense(i)
	return uu
}

// SetNillableDefense sets the "defense" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableDefense(i *int) *UsersUpdate {
	if i != nil {
		uu.SetDefense(*i)
	}
	return uu
}

// AddDefense adds i to the "defense" field.
func (uu *UsersUpdate) AddDefense(i int) *UsersUpdate {
	uu.mutation.AddDefense(i)
	return uu
}

// ClearDefense clears the value of the "defense" field.
func (uu *UsersUpdate) ClearDefense() *UsersUpdate {
	uu.mutation.ClearDefense()
	return uu
}

// SetCritical sets the "critical" field.
func (uu *UsersUpdate) SetCritical(i int) *UsersUpdate {
	uu.mutation.ResetCritical()
	uu.mutation.SetCritical(i)
	return uu
}

// SetNillableCritical sets the "critical" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableCritical(i *int) *UsersUpdate {
	if i != nil {
		uu.SetCritical(*i)
	}
	return uu
}

// AddCritical adds i to the "critical" field.
func (uu *UsersUpdate) AddCritical(i int) *UsersUpdate {
	uu.mutation.AddCritical(i)
	return uu
}

// ClearCritical clears the value of the "critical" field.
func (uu *UsersUpdate) ClearCritical() *UsersUpdate {
	uu.mutation.ClearCritical()
	return uu
}

// SetBattle sets the "battle" field.
func (uu *UsersUpdate) SetBattle(i int) *UsersUpdate {
	uu.mutation.ResetBattle()
	uu.mutation.SetBattle(i)
	return uu
}

// SetNillableBattle sets the "battle" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableBattle(i *int) *UsersUpdate {
	if i != nil {
		uu.SetBattle(*i)
	}
	return uu
}

// AddBattle adds i to the "battle" field.
func (uu *UsersUpdate) AddBattle(i int) *UsersUpdate {
	uu.mutation.AddBattle(i)
	return uu
}

// ClearBattle clears the value of the "battle" field.
func (uu *UsersUpdate) ClearBattle() *UsersUpdate {
	uu.mutation.ClearBattle()
	return uu
}

// SetWin sets the "win" field.
func (uu *UsersUpdate) SetWin(i int) *UsersUpdate {
	uu.mutation.ResetWin()
	uu.mutation.SetWin(i)
	return uu
}

// SetNillableWin sets the "win" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableWin(i *int) *UsersUpdate {
	if i != nil {
		uu.SetWin(*i)
	}
	return uu
}

// AddWin adds i to the "win" field.
func (uu *UsersUpdate) AddWin(i int) *UsersUpdate {
	uu.mutation.AddWin(i)
	return uu
}

// ClearWin clears the value of the "win" field.
func (uu *UsersUpdate) ClearWin() *UsersUpdate {
	uu.mutation.ClearWin()
	return uu
}

// SetDay sets the "day" field.
func (uu *UsersUpdate) SetDay(i int) *UsersUpdate {
	uu.mutation.ResetDay()
	uu.mutation.SetDay(i)
	return uu
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableDay(i *int) *UsersUpdate {
	if i != nil {
		uu.SetDay(*i)
	}
	return uu
}

// AddDay adds i to the "day" field.
func (uu *UsersUpdate) AddDay(i int) *UsersUpdate {
	uu.mutation.AddDay(i)
	return uu
}

// ClearDay clears the value of the "day" field.
func (uu *UsersUpdate) ClearDay() *UsersUpdate {
	uu.mutation.ClearDay()
	return uu
}

// SetPercentage sets the "percentage" field.
func (uu *UsersUpdate) SetPercentage(f float64) *UsersUpdate {
	uu.mutation.ResetPercentage()
	uu.mutation.SetPercentage(f)
	return uu
}

// SetNillablePercentage sets the "percentage" field if the given value is not nil.
func (uu *UsersUpdate) SetNillablePercentage(f *float64) *UsersUpdate {
	if f != nil {
		uu.SetPercentage(*f)
	}
	return uu
}

// AddPercentage adds f to the "percentage" field.
func (uu *UsersUpdate) AddPercentage(f float64) *UsersUpdate {
	uu.mutation.AddPercentage(f)
	return uu
}

// ClearPercentage clears the value of the "percentage" field.
func (uu *UsersUpdate) ClearPercentage() *UsersUpdate {
	uu.mutation.ClearPercentage()
	return uu
}

// SetLimit sets the "limit" field.
func (uu *UsersUpdate) SetLimit(b bool) *UsersUpdate {
	uu.mutation.SetLimit(b)
	return uu
}

// SetNillableLimit sets the "limit" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableLimit(b *bool) *UsersUpdate {
	if b != nil {
		uu.SetLimit(*b)
	}
	return uu
}

// ClearLimit clears the value of the "limit" field.
func (uu *UsersUpdate) ClearLimit() *UsersUpdate {
	uu.mutation.ClearLimit()
	return uu
}

// SetComment sets the "comment" field.
func (uu *UsersUpdate) SetComment(s string) *UsersUpdate {
	uu.mutation.SetComment(s)
	return uu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableComment(s *string) *UsersUpdate {
	if s != nil {
		uu.SetComment(*s)
	}
	return uu
}

// ClearComment clears the value of the "comment" field.
func (uu *UsersUpdate) ClearComment() *UsersUpdate {
	uu.mutation.ClearComment()
	return uu
}

// SetNext sets the "next" field.
func (uu *UsersUpdate) SetNext(s string) *UsersUpdate {
	uu.mutation.SetNext(s)
	return uu
}

// SetNillableNext sets the "next" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableNext(s *string) *UsersUpdate {
	if s != nil {
		uu.SetNext(*s)
	}
	return uu
}

// ClearNext clears the value of the "next" field.
func (uu *UsersUpdate) ClearNext() *UsersUpdate {
	uu.mutation.ClearNext()
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UsersUpdate) SetUpdatedAt(t time.Time) *UsersUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableUpdatedAt(t *time.Time) *UsersUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uu *UsersUpdate) ClearUpdatedAt() *UsersUpdate {
	uu.mutation.ClearUpdatedAt()
	return uu
}

// Mutation returns the UsersMutation object of the builder.
func (uu *UsersUpdate) Mutation() *UsersMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UsersUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsersUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsersUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsersUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UsersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   users.Table,
			Columns: users.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: users.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Chara(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldChara,
		})
	}
	if uu.mutation.CharaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldChara,
		})
	}
	if uu.mutation.SkillCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldSkill,
		})
	}
	if value, ok := uu.mutation.Hp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldHp,
		})
	}
	if value, ok := uu.mutation.AddedHp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldHp,
		})
	}
	if uu.mutation.HpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldHp,
		})
	}
	if value, ok := uu.mutation.Attack(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldAttack,
		})
	}
	if value, ok := uu.mutation.AddedAttack(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldAttack,
		})
	}
	if uu.mutation.AttackCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldAttack,
		})
	}
	if value, ok := uu.mutation.Defense(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDefense,
		})
	}
	if value, ok := uu.mutation.AddedDefense(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDefense,
		})
	}
	if uu.mutation.DefenseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldDefense,
		})
	}
	if value, ok := uu.mutation.Critical(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldCritical,
		})
	}
	if value, ok := uu.mutation.AddedCritical(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldCritical,
		})
	}
	if uu.mutation.CriticalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldCritical,
		})
	}
	if value, ok := uu.mutation.Battle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldBattle,
		})
	}
	if value, ok := uu.mutation.AddedBattle(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldBattle,
		})
	}
	if uu.mutation.BattleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldBattle,
		})
	}
	if value, ok := uu.mutation.Win(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldWin,
		})
	}
	if value, ok := uu.mutation.AddedWin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldWin,
		})
	}
	if uu.mutation.WinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldWin,
		})
	}
	if value, ok := uu.mutation.Day(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDay,
		})
	}
	if value, ok := uu.mutation.AddedDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDay,
		})
	}
	if uu.mutation.DayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldDay,
		})
	}
	if value, ok := uu.mutation.Percentage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: users.FieldPercentage,
		})
	}
	if value, ok := uu.mutation.AddedPercentage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: users.FieldPercentage,
		})
	}
	if uu.mutation.PercentageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: users.FieldPercentage,
		})
	}
	if value, ok := uu.mutation.Limit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: users.FieldLimit,
		})
	}
	if uu.mutation.LimitCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: users.FieldLimit,
		})
	}
	if uu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldStatus,
		})
	}
	if value, ok := uu.mutation.Comment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldComment,
		})
	}
	if uu.mutation.CommentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldComment,
		})
	}
	if uu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: users.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.Next(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldNext,
		})
	}
	if uu.mutation.NextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldNext,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldUpdatedAt,
		})
	}
	if uu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: users.FieldUpdatedAt,
		})
	}
	if uu.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldURL,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UsersUpdateOne is the builder for updating a single Users entity.
type UsersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersMutation
}

// SetChara sets the "chara" field.
func (uuo *UsersUpdateOne) SetChara(s string) *UsersUpdateOne {
	uuo.mutation.SetChara(s)
	return uuo
}

// SetNillableChara sets the "chara" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableChara(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetChara(*s)
	}
	return uuo
}

// ClearChara clears the value of the "chara" field.
func (uuo *UsersUpdateOne) ClearChara() *UsersUpdateOne {
	uuo.mutation.ClearChara()
	return uuo
}

// SetHp sets the "hp" field.
func (uuo *UsersUpdateOne) SetHp(i int) *UsersUpdateOne {
	uuo.mutation.ResetHp()
	uuo.mutation.SetHp(i)
	return uuo
}

// SetNillableHp sets the "hp" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableHp(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetHp(*i)
	}
	return uuo
}

// AddHp adds i to the "hp" field.
func (uuo *UsersUpdateOne) AddHp(i int) *UsersUpdateOne {
	uuo.mutation.AddHp(i)
	return uuo
}

// ClearHp clears the value of the "hp" field.
func (uuo *UsersUpdateOne) ClearHp() *UsersUpdateOne {
	uuo.mutation.ClearHp()
	return uuo
}

// SetAttack sets the "attack" field.
func (uuo *UsersUpdateOne) SetAttack(i int) *UsersUpdateOne {
	uuo.mutation.ResetAttack()
	uuo.mutation.SetAttack(i)
	return uuo
}

// SetNillableAttack sets the "attack" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableAttack(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetAttack(*i)
	}
	return uuo
}

// AddAttack adds i to the "attack" field.
func (uuo *UsersUpdateOne) AddAttack(i int) *UsersUpdateOne {
	uuo.mutation.AddAttack(i)
	return uuo
}

// ClearAttack clears the value of the "attack" field.
func (uuo *UsersUpdateOne) ClearAttack() *UsersUpdateOne {
	uuo.mutation.ClearAttack()
	return uuo
}

// SetDefense sets the "defense" field.
func (uuo *UsersUpdateOne) SetDefense(i int) *UsersUpdateOne {
	uuo.mutation.ResetDefense()
	uuo.mutation.SetDefense(i)
	return uuo
}

// SetNillableDefense sets the "defense" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableDefense(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetDefense(*i)
	}
	return uuo
}

// AddDefense adds i to the "defense" field.
func (uuo *UsersUpdateOne) AddDefense(i int) *UsersUpdateOne {
	uuo.mutation.AddDefense(i)
	return uuo
}

// ClearDefense clears the value of the "defense" field.
func (uuo *UsersUpdateOne) ClearDefense() *UsersUpdateOne {
	uuo.mutation.ClearDefense()
	return uuo
}

// SetCritical sets the "critical" field.
func (uuo *UsersUpdateOne) SetCritical(i int) *UsersUpdateOne {
	uuo.mutation.ResetCritical()
	uuo.mutation.SetCritical(i)
	return uuo
}

// SetNillableCritical sets the "critical" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableCritical(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetCritical(*i)
	}
	return uuo
}

// AddCritical adds i to the "critical" field.
func (uuo *UsersUpdateOne) AddCritical(i int) *UsersUpdateOne {
	uuo.mutation.AddCritical(i)
	return uuo
}

// ClearCritical clears the value of the "critical" field.
func (uuo *UsersUpdateOne) ClearCritical() *UsersUpdateOne {
	uuo.mutation.ClearCritical()
	return uuo
}

// SetBattle sets the "battle" field.
func (uuo *UsersUpdateOne) SetBattle(i int) *UsersUpdateOne {
	uuo.mutation.ResetBattle()
	uuo.mutation.SetBattle(i)
	return uuo
}

// SetNillableBattle sets the "battle" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableBattle(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetBattle(*i)
	}
	return uuo
}

// AddBattle adds i to the "battle" field.
func (uuo *UsersUpdateOne) AddBattle(i int) *UsersUpdateOne {
	uuo.mutation.AddBattle(i)
	return uuo
}

// ClearBattle clears the value of the "battle" field.
func (uuo *UsersUpdateOne) ClearBattle() *UsersUpdateOne {
	uuo.mutation.ClearBattle()
	return uuo
}

// SetWin sets the "win" field.
func (uuo *UsersUpdateOne) SetWin(i int) *UsersUpdateOne {
	uuo.mutation.ResetWin()
	uuo.mutation.SetWin(i)
	return uuo
}

// SetNillableWin sets the "win" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableWin(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetWin(*i)
	}
	return uuo
}

// AddWin adds i to the "win" field.
func (uuo *UsersUpdateOne) AddWin(i int) *UsersUpdateOne {
	uuo.mutation.AddWin(i)
	return uuo
}

// ClearWin clears the value of the "win" field.
func (uuo *UsersUpdateOne) ClearWin() *UsersUpdateOne {
	uuo.mutation.ClearWin()
	return uuo
}

// SetDay sets the "day" field.
func (uuo *UsersUpdateOne) SetDay(i int) *UsersUpdateOne {
	uuo.mutation.ResetDay()
	uuo.mutation.SetDay(i)
	return uuo
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableDay(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetDay(*i)
	}
	return uuo
}

// AddDay adds i to the "day" field.
func (uuo *UsersUpdateOne) AddDay(i int) *UsersUpdateOne {
	uuo.mutation.AddDay(i)
	return uuo
}

// ClearDay clears the value of the "day" field.
func (uuo *UsersUpdateOne) ClearDay() *UsersUpdateOne {
	uuo.mutation.ClearDay()
	return uuo
}

// SetPercentage sets the "percentage" field.
func (uuo *UsersUpdateOne) SetPercentage(f float64) *UsersUpdateOne {
	uuo.mutation.ResetPercentage()
	uuo.mutation.SetPercentage(f)
	return uuo
}

// SetNillablePercentage sets the "percentage" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillablePercentage(f *float64) *UsersUpdateOne {
	if f != nil {
		uuo.SetPercentage(*f)
	}
	return uuo
}

// AddPercentage adds f to the "percentage" field.
func (uuo *UsersUpdateOne) AddPercentage(f float64) *UsersUpdateOne {
	uuo.mutation.AddPercentage(f)
	return uuo
}

// ClearPercentage clears the value of the "percentage" field.
func (uuo *UsersUpdateOne) ClearPercentage() *UsersUpdateOne {
	uuo.mutation.ClearPercentage()
	return uuo
}

// SetLimit sets the "limit" field.
func (uuo *UsersUpdateOne) SetLimit(b bool) *UsersUpdateOne {
	uuo.mutation.SetLimit(b)
	return uuo
}

// SetNillableLimit sets the "limit" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableLimit(b *bool) *UsersUpdateOne {
	if b != nil {
		uuo.SetLimit(*b)
	}
	return uuo
}

// ClearLimit clears the value of the "limit" field.
func (uuo *UsersUpdateOne) ClearLimit() *UsersUpdateOne {
	uuo.mutation.ClearLimit()
	return uuo
}

// SetComment sets the "comment" field.
func (uuo *UsersUpdateOne) SetComment(s string) *UsersUpdateOne {
	uuo.mutation.SetComment(s)
	return uuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableComment(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetComment(*s)
	}
	return uuo
}

// ClearComment clears the value of the "comment" field.
func (uuo *UsersUpdateOne) ClearComment() *UsersUpdateOne {
	uuo.mutation.ClearComment()
	return uuo
}

// SetNext sets the "next" field.
func (uuo *UsersUpdateOne) SetNext(s string) *UsersUpdateOne {
	uuo.mutation.SetNext(s)
	return uuo
}

// SetNillableNext sets the "next" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableNext(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetNext(*s)
	}
	return uuo
}

// ClearNext clears the value of the "next" field.
func (uuo *UsersUpdateOne) ClearNext() *UsersUpdateOne {
	uuo.mutation.ClearNext()
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UsersUpdateOne) SetUpdatedAt(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableUpdatedAt(t *time.Time) *UsersUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uuo *UsersUpdateOne) ClearUpdatedAt() *UsersUpdateOne {
	uuo.mutation.ClearUpdatedAt()
	return uuo
}

// Mutation returns the UsersMutation object of the builder.
func (uuo *UsersUpdateOne) Mutation() *UsersMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UsersUpdateOne) Select(field string, fields ...string) *UsersUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Users entity.
func (uuo *UsersUpdateOne) Save(ctx context.Context) (*Users, error) {
	var (
		err  error
		node *Users
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsersUpdateOne) SaveX(ctx context.Context) *Users {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UsersUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsersUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UsersUpdateOne) sqlSave(ctx context.Context) (_node *Users, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   users.Table,
			Columns: users.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: users.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Users.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Chara(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldChara,
		})
	}
	if uuo.mutation.CharaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldChara,
		})
	}
	if uuo.mutation.SkillCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldSkill,
		})
	}
	if value, ok := uuo.mutation.Hp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldHp,
		})
	}
	if value, ok := uuo.mutation.AddedHp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldHp,
		})
	}
	if uuo.mutation.HpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldHp,
		})
	}
	if value, ok := uuo.mutation.Attack(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldAttack,
		})
	}
	if value, ok := uuo.mutation.AddedAttack(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldAttack,
		})
	}
	if uuo.mutation.AttackCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldAttack,
		})
	}
	if value, ok := uuo.mutation.Defense(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDefense,
		})
	}
	if value, ok := uuo.mutation.AddedDefense(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDefense,
		})
	}
	if uuo.mutation.DefenseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldDefense,
		})
	}
	if value, ok := uuo.mutation.Critical(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldCritical,
		})
	}
	if value, ok := uuo.mutation.AddedCritical(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldCritical,
		})
	}
	if uuo.mutation.CriticalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldCritical,
		})
	}
	if value, ok := uuo.mutation.Battle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldBattle,
		})
	}
	if value, ok := uuo.mutation.AddedBattle(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldBattle,
		})
	}
	if uuo.mutation.BattleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldBattle,
		})
	}
	if value, ok := uuo.mutation.Win(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldWin,
		})
	}
	if value, ok := uuo.mutation.AddedWin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldWin,
		})
	}
	if uuo.mutation.WinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldWin,
		})
	}
	if value, ok := uuo.mutation.Day(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDay,
		})
	}
	if value, ok := uuo.mutation.AddedDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: users.FieldDay,
		})
	}
	if uuo.mutation.DayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: users.FieldDay,
		})
	}
	if value, ok := uuo.mutation.Percentage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: users.FieldPercentage,
		})
	}
	if value, ok := uuo.mutation.AddedPercentage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: users.FieldPercentage,
		})
	}
	if uuo.mutation.PercentageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: users.FieldPercentage,
		})
	}
	if value, ok := uuo.mutation.Limit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: users.FieldLimit,
		})
	}
	if uuo.mutation.LimitCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: users.FieldLimit,
		})
	}
	if uuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldStatus,
		})
	}
	if value, ok := uuo.mutation.Comment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldComment,
		})
	}
	if uuo.mutation.CommentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldComment,
		})
	}
	if uuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: users.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.Next(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldNext,
		})
	}
	if uuo.mutation.NextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldNext,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldUpdatedAt,
		})
	}
	if uuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: users.FieldUpdatedAt,
		})
	}
	if uuo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: users.FieldURL,
		})
	}
	_node = &Users{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}