// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDictionaryDetailCreate is the builder for creating a SysDictionaryDetail entity.
type SysDictionaryDetailCreate struct {
	config
	mutation *SysDictionaryDetailMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (sddc *SysDictionaryDetailCreate) SetTitle(s string) *SysDictionaryDetailCreate {
	sddc.mutation.SetTitle(s)
	return sddc
}

// SetKey sets the "key" field.
func (sddc *SysDictionaryDetailCreate) SetKey(s string) *SysDictionaryDetailCreate {
	sddc.mutation.SetKey(s)
	return sddc
}

// SetValue sets the "value" field.
func (sddc *SysDictionaryDetailCreate) SetValue(s string) *SysDictionaryDetailCreate {
	sddc.mutation.SetValue(s)
	return sddc
}

// SetStatus sets the "status" field.
func (sddc *SysDictionaryDetailCreate) SetStatus(t types.Status) *SysDictionaryDetailCreate {
	sddc.mutation.SetStatus(t)
	return sddc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sddc *SysDictionaryDetailCreate) SetNillableStatus(t *types.Status) *SysDictionaryDetailCreate {
	if t != nil {
		sddc.SetStatus(*t)
	}
	return sddc
}

// SetRemark sets the "remark" field.
func (sddc *SysDictionaryDetailCreate) SetRemark(s string) *SysDictionaryDetailCreate {
	sddc.mutation.SetRemark(s)
	return sddc
}

// SetOrderNo sets the "order_no" field.
func (sddc *SysDictionaryDetailCreate) SetOrderNo(u uint32) *SysDictionaryDetailCreate {
	sddc.mutation.SetOrderNo(u)
	return sddc
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (sddc *SysDictionaryDetailCreate) SetNillableOrderNo(u *uint32) *SysDictionaryDetailCreate {
	if u != nil {
		sddc.SetOrderNo(*u)
	}
	return sddc
}

// SetCreatedAt sets the "created_at" field.
func (sddc *SysDictionaryDetailCreate) SetCreatedAt(t time.Time) *SysDictionaryDetailCreate {
	sddc.mutation.SetCreatedAt(t)
	return sddc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sddc *SysDictionaryDetailCreate) SetNillableCreatedAt(t *time.Time) *SysDictionaryDetailCreate {
	if t != nil {
		sddc.SetCreatedAt(*t)
	}
	return sddc
}

// SetUpdatedAt sets the "updated_at" field.
func (sddc *SysDictionaryDetailCreate) SetUpdatedAt(t time.Time) *SysDictionaryDetailCreate {
	sddc.mutation.SetUpdatedAt(t)
	return sddc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sddc *SysDictionaryDetailCreate) SetNillableUpdatedAt(t *time.Time) *SysDictionaryDetailCreate {
	if t != nil {
		sddc.SetUpdatedAt(*t)
	}
	return sddc
}

// SetDeletedAt sets the "deleted_at" field.
func (sddc *SysDictionaryDetailCreate) SetDeletedAt(t time.Time) *SysDictionaryDetailCreate {
	sddc.mutation.SetDeletedAt(t)
	return sddc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sddc *SysDictionaryDetailCreate) SetNillableDeletedAt(t *time.Time) *SysDictionaryDetailCreate {
	if t != nil {
		sddc.SetDeletedAt(*t)
	}
	return sddc
}

// SetID sets the "id" field.
func (sddc *SysDictionaryDetailCreate) SetID(u uint64) *SysDictionaryDetailCreate {
	sddc.mutation.SetID(u)
	return sddc
}

// Mutation returns the SysDictionaryDetailMutation object of the builder.
func (sddc *SysDictionaryDetailCreate) Mutation() *SysDictionaryDetailMutation {
	return sddc.mutation
}

// Save creates the SysDictionaryDetail in the database.
func (sddc *SysDictionaryDetailCreate) Save(ctx context.Context) (*SysDictionaryDetail, error) {
	var (
		err  error
		node *SysDictionaryDetail
	)
	sddc.defaults()
	if len(sddc.hooks) == 0 {
		if err = sddc.check(); err != nil {
			return nil, err
		}
		node, err = sddc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictionaryDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sddc.check(); err != nil {
				return nil, err
			}
			sddc.mutation = mutation
			if node, err = sddc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sddc.hooks) - 1; i >= 0; i-- {
			if sddc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sddc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sddc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysDictionaryDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysDictionaryDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sddc *SysDictionaryDetailCreate) SaveX(ctx context.Context) *SysDictionaryDetail {
	v, err := sddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sddc *SysDictionaryDetailCreate) Exec(ctx context.Context) error {
	_, err := sddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sddc *SysDictionaryDetailCreate) ExecX(ctx context.Context) {
	if err := sddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sddc *SysDictionaryDetailCreate) defaults() {
	if _, ok := sddc.mutation.Status(); !ok {
		v := sysdictionarydetail.DefaultStatus
		sddc.mutation.SetStatus(v)
	}
	if _, ok := sddc.mutation.OrderNo(); !ok {
		v := sysdictionarydetail.DefaultOrderNo
		sddc.mutation.SetOrderNo(v)
	}
	if _, ok := sddc.mutation.CreatedAt(); !ok {
		v := sysdictionarydetail.DefaultCreatedAt()
		sddc.mutation.SetCreatedAt(v)
	}
	if _, ok := sddc.mutation.UpdatedAt(); !ok {
		v := sysdictionarydetail.DefaultUpdatedAt()
		sddc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sddc *SysDictionaryDetailCreate) check() error {
	if _, ok := sddc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "SysDictionaryDetail.title"`)}
	}
	if _, ok := sddc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "SysDictionaryDetail.key"`)}
	}
	if _, ok := sddc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SysDictionaryDetail.value"`)}
	}
	if _, ok := sddc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "SysDictionaryDetail.remark"`)}
	}
	if _, ok := sddc.mutation.OrderNo(); !ok {
		return &ValidationError{Name: "order_no", err: errors.New(`ent: missing required field "SysDictionaryDetail.order_no"`)}
	}
	if _, ok := sddc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysDictionaryDetail.created_at"`)}
	}
	if _, ok := sddc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysDictionaryDetail.updated_at"`)}
	}
	return nil
}

