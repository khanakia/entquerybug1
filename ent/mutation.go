// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/bug/ent/category"
	"entgo.io/bug/ent/post"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/user"

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
	TypeCategory = "Category"
	TypePost     = "Post"
	TypeUser     = "User"
)

// CategoryMutation represents an operation that mutates the Category nodes in the graph.
type CategoryMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	posts         map[int]struct{}
	removedposts  map[int]struct{}
	clearedposts  bool
	done          bool
	oldValue      func(context.Context) (*Category, error)
	predicates    []predicate.Category
}

var _ ent.Mutation = (*CategoryMutation)(nil)

// categoryOption allows management of the mutation configuration using functional options.
type categoryOption func(*CategoryMutation)

// newCategoryMutation creates new mutation for the Category entity.
func newCategoryMutation(c config, op Op, opts ...categoryOption) *CategoryMutation {
	m := &CategoryMutation{
		config:        c,
		op:            op,
		typ:           TypeCategory,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCategoryID sets the ID field of the mutation.
func withCategoryID(id int) categoryOption {
	return func(m *CategoryMutation) {
		var (
			err   error
			once  sync.Once
			value *Category
		)
		m.oldValue = func(ctx context.Context) (*Category, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Category.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCategory sets the old Category of the mutation.
func withCategory(node *Category) categoryOption {
	return func(m *CategoryMutation) {
		m.oldValue = func(context.Context) (*Category, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CategoryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CategoryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CategoryMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CategoryMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Category.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *CategoryMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CategoryMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *CategoryMutation) ClearName() {
	m.name = nil
	m.clearedFields[category.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *CategoryMutation) NameCleared() bool {
	_, ok := m.clearedFields[category.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *CategoryMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, category.FieldName)
}

// AddPostIDs adds the "posts" edge to the Post entity by ids.
func (m *CategoryMutation) AddPostIDs(ids ...int) {
	if m.posts == nil {
		m.posts = make(map[int]struct{})
	}
	for i := range ids {
		m.posts[ids[i]] = struct{}{}
	}
}

// ClearPosts clears the "posts" edge to the Post entity.
func (m *CategoryMutation) ClearPosts() {
	m.clearedposts = true
}

// PostsCleared reports if the "posts" edge to the Post entity was cleared.
func (m *CategoryMutation) PostsCleared() bool {
	return m.clearedposts
}

// RemovePostIDs removes the "posts" edge to the Post entity by IDs.
func (m *CategoryMutation) RemovePostIDs(ids ...int) {
	if m.removedposts == nil {
		m.removedposts = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.posts, ids[i])
		m.removedposts[ids[i]] = struct{}{}
	}
}

// RemovedPosts returns the removed IDs of the "posts" edge to the Post entity.
func (m *CategoryMutation) RemovedPostsIDs() (ids []int) {
	for id := range m.removedposts {
		ids = append(ids, id)
	}
	return
}

// PostsIDs returns the "posts" edge IDs in the mutation.
func (m *CategoryMutation) PostsIDs() (ids []int) {
	for id := range m.posts {
		ids = append(ids, id)
	}
	return
}

// ResetPosts resets all changes to the "posts" edge.
func (m *CategoryMutation) ResetPosts() {
	m.posts = nil
	m.clearedposts = false
	m.removedposts = nil
}

// Where appends a list predicates to the CategoryMutation builder.
func (m *CategoryMutation) Where(ps ...predicate.Category) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CategoryMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Category).
func (m *CategoryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CategoryMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, category.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CategoryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case category.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CategoryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case category.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Category field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CategoryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case category.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Category field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CategoryMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CategoryMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CategoryMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Category numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CategoryMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(category.FieldName) {
		fields = append(fields, category.FieldName)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CategoryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CategoryMutation) ClearField(name string) error {
	switch name {
	case category.FieldName:
		m.ClearName()
		return nil
	}
	return fmt.Errorf("unknown Category nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CategoryMutation) ResetField(name string) error {
	switch name {
	case category.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Category field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CategoryMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.posts != nil {
		edges = append(edges, category.EdgePosts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CategoryMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case category.EdgePosts:
		ids := make([]ent.Value, 0, len(m.posts))
		for id := range m.posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CategoryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedposts != nil {
		edges = append(edges, category.EdgePosts)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CategoryMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case category.EdgePosts:
		ids := make([]ent.Value, 0, len(m.removedposts))
		for id := range m.removedposts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CategoryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedposts {
		edges = append(edges, category.EdgePosts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CategoryMutation) EdgeCleared(name string) bool {
	switch name {
	case category.EdgePosts:
		return m.clearedposts
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CategoryMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Category unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CategoryMutation) ResetEdge(name string) error {
	switch name {
	case category.EdgePosts:
		m.ResetPosts()
		return nil
	}
	return fmt.Errorf("unknown Category edge %s", name)
}

// PostMutation represents an operation that mutates the Post nodes in the graph.
type PostMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	clearedFields   map[string]struct{}
	category        *int
	clearedcategory bool
	done            bool
	oldValue        func(context.Context) (*Post, error)
	predicates      []predicate.Post
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
func withPostID(id int) postOption {
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

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PostMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PostMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Post.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *PostMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PostMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *PostMutation) ClearName() {
	m.name = nil
	m.clearedFields[post.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *PostMutation) NameCleared() bool {
	_, ok := m.clearedFields[post.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *PostMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, post.FieldName)
}

// SetCategoryID sets the "category_id" field.
func (m *PostMutation) SetCategoryID(i int) {
	m.category = &i
}

// CategoryID returns the value of the "category_id" field in the mutation.
func (m *PostMutation) CategoryID() (r int, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategoryID returns the old "category_id" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldCategoryID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCategoryID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCategoryID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategoryID: %w", err)
	}
	return oldValue.CategoryID, nil
}

// ClearCategoryID clears the value of the "category_id" field.
func (m *PostMutation) ClearCategoryID() {
	m.category = nil
	m.clearedFields[post.FieldCategoryID] = struct{}{}
}

// CategoryIDCleared returns if the "category_id" field was cleared in this mutation.
func (m *PostMutation) CategoryIDCleared() bool {
	_, ok := m.clearedFields[post.FieldCategoryID]
	return ok
}

// ResetCategoryID resets all changes to the "category_id" field.
func (m *PostMutation) ResetCategoryID() {
	m.category = nil
	delete(m.clearedFields, post.FieldCategoryID)
}

// ClearCategory clears the "category" edge to the Category entity.
func (m *PostMutation) ClearCategory() {
	m.clearedcategory = true
}

// CategoryCleared reports if the "category" edge to the Category entity was cleared.
func (m *PostMutation) CategoryCleared() bool {
	return m.CategoryIDCleared() || m.clearedcategory
}

// CategoryIDs returns the "category" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// CategoryID instead. It exists only for internal usage by the builders.
func (m *PostMutation) CategoryIDs() (ids []int) {
	if id := m.category; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCategory resets all changes to the "category" edge.
func (m *PostMutation) ResetCategory() {
	m.category = nil
	m.clearedcategory = false
}

// Where appends a list predicates to the PostMutation builder.
func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PostMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Post).
func (m *PostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, post.FieldName)
	}
	if m.category != nil {
		fields = append(fields, post.FieldCategoryID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldName:
		return m.Name()
	case post.FieldCategoryID:
		return m.CategoryID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldName:
		return m.OldName(ctx)
	case post.FieldCategoryID:
		return m.OldCategoryID(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case post.FieldCategoryID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategoryID(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PostMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
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
	var fields []string
	if m.FieldCleared(post.FieldName) {
		fields = append(fields, post.FieldName)
	}
	if m.FieldCleared(post.FieldCategoryID) {
		fields = append(fields, post.FieldCategoryID)
	}
	return fields
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
	switch name {
	case post.FieldName:
		m.ClearName()
		return nil
	case post.FieldCategoryID:
		m.ClearCategoryID()
		return nil
	}
	return fmt.Errorf("unknown Post nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldName:
		m.ResetName()
		return nil
	case post.FieldCategoryID:
		m.ResetCategoryID()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.category != nil {
		edges = append(edges, post.EdgeCategory)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PostMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeCategory:
		if id := m.category; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcategory {
		edges = append(edges, post.EdgeCategory)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PostMutation) EdgeCleared(name string) bool {
	switch name {
	case post.EdgeCategory:
		return m.clearedcategory
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PostMutation) ClearEdge(name string) error {
	switch name {
	case post.EdgeCategory:
		m.ClearCategory()
		return nil
	}
	return fmt.Errorf("unknown Post unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PostMutation) ResetEdge(name string) error {
	switch name {
	case post.EdgeCategory:
		m.ResetCategory()
		return nil
	}
	return fmt.Errorf("unknown Post edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
