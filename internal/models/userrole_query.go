// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserRoleQuery is the builder for querying UserRole entities.
type UserRoleQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.UserRole
	// eager-loading edges.
	withUsers *UserQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (urq *UserRoleQuery) Where(ps ...predicate.UserRole) *UserRoleQuery {
	urq.predicates = append(urq.predicates, ps...)
	return urq
}

// Limit adds a limit step to the query.
func (urq *UserRoleQuery) Limit(limit int) *UserRoleQuery {
	urq.limit = &limit
	return urq
}

// Offset adds an offset step to the query.
func (urq *UserRoleQuery) Offset(offset int) *UserRoleQuery {
	urq.offset = &offset
	return urq
}

// Order adds an order step to the query.
func (urq *UserRoleQuery) Order(o ...OrderFunc) *UserRoleQuery {
	urq.order = append(urq.order, o...)
	return urq
}

// QueryUsers chains the current query on the users edge.
func (urq *UserRoleQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: urq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := urq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userrole.UsersTable, userrole.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(urq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserRole entity in the query. Returns *NotFoundError when no userrole was found.
func (urq *UserRoleQuery) First(ctx context.Context) (*UserRole, error) {
	nodes, err := urq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urq *UserRoleQuery) FirstX(ctx context.Context) *UserRole {
	node, err := urq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserRole id in the query. Returns *NotFoundError when no id was found.
func (urq *UserRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urq *UserRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := urq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserRole entity in the query, returns an error if not exactly one entity was returned.
func (urq *UserRoleQuery) Only(ctx context.Context) (*UserRole, error) {
	nodes, err := urq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userrole.Label}
	default:
		return nil, &NotSingularError{userrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urq *UserRoleQuery) OnlyX(ctx context.Context) *UserRole {
	node, err := urq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only UserRole id in the query, returns an error if not exactly one id was returned.
func (urq *UserRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = &NotSingularError{userrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urq *UserRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := urq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserRoles.
func (urq *UserRoleQuery) All(ctx context.Context) ([]*UserRole, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return urq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (urq *UserRoleQuery) AllX(ctx context.Context) []*UserRole {
	nodes, err := urq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserRole ids.
func (urq *UserRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := urq.Select(userrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urq *UserRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := urq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urq *UserRoleQuery) Count(ctx context.Context) (int, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return urq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (urq *UserRoleQuery) CountX(ctx context.Context) int {
	count, err := urq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urq *UserRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return urq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (urq *UserRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := urq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urq *UserRoleQuery) Clone() *UserRoleQuery {
	if urq == nil {
		return nil
	}
	return &UserRoleQuery{
		config:     urq.config,
		limit:      urq.limit,
		offset:     urq.offset,
		order:      append([]OrderFunc{}, urq.order...),
		predicates: append([]predicate.UserRole{}, urq.predicates...),
		withUsers:  urq.withUsers.Clone(),
		// clone intermediate query.
		sql:  urq.sql.Clone(),
		path: urq.path,
	}
}

//  WithUsers tells the query-builder to eager-loads the nodes that are connected to
// the "users" edge. The optional arguments used to configure the query builder of the edge.
func (urq *UserRoleQuery) WithUsers(opts ...func(*UserQuery)) *UserRoleQuery {
	query := &UserQuery{config: urq.config}
	for _, opt := range opts {
		opt(query)
	}
	urq.withUsers = query
	return urq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserRole.Query().
//		GroupBy(userrole.FieldName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (urq *UserRoleQuery) GroupBy(field string, fields ...string) *UserRoleGroupBy {
	group := &UserRoleGroupBy{config: urq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return urq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.UserRole.Query().
//		Select(userrole.FieldName).
//		Scan(ctx, &v)
//
func (urq *UserRoleQuery) Select(field string, fields ...string) *UserRoleSelect {
	selector := &UserRoleSelect{config: urq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return urq.sqlQuery(), nil
	}
	return selector
}

func (urq *UserRoleQuery) prepareQuery(ctx context.Context) error {
	if urq.path != nil {
		prev, err := urq.path(ctx)
		if err != nil {
			return err
		}
		urq.sql = prev
	}
	return nil
}

func (urq *UserRoleQuery) sqlAll(ctx context.Context) ([]*UserRole, error) {
	var (
		nodes       = []*UserRole{}
		withFKs     = urq.withFKs
		_spec       = urq.querySpec()
		loadedTypes = [1]bool{
			urq.withUsers != nil,
		}
	)
	if urq.withUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userrole.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &UserRole{config: urq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("models: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, urq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := urq.withUsers; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*UserRole)
		for i := range nodes {
			if fk := nodes[i].user_user_roles; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_user_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Users = n
			}
		}
	}

	return nodes, nil
}

func (urq *UserRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urq.querySpec()
	return sqlgraph.CountNodes(ctx, urq.driver, _spec)
}

func (urq *UserRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := urq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %v", err)
	}
	return n > 0, nil
}

func (urq *UserRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
		From:   urq.sql,
		Unique: true,
	}
	if ps := urq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, userrole.ValidColumn)
			}
		}
	}
	return _spec
}

func (urq *UserRoleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(urq.driver.Dialect())
	t1 := builder.Table(userrole.Table)
	selector := builder.Select(t1.Columns(userrole.Columns...)...).From(t1)
	if urq.sql != nil {
		selector = urq.sql
		selector.Select(selector.Columns(userrole.Columns...)...)
	}
	for _, p := range urq.predicates {
		p(selector)
	}
	for _, p := range urq.order {
		p(selector, userrole.ValidColumn)
	}
	if offset := urq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserRoleGroupBy is the builder for group-by UserRole entities.
type UserRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urgb *UserRoleGroupBy) Aggregate(fns ...AggregateFunc) *UserRoleGroupBy {
	urgb.fns = append(urgb.fns, fns...)
	return urgb
}

// Scan applies the group-by query and scan the result into the given value.
func (urgb *UserRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := urgb.path(ctx)
	if err != nil {
		return err
	}
	urgb.sql = query
	return urgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (urgb *UserRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := urgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(urgb.fields) > 1 {
		return nil, errors.New("models: UserRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := urgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (urgb *UserRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := urgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = urgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (urgb *UserRoleGroupBy) StringX(ctx context.Context) string {
	v, err := urgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(urgb.fields) > 1 {
		return nil, errors.New("models: UserRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := urgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (urgb *UserRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := urgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = urgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (urgb *UserRoleGroupBy) IntX(ctx context.Context) int {
	v, err := urgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(urgb.fields) > 1 {
		return nil, errors.New("models: UserRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := urgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (urgb *UserRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := urgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = urgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (urgb *UserRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := urgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(urgb.fields) > 1 {
		return nil, errors.New("models: UserRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := urgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (urgb *UserRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := urgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (urgb *UserRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = urgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (urgb *UserRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := urgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (urgb *UserRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range urgb.fields {
		if !userrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := urgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (urgb *UserRoleGroupBy) sqlQuery() *sql.Selector {
	selector := urgb.sql
	columns := make([]string, 0, len(urgb.fields)+len(urgb.fns))
	columns = append(columns, urgb.fields...)
	for _, fn := range urgb.fns {
		columns = append(columns, fn(selector, userrole.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(urgb.fields...)
}

// UserRoleSelect is the builder for select fields of UserRole entities.
type UserRoleSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (urs *UserRoleSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := urs.path(ctx)
	if err != nil {
		return err
	}
	urs.sql = query
	return urs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (urs *UserRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := urs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(urs.fields) > 1 {
		return nil, errors.New("models: UserRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := urs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (urs *UserRoleSelect) StringsX(ctx context.Context) []string {
	v, err := urs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = urs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (urs *UserRoleSelect) StringX(ctx context.Context) string {
	v, err := urs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(urs.fields) > 1 {
		return nil, errors.New("models: UserRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := urs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (urs *UserRoleSelect) IntsX(ctx context.Context) []int {
	v, err := urs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = urs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (urs *UserRoleSelect) IntX(ctx context.Context) int {
	v, err := urs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(urs.fields) > 1 {
		return nil, errors.New("models: UserRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := urs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (urs *UserRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := urs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = urs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (urs *UserRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := urs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(urs.fields) > 1 {
		return nil, errors.New("models: UserRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := urs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (urs *UserRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := urs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (urs *UserRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = urs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = fmt.Errorf("models: UserRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (urs *UserRoleSelect) BoolX(ctx context.Context) bool {
	v, err := urs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (urs *UserRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range urs.fields {
		if !userrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := urs.sqlQuery().Query()
	if err := urs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (urs *UserRoleSelect) sqlQuery() sql.Querier {
	selector := urs.sql
	selector.Select(selector.Columns(urs.fields...)...)
	return selector
}