func (sddc *SysDictionaryDetailCreate) sqlSave(ctx context.Context) (*SysDictionaryDetail, error) {
	_node, _spec := sddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (sddc *SysDictionaryDetailCreate) createSpec() (*SysDictionaryDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDictionaryDetail{config: sddc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysdictionarydetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysdictionarydetail.FieldID,
			},
		}
	)
	_spec.OnConflict = sddc.conflict
	if id, ok := sddc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sddc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := sddc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := sddc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := sddc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionarydetail.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sddc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := sddc.mutation.OrderNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysdictionarydetail.FieldOrderNo,
		})
		_node.OrderNo = value
	}
	if value, ok := sddc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sddc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sddc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysDictionaryDetail.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysDictionaryDetailUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (sddc *SysDictionaryDetailCreate) OnConflict(opts ...sql.ConflictOption) *SysDictionaryDetailUpsertOne {
	sddc.conflict = opts
	return &SysDictionaryDetailUpsertOne{
		create: sddc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysDictionaryDetail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sddc *SysDictionaryDetailCreate) OnConflictColumns(columns ...string) *SysDictionaryDetailUpsertOne {
	sddc.conflict = append(sddc.conflict, sql.ConflictColumns(columns...))
	return &SysDictionaryDetailUpsertOne{
		create: sddc,
	}
}

type (
	// SysDictionaryDetailUpsertOne is the builder for "upsert"-ing
	//  one SysDictionaryDetail node.
	SysDictionaryDetailUpsertOne struct {
		create *SysDictionaryDetailCreate
	}

	// SysDictionaryDetailUpsert is the "OnConflict" setter.
	SysDictionaryDetailUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *SysDictionaryDetailUpsert) SetTitle(v string) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateTitle() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldTitle)
	return u
}

// SetKey sets the "key" field.
func (u *SysDictionaryDetailUpsert) SetKey(v string) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateKey() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *SysDictionaryDetailUpsert) SetValue(v string) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateValue() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldValue)
	return u
}

