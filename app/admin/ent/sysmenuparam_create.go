// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/sysmenuparam"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysMenuParamCreate is the builder for creating a SysMenuParam entity.
type SysMenuParamCreate struct {
	config
	mutation *SysMenuParamMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetMenuID sets the "menu_id" field.
func (smpc *SysMenuParamCreate) SetMenuID(u uint64) *SysMenuParamCreate {
	smpc.mutation.SetMenuID(u)
	return smpc
}

// SetType sets the "type" field.
func (smpc *SysMenuParamCreate) SetType(s string) *SysMenuParamCreate {
	smpc.mutation.SetType(s)
	return smpc
}

// SetKey sets the "key" field.
func (smpc *SysMenuParamCreate) SetKey(s string) *SysMenuParamCreate {
	smpc.mutation.SetKey(s)
	return smpc
}

// SetValue sets the "value" field.
func (smpc *SysMenuParamCreate) SetValue(s string) *SysMenuParamCreate {
	smpc.mutation.SetValue(s)
	return smpc
}

// SetCreatedAt sets the "created_at" field.
func (smpc *SysMenuParamCreate) SetCreatedAt(t time.Time) *SysMenuParamCreate {
	smpc.mutation.SetCreatedAt(t)
	return smpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smpc *SysMenuParamCreate) SetNillableCreatedAt(t *time.Time) *SysMenuParamCreate {
	if t != nil {
		smpc.SetCreatedAt(*t)
	}
	return smpc
}

// SetUpdatedAt sets the "updated_at" field.
func (smpc *SysMenuParamCreate) SetUpdatedAt(t time.Time) *SysMenuParamCreate {
	smpc.mutation.SetUpdatedAt(t)
	return smpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smpc *SysMenuParamCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuParamCreate {
	if t != nil {
		smpc.SetUpdatedAt(*t)
	}
	return smpc
}

// SetDeletedAt sets the "deleted_at" field.
func (smpc *SysMenuParamCreate) SetDeletedAt(t time.Time) *SysMenuParamCreate {
	smpc.mutation.SetDeletedAt(t)
	return smpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smpc *SysMenuParamCreate) SetNillableDeletedAt(t *time.Time) *SysMenuParamCreate {
	if t != nil {
		smpc.SetDeletedAt(*t)
	}
	return smpc
}

// SetID sets the "id" field.
func (smpc *SysMenuParamCreate) SetID(u uint64) *SysMenuParamCreate {
	smpc.mutation.SetID(u)
	return smpc
}

// Mutation returns the SysMenuParamMutation object of the builder.
func (smpc *SysMenuParamCreate) Mutation() *SysMenuParamMutation {
	return smpc.mutation
}

// Save creates the SysMenuParam in the database.
func (smpc *SysMenuParamCreate) Save(ctx context.Context) (*SysMenuParam, error) {
	var (
		err  error
		node *SysMenuParam
	)
	smpc.defaults()
	if len(smpc.hooks) == 0 {
		if err = smpc.check(); err != nil {
			return nil, err
		}
		node, err = smpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuParamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smpc.check(); err != nil {
				return nil, err
			}
			smpc.mutation = mutation
			if node, err = smpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(smpc.hooks) - 1; i >= 0; i-- {
			if smpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, smpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysMenuParam)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysMenuParamMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smpc *SysMenuParamCreate) SaveX(ctx context.Context) *SysMenuParam {
	v, err := smpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smpc *SysMenuParamCreate) Exec(ctx context.Context) error {
	_, err := smpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smpc *SysMenuParamCreate) ExecX(ctx context.Context) {
	if err := smpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smpc *SysMenuParamCreate) defaults() {
	if _, ok := smpc.mutation.CreatedAt(); !ok {
		v := sysmenuparam.DefaultCreatedAt()
		smpc.mutation.SetCreatedAt(v)
	}
	if _, ok := smpc.mutation.UpdatedAt(); !ok {
		v := sysmenuparam.DefaultUpdatedAt()
		smpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smpc *SysMenuParamCreate) check() error {
	if _, ok := smpc.mutation.MenuID(); !ok {
		return &ValidationError{Name: "menu_id", err: errors.New(`ent: missing required field "SysMenuParam.menu_id"`)}
	}
	if _, ok := smpc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "SysMenuParam.type"`)}
	}
	if _, ok := smpc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "SysMenuParam.key"`)}
	}
	if _, ok := smpc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SysMenuParam.value"`)}
	}
	if _, ok := smpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysMenuParam.created_at"`)}
	}
	if _, ok := smpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysMenuParam.updated_at"`)}
	}
	return nil
}

