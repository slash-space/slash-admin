// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysapi"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysApiUpdate is the builder for updating SysApi entities.
type SysApiUpdate struct {
	config
	hooks     []Hook
	mutation  *SysApiMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysApiUpdate builder.
func (sau *SysApiUpdate) Where(ps ...predicate.SysApi) *SysApiUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetPath sets the "path" field.
func (sau *SysApiUpdate) SetPath(s string) *SysApiUpdate {
	sau.mutation.SetPath(s)
	return sau
}

// SetDescription sets the "description" field.
func (sau *SysApiUpdate) SetDescription(s string) *SysApiUpdate {
	sau.mutation.SetDescription(s)
	return sau
}

// SetAPIGroup sets the "api_group" field.
func (sau *SysApiUpdate) SetAPIGroup(s string) *SysApiUpdate {
	sau.mutation.SetAPIGroup(s)
	return sau
}

// SetMethod sets the "method" field.
func (sau *SysApiUpdate) SetMethod(s string) *SysApiUpdate {
	sau.mutation.SetMethod(s)
	return sau
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (sau *SysApiUpdate) SetNillableMethod(s *string) *SysApiUpdate {
	if s != nil {
		sau.SetMethod(*s)
	}
	return sau
}

// SetUpdatedAt sets the "updated_at" field.
func (sau *SysApiUpdate) SetUpdatedAt(t time.Time) *SysApiUpdate {
	sau.mutation.SetUpdatedAt(t)
	return sau
}

// SetDeletedAt sets the "deleted_at" field.
func (sau *SysApiUpdate) SetDeletedAt(t time.Time) *SysApiUpdate {
	sau.mutation.SetDeletedAt(t)
	return sau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sau *SysApiUpdate) SetNillableDeletedAt(t *time.Time) *SysApiUpdate {
	if t != nil {
		sau.SetDeletedAt(*t)
	}
	return sau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sau *SysApiUpdate) ClearDeletedAt() *SysApiUpdate {
	sau.mutation.ClearDeletedAt()
	return sau
}

// Mutation returns the SysApiMutation object of the builder.
func (sau *SysApiUpdate) Mutation() *SysApiMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SysApiUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sau.defaults()
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysApiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SysApiUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SysApiUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SysApiUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *SysApiUpdate) defaults() {
	if _, ok := sau.mutation.UpdatedAt(); !ok {
		v := sysapi.UpdateDefaultUpdatedAt()
		sau.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sau *SysApiUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysApiUpdate {
	sau.modifiers = append(sau.modifiers, modifiers...)
	return sau
}

func (sau *SysApiUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysapi.Table,
			Columns: sysapi.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysapi.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldPath,
		})
	}
	if value, ok := sau.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldDescription,
		})
	}
	if value, ok := sau.mutation.APIGroup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldAPIGroup,
		})
	}
	if value, ok := sau.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldMethod,
		})
	}
	if value, ok := sau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysapi.FieldUpdatedAt,
		})
	}
	if value, ok := sau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysapi.FieldDeletedAt,
		})
	}
	if sau.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysapi.FieldDeletedAt,
		})
	}
	_spec.AddModifiers(sau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysapi.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SysApiUpdateOne is the builder for updating a single SysApi entity.
type SysApiUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysApiMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPath sets the "path" field.
func (sauo *SysApiUpdateOne) SetPath(s string) *SysApiUpdateOne {
	sauo.mutation.SetPath(s)
	return sauo
}

// SetDescription sets the "description" field.
func (sauo *SysApiUpdateOne) SetDescription(s string) *SysApiUpdateOne {
	sauo.mutation.SetDescription(s)
	return sauo
}

// SetAPIGroup sets the "api_group" field.
func (sauo *SysApiUpdateOne) SetAPIGroup(s string) *SysApiUpdateOne {
	sauo.mutation.SetAPIGroup(s)
	return sauo
}

// SetMethod sets the "method" field.
func (sauo *SysApiUpdateOne) SetMethod(s string) *SysApiUpdateOne {
	sauo.mutation.SetMethod(s)
	return sauo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (sauo *SysApiUpdateOne) SetNillableMethod(s *string) *SysApiUpdateOne {
	if s != nil {
		sauo.SetMethod(*s)
	}
	return sauo
}

// SetUpdatedAt sets the "updated_at" field.
func (sauo *SysApiUpdateOne) SetUpdatedAt(t time.Time) *SysApiUpdateOne {
	sauo.mutation.SetUpdatedAt(t)
	return sauo
}

// SetDeletedAt sets the "deleted_at" field.
func (sauo *SysApiUpdateOne) SetDeletedAt(t time.Time) *SysApiUpdateOne {
	sauo.mutation.SetDeletedAt(t)
	return sauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sauo *SysApiUpdateOne) SetNillableDeletedAt(t *time.Time) *SysApiUpdateOne {
	if t != nil {
		sauo.SetDeletedAt(*t)
	}
	return sauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sauo *SysApiUpdateOne) ClearDeletedAt() *SysApiUpdateOne {
	sauo.mutation.ClearDeletedAt()
	return sauo
}

// Mutation returns the SysApiMutation object of the builder.
func (sauo *SysApiUpdateOne) Mutation() *SysApiMutation {
	return sauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *SysApiUpdateOne) Select(field string, fields ...string) *SysApiUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated SysApi entity.
func (sauo *SysApiUpdateOne) Save(ctx context.Context) (*SysApi, error) {
	var (
		err  error
		node *SysApi
	)
	sauo.defaults()
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysApiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysApi)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysApiMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SysApiUpdateOne) SaveX(ctx context.Context) *SysApi {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SysApiUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SysApiUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *SysApiUpdateOne) defaults() {
	if _, ok := sauo.mutation.UpdatedAt(); !ok {
		v := sysapi.UpdateDefaultUpdatedAt()
		sauo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sauo *SysApiUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysApiUpdateOne {
	sauo.modifiers = append(sauo.modifiers, modifiers...)
	return sauo
}

func (sauo *SysApiUpdateOne) sqlSave(ctx context.Context) (_node *SysApi, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysapi.Table,
			Columns: sysapi.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysapi.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysApi.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysapi.FieldID)
		for _, f := range fields {
			if !sysapi.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysapi.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldPath,
		})
	}
	if value, ok := sauo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldDescription,
		})
	}
	if value, ok := sauo.mutation.APIGroup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldAPIGroup,
		})
	}
	if value, ok := sauo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysapi.FieldMethod,
		})
	}
	if value, ok := sauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysapi.FieldUpdatedAt,
		})
	}
	if value, ok := sauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysapi.FieldDeletedAt,
		})
	}
	if sauo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysapi.FieldDeletedAt,
		})
	}
	_spec.AddModifiers(sauo.modifiers...)
	_node = &SysApi{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysapi.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}