// SetStatus sets the "status" field.
func (u *SysDictionaryDetailUpsert) SetStatus(v types.Status) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateStatus() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *SysDictionaryDetailUpsert) AddStatus(v types.Status) *SysDictionaryDetailUpsert {
	u.Add(sysdictionarydetail.FieldStatus, v)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *SysDictionaryDetailUpsert) ClearStatus() *SysDictionaryDetailUpsert {
	u.SetNull(sysdictionarydetail.FieldStatus)
	return u
}

// SetRemark sets the "remark" field.
func (u *SysDictionaryDetailUpsert) SetRemark(v string) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateRemark() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldRemark)
	return u
}

// SetOrderNo sets the "order_no" field.
func (u *SysDictionaryDetailUpsert) SetOrderNo(v uint32) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldOrderNo, v)
	return u
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateOrderNo() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldOrderNo)
	return u
}

// AddOrderNo adds v to the "order_no" field.
func (u *SysDictionaryDetailUpsert) AddOrderNo(v uint32) *SysDictionaryDetailUpsert {
	u.Add(sysdictionarydetail.FieldOrderNo, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SysDictionaryDetailUpsert) SetCreatedAt(v time.Time) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateCreatedAt() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictionaryDetailUpsert) SetUpdatedAt(v time.Time) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateUpdatedAt() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictionaryDetailUpsert) SetDeletedAt(v time.Time) *SysDictionaryDetailUpsert {
	u.Set(sysdictionarydetail.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsert) UpdateDeletedAt() *SysDictionaryDetailUpsert {
	u.SetExcluded(sysdictionarydetail.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictionaryDetailUpsert) ClearDeletedAt() *SysDictionaryDetailUpsert {
	u.SetNull(sysdictionarydetail.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysDictionaryDetail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysdictionarydetail.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysDictionaryDetailUpsertOne) UpdateNewValues() *SysDictionaryDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysdictionarydetail.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysDictionaryDetail.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysDictionaryDetailUpsertOne) Ignore() *SysDictionaryDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysDictionaryDetailUpsertOne) DoNothing() *SysDictionaryDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysDictionaryDetailCreate.OnConflict
// documentation for more info.
func (u *SysDictionaryDetailUpsertOne) Update(set func(*SysDictionaryDetailUpsert)) *SysDictionaryDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysDictionaryDetailUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *SysDictionaryDetailUpsertOne) SetTitle(v string) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateTitle() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateTitle()
	})
}

// SetKey sets the "key" field.
func (u *SysDictionaryDetailUpsertOne) SetKey(v string) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateKey() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *SysDictionaryDetailUpsertOne) SetValue(v string) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateValue() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateValue()
	})
}

// SetStatus sets the "status" field.
func (u *SysDictionaryDetailUpsertOne) SetStatus(v types.Status) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *SysDictionaryDetailUpsertOne) AddStatus(v types.Status) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateStatus() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *SysDictionaryDetailUpsertOne) ClearStatus() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.ClearStatus()
	})
}

