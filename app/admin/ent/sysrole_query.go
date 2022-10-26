// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleQuery is the builder for querying SysRole entities.
type SysRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysRole
	loadTotal  []func(context.Context, []*SysRole) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysRoleQuery builder.
func (srq *SysRoleQuery) Where(ps ...predicate.SysRole) *SysRoleQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit adds a limit step to the query.
func (srq *SysRoleQuery) Limit(limit int) *SysRoleQuery {
	srq.limit = &limit
	return srq
}

// Offset adds an offset step to the query.
func (srq *SysRoleQuery) Offset(offset int) *SysRoleQuery {
	srq.offset = &offset
	return srq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srq *SysRoleQuery) Unique(unique bool) *SysRoleQuery {
	srq.unique = &unique
	return srq
}

// Order adds an order step to the query.
func (srq *SysRoleQuery) Order(o ...OrderFunc) *SysRoleQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// First returns the first SysRole entity from the query.
// Returns a *NotFoundError when no SysRole was found.
func (srq *SysRoleQuery) First(ctx context.Context) (*SysRole, error) {
	nodes, err := srq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *SysRoleQuery) FirstX(ctx context.Context) *SysRole {
	node, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysRole ID from the query.
// Returns a *NotFoundError when no SysRole ID was found.
func (srq *SysRoleQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = srq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srq *SysRoleQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysRole entity is found.
// Returns a *NotFoundError when no SysRole entities are found.
func (srq *SysRoleQuery) Only(ctx context.Context) (*SysRole, error) {
	nodes, err := srq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysrole.Label}
	default:
		return nil, &NotSingularError{sysrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *SysRoleQuery) OnlyX(ctx context.Context) *SysRole {
	node, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysRole ID in the query.
// Returns a *NotSingularError when more than one SysRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (srq *SysRoleQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = srq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = &NotSingularError{sysrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *SysRoleQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysRoles.
func (srq *SysRoleQuery) All(ctx context.Context) ([]*SysRole, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return srq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (srq *SysRoleQuery) AllX(ctx context.Context) []*SysRole {
	nodes, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysRole IDs.
func (srq *SysRoleQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := srq.Select(sysrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *SysRoleQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *SysRoleQuery) Count(ctx context.Context) (int, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return srq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (srq *SysRoleQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *SysRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return srq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *SysRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *SysRoleQuery) Clone() *SysRoleQuery {
	if srq == nil {
		return nil
	}
	return &SysRoleQuery{
		config:     srq.config,
		limit:      srq.limit,
		offset:     srq.offset,
		order:      append([]OrderFunc{}, srq.order...),
		predicates: append([]predicate.SysRole{}, srq.predicates...),
		// clone intermediate query.
		sql:    srq.sql.Clone(),
		path:   srq.path,
		unique: srq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysRole.Query().
//		GroupBy(sysrole.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (srq *SysRoleQuery) GroupBy(field string, fields ...string) *SysRoleGroupBy {
	grbuild := &SysRoleGroupBy{config: srq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srq.sqlQuery(ctx), nil
	}
	grbuild.label = sysrole.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.SysRole.Query().
//		Select(sysrole.FieldName).
//		Scan(ctx, &v)
func (srq *SysRoleQuery) Select(fields ...string) *SysRoleSelect {
	srq.fields = append(srq.fields, fields...)
	selbuild := &SysRoleSelect{SysRoleQuery: srq}
	selbuild.label = sysrole.Label
	selbuild.flds, selbuild.scan = &srq.fields, selbuild.Scan
	return selbuild
}

func (srq *SysRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range srq.fields {
		if !sysrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *SysRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysRole, error) {
	var (
		nodes = []*SysRole{}
		_spec = srq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysRole{config: srq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(srq.modifiers) > 0 {
		_spec.Modifiers = srq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range srq.loadTotal {
		if err := srq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (srq *SysRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	if len(srq.modifiers) > 0 {
		_spec.Modifiers = srq.modifiers
	}
	_spec.Node.Columns = srq.fields
	if len(srq.fields) > 0 {
		_spec.Unique = srq.unique != nil && *srq.unique
	}
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *SysRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := srq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (srq *SysRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysrole.FieldID,
			},
		},
		From:   srq.sql,
		Unique: true,
	}
	if unique := srq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := srq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for i := range fields {
			if fields[i] != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *SysRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(sysrole.Table)
	columns := srq.fields
	if len(columns) == 0 {
		columns = sysrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if srq.unique != nil && *srq.unique {
		selector.Distinct()
	}
	for _, m := range srq.modifiers {
		m(selector)
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (srq *SysRoleQuery) Modify(modifiers ...func(s *sql.Selector)) *SysRoleSelect {
	srq.modifiers = append(srq.modifiers, modifiers...)
	return srq.Select()
}

// SysRoleGroupBy is the group-by builder for SysRole entities.
type SysRoleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *SysRoleGroupBy) Aggregate(fns ...AggregateFunc) *SysRoleGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the group-by query and scans the result into the given value.
func (srgb *SysRoleGroupBy) Scan(ctx context.Context, v any) error {
	query, err := srgb.path(ctx)
	if err != nil {
		return err
	}
	srgb.sql = query
	return srgb.sqlScan(ctx, v)
}

func (srgb *SysRoleGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range srgb.fields {
		if !sysrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := srgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srgb *SysRoleGroupBy) sqlQuery() *sql.Selector {
	selector := srgb.sql.Select()
	aggregation := make([]string, 0, len(srgb.fns))
	for _, fn := range srgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(srgb.fields)+len(srgb.fns))
		for _, f := range srgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(srgb.fields...)...)
}

// SysRoleSelect is the builder for selecting fields of SysRole entities.
type SysRoleSelect struct {
	*SysRoleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (srs *SysRoleSelect) Scan(ctx context.Context, v any) error {
	if err := srs.prepareQuery(ctx); err != nil {
		return err
	}
	srs.sql = srs.SysRoleQuery.sqlQuery(ctx)
	return srs.sqlScan(ctx, v)
}

func (srs *SysRoleSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := srs.sql.Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (srs *SysRoleSelect) Modify(modifiers ...func(s *sql.Selector)) *SysRoleSelect {
	srs.modifiers = append(srs.modifiers, modifiers...)
	return srs
}