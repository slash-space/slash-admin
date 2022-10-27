// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/pkg/types"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SysMenu is the model entity for the SysMenu schema.
type SysMenu struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// parent menu ID
	ParentID uint64 `json:"parent_id,omitempty"`
	// menu level
	MenuLevel uint8 `json:"menu_level,omitempty"`
	// menu type: 0. group 1. menu
	MenuType uint8 `json:"menu_type,omitempty"`
	// index path
	Path string `json:"path,omitempty"`
	// index name
	Name string `json:"name,omitempty"`
	// redirect path
	Redirect string `json:"redirect,omitempty"`
	// the path of vue file
	Component string `json:"component,omitempty"`
	// numbers for sorting
	OrderNo uint8 `json:"order_no,omitempty"`
	// if true, disable
	Disabled bool `json:"disabled,omitempty"`
	// extra parameters
	Meta types.MenuMeta `json:"meta,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysMenuQuery when eager-loading is set.
	Edges SysMenuEdges `json:"edges"`
}

// SysMenuEdges holds the relations/edges for other nodes in the graph.
type SysMenuEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*SysRole `json:"roles,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *SysMenu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*SysMenu `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e SysMenuEdges) RolesOrErr() ([]*SysRole, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysMenuEdges) ParentOrErr() (*SysMenu, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sysmenu.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e SysMenuEdges) ChildrenOrErr() ([]*SysMenu, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldDisabled:
			values[i] = new(sql.NullBool)
		case sysmenu.FieldID, sysmenu.FieldParentID, sysmenu.FieldMenuLevel, sysmenu.FieldMenuType, sysmenu.FieldOrderNo:
			values[i] = new(sql.NullInt64)
		case sysmenu.FieldPath, sysmenu.FieldName, sysmenu.FieldRedirect, sysmenu.FieldComponent:
			values[i] = new(sql.NullString)
		case sysmenu.FieldCreatedAt, sysmenu.FieldUpdatedAt, sysmenu.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case sysmenu.FieldMeta:
			values[i] = new(types.MenuMeta)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysMenu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenu fields.
func (sm *SysMenu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = uint64(value.Int64)
		case sysmenu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				sm.ParentID = uint64(value.Int64)
			}
		case sysmenu.FieldMenuLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_level", values[i])
			} else if value.Valid {
				sm.MenuLevel = uint8(value.Int64)
			}
		case sysmenu.FieldMenuType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_type", values[i])
			} else if value.Valid {
				sm.MenuType = uint8(value.Int64)
			}
		case sysmenu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				sm.Path = value.String
			}
		case sysmenu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sm.Name = value.String
			}
		case sysmenu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				sm.Redirect = value.String
			}
		case sysmenu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				sm.Component = value.String
			}
		case sysmenu.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				sm.OrderNo = uint8(value.Int64)
			}
		case sysmenu.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				sm.Disabled = value.Bool
			}
		case sysmenu.FieldMeta:
			if value, ok := values[i].(*types.MenuMeta); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil {
				sm.Meta = *value
			}
		case sysmenu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = value.Time
			}
		case sysmenu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = value.Time
			}
		case sysmenu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sm.DeletedAt = new(time.Time)
				*sm.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryRoles queries the "roles" edge of the SysMenu entity.
func (sm *SysMenu) QueryRoles() *SysRoleQuery {
	return (&SysMenuClient{config: sm.config}).QueryRoles(sm)
}

// QueryParent queries the "parent" edge of the SysMenu entity.
func (sm *SysMenu) QueryParent() *SysMenuQuery {
	return (&SysMenuClient{config: sm.config}).QueryParent(sm)
}

// QueryChildren queries the "children" edge of the SysMenu entity.
func (sm *SysMenu) QueryChildren() *SysMenuQuery {
	return (&SysMenuClient{config: sm.config}).QueryChildren(sm)
}

// Update returns a builder for updating this SysMenu.
// Note that you need to call SysMenu.Unwrap() before calling this method if this SysMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SysMenu) Update() *SysMenuUpdateOne {
	return (&SysMenuClient{config: sm.config}).UpdateOne(sm)
}

// Unwrap unwraps the SysMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SysMenu) Unwrap() *SysMenu {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysMenu is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SysMenu) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.ParentID))
	builder.WriteString(", ")
	builder.WriteString("menu_level=")
	builder.WriteString(fmt.Sprintf("%v", sm.MenuLevel))
	builder.WriteString(", ")
	builder.WriteString("menu_type=")
	builder.WriteString(fmt.Sprintf("%v", sm.MenuType))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(sm.Path)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sm.Name)
	builder.WriteString(", ")
	builder.WriteString("redirect=")
	builder.WriteString(sm.Redirect)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(sm.Component)
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(fmt.Sprintf("%v", sm.OrderNo))
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", sm.Disabled))
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(fmt.Sprintf("%v", sm.Meta))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sm.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysMenus is a parsable slice of SysMenu.
type SysMenus []*SysMenu

func (sm SysMenus) config(cfg config) {
	for _i := range sm {
		sm[_i].config = cfg
	}
}