// SetRemark sets the "remark" field.
func (u *SysDictionaryDetailUpsertOne) SetRemark(v string) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateRemark() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateRemark()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *SysDictionaryDetailUpsertOne) SetOrderNo(v uint32) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *SysDictionaryDetailUpsertOne) AddOrderNo(v uint32) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateOrderNo() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateOrderNo()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SysDictionaryDetailUpsertOne) SetCreatedAt(v time.Time) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateCreatedAt() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictionaryDetailUpsertOne) SetUpdatedAt(v time.Time) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateUpdatedAt() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictionaryDetailUpsertOne) SetDeletedAt(v time.Time) *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertOne) UpdateDeletedAt() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictionaryDetailUpsertOne) ClearDeletedAt() *SysDictionaryDetailUpsertOne {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SysDictionaryDetailUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysDictionaryDetailCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysDictionaryDetailUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysDictionaryDetailUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysDictionaryDetailUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysDictionaryDetailCreateBulk is the builder for creating many SysDictionaryDetail entities in bulk.
type SysDictionaryDetailCreateBulk struct {
	config
	builders []*SysDictionaryDetailCreate
	conflict []sql.ConflictOption
}

// Save creates the SysDictionaryDetail entities in the database.
func (sddcb *SysDictionaryDetailCreateBulk) Save(ctx context.Context) ([]*SysDictionaryDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sddcb.builders))
	nodes := make([]*SysDictionaryDetail, len(sddcb.builders))
	mutators := make([]Mutator, len(sddcb.builders))
	for i := range sddcb.builders {
		func(i int, root context.Context) {
			builder := sddcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDictionaryDetailMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sddcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sddcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sddcb *SysDictionaryDetailCreateBulk) SaveX(ctx context.Context) []*SysDictionaryDetail {
	v, err := sddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sddcb *SysDictionaryDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := sddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sddcb *SysDictionaryDetailCreateBulk) ExecX(ctx context.Context) {
	if err := sddcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysDictionaryDetail.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysDictionaryDetailUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (sddcb *SysDictionaryDetailCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysDictionaryDetailUpsertBulk {
	sddcb.conflict = opts
	return &SysDictionaryDetailUpsertBulk{
		create: sddcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysDictionaryDetail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sddcb *SysDictionaryDetailCreateBulk) OnConflictColumns(columns ...string) *SysDictionaryDetailUpsertBulk {
	sddcb.conflict = append(sddcb.conflict, sql.ConflictColumns(columns...))
	return &SysDictionaryDetailUpsertBulk{
		create: sddcb,
	}
}

// SysDictionaryDetailUpsertBulk is the builder for "upsert"-ing
// a bulk of SysDictionaryDetail nodes.
type SysDictionaryDetailUpsertBulk struct {
	create *SysDictionaryDetailCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysDictionaryDetail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysdictionarydetail.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysDictionaryDetailUpsertBulk) UpdateNewValues() *SysDictionaryDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysdictionarydetail.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysDictionaryDetail.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysDictionaryDetailUpsertBulk) Ignore() *SysDictionaryDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysDictionaryDetailUpsertBulk) DoNothing() *SysDictionaryDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysDictionaryDetailCreateBulk.OnConflict
// documentation for more info.
func (u *SysDictionaryDetailUpsertBulk) Update(set func(*SysDictionaryDetailUpsert)) *SysDictionaryDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysDictionaryDetailUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *SysDictionaryDetailUpsertBulk) SetTitle(v string) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateTitle() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateTitle()
	})
}

// SetKey sets the "key" field.
func (u *SysDictionaryDetailUpsertBulk) SetKey(v string) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateKey() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *SysDictionaryDetailUpsertBulk) SetValue(v string) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateValue() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateValue()
	})
}

// SetStatus sets the "status" field.
func (u *SysDictionaryDetailUpsertBulk) SetStatus(v types.Status) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *SysDictionaryDetailUpsertBulk) AddStatus(v types.Status) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateStatus() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *SysDictionaryDetailUpsertBulk) ClearStatus() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.ClearStatus()
	})
}

// SetRemark sets the "remark" field.
func (u *SysDictionaryDetailUpsertBulk) SetRemark(v string) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateRemark() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateRemark()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *SysDictionaryDetailUpsertBulk) SetOrderNo(v uint32) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *SysDictionaryDetailUpsertBulk) AddOrderNo(v uint32) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateOrderNo() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateOrderNo()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SysDictionaryDetailUpsertBulk) SetCreatedAt(v time.Time) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateCreatedAt() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictionaryDetailUpsertBulk) SetUpdatedAt(v time.Time) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateUpdatedAt() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictionaryDetailUpsertBulk) SetDeletedAt(v time.Time) *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictionaryDetailUpsertBulk) UpdateDeletedAt() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictionaryDetailUpsertBulk) ClearDeletedAt() *SysDictionaryDetailUpsertBulk {
	return u.Update(func(s *SysDictionaryDetailUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SysDictionaryDetailUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysDictionaryDetailCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysDictionaryDetailCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysDictionaryDetailUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
