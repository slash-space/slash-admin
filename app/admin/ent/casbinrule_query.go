// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"slash-admin/app/admin/ent/casbinrule"
	"slash-admin/app/admin/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CasbinRuleQuery is the builder for querying CasbinRule entities.
type CasbinRuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CasbinRule
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CasbinRuleQuery builder.
func (crq *CasbinRuleQuery) Where(ps ...predicate.CasbinRule) *CasbinRuleQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit adds a limit step to the query.
func (crq *CasbinRuleQuery) Limit(limit int) *CasbinRuleQuery {
	crq.limit = &limit
	return crq
}

// Offset adds an offset step to the query.
func (crq *CasbinRuleQuery) Offset(offset int) *CasbinRuleQuery {
	crq.offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *CasbinRuleQuery) Unique(unique bool) *CasbinRuleQuery {
	crq.unique = &unique
	return crq
}

// Order adds an order step to the query.
func (crq *CasbinRuleQuery) Order(o ...OrderFunc) *CasbinRuleQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// First returns the first CasbinRule entity from the query.
// Returns a *NotFoundError when no CasbinRule was found.
func (crq *CasbinRuleQuery) First(ctx context.Context) (*CasbinRule, error) {
	nodes, err := crq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{casbinrule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *CasbinRuleQuery) FirstX(ctx context.Context) *CasbinRule {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CasbinRule ID from the query.
// Returns a *NotFoundError when no CasbinRule ID was found.
func (crq *CasbinRuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{casbinrule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *CasbinRuleQuery) FirstIDX(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CasbinRule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CasbinRule entity is found.
// Returns a *NotFoundError when no CasbinRule entities are found.
func (crq *CasbinRuleQuery) Only(ctx context.Context) (*CasbinRule, error) {
	nodes, err := crq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{casbinrule.Label}
	default:
		return nil, &NotSingularError{casbinrule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *CasbinRuleQuery) OnlyX(ctx context.Context) *CasbinRule {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CasbinRule ID in the query.
// Returns a *NotSingularError when more than one CasbinRule ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *CasbinRuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{casbinrule.Label}
	default:
		err = &NotSingularError{casbinrule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *CasbinRuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CasbinRules.
func (crq *CasbinRuleQuery) All(ctx context.Context) ([]*CasbinRule, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return crq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (crq *CasbinRuleQuery) AllX(ctx context.Context) []*CasbinRule {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CasbinRule IDs.
func (crq *CasbinRuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := crq.Select(casbinrule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *CasbinRuleQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *CasbinRuleQuery) Count(ctx context.Context) (int, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return crq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (crq *CasbinRuleQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *CasbinRuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return crq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *CasbinRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CasbinRuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *CasbinRuleQuery) Clone() *CasbinRuleQuery {
	if crq == nil {
		return nil
	}
	return &CasbinRuleQuery{
		config:     crq.config,
		limit:      crq.limit,
		offset:     crq.offset,
		order:      append([]OrderFunc{}, crq.order...),
		predicates: append([]predicate.CasbinRule{}, crq.predicates...),
		// clone intermediate query.
		sql:    crq.sql.Clone(),
		path:   crq.path,
		unique: crq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Ptype string `json:"Ptype,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CasbinRule.Query().
//		GroupBy(casbinrule.FieldPtype).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *CasbinRuleQuery) GroupBy(field string, fields ...string) *CasbinRuleGroupBy {
	grbuild := &CasbinRuleGroupBy{config: crq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return crq.sqlQuery(ctx), nil
	}
	grbuild.label = casbinrule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Ptype string `json:"Ptype,omitempty"`
//	}
//
//	client.CasbinRule.Query().
//		Select(casbinrule.FieldPtype).
//		Scan(ctx, &v)
func (crq *CasbinRuleQuery) Select(fields ...string) *CasbinRuleSelect {
	crq.fields = append(crq.fields, fields...)
	selbuild := &CasbinRuleSelect{CasbinRuleQuery: crq}
	selbuild.label = casbinrule.Label
	selbuild.flds, selbuild.scan = &crq.fields, selbuild.Scan
	return selbuild
}

func (crq *CasbinRuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range crq.fields {
		if !casbinrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *CasbinRuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CasbinRule, error) {
	var (
		nodes = []*CasbinRule{}
		_spec = crq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CasbinRule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CasbinRule{config: crq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(crq.modifiers) > 0 {
		_spec.Modifiers = crq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (crq *CasbinRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	if len(crq.modifiers) > 0 {
		_spec.Modifiers = crq.modifiers
	}
	_spec.Node.Columns = crq.fields
	if len(crq.fields) > 0 {
		_spec.Unique = crq.unique != nil && *crq.unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *CasbinRuleQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (crq *CasbinRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   casbinrule.Table,
			Columns: casbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: casbinrule.FieldID,
			},
		},
		From:   crq.sql,
		Unique: true,
	}
	if unique := crq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := crq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casbinrule.FieldID)
		for i := range fields {
			if fields[i] != casbinrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *CasbinRuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(casbinrule.Table)
	columns := crq.fields
	if len(columns) == 0 {
		columns = casbinrule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.unique != nil && *crq.unique {
		selector.Distinct()
	}
	for _, m := range crq.modifiers {
		m(selector)
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (crq *CasbinRuleQuery) Modify(modifiers ...func(s *sql.Selector)) *CasbinRuleSelect {
	crq.modifiers = append(crq.modifiers, modifiers...)
	return crq.Select()
}

// CasbinRuleGroupBy is the group-by builder for CasbinRule entities.
type CasbinRuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *CasbinRuleGroupBy) Aggregate(fns ...AggregateFunc) *CasbinRuleGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the group-by query and scans the result into the given value.
func (crgb *CasbinRuleGroupBy) Scan(ctx context.Context, v any) error {
	query, err := crgb.path(ctx)
	if err != nil {
		return err
	}
	crgb.sql = query
	return crgb.sqlScan(ctx, v)
}

func (crgb *CasbinRuleGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range crgb.fields {
		if !casbinrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := crgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crgb *CasbinRuleGroupBy) sqlQuery() *sql.Selector {
	selector := crgb.sql.Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(crgb.fields)+len(crgb.fns))
		for _, f := range crgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(crgb.fields...)...)
}

// CasbinRuleSelect is the builder for selecting fields of CasbinRule entities.
type CasbinRuleSelect struct {
	*CasbinRuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (crs *CasbinRuleSelect) Scan(ctx context.Context, v any) error {
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	crs.sql = crs.CasbinRuleQuery.sqlQuery(ctx)
	return crs.sqlScan(ctx, v)
}

func (crs *CasbinRuleSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := crs.sql.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (crs *CasbinRuleSelect) Modify(modifiers ...func(s *sql.Selector)) *CasbinRuleSelect {
	crs.modifiers = append(crs.modifiers, modifiers...)
	return crs
}
