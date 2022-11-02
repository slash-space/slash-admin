// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sa *SysApiQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysApiQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sa, nil
	}
	if err := sa.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sa, nil
}

func (sa *SysApiQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type sysapiPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysApiPaginateOption
}

func newSysApiPaginateArgs(rv map[string]interface{}) *sysapiPaginateArgs {
	args := &sysapiPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sd *SysDictionaryQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysDictionaryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sd, nil
	}
	if err := sd.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *SysDictionaryQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "details":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysDictionaryDetailQuery{config: sd.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sd.WithNamedDetails(alias, func(wq *SysDictionaryDetailQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type sysdictionaryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysDictionaryPaginateOption
}

func newSysDictionaryPaginateArgs(rv map[string]interface{}) *sysdictionaryPaginateArgs {
	args := &sysdictionaryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sdd *SysDictionaryDetailQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysDictionaryDetailQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sdd, nil
	}
	if err := sdd.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sdd, nil
}

func (sdd *SysDictionaryDetailQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysDictionaryQuery{config: sdd.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sdd.withParent = query
		}
	}
	return nil
}

type sysdictionarydetailPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysDictionaryDetailPaginateOption
}

func newSysDictionaryDetailPaginateArgs(rv map[string]interface{}) *sysdictionarydetailPaginateArgs {
	args := &sysdictionarydetailPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sm *SysMenuQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysMenuQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sm, nil
	}
	if err := sm.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sm, nil
}

func (sm *SysMenuQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysRoleQuery{config: sm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sm.WithNamedRoles(alias, func(wq *SysRoleQuery) {
				*wq = *query
			})
		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysMenuQuery{config: sm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sm.withParent = query
		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysMenuQuery{config: sm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sm.WithNamedChildren(alias, func(wq *SysMenuQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type sysmenuPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysMenuPaginateOption
}

func newSysMenuPaginateArgs(rv map[string]interface{}) *sysmenuPaginateArgs {
	args := &sysmenuPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (smp *SysMenuParamQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysMenuParamQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return smp, nil
	}
	if err := smp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return smp, nil
}

func (smp *SysMenuParamQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type sysmenuparamPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysMenuParamPaginateOption
}

func newSysMenuParamPaginateArgs(rv map[string]interface{}) *sysmenuparamPaginateArgs {
	args := &sysmenuparamPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sop *SysOauthProviderQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysOauthProviderQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sop, nil
	}
	if err := sop.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sop, nil
}

func (sop *SysOauthProviderQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type sysoauthproviderPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysOauthProviderPaginateOption
}

func newSysOauthProviderPaginateArgs(rv map[string]interface{}) *sysoauthproviderPaginateArgs {
	args := &sysoauthproviderPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sr *SysRoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysRoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sr, nil
	}
	if err := sr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sr, nil
}

func (sr *SysRoleQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "menus":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysMenuQuery{config: sr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sr.WithNamedMenus(alias, func(wq *SysMenuQuery) {
				*wq = *query
			})
		case "role":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysUserQuery{config: sr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sr.WithNamedRole(alias, func(wq *SysUserQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type sysrolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysRolePaginateOption
}

func newSysRolePaginateArgs(rv map[string]interface{}) *sysrolePaginateArgs {
	args := &sysrolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (st *SysTokenQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysTokenQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return st, nil
	}
	if err := st.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return st, nil
}

func (st *SysTokenQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type systokenPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysTokenPaginateOption
}

func newSysTokenPaginateArgs(rv map[string]interface{}) *systokenPaginateArgs {
	args := &systokenPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (su *SysUserQuery) CollectFields(ctx context.Context, satisfies ...string) (*SysUserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return su, nil
	}
	if err := su.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return su, nil
}

func (su *SysUserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "role":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &SysRoleQuery{config: su.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			su.withRole = query
		}
	}
	return nil
}

type sysuserPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SysUserPaginateOption
}

func newSysUserPaginateArgs(rv map[string]interface{}) *sysuserPaginateArgs {
	args := &sysuserPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
