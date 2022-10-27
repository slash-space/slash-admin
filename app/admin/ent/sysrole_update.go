// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/app/admin/ent/sysrole"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleUpdate is the builder for updating SysRole entities.
type SysRoleUpdate struct {
	config
	hooks     []Hook
	mutation  *SysRoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysRoleUpdate builder.
func (sru *SysRoleUpdate) Where(ps ...predicate.SysRole) *SysRoleUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetName sets the "name" field.
func (sru *SysRoleUpdate) SetName(s string) *SysRoleUpdate {
	sru.mutation.SetName(s)
	return sru
}

// SetValue sets the "value" field.
func (sru *SysRoleUpdate) SetValue(s string) *SysRoleUpdate {
	sru.mutation.SetValue(s)
	return sru
}

// SetDefaultRouter sets the "default_router" field.
func (sru *SysRoleUpdate) SetDefaultRouter(s string) *SysRoleUpdate {
	sru.mutation.SetDefaultRouter(s)
	return sru
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableDefaultRouter(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetDefaultRouter(*s)
	}
	return sru
}

// SetStatus sets the "status" field.
func (sru *SysRoleUpdate) SetStatus(t types.Status) *SysRoleUpdate {
	sru.mutation.ResetStatus()
	sru.mutation.SetStatus(t)
	return sru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableStatus(t *types.Status) *SysRoleUpdate {
	if t != nil {
		sru.SetStatus(*t)
	}
	return sru
}

// AddStatus adds t to the "status" field.
func (sru *SysRoleUpdate) AddStatus(t types.Status) *SysRoleUpdate {
	sru.mutation.AddStatus(t)
	return sru
}

// ClearStatus clears the value of the "status" field.
func (sru *SysRoleUpdate) ClearStatus() *SysRoleUpdate {
	sru.mutation.ClearStatus()
	return sru
}

// SetRemark sets the "remark" field.
func (sru *SysRoleUpdate) SetRemark(s string) *SysRoleUpdate {
	sru.mutation.SetRemark(s)
	return sru
}

// SetOrderNo sets the "order_no" field.
func (sru *SysRoleUpdate) SetOrderNo(u uint32) *SysRoleUpdate {
	sru.mutation.ResetOrderNo()
	sru.mutation.SetOrderNo(u)
	return sru
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableOrderNo(u *uint32) *SysRoleUpdate {
	if u != nil {
		sru.SetOrderNo(*u)
	}
	return sru
}

// AddOrderNo adds u to the "order_no" field.
func (sru *SysRoleUpdate) AddOrderNo(u int32) *SysRoleUpdate {
	sru.mutation.AddOrderNo(u)
	return sru
}

// SetCreatedAt sets the "created_at" field.
func (sru *SysRoleUpdate) SetCreatedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetCreatedAt(t)
	return sru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableCreatedAt(t *time.Time) *SysRoleUpdate {
	if t != nil {
		sru.SetCreatedAt(*t)
	}
	return sru
}

// SetUpdatedAt sets the "updated_at" field.
func (sru *SysRoleUpdate) SetUpdatedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// SetDeletedAt sets the "deleted_at" field.
func (sru *SysRoleUpdate) SetDeletedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetDeletedAt(t)
	return sru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableDeletedAt(t *time.Time) *SysRoleUpdate {
	if t != nil {
		sru.SetDeletedAt(*t)
	}
	return sru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sru *SysRoleUpdate) ClearDeletedAt() *SysRoleUpdate {
	sru.mutation.ClearDeletedAt()
	return sru
}

// AddMenuIDs adds the "menus" edge to the SysMenu entity by IDs.
func (sru *SysRoleUpdate) AddMenuIDs(ids ...uint64) *SysRoleUpdate {
	sru.mutation.AddMenuIDs(ids...)
	return sru
}

// AddMenus adds the "menus" edges to the SysMenu entity.
func (sru *SysRoleUpdate) AddMenus(s ...*SysMenu) *SysRoleUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddMenuIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (sru *SysRoleUpdate) Mutation() *SysRoleMutation {
	return sru.mutation
}

// ClearMenus clears all "menus" edges to the SysMenu entity.
func (sru *SysRoleUpdate) ClearMenus() *SysRoleUpdate {
	sru.mutation.ClearMenus()
	return sru
}

// RemoveMenuIDs removes the "menus" edge to SysMenu entities by IDs.
func (sru *SysRoleUpdate) RemoveMenuIDs(ids ...uint64) *SysRoleUpdate {
	sru.mutation.RemoveMenuIDs(ids...)
	return sru
}

// RemoveMenus removes "menus" edges to SysMenu entities.
func (sru *SysRoleUpdate) RemoveMenus(s ...*SysMenu) *SysRoleUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveMenuIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *SysRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sru.defaults()
	if len(sru.hooks) == 0 {
		affected, err = sru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sru.mutation = mutation
			affected, err = sru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sru.hooks) - 1; i >= 0; i-- {
			if sru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sru *SysRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *SysRoleUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *SysRoleUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *SysRoleUpdate) defaults() {
	if _, ok := sru.mutation.UpdatedAt(); !ok {
		v := sysrole.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sru *SysRoleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysRoleUpdate {
	sru.modifiers = append(sru.modifiers, modifiers...)
	return sru
}

func (sru *SysRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysrole.FieldID,
			},
		},
	}
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldName,
		})
	}
	if value, ok := sru.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldValue,
		})
	}
	if value, ok := sru.mutation.DefaultRouter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldDefaultRouter,
		})
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sru.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if sru.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sru.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRemark,
		})
	}
	if value, ok := sru.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysrole.FieldOrderNo,
		})
	}
	if value, ok := sru.mutation.AddedOrderNo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysrole.FieldOrderNo,
		})
	}
	if value, ok := sru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldCreatedAt,
		})
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
	}
	if value, ok := sru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if sru.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if sru.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedMenusIDs(); len(nodes) > 0 && !sru.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SysRoleUpdateOne is the builder for updating a single SysRole entity.
type SysRoleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysRoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (sruo *SysRoleUpdateOne) SetName(s string) *SysRoleUpdateOne {
	sruo.mutation.SetName(s)
	return sruo
}

// SetValue sets the "value" field.
func (sruo *SysRoleUpdateOne) SetValue(s string) *SysRoleUpdateOne {
	sruo.mutation.SetValue(s)
	return sruo
}

// SetDefaultRouter sets the "default_router" field.
func (sruo *SysRoleUpdateOne) SetDefaultRouter(s string) *SysRoleUpdateOne {
	sruo.mutation.SetDefaultRouter(s)
	return sruo
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableDefaultRouter(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetDefaultRouter(*s)
	}
	return sruo
}

// SetStatus sets the "status" field.
func (sruo *SysRoleUpdateOne) SetStatus(t types.Status) *SysRoleUpdateOne {
	sruo.mutation.ResetStatus()
	sruo.mutation.SetStatus(t)
	return sruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableStatus(t *types.Status) *SysRoleUpdateOne {
	if t != nil {
		sruo.SetStatus(*t)
	}
	return sruo
}

// AddStatus adds t to the "status" field.
func (sruo *SysRoleUpdateOne) AddStatus(t types.Status) *SysRoleUpdateOne {
	sruo.mutation.AddStatus(t)
	return sruo
}

// ClearStatus clears the value of the "status" field.
func (sruo *SysRoleUpdateOne) ClearStatus() *SysRoleUpdateOne {
	sruo.mutation.ClearStatus()
	return sruo
}

// SetRemark sets the "remark" field.
func (sruo *SysRoleUpdateOne) SetRemark(s string) *SysRoleUpdateOne {
	sruo.mutation.SetRemark(s)
	return sruo
}

// SetOrderNo sets the "order_no" field.
func (sruo *SysRoleUpdateOne) SetOrderNo(u uint32) *SysRoleUpdateOne {
	sruo.mutation.ResetOrderNo()
	sruo.mutation.SetOrderNo(u)
	return sruo
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableOrderNo(u *uint32) *SysRoleUpdateOne {
	if u != nil {
		sruo.SetOrderNo(*u)
	}
	return sruo
}

// AddOrderNo adds u to the "order_no" field.
func (sruo *SysRoleUpdateOne) AddOrderNo(u int32) *SysRoleUpdateOne {
	sruo.mutation.AddOrderNo(u)
	return sruo
}

// SetCreatedAt sets the "created_at" field.
func (sruo *SysRoleUpdateOne) SetCreatedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetCreatedAt(t)
	return sruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableCreatedAt(t *time.Time) *SysRoleUpdateOne {
	if t != nil {
		sruo.SetCreatedAt(*t)
	}
	return sruo
}

// SetUpdatedAt sets the "updated_at" field.
func (sruo *SysRoleUpdateOne) SetUpdatedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// SetDeletedAt sets the "deleted_at" field.
func (sruo *SysRoleUpdateOne) SetDeletedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetDeletedAt(t)
	return sruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableDeletedAt(t *time.Time) *SysRoleUpdateOne {
	if t != nil {
		sruo.SetDeletedAt(*t)
	}
	return sruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sruo *SysRoleUpdateOne) ClearDeletedAt() *SysRoleUpdateOne {
	sruo.mutation.ClearDeletedAt()
	return sruo
}

// AddMenuIDs adds the "menus" edge to the SysMenu entity by IDs.
func (sruo *SysRoleUpdateOne) AddMenuIDs(ids ...uint64) *SysRoleUpdateOne {
	sruo.mutation.AddMenuIDs(ids...)
	return sruo
}

// AddMenus adds the "menus" edges to the SysMenu entity.
func (sruo *SysRoleUpdateOne) AddMenus(s ...*SysMenu) *SysRoleUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddMenuIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (sruo *SysRoleUpdateOne) Mutation() *SysRoleMutation {
	return sruo.mutation
}

// ClearMenus clears all "menus" edges to the SysMenu entity.
func (sruo *SysRoleUpdateOne) ClearMenus() *SysRoleUpdateOne {
	sruo.mutation.ClearMenus()
	return sruo
}

// RemoveMenuIDs removes the "menus" edge to SysMenu entities by IDs.
func (sruo *SysRoleUpdateOne) RemoveMenuIDs(ids ...uint64) *SysRoleUpdateOne {
	sruo.mutation.RemoveMenuIDs(ids...)
	return sruo
}

// RemoveMenus removes "menus" edges to SysMenu entities.
func (sruo *SysRoleUpdateOne) RemoveMenus(s ...*SysMenu) *SysRoleUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveMenuIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *SysRoleUpdateOne) Select(field string, fields ...string) *SysRoleUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated SysRole entity.
func (sruo *SysRoleUpdateOne) Save(ctx context.Context) (*SysRole, error) {
	var (
		err  error
		node *SysRole
	)
	sruo.defaults()
	if len(sruo.hooks) == 0 {
		node, err = sruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sruo.mutation = mutation
			node, err = sruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sruo.hooks) - 1; i >= 0; i-- {
			if sruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysRole)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysRoleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) SaveX(ctx context.Context) *SysRole {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *SysRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *SysRoleUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdatedAt(); !ok {
		v := sysrole.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sruo *SysRoleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysRoleUpdateOne {
	sruo.modifiers = append(sruo.modifiers, modifiers...)
	return sruo
}

func (sruo *SysRoleUpdateOne) sqlSave(ctx context.Context) (_node *SysRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysrole.FieldID,
			},
		},
	}
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for _, f := range fields {
			if !sysrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldName,
		})
	}
	if value, ok := sruo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldValue,
		})
	}
	if value, ok := sruo.mutation.DefaultRouter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldDefaultRouter,
		})
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sruo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if sruo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sruo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRemark,
		})
	}
	if value, ok := sruo.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysrole.FieldOrderNo,
		})
	}
	if value, ok := sruo.mutation.AddedOrderNo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysrole.FieldOrderNo,
		})
	}
	if value, ok := sruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldCreatedAt,
		})
	}
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
	}
	if value, ok := sruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if sruo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if sruo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedMenusIDs(); len(nodes) > 0 && !sruo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sruo.modifiers...)
	_node = &SysRole{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
