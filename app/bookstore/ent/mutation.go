// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/bookstore/ent/bsuser"
	"slash-admin/app/bookstore/ent/predicate"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBsUser = "BsUser"
)

// BsUserMutation represents an operation that mutates the BsUser nodes in the graph.
type BsUserMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	username      *string
	points        *uint64
	addpoints     *int64
	status        *bsuser.Status
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*BsUser, error)
	predicates    []predicate.BsUser
}

var _ ent.Mutation = (*BsUserMutation)(nil)

// bsuserOption allows management of the mutation configuration using functional options.
type bsuserOption func(*BsUserMutation)

// newBsUserMutation creates new mutation for the BsUser entity.
func newBsUserMutation(c config, op Op, opts ...bsuserOption) *BsUserMutation {
	m := &BsUserMutation{
		config:        c,
		op:            op,
		typ:           TypeBsUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBsUserID sets the ID field of the mutation.
func withBsUserID(id uint64) bsuserOption {
	return func(m *BsUserMutation) {
		var (
			err   error
			once  sync.Once
			value *BsUser
		)
		m.oldValue = func(ctx context.Context) (*BsUser, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().BsUser.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBsUser sets the old BsUser of the mutation.
func withBsUser(node *BsUser) bsuserOption {
	return func(m *BsUserMutation) {
		m.oldValue = func(context.Context) (*BsUser, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BsUserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BsUserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of BsUser entities.
func (m *BsUserMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BsUserMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BsUserMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().BsUser.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *BsUserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *BsUserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the BsUser entity.
// If the BsUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsUserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *BsUserMutation) ResetUsername() {
	m.username = nil
}

// SetPoints sets the "points" field.
func (m *BsUserMutation) SetPoints(u uint64) {
	m.points = &u
	m.addpoints = nil
}

// Points returns the value of the "points" field in the mutation.
func (m *BsUserMutation) Points() (r uint64, exists bool) {
	v := m.points
	if v == nil {
		return
	}
	return *v, true
}

// OldPoints returns the old "points" field's value of the BsUser entity.
// If the BsUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsUserMutation) OldPoints(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPoints is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPoints requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPoints: %w", err)
	}
	return oldValue.Points, nil
}

// AddPoints adds u to the "points" field.
func (m *BsUserMutation) AddPoints(u int64) {
	if m.addpoints != nil {
		*m.addpoints += u
	} else {
		m.addpoints = &u
	}
}

// AddedPoints returns the value that was added to the "points" field in this mutation.
func (m *BsUserMutation) AddedPoints() (r int64, exists bool) {
	v := m.addpoints
	if v == nil {
		return
	}
	return *v, true
}

// ResetPoints resets all changes to the "points" field.
func (m *BsUserMutation) ResetPoints() {
	m.points = nil
	m.addpoints = nil
}

// SetStatus sets the "status" field.
func (m *BsUserMutation) SetStatus(b bsuser.Status) {
	m.status = &b
}

// Status returns the value of the "status" field in the mutation.
func (m *BsUserMutation) Status() (r bsuser.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the BsUser entity.
// If the BsUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsUserMutation) OldStatus(ctx context.Context) (v bsuser.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *BsUserMutation) ResetStatus() {
	m.status = nil
}

// Where appends a list predicates to the BsUserMutation builder.
func (m *BsUserMutation) Where(ps ...predicate.BsUser) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *BsUserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (BsUser).
func (m *BsUserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BsUserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.username != nil {
		fields = append(fields, bsuser.FieldUsername)
	}
	if m.points != nil {
		fields = append(fields, bsuser.FieldPoints)
	}
	if m.status != nil {
		fields = append(fields, bsuser.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BsUserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case bsuser.FieldUsername:
		return m.Username()
	case bsuser.FieldPoints:
		return m.Points()
	case bsuser.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BsUserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case bsuser.FieldUsername:
		return m.OldUsername(ctx)
	case bsuser.FieldPoints:
		return m.OldPoints(ctx)
	case bsuser.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown BsUser field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BsUserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case bsuser.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case bsuser.FieldPoints:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPoints(v)
		return nil
	case bsuser.FieldStatus:
		v, ok := value.(bsuser.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown BsUser field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BsUserMutation) AddedFields() []string {
	var fields []string
	if m.addpoints != nil {
		fields = append(fields, bsuser.FieldPoints)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BsUserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case bsuser.FieldPoints:
		return m.AddedPoints()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BsUserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case bsuser.FieldPoints:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPoints(v)
		return nil
	}
	return fmt.Errorf("unknown BsUser numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BsUserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BsUserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BsUserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown BsUser nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BsUserMutation) ResetField(name string) error {
	switch name {
	case bsuser.FieldUsername:
		m.ResetUsername()
		return nil
	case bsuser.FieldPoints:
		m.ResetPoints()
		return nil
	case bsuser.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown BsUser field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BsUserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BsUserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BsUserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BsUserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BsUserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BsUserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BsUserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown BsUser unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BsUserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown BsUser edge %s", name)
}