func (smpc *SysMenuParamCreate) sqlSave(ctx context.Context) (*SysMenuParam, error) {
	_node, _spec := smpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smpc.driver, _spec); err != nil {
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

func (smpc *SysMenuParamCreate) createSpec() (*SysMenuParam, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenuParam{config: smpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysmenuparam.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysmenuparam.FieldID,
			},
		}
	)
	_spec.OnConflict = smpc.conflict
	if id, ok := smpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smpc.mutation.MenuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: sysmenuparam.FieldMenuID,
		})
		_node.MenuID = value
	}
	if value, ok := smpc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuparam.FieldType,
		})
		_node.Type = value
	}
	if value, ok := smpc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuparam.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := smpc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuparam.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := smpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuparam.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuparam.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuparam.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysMenuParam.Create().
//		SetMenuID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysMenuParamUpsert) {
//			SetMenuID(v+v).
//		}).
//		Exec(ctx)
func (smpc *SysMenuParamCreate) OnConflict(opts ...sql.ConflictOption) *SysMenuParamUpsertOne {
	smpc.conflict = opts
	return &SysMenuParamUpsertOne{
		create: smpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysMenuParam.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (smpc *SysMenuParamCreate) OnConflictColumns(columns ...string) *SysMenuParamUpsertOne {
	smpc.conflict = append(smpc.conflict, sql.ConflictColumns(columns...))
	return &SysMenuParamUpsertOne{
		create: smpc,
	}
}

type (
	// SysMenuParamUpsertOne is the builder for "upsert"-ing
	//  one SysMenuParam node.
	SysMenuParamUpsertOne struct {
		create *SysMenuParamCreate
	}

	// SysMenuParamUpsert is the "OnConflict" setter.
	SysMenuParamUpsert struct {
		*sql.UpdateSet
	}
)

// SetMenuID sets the "menu_id" field.
func (u *SysMenuParamUpsert) SetMenuID(v uint64) *SysMenuParamUpsert {
	u.Set(sysmenuparam.FieldMenuID, v)
	return u
}

// UpdateMenuID sets the "menu_id" field to the value that was provided on create.
func (u *SysMenuParamUpsert) UpdateMenuID() *SysMenuParamUpsert {
	u.SetExcluded(sysmenuparam.FieldMenuID)
	return u
}

// AddMenuID adds v to the "menu_id" field.
func (u *SysMenuParamUpsert) AddMenuID(v uint64) *SysMenuParamUpsert {
	u.Add(sysmenuparam.FieldMenuID, v)
	return u
}

// SetType sets the "type" field.
func (u *SysMenuParamUpsert) SetType(v string) *SysMenuParamUpsert {
	u.Set(sysmenuparam.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *SysMenuParamUpsert) UpdateType() *SysMenuParamUpsert {
	u.SetExcluded(sysmenuparam.FieldType)
	return u
}

// SetKey sets the "key" field.
func (u *SysMenuParamUpsert) SetKey(v string) *SysMenuParamUpsert {
	u.Set(sysmenuparam.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SysMenuParamUpsert) UpdateKey() *SysMenuParamUpsert {
	u.SetExcluded(sysmenuparam.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *SysMenuParamUpsert) SetValue(v string) *SysMenuParamUpsert {
	u.Set(sysmenuparam.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysMenuParamUpsert) UpdateValue() *SysMenuParamUpsert {
	u.SetExcluded(sysmenuparam.FieldValue)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuParamUpsert) SetUpdatedAt(v time.Time) *SysMenuParamUpsert {
	u.Set(sysmenuparam.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuParamUpsert) UpdateUpdatedAt() *SysMenuParamUpsert {
	u.SetExcluded(sysmenuparam.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuParamUpsert) SetDeletedAt(v time.Time) *SysMenuParamUpsert {
	u.Set(sysmenuparam.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuParamUpsert) UpdateDeletedAt() *SysMenuParamUpsert {
	u.SetExcluded(sysmenuparam.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuParamUpsert) ClearDeletedAt() *SysMenuParamUpsert {
	u.SetNull(sysmenuparam.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysMenuParam.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysmenuparam.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysMenuParamUpsertOne) UpdateNewValues() *SysMenuParamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysmenuparam.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysmenuparam.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysMenuParam.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysMenuParamUpsertOne) Ignore() *SysMenuParamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysMenuParamUpsertOne) DoNothing() *SysMenuParamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysMenuParamCreate.OnConflict
// documentation for more info.
func (u *SysMenuParamUpsertOne) Update(set func(*SysMenuParamUpsert)) *SysMenuParamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysMenuParamUpsert{UpdateSet: update})
	}))
	return u
}

// SetMenuID sets the "menu_id" field.
func (u *SysMenuParamUpsertOne) SetMenuID(v uint64) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetMenuID(v)
	})
}

// AddMenuID adds v to the "menu_id" field.
func (u *SysMenuParamUpsertOne) AddMenuID(v uint64) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.AddMenuID(v)
	})
}

// UpdateMenuID sets the "menu_id" field to the value that was provided on create.
func (u *SysMenuParamUpsertOne) UpdateMenuID() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateMenuID()
	})
}

