// Code generated by ent, DO NOT EDIT.

package sysapi

import (
	"slash-admin/app/admin/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// APIGroup applies equality check predicate on the "api_group" field. It's identical to APIGroupEQ.
func APIGroup(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIGroup), v))
	})
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// APIGroupEQ applies the EQ predicate on the "api_group" field.
func APIGroupEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIGroup), v))
	})
}

// APIGroupNEQ applies the NEQ predicate on the "api_group" field.
func APIGroupNEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAPIGroup), v))
	})
}

// APIGroupIn applies the In predicate on the "api_group" field.
func APIGroupIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAPIGroup), v...))
	})
}

// APIGroupNotIn applies the NotIn predicate on the "api_group" field.
func APIGroupNotIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAPIGroup), v...))
	})
}

// APIGroupGT applies the GT predicate on the "api_group" field.
func APIGroupGT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAPIGroup), v))
	})
}

// APIGroupGTE applies the GTE predicate on the "api_group" field.
func APIGroupGTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAPIGroup), v))
	})
}

// APIGroupLT applies the LT predicate on the "api_group" field.
func APIGroupLT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAPIGroup), v))
	})
}

// APIGroupLTE applies the LTE predicate on the "api_group" field.
func APIGroupLTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAPIGroup), v))
	})
}

// APIGroupContains applies the Contains predicate on the "api_group" field.
func APIGroupContains(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAPIGroup), v))
	})
}

// APIGroupHasPrefix applies the HasPrefix predicate on the "api_group" field.
func APIGroupHasPrefix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAPIGroup), v))
	})
}

// APIGroupHasSuffix applies the HasSuffix predicate on the "api_group" field.
func APIGroupHasSuffix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAPIGroup), v))
	})
}

// APIGroupEqualFold applies the EqualFold predicate on the "api_group" field.
func APIGroupEqualFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAPIGroup), v))
	})
}

// APIGroupContainsFold applies the ContainsFold predicate on the "api_group" field.
func APIGroupContainsFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAPIGroup), v))
	})
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMethod), v))
	})
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMethod), v...))
	})
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMethod), v...))
	})
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMethod), v))
	})
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMethod), v))
	})
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMethod), v))
	})
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMethod), v))
	})
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMethod), v))
	})
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMethod), v))
	})
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMethod), v))
	})
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMethod), v))
	})
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMethod), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysApi {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysApi) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysApi) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysApi) predicate.SysApi {
	return predicate.SysApi(func(s *sql.Selector) {
		p(s.Not())
	})
}