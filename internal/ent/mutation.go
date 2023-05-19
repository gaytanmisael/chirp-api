// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chirp-api/internal/ent/post"
	"chirp-api/internal/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePost = "Post"
)

// PostMutation represents an operation that mutates the Post nodes in the graph.
type PostMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	content       *string
	authorId      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Post, error)
	predicates    []predicate.Post
}

var _ ent.Mutation = (*PostMutation)(nil)

// postOption allows management of the mutation configuration using functional options.
type postOption func(*PostMutation)

// newPostMutation creates new mutation for the Post entity.
func newPostMutation(c config, op Op, opts ...postOption) *PostMutation {
	m := &PostMutation{
		config:        c,
		op:            op,
		typ:           TypePost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPostID sets the ID field of the mutation.
func withPostID(id uuid.UUID) postOption {
	return func(m *PostMutation) {
		var (
			err   error
			once  sync.Once
			value *Post
		)
		m.oldValue = func(ctx context.Context) (*Post, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Post.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPost sets the old Post of the mutation.
func withPost(node *Post) postOption {
	return func(m *PostMutation) {
		m.oldValue = func(context.Context) (*Post, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Post entities.
func (m *PostMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PostMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PostMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Post.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *PostMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PostMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PostMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *PostMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *PostMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *PostMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetContent sets the "content" field.
func (m *PostMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *PostMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *PostMutation) ResetContent() {
	m.content = nil
}

// SetAuthorId sets the "authorId" field.
func (m *PostMutation) SetAuthorId(s string) {
	m.authorId = &s
}

// AuthorId returns the value of the "authorId" field in the mutation.
func (m *PostMutation) AuthorId() (r string, exists bool) {
	v := m.authorId
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthorId returns the old "authorId" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldAuthorId(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAuthorId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAuthorId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthorId: %w", err)
	}
	return oldValue.AuthorId, nil
}

// ResetAuthorId resets all changes to the "authorId" field.
func (m *PostMutation) ResetAuthorId() {
	m.authorId = nil
}

// Where appends a list predicates to the PostMutation builder.
func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PostMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PostMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Post, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PostMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PostMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Post).
func (m *PostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, post.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, post.FieldUpdatedAt)
	}
	if m.content != nil {
		fields = append(fields, post.FieldContent)
	}
	if m.authorId != nil {
		fields = append(fields, post.FieldAuthorId)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldCreatedAt:
		return m.CreatedAt()
	case post.FieldUpdatedAt:
		return m.UpdatedAt()
	case post.FieldContent:
		return m.Content()
	case post.FieldAuthorId:
		return m.AuthorId()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case post.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case post.FieldContent:
		return m.OldContent(ctx)
	case post.FieldAuthorId:
		return m.OldAuthorId(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case post.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case post.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case post.FieldAuthorId:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthorId(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PostMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Post numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PostMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PostMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Post nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case post.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case post.FieldContent:
		m.ResetContent()
		return nil
	case post.FieldAuthorId:
		m.ResetAuthorId()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PostMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PostMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PostMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Post unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PostMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Post edge %s", name)
}