// SetType sets the "type" field.
func (u *SysMenuParamUpsertOne) SetType(v string) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *SysMenuParamUpsertOne) UpdateType() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateType()
	})
}

// SetKey sets the "key" field.
func (u *SysMenuParamUpsertOne) SetKey(v string) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SysMenuParamUpsertOne) UpdateKey() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *SysMenuParamUpsertOne) SetValue(v string) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysMenuParamUpsertOne) UpdateValue() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateValue()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuParamUpsertOne) SetUpdatedAt(v time.Time) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuParamUpsertOne) UpdateUpdatedAt() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuParamUpsertOne) SetDeletedAt(v time.Time) *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuParamUpsertOne) UpdateDeletedAt() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuParamUpsertOne) ClearDeletedAt() *SysMenuParamUpsertOne {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SysMenuParamUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysMenuParamCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysMenuParamUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysMenuParamUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysMenuParamUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysMenuParamCreateBulk is the builder for creating many SysMenuParam entities in bulk.
type SysMenuParamCreateBulk struct {
	config
	builders []*SysMenuParamCreate
	conflict []sql.ConflictOption
}

// Save creates the SysMenuParam entities in the database.
func (smpcb *SysMenuParamCreateBulk) Save(ctx context.Context) ([]*SysMenuParam, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smpcb.builders))
	nodes := make([]*SysMenuParam, len(smpcb.builders))
	mutators := make([]Mutator, len(smpcb.builders))
	for i := range smpcb.builders {
		func(i int, root context.Context) {
			builder := smpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuParamMutation)
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
					_, err = mutators[i+1].Mutate(root, smpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = smpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smpcb *SysMenuParamCreateBulk) SaveX(ctx context.Context) []*SysMenuParam {
	v, err := smpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smpcb *SysMenuParamCreateBulk) Exec(ctx context.Context) error {
	_, err := smpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smpcb *SysMenuParamCreateBulk) ExecX(ctx context.Context) {
	if err := smpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysMenuParam.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysMenuParamUpsert) {
//			SetMenuID(v+v).
//		}).
//		Exec(ctx)
func (smpcb *SysMenuParamCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysMenuParamUpsertBulk {
	smpcb.conflict = opts
	return &SysMenuParamUpsertBulk{
		create: smpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysMenuParam.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (smpcb *SysMenuParamCreateBulk) OnConflictColumns(columns ...string) *SysMenuParamUpsertBulk {
	smpcb.conflict = append(smpcb.conflict, sql.ConflictColumns(columns...))
	return &SysMenuParamUpsertBulk{
		create: smpcb,
	}
}

// SysMenuParamUpsertBulk is the builder for "upsert"-ing
// a bulk of SysMenuParam nodes.
type SysMenuParamUpsertBulk struct {
	create *SysMenuParamCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysMenuParam.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysmenuparam.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysMenuParamUpsertBulk) UpdateNewValues() *SysMenuParamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysmenuparam.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysmenuparam.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysMenuParam.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysMenuParamUpsertBulk) Ignore() *SysMenuParamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysMenuParamUpsertBulk) DoNothing() *SysMenuParamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysMenuParamCreateBulk.OnConflict
// documentation for more info.
func (u *SysMenuParamUpsertBulk) Update(set func(*SysMenuParamUpsert)) *SysMenuParamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysMenuParamUpsert{UpdateSet: update})
	}))
	return u
}

// SetMenuID sets the "menu_id" field.
func (u *SysMenuParamUpsertBulk) SetMenuID(v uint64) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetMenuID(v)
	})
}

// AddMenuID adds v to the "menu_id" field.
func (u *SysMenuParamUpsertBulk) AddMenuID(v uint64) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.AddMenuID(v)
	})
}

// UpdateMenuID sets the "menu_id" field to the value that was provided on create.
func (u *SysMenuParamUpsertBulk) UpdateMenuID() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateMenuID()
	})
}

// SetType sets the "type" field.
func (u *SysMenuParamUpsertBulk) SetType(v string) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *SysMenuParamUpsertBulk) UpdateType() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateType()
	})
}

// SetKey sets the "key" field.
func (u *SysMenuParamUpsertBulk) SetKey(v string) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SysMenuParamUpsertBulk) UpdateKey() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *SysMenuParamUpsertBulk) SetValue(v string) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysMenuParamUpsertBulk) UpdateValue() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateValue()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuParamUpsertBulk) SetUpdatedAt(v time.Time) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuParamUpsertBulk) UpdateUpdatedAt() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuParamUpsertBulk) SetDeletedAt(v time.Time) *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuParamUpsertBulk) UpdateDeletedAt() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuParamUpsertBulk) ClearDeletedAt() *SysMenuParamUpsertBulk {
	return u.Update(func(s *SysMenuParamUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SysMenuParamUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysMenuParamCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysMenuParamCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysMenuParamUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}