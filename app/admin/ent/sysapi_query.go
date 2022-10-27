// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysapi"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysApiQuery is the builder for querying SysApi entities.
type SysApiQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysApi
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysApiQuery builder.
func (saq *SysApiQuery) Where(ps ...predicate.SysApi) *SysApiQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit adds a limit step to the query.
func (saq *SysApiQuery) Limit(limit int) *SysApiQuery {
	saq.limit = &limit
	return saq
}

// Offset adds an offset step to the query.
func (saq *SysApiQuery) Offset(offset int) *SysApiQuery {
	saq.offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *SysApiQuery) Unique(unique bool) *SysApiQuery {
	saq.unique = &unique
	return saq
}

// Order adds an order step to the query.
func (saq *SysApiQuery) Order(o ...OrderFunc) *SysApiQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// First returns the first SysApi entity from the query.
// Returns a *NotFoundError when no SysApi was found.
func (saq *SysApiQuery) First(ctx context.Context) (*SysApi, error) {
	nodes, err := saq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysapi.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *SysApiQuery) FirstX(ctx context.Context) *SysApi {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysApi ID from the query.
// Returns a *NotFoundError when no SysApi ID was found.
func (saq *SysApiQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = saq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysapi.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *SysApiQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysApi entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysApi entity is found.
// Returns a *NotFoundError when no SysApi entities are found.
func (saq *SysApiQuery) Only(ctx context.Context) (*SysApi, error) {
	nodes, err := saq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysapi.Label}
	default:
		return nil, &NotSingularError{sysapi.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *SysApiQuery) OnlyX(ctx context.Context) *SysApi {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysApi ID in the query.
// Returns a *NotSingularError when more than one SysApi ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *SysApiQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = saq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysapi.Label}
	default:
		err = &NotSingularError{sysapi.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *SysApiQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysApis.
func (saq *SysApiQuery) All(ctx context.Context) ([]*SysApi, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return saq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (saq *SysApiQuery) AllX(ctx context.Context) []*SysApi {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysApi IDs.
func (saq *SysApiQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := saq.Select(sysapi.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *SysApiQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *SysApiQuery) Count(ctx context.Context) (int, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return saq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (saq *SysApiQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *SysApiQuery) Exist(ctx context.Context) (bool, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return saq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *SysApiQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysApiQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *SysApiQuery) Clone() *SysApiQuery {
	if saq == nil {
		return nil
	}
	return &SysApiQuery{
		config:     saq.config,
		limit:      saq.limit,
		offset:     saq.offset,
		order:      append([]OrderFunc{}, saq.order...),
		predicates: append([]predicate.SysApi{}, saq.predicates...),
		// clone intermediate query.
		sql:    saq.sql.Clone(),
		path:   saq.path,
		unique: saq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysApi.Query().
//		GroupBy(sysapi.FieldPath).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (saq *SysApiQuery) GroupBy(field string, fields ...string) *SysApiGroupBy {
	grbuild := &SysApiGroupBy{config: saq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return saq.sqlQuery(ctx), nil
	}
	grbuild.label = sysapi.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//	}
//
//	client.SysApi.Query().
//		Select(sysapi.FieldPath).
//		Scan(ctx, &v)
func (saq *SysApiQuery) Select(fields ...string) *SysApiSelect {
	saq.fields = append(saq.fields, fields...)
	selbuild := &SysApiSelect{SysApiQuery: saq}
	selbuild.label = sysapi.Label
	selbuild.flds, selbuild.scan = &saq.fields, selbuild.Scan
	return selbuild
}

func (saq *SysApiQuery) prepareQuery(ctx context.Context) error {
	for _, f := range saq.fields {
		if !sysapi.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *SysApiQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysApi, error) {
	var (
		nodes = []*SysApi{}
		_spec = saq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysApi).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysApi{config: saq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (saq *SysApiQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	_spec.Node.Columns = saq.fields
	if len(saq.fields) > 0 {
		_spec.Unique = saq.unique != nil && *saq.unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *SysApiQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := saq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (saq *SysApiQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysapi.Table,
			Columns: sysapi.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysapi.FieldID,
			},
		},
		From:   saq.sql,
		Unique: true,
	}
	if unique := saq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := saq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysapi.FieldID)
		for i := range fields {
			if fields[i] != sysapi.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *SysApiQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(sysapi.Table)
	columns := saq.fields
	if len(columns) == 0 {
		columns = sysapi.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.unique != nil && *saq.unique {
		selector.Distinct()
	}
	for _, m := range saq.modifiers {
		m(selector)
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (saq *SysApiQuery) Modify(modifiers ...func(s *sql.Selector)) *SysApiSelect {
	saq.modifiers = append(saq.modifiers, modifiers...)
	return saq.Select()
}

// SysApiGroupBy is the group-by builder for SysApi entities.
type SysApiGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *SysApiGroupBy) Aggregate(fns ...AggregateFunc) *SysApiGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sagb *SysApiGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sagb.path(ctx)
	if err != nil {
		return err
	}
	sagb.sql = query
	return sagb.sqlScan(ctx, v)
}

func (sagb *SysApiGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sagb.fields {
		if !sysapi.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sagb *SysApiGroupBy) sqlQuery() *sql.Selector {
	selector := sagb.sql.Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sagb.fields)+len(sagb.fns))
		for _, f := range sagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sagb.fields...)...)
}

// SysApiSelect is the builder for selecting fields of SysApi entities.
type SysApiSelect struct {
	*SysApiQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sas *SysApiSelect) Scan(ctx context.Context, v any) error {
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	sas.sql = sas.SysApiQuery.sqlQuery(ctx)
	return sas.sqlScan(ctx, v)
}

func (sas *SysApiSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sas.sql.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sas *SysApiSelect) Modify(modifiers ...func(s *sql.Selector)) *SysApiSelect {
	sas.modifiers = append(sas.modifiers, modifiers...)
	return sas
}
